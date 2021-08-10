package posts

import (
	"errors"
	"sort"

	"github.com/sajir-dev/go-crowdfire/domain"
	"github.com/sajir-dev/go-crowdfire/services/follow"
	followcontract "github.com/sajir-dev/go-crowdfire/services/follow/contract"
	"github.com/sajir-dev/go-crowdfire/services/posts/contract"
)

// Update ...
func Update(req *contract.UpdatePostReq) (*contract.PostModel, error) {
	res, err := domain.UpdatePost(req)
	return res, err
}

// GetPosts ...
func GetPosts(req *contract.GetPostsReq) ([]contract.PostModel, error) {
	posts := make([]contract.PostModel, 0)

	following, err := follow.GetFollowers(&followcontract.GetFollowersReq{
		Userid: req.Id,
	})

	if err != nil {
		return nil, errors.New("could not perform the operation")
	}

	for _, user := range following.Userids {
		post, _ := domain.GetPosts(&contract.GetPostsReq{Id: user})
		posts = append(posts, post...)
	}

	// Sort posts
	sorted_posts := make(timeSlice, 0, len(posts))
	for _, d := range posts {
		sorted_posts = append(sorted_posts, d)
	}
	sort.Sort(sorted_posts)

	return sorted_posts, nil
}

// for sorting the posts
type timeSlice []contract.PostModel

func (p timeSlice) Len() int {
	return len(p)
}

func (p timeSlice) Less(i, j int) bool {
	return p[i].CreatedAt.After(p[j].CreatedAt)
}

func (p timeSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
