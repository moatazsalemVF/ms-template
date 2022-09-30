package utils

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	// swagger embed files
	swaggerfiles "github.com/swaggo/files"
)

var Router *gin.Engine

func InitServer() {
	fmt.Println("Debug State: ", Debug)
	if Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	Router = gin.New()

	Router.GET("/health", checkHealth)
	Router.SetTrustedProxies(nil)

	attachPromo(Router)
	attachSwagger(Router)

}

func attachSwagger(Router *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/"
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}

func attachPromo(router *gin.Engine) {
	p := ginprometheus.NewPrometheus(Conf.App.Name)
	p.Use(router)
}

// @BasePath /
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /health [get]
func checkHealth(c *gin.Context) {
	m := make(map[string]string)
	m["status"] = "OK"
	c.IndentedJSON(http.StatusOK, m)
}
