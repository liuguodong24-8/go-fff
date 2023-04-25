package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
	"time"

	pb "helloworld/api/pledge/v1"
)

type PledgeService struct {
	pb.UnimplementedPledgeServer
	ple *biz.PledgeCase
	log *log.Helper
}

func NewPledgeService(ple *biz.PledgeCase, logger log.Logger) *PledgeService {
	return &PledgeService{
		ple: ple,
		log: log.NewHelper(logger),
	}
}

func (s *PledgeService) CreatePledge(ctx context.Context, req *pb.CreatePledgeRequest) (*pb.CreatePledgeReply, error) {
	_, err := s.ple.CreatePledge(ctx, &biz.Pledge{
		TokenID:   1,
		Address:   req.Address,
		Price:     req.Price,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
		Status:    1,
	})
	if err != nil {
		return &pb.CreatePledgeReply{
			Code: 410,
			Msg:  err.Error(),
		}, err
	}
	return &pb.CreatePledgeReply{
		Code: 0,
		Msg:  "",
	}, nil
}
func (s *PledgeService) UpdatePledge(ctx context.Context, req *pb.UpdatePledgeRequest) (*pb.UpdatePledgeReply, error) {
	return &pb.UpdatePledgeReply{}, nil
}
func (s *PledgeService) DeletePledge(ctx context.Context, req *pb.DeletePledgeRequest) (*pb.DeletePledgeReply, error) {
	return &pb.DeletePledgeReply{}, nil
}
func (s *PledgeService) GetPledge(ctx context.Context, req *pb.GetPledgeRequest) (*pb.GetPledgeReply, error) {
	ple, err := s.ple.GetPledge(ctx, req.Address)
	if err != nil {
		return &pb.GetPledgeReply{
			Code: 0,
			Msg:  "",
			Data: nil,
		}, nil
	}
	return &pb.GetPledgeReply{
		Code: 0,
		Msg:  "",
		Data: &pb.PledgeInfo{
			Id:      ple.ID,
			TokenId: ple.TokenID,
			Price:   ple.Price,
			Status:  ple.Status,
			Times:   ple.Times,
			Address: ple.Address,
		},
	}, nil
	return &pb.GetPledgeReply{}, nil
}
func (s *PledgeService) ListPledge(ctx context.Context, req *pb.ListPledgeRequest) (*pb.ListPledgeReply, error) {
	return &pb.ListPledgeReply{}, nil
}
