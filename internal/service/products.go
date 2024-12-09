package service

import (
	"context"

	pb "appix/api/appix/v1"

	"github.com/go-kratos/kratos/v2/log"

	//  TODO: modify project name
	biz "appix/internal/biz"
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

func (s *ProductsService) CreateProducts(ctx context.Context, req *pb.CreateProductsRequest) (*pb.CreateProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}

	ps := make([]biz.Product, len(req.Products))
	for i, p := range req.Products {
		ps[i] = biz.Product{
			Code:        p.Code,
			Description: p.Description,
			Id:          p.Id,
			Name:        p.Name,
		}
	}
	err := s.usecase.CreateProducts(ctx, ps)
	reply := &pb.CreateProductsReply{
		Action:  "createProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, err
	}
	return reply, nil
}

func (s *ProductsService) UpdateProducts(ctx context.Context, req *pb.UpdateProductsRequest) (*pb.UpdateProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	ps := make([]biz.Product, len(req.Products))
	for i, p := range req.Products {
		ps[i] = biz.Product{
			Code:        p.Code,
			Description: p.Description,
			Id:          p.Id,
			Name:        p.Name,
		}
	}
	err := s.usecase.UpdateProducts(ctx, ps)
	reply := &pb.UpdateProductsReply{
		Action:  "updateProducts",
		Code:    0,
		Message: "success",
	}
	if err != nil {
		reply.Code = 1
		reply.Message = err.Error()
		return reply, err
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
		return reply, err
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
		return reply, err
	}
	reply.Product = &pb.Product{
		Code:        p.Code,
		Description: p.Description,
		Id:          p.Id,
		Name:        p.Name,
	}
	return reply, nil
}
func (s *ProductsService) ListProducts(ctx context.Context, req *pb.ListProductsRequest) (*pb.ListProductsReply, error) {
	if req == nil {
		return nil, ErrRequestNil
	}
	var filter *biz.ListProductsFilter
	if req.Filter != nil {
		filter = &biz.ListProductsFilter{
			Page:     req.Filter.Page,
			PageSize: req.Filter.PageSize,
		}

		filter.Filters = make([]biz.ProductFilter, len(req.Filter.Filters))
		for i, f := range req.Filter.Filters {
			filter.Filters[i] = biz.ProductFilter{
				Code: f.Code,
				Name: f.Name,
			}
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
		return reply, err
	}
	reply.Products = make([]*pb.Product, len(ps))
	for i, p := range ps {
		reply.Products[i] = &pb.Product{
			Code:        p.Code,
			Description: p.Description,
			Id:          p.Id,
			Name:        p.Name,
		}
	}
	return reply, nil

}
