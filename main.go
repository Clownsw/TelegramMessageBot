package main

import (
	"fmt"
	"github.com/Clownsw/TelegramMessageBot/common"
	"github.com/Clownsw/TelegramMessageBot/rpc"
	"github.com/Clownsw/TelegramMessageBot/rpc/service"
	botApi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gobuffalo/envy"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

var err error

func init() {
	logrus.SetReportCaller(true)

	// 初始化配置
	if err := envy.Load(".env"); err != nil {
		logrus.Panic(fmt.Sprintf("can not load .env, err: %s", err.Error()))
	}

	common.Config.Token, err = envy.MustGet("token")
	if err != nil || common.Config.Token == "" {
		logrus.Panic(fmt.Sprintf("can not read env token field, err: %s", err.Error()))
	}

	common.Config.Addr, err = envy.MustGet("addr")
	if err != nil || common.Config.Addr == "" {
		logrus.Panic(fmt.Sprintf("can not read env addr field, err: %s", err.Error()))
	}

	common.Config.Port, err = envy.MustGet("port")
	if err != nil || common.Config.Port == "" {
		logrus.Panic(fmt.Sprintf("can not read env port field, err: %s", err.Error()))
	}
}

func main() {
	// 初始化Telegram Bot
	common.BotApi, err = botApi.NewBotAPI(common.Config.Token)
	if err != nil {
		logrus.Panic(err)
	}

	common.BotApi.Debug = true

	logrus.Info("Authorized on account", common.BotApi.Self.UserName)

	// 初始化GRPC服务
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", common.Config.Addr, common.Config.Port))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	rpc.RegisterSendServiceServer(grpcServer, &service.SendService{})

	if err := grpcServer.Serve(listener); err != nil {
		logrus.Panic(err)
	}
}
