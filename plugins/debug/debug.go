package chatbotdebugplugin

import (
	"context"

	"github.com/golang/protobuf/proto"
	chatbot "github.com/zhs007/chatbot"
	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/proto"
	"go.uber.org/zap"
)

type debugPlugin struct {
}

// OnMessage - get message
func (dbp *debugPlugin) OnMessage(ctx context.Context, serv *chatbot.Serv, chat *chatbotpb.ChatMsg,
	ui *chatbotpb.UserInfo, ud proto.Message) ([]*chatbotpb.ChatMsg, error) {

	if serv == nil {
		return nil, chatbotbase.ErrPluginInvalidServ
	}

	if serv.MgrText == nil {
		return nil, chatbotbase.ErrPluginInvalidServMgrText
	}

	lang := serv.GetChatMsgLang(chat)

	mParams, err := serv.BuildBasicParamsMap(chat, ui, lang)
	if err != nil {
		return nil, err
	}

	locale, err := serv.MgrText.GetLocalizer(lang)
	if err != nil {
		return nil, err
	}

	var lst []*chatbotpb.ChatMsg

	msgigetit, err := chatbot.NewChatMsgWithText(locale, "igetit", mParams, chat.Uai)
	if err != nil {
		return nil, err
	}

	lst = append(lst, msgigetit)

	msgyousaid, err := chatbot.NewChatMsgWithText(locale, "yousaid", mParams, chat.Uai)
	if err != nil {
		return nil, err
	}

	lst = append(lst, msgyousaid)

	strui, err := chatbotbase.JSONFormat(ui)
	if err != nil {
		chatbotbase.Warn("debugPlugin.OnMessage:UI",
			zap.Error(err))
	} else {
		msgui := &chatbotpb.ChatMsg{
			Msg: strui,
			Uai: chat.Uai,
		}

		lst = append(lst, msgui)
	}

	return lst, nil
}

// OnStart - on start
func (dbp *debugPlugin) OnStart(ctx context.Context) error {
	return nil
}

// GetPluginName - get plugin name
func (dbp *debugPlugin) GetPluginName() string {
	return "debug"
}

// RegisterPlugin - register debug plugin
func RegisterPlugin() error {
	return chatbot.RegPlugin(&debugPlugin{})
}
