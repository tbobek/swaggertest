package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	models "github.com/tbobek/swaggertest/models"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping hello world
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description get dataset with specific id
// @Tags data
// @Accept json
// @Produce json
// @Success 200 {} models.Dataset "ok"
// @Router /data/id [get]
func GetId(g *gin.Context) {
	d := models.Dataset{
		Id:   "130670",
		Data: 123,
	}
	g.JSON(http.StatusOK, d)
}

// Id range godoc
// @Summary id range
// @Schemes
// @Description get a range of ids in a given time range
// @Tags data
// @Accept json
// @Produce json
// @Param start query string true "start date"
// @Param end query string true "end date"
// @Success      200  {array}   models.Dataset
// @Failure      400  {object}  models.ErrorMsg
// @Router /data/idrange [get]
func GetIdRange(g *gin.Context) {
	d := models.Dataset{
		Id:   "130670",
		Data: 123,
	}
	g.JSON(http.StatusOK, d)
}

// Time range godoc
// @Summary get time range
// @Schemes
// @Description get a range of ids in a given time range
// @Tags data
// @Accept json
// @Produce json
// @Success 200 {} models.Dataset "ok"
// @Router /data/timerange [get]
func GetTimeRange(g *gin.Context) {
	d := models.Dataset{
		Id:   "130670",
		Data: 123,
	}
	g.JSON(http.StatusOK, d)
}
