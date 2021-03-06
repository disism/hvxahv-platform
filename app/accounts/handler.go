/**
	实现 Accounts GRPC 服务的方法
 */
package main

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/app/accounts/services"
	"log"
)



// NewAccount 创建账户 将接收到的用户数据存储到数据库
func (s *server) NewAccount(ctx context.Context, in *pb.AccountData) (*pb.NewAccountReply, error) {
	r := services.NewAccount(in)
	log.Println(in)
	i := int32(r)

	return &pb.NewAccountReply{Reply: i}, nil
}

// GetAccount 获取账户资料
func (s *server) GetAccount(ctx context.Context, in *pb.AccountName) (*pb.AccountData, error) {
	r := services.GetAccountData(in.Username)

	ad := &pb.AccountData{
		Name: r.Name,
		Username: r.Username,
		Password: r.Password,
	}
	return ad, nil
}

// GetActor 获取 Activitypub 的 Actor 信息
func (s *server) GetActor(ctx context.Context, in *pb.AccountName) (*pb.AccountData, error) {
	r := services.GetActorData(in.Username)

	ad := &pb.AccountData{
		Username: r.Username,
		Name: r.Name,
		PrivateKey: r.PrivateKey,
		PublicKey: r.PublicKey,
	}
	return ad, nil
}

// VerifyAccount 用于登陆时用于验证账户的服务
func (s *server) VerifyAccount(ctx context.Context, in *pb.AccountName) (*pb.AccountData, error) {
	r := services.VerifyAccounts(in.Username)

	ad := &pb.AccountData{
		Name: r.Name,
		Username: r.Username,
		Password: r.Password,
	}
	return ad, nil
}