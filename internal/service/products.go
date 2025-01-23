package service

import (
	"context"

	pb "opspillar/api/opspillar/v1"

	"github.com/go-kratos/kratos/v2/log"

	biz "opspillar/internal/biz"
)

type ProductsService struct {
	pb.UnimplementedProductsServer
	usecase *biz.ProductsUsecase
	log     *log.Helper
}

func NewProductsService(uc *biz.ProductsUsecase, logger log.Logger) *ProductsService {
	return &ProductsService{
		usecase: uc,
		log:     log.NewHelper(logger),
	}
}

func toBizProduct(p *pb.Product) (*biz.Product, error) {
	if p == nil {
		return nil, nil
	}
	return &biz.Product{
		Id:          p.Id,
		Name:        p.Name,
		Code:        p.Code,
		Description: p.Description,
	}, nil
}

func toBizProducts(ps []*pb.Product) ([]*biz.Product, error) {
	if ps == nil {
		return nil, nil
	}
	bizProducts := make([]*biz.Product, len(ps))
	for i, p := range ps {
		bizProduct, err := toBizProduct(p)
		if err != nil {
			return nil, err
		}
		bizProducts[i] = bizProduct
	}
	return bizProducts, nil
}

func (s *ProductsService) CreateProducts(ctx context.Context, req *pb.CreateProductsRequest) (*pb.CreateProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}

	ps, err := toBizProducts(req.Products)
	if err == nil {
		err = s.usecase.CreateProducts(ctx, ps)
	}
	reply := &pb.CreateProductsReply{
		Action:  "createProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *ProductsService) UpdateProducts(ctx context.Context, req *pb.UpdateProductsRequest) (*pb.UpdateProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	ps, err := toBizProducts(req.Products)
	if err == nil {
		err = s.usecase.UpdateProducts(ctx, ps)
	}

	reply := &pb.UpdateProductsReply{
		Action:  "updateProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *ProductsService) DeleteProducts(ctx context.Context, req *pb.DeleteProductsRequest) (*pb.DeleteProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	err := s.usecase.DeleteProducts(ctx, req.Ids)
	reply := &pb.DeleteProductsReply{
		Action:  "DeleteProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	return reply, nil
}

func (s *ProductsService) GetProducts(ctx context.Context, req *pb.GetProductsRequest) (*pb.GetProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	p, err := s.usecase.GetProducts(ctx, req.Id)
	reply := &pb.GetProductsReply{
		Action:  "GetProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Product = toPbProduct(p)
	return reply, nil
}
func (s *ProductsService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsReply, error) {
	filter := biz.DefaultProductsFilter()
	if req != nil {
		if len(req.Names) > 0 {
			filter.Names = req.Names
		}
		if len(req.Codes) > 0 {
			filter.Codes = req.Codes
		}
		if len(req.Ids) > 0 {
			filter.Ids = req.Ids
		}
		if req.PageSize > 0 {
			filter.PageSize = req.PageSize
		}
		if req.Page > 0 {
			filter.Page = req.Page
		}
	}

	ps, err := s.usecase.ListProducts(ctx, filter)
	reply := &pb.ListProductsReply{
		Action:  "ListProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, nil
	}
	reply.Products = toPbProducts(ps)
	return reply, nil

}

func toPbProduct(p *biz.Product) *pb.Product {
	if p == nil {
		return nil
	}
	return &pb.Product{
		Code:        p.Code,
		Description: p.Description,
		Id:          p.Id,
		Name:        p.Name,
	}
}

func toPbProducts(ps []*biz.Product) []*pb.Product {
	if ps == nil {
		return nil
	}
	reply := []*pb.Product{}
	for _, p := range ps {
		if p != nil {
			reply = append(reply, toPbProduct(p))
		}
	}
	return reply
}
