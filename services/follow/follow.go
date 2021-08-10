package follow

import (
	"fmt"

	"github.com/sajir-dev/go-crowdfire/domain"
	"github.com/sajir-dev/go-crowdfire/services/follow/contract"
)

func Follow(req *contract.FollowReq) error {

	err := domain.Follow(req)
	fmt.Println("req:", req, "err:", err)
	return err
}

func GetFollowers(req *contract.GetFollowersReq) (*contract.Following, error) {
	res, err := domain.GetFollowers(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
