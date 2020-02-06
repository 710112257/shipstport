package main

import (
	pb "consignment-cli/proto/consignment"
	"context"
	"encoding/json"
	"errors"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"io/ioutil"
	"log"
	"sync"
	"time"
)

const (
	DEFAULT_INFO_FILE = "consignment.json"
)

// 读取 consignment.json 中记录的货物信息
func parseFile(fileName string) (*pb.Consignment, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var consignment *pb.Consignment
	err = json.Unmarshal(data, &consignment)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return consignment, nil
}

var (
	s sync.WaitGroup
	num = 3
)

func main() {
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"192.168.16.100:2888",
		}
	})
	service := micro.NewService(
		micro.Registry(reg),
	)
	service.Init()

	client := pb.NewShippingService("go.micro.srv.consignment", service.Client())

	// 解析货物信息
	consignment, err := parseFile(DEFAULT_INFO_FILE)
	if err != nil {
		log.Fatalf("parse info file error: %v", err)
	}


	s.Add(num)
	for i := 0; i < num; i++ {
		go func(j int) {
			r, err := client.CreateConsignment(context.TODO(), consignment)
			if err != nil {
				log.Fatalf("Could not create: %v", err)
			}
			log.Printf("Created: %t %d", r.Created, j)
			getAll, err := client.GetConsignments(context.TODO(), &pb.GetRequest{})
			if err != nil {
				log.Fatalf("Could not list consignments: %v", err)
			}
			for _, v := range getAll.Consignments {
				log.Println(v)
			}
			s.Done()
		}(i)
		time.Sleep(time.Second / 2)
	}

	s.Wait()
}
