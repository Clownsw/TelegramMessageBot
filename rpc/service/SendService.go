package service

import (
	"errors"
	"github.com/Clownsw/TelegramMessageBot/common"
	"github.com/Clownsw/TelegramMessageBot/rpc"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
)

type SendService struct {
	rpc.SendServiceServer
}

func (sendService *SendService) Send(stream rpc.SendService_SendServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&rpc.ResponseMessage{
				Code: common.RpcSendStatusOk,
				Msg:  []string{common.RpcSendStatusOkMsg},
			})
		}

		if err != nil {
			return err
		}

		message := tgbotapi.NewMessage(req.ChatId, req.SendMessage)

		switch req.Type {
		case tgbotapi.ModeHTML:
			message.ParseMode = tgbotapi.ModeHTML

		case tgbotapi.ModeMarkdown:
			message.ParseMode = tgbotapi.ModeMarkdown

		case tgbotapi.ModeMarkdownV2:
			message.ParseMode = tgbotapi.ModeMarkdownV2

		default:
			return errors.New("unknown type")
		}

		_, err = common.BotApi.Send(&message)
		if err != nil {
			return err
		}
	}
}
