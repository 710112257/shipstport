package main

import (
	"consignment-service/handler"
	pb "consignment-service/proto/consignment"
	vesselPb "consignment-service/proto/vessel"
	"consignment-service/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options){
		op.Addrs = []string{
			"192.168.16.100:2888",
		}
	})
	server := micro.NewService(
		micro.Name("go.micro.srv.consignment"),
		micro.Registry(reg),
	)
	server.Init()

	client := vesselPb.NewVesselService("go.micro.srv.vessel", server.Client())

	service := handler.Service{
		subscriber.Repository{},
		client,
	}

	pb.RegisterShippingServiceHandler(server.Server(), &service)

	if err := server.Run(); err != nil {
		panic(err)
	}
}
