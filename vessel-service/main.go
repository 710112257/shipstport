package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"vessel-service/handler"
	pb "vessel-service/proto/vessel"
	"vessel-service/subscriber"
)

const (
	DEFAULT_INFO_FILE = "./data/vessel.json"
)
func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{"192.168.16.100:2888"}
	})
	server:= micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Registry(reg),
	)
	server.Init()

	service:=new(handler.Service)

	//du shuju
	service.Repo.Vessels,_=subscriber.ParseFile(DEFAULT_INFO_FILE)

	if err:=pb.RegisterVesselServiceHandler(server.Server(), service);err!=nil{
		panic(err)
	}

	if err := server.Run(); err != nil {
		panic(err)
	}
}
