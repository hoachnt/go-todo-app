package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/hoachnt/go-todo-app/docs"
)

type Swagger struct {
	router *gin.Engine
}

func NewSwagger(router *gin.Engine) *Swagger {
	return &Swagger{router: router}
}

func (swagger *Swagger) Setup() *gin.Engine {

	swagger.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return swagger.router
}
