package initialize

import (
	"fmt"
	retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	payProto "mall/api/qvbilam/pay/v1"
	userProto "mall/api/qvbilam/user/v1"
	"mall/global"
	"time"
)

type dialConfig struct {
	host string
	port int64
	name string
}

type serverClientConfig struct {
	userDialConfig *dialConfig
	payDialConfig  *dialConfig
}

func InitServer() {
	s := serverClientConfig{
		userDialConfig: &dialConfig{
			host: global.ServerConfig.UserServerConfig.Host,
			port: global.ServerConfig.UserServerConfig.Port,
			name: global.ServerConfig.UserServerConfig.Name,
		},
		payDialConfig: &dialConfig{
			host: global.ServerConfig.PayServerConfig.Host,
			port: global.ServerConfig.PayServerConfig.Port,
			name: global.ServerConfig.PayServerConfig.Name,
		},
	}

	s.initUserServer()
	s.initPayServer()
}

func clientOption() []retry.CallOption {
	opts := []retry.CallOption{
		retry.WithBackoff(retry.BackoffLinear(100 * time.Millisecond)), // 重试间隔
		retry.WithMax(3), // 最大重试次数
		retry.WithPerRetryTimeout(1 * time.Second),                                 // 请求超时时间
		retry.WithCodes(codes.NotFound, codes.DeadlineExceeded, codes.Unavailable), // 指定返回码重试
	}
	return opts
}

func (s *serverClientConfig) initUserServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.userDialConfig.host, s.userDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)))
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", s.userDialConfig.name, err)
	}

	userClient := userProto.NewUserClient(conn)

	global.UserServerClient = userClient
}

func (s *serverClientConfig) initPayServer() {
	opts := clientOption()

	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%d", s.payDialConfig.host, s.payDialConfig.port),
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(opts...)),
	)
	if err != nil {
		zap.S().Fatalf("%s dial error: %s", s.payDialConfig.name, err)
	}

	payClient := payProto.NewPayClient(conn)
	global.PayServerClient = payClient
}
