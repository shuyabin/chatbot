package chatbot

import (
	"context"
	"io"
	"net"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotdb "github.com/zhs007/chatbot/db"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// Serv - service
type Serv struct {
	cfg         *Config
	lis         net.Listener
	grpcServ    *grpc.Server
	dbAppServ   *chatbotdb.AppServDB
	lstPlugins  *PluginsList
	lstPlugins2 *PluginsList
	mgrUser     UserMgr
	mgrText     *TextMgr
	cmds        *CommondsList
}

// NewChatBotServ -
func NewChatBotServ(cfg *Config, mgr UserMgr) (*Serv, error) {
	if cfg == nil {
		return nil, chatbotbase.ErrNoConfig
	}

	db, err := chatbotdb.NewAppServDB(cfg.DBPath, "", cfg.DBEngine)
	if err != nil {
		return nil, err
	}

	mgrText, err := NewTextMgr(cfg)
	if err != nil {
		return nil, err
	}

	lis, err := net.Listen("tcp", cfg.BindAddr)
	if err != nil {
		chatbotbase.Error("NewChatBotServ", zap.Error(err))

		return nil, err
	}

	chatbotbase.Info("Listen", zap.String("addr", cfg.BindAddr))

	grpcServ := grpc.NewServer()

	serv := &Serv{
		cfg:         cfg,
		lis:         lis,
		grpcServ:    grpcServ,
		lstPlugins:  NewPluginsList(),
		lstPlugins2: NewPluginsList(),
		mgrUser:     mgr,
		mgrText:     mgrText,
		cmds:        NewCommondsList(),
	}

	for _, v := range cfg.Plugins {
		err = serv.lstPlugins.AddPlugin(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cfg.PluginsSecondLine {
		err = serv.lstPlugins2.AddPlugin(v)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range cfg.Commands {
		err = serv.cmds.AddCommand(v)
		if err != nil {
			return nil, err
		}
	}

	chatbotpb.RegisterChatBotServiceServer(grpcServ, serv)

	serv.dbAppServ = db

	chatbotbase.Info("NewChatBotServ OK.")

	return serv, nil
}

// Init - initial service
func (serv *Serv) Init(ctx context.Context) error {
	return serv.dbAppServ.Init(ctx, serv.cfg.AppServ)
}

// Start - start a service
func (serv *Serv) Start(ctx context.Context) error {
	return serv.grpcServ.Serve(serv.lis)
}

// Stop - stop service
func (serv *Serv) Stop() {
	serv.lis.Close()

	return
}

// RegisterAppService - register app service
func (serv *Serv) RegisterAppService(ctx context.Context, ras *chatbotpb.RegisterAppService) (
	*chatbotpb.ReplyRegisterAppService, error) {

	asi, err := serv.dbAppServ.GetAppServ(ctx, ras.AppServ.Token)
	if err != nil {
		return &chatbotpb.ReplyRegisterAppService{
			Error: err.Error(),
		}, nil
	}

	if asi.AppType != ras.AppServ.AppType {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServType.Error(),
		}, nil
	}

	if asi.Username != ras.AppServ.Username {
		return &chatbotpb.ReplyRegisterAppService{
			Error: chatbotbase.ErrInvalidAppServUserName.Error(),
		}, nil
	}

	asi, err = serv.dbAppServ.GenerateSessionID(ctx, asi)
	if err != nil {
		return &chatbotpb.ReplyRegisterAppService{
			Error: err.Error(),
		}, nil
	}

	return &chatbotpb.ReplyRegisterAppService{
		AppType:   ras.AppServ.AppType,
		SessionID: asi.Sessionid,
	}, nil
}

// SendChat - send chat
func (serv *Serv) SendChat(scs chatbotpb.ChatBotService_SendChatServer) error {
	var lst []*chatbotpb.ChatMsgStream

	for {
		in, err := scs.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			serv.replySendChatErr(scs, err)

			return err
		}

		isvalidtoken, err := serv.dbAppServ.CheckTokenSessionID(scs.Context(), in.Token, in.SessionID)
		if err != nil {
			serv.replySendChatErr(scs, err)

			return err
		}

		if !isvalidtoken {
			serv.replySendChatErr(scs, chatbotbase.ErrServInvalidToken)

			return chatbotbase.ErrServInvalidToken
		}

		lst = append(lst, in)
	}

	cd, err := BuildChatMsg(lst)
	if err != nil {
		serv.replySendChatErr(scs, err)

		return err
	}

	if cd == nil {
		serv.replySendChatErr(scs, chatbotbase.ErrInvalidStream)

		return chatbotbase.ErrInvalidStream
	}

	ui, ud, err := serv.mgrUser.GetAppUserInfo(scs.Context(), cd.Token, cd.Uai)
	if err != nil || ui == nil {
		serv.replySendChatErr(scs, chatbotbase.ErrServInvalidUserInfo)

		return chatbotbase.ErrServInvalidUserInfo
	}

	lstret, err := serv.lstPlugins.OnMessage(scs.Context(), serv, cd, ui, ud)
	if err != nil {
		serv.replySendChatErr(scs, err)

		return err
	}

	lstret2, err := serv.lstPlugins2.OnMessageEx(scs.Context(), serv, cd, ui, ud)
	if err != nil {
		serv.replySendChatErr(scs, err)

		return err
	}

	if lstret2 != nil {
		lstret = append(lstret, lstret2...)
	}

	for _, v := range lstret {
		lststream, err := BuildChatMsgStream(v)
		if err != nil {
			serv.replySendChatErr(scs, err)

			return err
		}

		for _, sv := range lststream {
			err = scs.Send(sv)
			if err != nil {
				serv.replySendChatErr(scs, err)

				return err
			}
		}
	}

	return nil
}

// replySendChatErr - reply a error for SendChat
func (serv *Serv) replySendChatErr(scs chatbotpb.ChatBotService_SendChatServer, err error) error {
	if err == nil {
		return serv.replySendChatErr(scs, chatbotbase.ErrServInvalidErr)
	}

	reply := &chatbotpb.ChatMsgStream{
		Error: err.Error(),
	}

	return scs.Send(reply)
}

// RequestChat - request chat
func (serv *Serv) RequestChat(req *chatbotpb.RequestChatData,
	ecs chatbotpb.ChatBotService_RequestChatServer) error {

	return nil
}
