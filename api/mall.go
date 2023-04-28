package api

import (
	"context"
	proto "mall/api/qvbilam/mall/v1"
)

type MallServer struct {
	proto.UnimplementedMallServer
}

func GetCategory(context.Context, *proto.CategoryListRequest) (*proto.CategoryListResponse, error) {
	return nil, nil
}
