package app

import (
	"github.com/gin-gonic/gin"
	"github.com/sajir-dev/go-crowdfire/domain"
)

var (
	router = gin.Default()
)

func StartApp() {
	// initialising the DB
	if err := domain.InitDB(); err != nil {
		panic(err)
	}

	// Mounting the routes
	mapUrls()
	router.Run(":3030")
}
