package data

import (
	"time"

	follow "github.com/sajir-dev/go-crowdfire/services/follow/contract"
	post "github.com/sajir-dev/go-crowdfire/services/posts/contract"
	user "github.com/sajir-dev/go-crowdfire/services/users/contract"
)

var UsersMap = map[string]*user.UserModel{
	"u0": {Id: "u0", Name: "John", Email: "john@gmail.com", Password: "password"},
	"u1": {Id: "u1", Name: "Snow", Email: "snow@gmail.com", Password: "password1"},
	"u2": {Id: "u2", Name: "Kyte", Email: "kyte@gmail.com", Password: "password2"},
	"u3": {Id: "u3", Name: "Dale", Email: "dale@gmail.com", Password: "password3"},
}

var PostsMap = map[string]*post.PostModel{
	"p1": {Id: "p1", Content: "sample post one", CreatedAt: time.Now().Add(12 * time.Hour), CreatedBy: "u0"},
	"p2": {Id: "p2", Content: "sample post two", CreatedAt: time.Now().Add(-2 * time.Hour), CreatedBy: "u0"},
	"p3": {Id: "p3", Content: "sample post three", CreatedAt: time.Now().Add(-1 * time.Hour), CreatedBy: "u2"},
	"p4": {Id: "p4", Content: "sample post four", CreatedAt: time.Now().Add(10 * time.Hour), CreatedBy: "u2"},
	"p5": {Id: "p5", Content: "sample post five", CreatedAt: time.Now().Add(13 * time.Hour), CreatedBy: "u2"},
}

var FollowingMap = map[string]*follow.Following{
	"u0": {Userids: []string{"u1", "u2"}},
	"u1": {Userids: []string{"u3", "u2"}},
}
