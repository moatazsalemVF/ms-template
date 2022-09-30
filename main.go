package main

import (
	"fmt"
	"net/http"

	"github.com/moatazsalemVF/ms-template/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.Initialize()
	addPingController()
	utils.Router.Run(utils.Conf.Server.Address + ":" + fmt.Sprint(utils.Conf.Server.Port))
}

// @BasePath /ping
// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} OK
// @Router /ping [get]
func addPingController() {
	utils.Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
}
