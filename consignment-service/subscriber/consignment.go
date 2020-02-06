package subscriber

import (
	pb "consignment-service/proto/consignment"
	"sync"
)
//
// 仓库接口
//
type IRepository interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error) // 存放新货物
	GetAll() []*pb.Consignment
}

//
// 我们存放多批货物的仓库，实现了 IRepository 接口
//
type Repository struct {
	consignments []*pb.Consignment
	mux sync.Mutex
}

func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mux.Lock()
	defer repo.mux.Unlock()
	repo.consignments = append(repo.consignments, consignment)
	return consignment, nil
}

func (repo *Repository) GetAll() []*pb.Consignment {
	repo.mux.Lock()
	defer repo.mux.Unlock()
	return repo.consignments
}
