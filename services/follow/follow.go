package follow

import (
	"errors"

	"github.com/sajir-dev/go-crowdfire/data"
	"github.com/sajir-dev/go-crowdfire/services/follow/contract"
)

func Follow(req *contract.FollowReq) error {
	for user, followers := range data.FollowingMap {
		if user == req.Userid {
			// TODO Validation:
			// 1.  follower is already following,
			// 2. user-id is valid
			followers.Userids = append(followers.Userids, req.Following)
			return nil
		}
	}

	return errors.New("could not perform the operation")
}

func GetFollowers(req *contract.GetFollowersReq) (*contract.Following, error) {
	for user, followers := range data.FollowingMap {
		if user == req.Userid {
			return followers, nil
		}
	}
	return nil, errors.New("could not perform the operation")
}
