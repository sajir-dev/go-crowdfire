package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sajir-dev/go-crowdfire/controllers/dto"
	"github.com/sajir-dev/go-crowdfire/services/follow"
	followcontract "github.com/sajir-dev/go-crowdfire/services/follow/contract"
	"github.com/sajir-dev/go-crowdfire/services/posts"
	postscontract "github.com/sajir-dev/go-crowdfire/services/posts/contract"
	"github.com/sajir-dev/go-crowdfire/services/users"
	usercontract "github.com/sajir-dev/go-crowdfire/services/users/contract"
)

func CreateUser(c *gin.Context) {

	u := &usercontract.CreateUser{}

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"messge":  "invalid req body",
		})
		return
	}

	res, err := users.Create(u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"sucess": false,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    res,
	})
}

func UpdatePost(c *gin.Context) {
	body := dto.UpdatePostReq{
		Id: c.Param("id"),
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid req body",
		})
		return
	}

	user, err := LoginController(body.Email, body.Password)
	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": err,
		})
		return
	}

	req := &postscontract.UpdatePostReq{
		Id:      body.Id,
		Content: body.Content,
		UserId:  user.Id,
	}

	res, err := posts.Update(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"sucess":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}

func FollowPeople(c *gin.Context) {
	body := &dto.FollowReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid req body",
		})
		return
	}

	user, err := LoginController(body.Email, body.Password)
	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid credentials",
		})
		return
	}

	req := &followcontract.FollowReq{
		Userid:    user.Id,
		Following: body.Following,
	}

	err = follow.Follow(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("cannot perform the operation %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func GetPosts(c *gin.Context) {
	body := &dto.ListPostsReq{}

	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "invalid req body",
		})
		return
	}

	user, err := LoginController(body.Email, body.Password)
	if user == nil || err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid credentials",
		})
		return
	}

	res, err := posts.GetPosts(&postscontract.GetPostsReq{
		Id: user.Id,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": fmt.Sprintf("cannot perform the operation %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    res,
	})
}
