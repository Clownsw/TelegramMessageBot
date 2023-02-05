package service

import (
	"github.com/Clownsw/TelegramMessageBot/common"
	"github.com/Clownsw/TelegramMessageBot/rpc"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
)

type SendService struct {
	rpc.SendServiceServer
}

func (sendService *SendService) Send(stream rpc.SendService_SendServer) error {
	var resultMsgSlice []string

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			code := common.RpcSendStatusError

			if len(resultMsgSlice) == 0 {
				code = common.RpcSendStatusOk
				resultMsgSlice = append(resultMsgSlice, common.RpcSendStatusOkMsg)
			}

			return stream.SendAndClose(&rpc.ResponseMessage{
				Code: code,
				Msg:  resultMsgSlice,
			})
		}

		if err != nil {
			return err
		}

		_, err = common.BotApi.Send(tgbotapi.NewMessage(req.ChatId, req.SendMessage))
		if err != nil {
			resultMsgSlice = append(resultMsgSlice, err.Error())
		}
	}
}
