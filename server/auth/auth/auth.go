// type AuthServiceServer interface {
// 	Login(context.Context, *LoginRequest) (*LoginResponse, error)
// }
package auth

import (
	"context"
	authpb "coolcar/auth/api/gen/v1"
)

type Service struct {

}

func (*Service) Login(c context.Context,request *authpb.LoginRequest) (*authpb.LoginResponse, error)