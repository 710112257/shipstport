package subscriber

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	pb "vessel-service/proto/vessel"
)

// 读取 consignment.json 中记录的货物信息
func ParseFile(fileName string) ([]*pb.Vessel, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var vessel []*pb.Vessel
	err = json.Unmarshal(data, &vessel)
	if err != nil {
		return nil, errors.New("consignment.json file content error")
	}
	return vessel, nil
}