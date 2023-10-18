package global

import (
	"gorm.io/gorm"
	payProto "mall/api/qvbilam/pay/v1"
	userProto "mall/api/qvbilam/user/v1"
	"mall/config"
)

var (
	DB               *gorm.DB
	ServerConfig     *config.ServerConfig
	UserServerClient userProto.UserClient
	PayServerClient  payProto.PayClient
)
