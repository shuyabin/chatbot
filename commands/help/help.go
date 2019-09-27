package chatbotcmdhelp

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/zhs007/chatbot"
	chatbotpb "github.com/zhs007/chatbot/proto"
)

// cmdHelp - command help
type cmdHelp struct {
}

// RunCommand - run command
func (cmd *cmdHelp) RunCommand(ctx context.Context, serv *chatbot.Serv, params proto.Message,
	chat *chatbotpb.ChatMsg) ([]*chatbotpb.ChatMsg, error) {

	return nil, nil
}

// ParseCommandLine - parse command line
func (cmd *cmdHelp) ParseCommandLine(cmdline []string, chat *chatbotpb.ChatMsg) (
	proto.Message, error) {

	return nil, nil
}

// RegisterCommand - register command
func RegisterCommand() {
	chatbot.RegisterCommand("help", &cmdHelp{})
}
