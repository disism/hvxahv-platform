package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/db"
	"log"
	"net"
)

type server struct {
	pb.AccountsServer
}

/**
 Accounts 功能的 GRPC 微服务的服务端实现
*/
func main() {

	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	db.InitRedis()
	if err := db.InitMariaDB(); err != nil {
		log.Println("数据库初始化失败：", err)
	}
	p := viper.GetString("port.accounts")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
	if err != nil {
		fmt.Printf("Accounts gRPC Services Failed to Listen: %v", err)
		return
	}

	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &server{})
	reflection.Register(s)

	// 在控制台中打印服务启动 log 通知 并启动一个协程通知 Bot Account 服务已经开启
	log.Println("Accounts gRPC Services is running....", p)
	go bot.ServicesRunningNotice("account", p)
	if err := s.Serve(lis); err != nil {
		fmt.Printf("Accounts gRPC Services failed to start: %v", err)
	}
}
