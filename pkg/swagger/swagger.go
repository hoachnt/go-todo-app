package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/hoachnt/go-todo-app/docs"
)

type Swagger struct{}

func NewSwagger() *Swagger {
	return &Swagger{}
}

func (s *Swagger) Setup(router *gin.Engine) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
