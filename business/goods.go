package business

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	proto "mall/api/qvbilam/mall/v1"
	payProto "mall/api/qvbilam/pay/v1"
	"mall/enum"
	"mall/global"
	"mall/model"
)

type GoodsBusiness struct {
	UserId      int64
	ID          int64
	IDs         []int64
	CategoryID  int64
	OnSale      *bool
	NeedProduct bool
	Page        int64
	PerPage     int64
	Count       int64
}

type ShellResponse struct {
	PayType proto.PayType
	OrderSn string
}

// Detail 商品详情
func (b *GoodsBusiness) Detail() *model.Goods {
	entity := model.Goods{}
	global.DB.Where(&model.Goods{IDModel: model.IDModel{ID: b.ID}}).Preload("Products.Product").First(&entity)
	return &entity
}

// List 商品列表
func (b *GoodsBusiness) List() ([]*model.Goods, int64) {
	query := global.DB.Model(&model.Goods{})

	var goods []*model.Goods
	condition := model.Goods{}
	cb := CategoryBusiness{}
	var categoryIds []interface{}

	// 分类查询
	if b.CategoryID != 0 {
		cb.ID = b.CategoryID
		categoryIds, _ = cb.GetMultistageCategoryIds()

		query = query.Where("category_id in ?", categoryIds)
	}

	// ids 查询
	if b.IDs != nil {
		query = query.Where("id in ?", b.IDs)
	}

	// 上架状态查询
	if b.OnSale != nil {
		condition.OnSale = *b.OnSale
	}

	// 分页查询
	if b.Page != 0 {
		query = query.Scopes(model.Paginate(int(b.Page), int(b.PerPage)))
	}

	var count int64

	// 数量结果
	query = query.Where(&condition)
	query.Count(&count)

	// 详情查询
	if b.NeedProduct == true {
		query = query.Preload("Products.Product")
	}

	query.Find(&goods)

	return goods, count
}

// Sell 售卖
func (b *GoodsBusiness) Sell() (*ShellResponse, error) {
	count := b.Count
	if count <= 0 {
		return nil, status.Errorf(codes.InvalidArgument, "数量错误")
	}

	// 获取商品信息
	var goods *model.Goods
	if goods = b.Detail(); goods == nil {
		return nil, status.Errorf(codes.NotFound, "商品不存在")
	}

	//gJson, _ := json.Marshal(goods)
	//fmt.Println(string(gJson))

	var res string
	// 获取支付类型
	payType := proto.PayType_None
	if goods.PayType == enum.PayTypeMoney { // 金钱支付方式
		payType = proto.PayType_Money
		res = b.sellByMoney(goods)
	} else if goods.PayType == enum.PayTypeCoin { // 金币支付方式
		payType = proto.PayType_Coin
		return nil, status.Errorf(codes.InvalidArgument, "暂不支持当前支付方式")
	} else { // todo return error payType
		payType = proto.PayType_None
		return nil, status.Errorf(codes.InvalidArgument, "暂不支持当前支付方式")
	}

	return &ShellResponse{
		PayType: payType,
		OrderSn: res,
	}, nil
}

func (b *GoodsBusiness) sellByMoney(goods *model.Goods) string {
	var goodsRequest []*payProto.CreateGoodsRequest
	for _, g := range goods.Products {
		goodsRequest = append(goodsRequest, &payProto.CreateGoodsRequest{
			Type:  g.Product.Type,
			Id:    g.GoodsID,
			Name:  g.Product.Name,
			Icon:  g.Product.Icon,
			Price: float32(g.Product.Price),
			Count: g.Count,
		})
	}

	res, _ := global.PayServerClient.CreateOrder(context.Background(), &payProto.CreateOrderRequest{
		UserId:  b.UserId,
		Count:   b.Count,
		Amount:  float32(goods.Price / 100),
		Subject: fmt.Sprintf("购买【%s】", goods.Name),
		Goods:   goodsRequest,
	})
	return res.OrderSn
}
