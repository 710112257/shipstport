package subscriber

import (
	"errors"
	"fmt"
	pb "vessel-service/proto/vessel"
)

type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

type VesselRepository struct {
	Vessels []*pb.Vessel
}

// 接口实现
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	// 选择最近一条容量、载重都符合的货轮
	for _, v := range repo.Vessels {
		if v.Capacity >= spec.Capacity && v.MaxWeight >= spec.MaxWeight {
			fmt.Println("有可运输的货船")
			return v, nil
		}
	}
	return nil, errors.New("No vessel can't be use")
}