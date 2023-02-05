package common

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

const (
	RpcSendStatusOk    int32 = 200
	RpcSendStatusOkMsg       = "OK"
	RpcSendStatusError int32 = 500
)

var BotApi *tgbotapi.BotAPI
var Config GlobalConfig

type GlobalConfig struct {
	Token string
	Addr  string
	Port  string
}
