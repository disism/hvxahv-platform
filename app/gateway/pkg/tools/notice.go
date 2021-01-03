package tools

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/util/v1"
	"log"
)

// 通知给 Telegram Admin
func SendTGNotice(m string) {
	conn, err := grpc.Dial(viper.GetString("port.bot"), grpc.WithInsecure())
	if err != nil {
		fmt.Printf("链接到 ，，Faild to connect to Bot Services app: %v", err)
	}
	defer conn.Close()

	client := pb.NewBotNoticeClient(conn)
	r, err := client.Notice(context.Background(), &pb.BotNoticeSend{Message: m})
	if err != nil {
		log.Printf("发送消息给 Bot Services 服务端失败: %v", err)
	}
	fmt.Printf("获取返回的消息: %s !\n", r.Reply)
}
