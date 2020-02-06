package handler

import (
	pb "consignment-service/proto/consignment"
	vesselPb "consignment-service/proto/vessel"
	"consignment-service/subscriber"
	"context"
	"fmt"
	"log"
)

//
// 定义微服务
//
type Service struct {
	Repo subscriber.Repository
	VesselClient vesselPb.VesselService
}

//
// service 实现 consignment.pb.go 中的 ShippingServiceServer 接口
// 使 service 作为 gRPC 的服务端
//
// 托运新的货物
func (s *Service) CreateConsignment(ctx context.Context, req *pb.Consignment, resp *pb.Response) (error) {
	fmt.Println("进来了")
	// 检查是否有适合的货轮
	vReq := &vesselPb.Specification{
		Capacity:  int32(len(req.Containers)),
		MaxWeight: req.Weight,
	}
	vResp, err := s.VesselClient.FindAvailable(context.Background(), vReq)
	if err != nil {
		return err
	}

	// 货物被承运
	log.Printf("found vessel: %s\n", vResp.Vessel.Name)

	req.VesselId = vResp.Vessel.Id

	consignment, err := s.Repo.Create(req)
	if err != nil {
		return err
	}
	resp.Created = true
	resp.Consignment = consignment
	fmt.Printf("正在托运：%s\n",consignment)
	return nil
}

// 获取目前所有托运的货物
func (s *Service) GetConsignments(ctx context.Context, req *pb.GetRequest, resp *pb.Response) (error) {
	allConsignments := s.Repo.GetAll()
	//resp = &pb.Response{Consignments: allConsignments}
	resp.Consignments=allConsignments
	return  nil
}
