package handler

import (
	"context"
	pb "vessel-service/proto/vessel"
	"vessel-service/subscriber"
)

// 定义货船服务
type Service struct {
	Repo subscriber.VesselRepository
}

// 实现服务端
func (s *Service) FindAvailable(ctx context.Context, spec *pb.Specification, resp *pb.Response) error {
	// 调用内部方法查找
	v, err := s.Repo.FindAvailable(spec)
	if err != nil {
		return err
	}
	resp.Vessel = v
	return nil
}
