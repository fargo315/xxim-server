// Code generated by goctl. DO NOT EDIT.
// Source: third.proto

package server

import (
	"context"

	"github.com/cherish-chat/xxim-server/app/third/internal/logic/emailservice"
	"github.com/cherish-chat/xxim-server/app/third/internal/svc"
	"github.com/cherish-chat/xxim-server/common/pb"
)

type EmailServiceServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedEmailServiceServer
}

func NewEmailServiceServer(svcCtx *svc.ServiceContext) *EmailServiceServer {
	return &EmailServiceServer{
		svcCtx: svcCtx,
	}
}

// EmailCodeSend 发送邮件
func (s *EmailServiceServer) EmailCodeSend(ctx context.Context, in *pb.EmailCodeSendReq) (*pb.EmailCodeSendResp, error) {
	l := emailservicelogic.NewEmailCodeSendLogic(ctx, s.svcCtx)
	return l.EmailCodeSend(in)
}

// EmailCodeVerify 验证邮件
func (s *EmailServiceServer) EmailCodeVerify(ctx context.Context, in *pb.EmailCodeVerifyReq) (*pb.EmailCodeVerifyResp, error) {
	l := emailservicelogic.NewEmailCodeVerifyLogic(ctx, s.svcCtx)
	return l.EmailCodeVerify(in)
}