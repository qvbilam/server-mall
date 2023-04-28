package global

import (
	"gorm.io/gorm"
	userProto "mall/api/qvbilam/user/v1"
	"mall/config"
)

var (
	DB               *gorm.DB
	ServerConfig     *config.ServerConfig
	UserServerClient userProto.UserClient
)
