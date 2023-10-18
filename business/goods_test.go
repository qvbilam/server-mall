package business

import (
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	payProto "mall/api/qvbilam/pay/v1"
	"mall/global"
	"testing"
)

func initDBClient() {
	user := "root"
	password := "root"
	host := "127.0.0.1"
	port := 3306
	database := "qvbilam_mall"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不带表名
		},
	})
	if err != nil {
		panic(any(err))
	}
	global.DB = db
}

func initPayServer() {
	host := "127.0.0.1"
	port := 9800
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithInsecure())
	if err != nil {
		panic(any(err))
	}
	payClient := payProto.NewPayClient(conn)
	global.PayServerClient = payClient
}

func TestGoodsBusiness_GoodsList(t *testing.T) {
	initDBClient()
	b := GoodsBusiness{}
	b.CategoryID = 2
	goods, count := b.List()
	fmt.Printf("总数: %d\n", count)
	for _, p := range goods {
		fmt.Printf("%+v\n", p)
	}
}

func TestGoodsBusiness_Sell(t *testing.T) {
	initDBClient()
	initPayServer()
	b := GoodsBusiness{}
	b.UserId = 1
	b.ID = 1
	b.Count = 10
	res, err := b.Sell()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("orderSn: %s\n", res.OrderSn)
}
