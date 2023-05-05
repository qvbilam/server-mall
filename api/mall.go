package api

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "mall/api/qvbilam/mall/v1"
	"mall/business"
	"mall/model"
)

type MallServer struct {
	proto.UnimplementedMallServer
}

// GetCategory 分类列表
func (s *MallServer) GetCategory(ctx context.Context, request *proto.CategoryListRequest) (*proto.CategoryListResponse, error) {
	b := business.CategoryBusiness{PID: &request.Pid}
	var list []*proto.CategoryResponse
	cs := b.Get()
	for _, c := range cs {
		list = append(list, &proto.CategoryResponse{
			Id:   c.ID,
			Pid:  c.PID,
			Name: c.Name,
			Tag:  "",
		})
	}

	return &proto.CategoryListResponse{List: list}, nil
}

// GetGoodsList 商品列表
func (s *MallServer) GetGoodsList(ctx context.Context, request *proto.GoodsListRequest) (*proto.GoodsListResponse, error) {
	b := business.GoodsBusiness{}
	if request.Ids != nil {
		b.IDs = request.Ids
	}
	if request.CategoryId != 0 {
		b.CategoryID = request.CategoryId
	}
	if request.OnSale != false {
		b.OnSale = &request.OnSale
	}
	if request.Page != nil && request.Page.Page != 0 {
		b.Page = request.Page.Page
		b.PerPage = request.Page.PerPage
	}
	if request.NeedProduct == true {
		b.NeedProduct = request.NeedProduct
	}

	goods, total := b.List()
	var list []*proto.GoodsDetailResponse
	for _, g := range goods {
		list = append(list, s.goodsToResponse(g))
	}

	return &proto.GoodsListResponse{Total: total, List: list}, nil
}

// GetGoodsDetail 商品详情
func (s *MallServer) GetGoodsDetail(ctx context.Context, request *proto.GoodsDetailRequest) (*proto.GoodsDetailResponse, error) {
	b := business.GoodsBusiness{ID: request.Id}
	goods := b.Detail()
	return s.goodsToResponse(goods), nil
}

// Sell 售卖 todo
func (s *MallServer) Sell(ctx context.Context, request *proto.SellRequest) (*emptypb.Empty, error) {
	return nil, nil
}

// Rollback 回滚 todo
func (s *MallServer) Rollback(ctx context.Context, request *proto.SellRequest) (*emptypb.Empty, error) {
	return nil, nil
}

func (s *MallServer) goodsToResponse(goods *model.Goods) *proto.GoodsDetailResponse {
	if goods == nil {
		return nil
	}

	res := proto.GoodsDetailResponse{
		Id:            goods.ID,
		CategoryId:    goods.CategoryID,
		Name:          goods.Name,
		Introduce:     goods.Introduce,
		Icon:          goods.Icon,
		PayType:       goods.PayType,
		Price:         goods.Price,
		OriginalPrice: goods.OriginalPrice,
		Stocks:        goods.Stocks,
		SoldCount:     goods.SoldCount,
		IsHot:         goods.IsHot,
		IsUnlimited:   goods.IsUnlimited,
		OnSale:        goods.OnSale,
	}
	var goodsProducts []*proto.GoodsProductResponse
	for _, ps := range goods.Products {
		gpRes := &proto.GoodsProductResponse{
			Id:        ps.ID,
			GoodsId:   ps.GoodsID,
			ProductId: ps.ProductID,
			Count:     ps.Count,
		}
		if ps.Product != nil {
			gpRes.Product = &proto.ProductResponse{
				Id:        ps.Product.ID,
				Name:      ps.Product.Name,
				Icon:      ps.Product.Icon,
				Price:     ps.Product.Price,
				Introduce: ps.Product.Introduce,
				Type:      ps.Product.Type,
				Tag:       ps.Product.Tag,
			}
		}
		goodsProducts = append(goodsProducts, gpRes)
	}

	res.Products = goodsProducts
	return &res
}
