package routes

import (
	"Goweb_cli/Controller"
	"Goweb_cli/logger"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.POST("/api/v1/signup", Controller.SignUpHandler)
	r.POST("/api/v1/login", Controller.LoginHandler)
	r.GET("/api/v1/anime_rate", Controller.GetAnimeHandler)
	r.GET("/api/v1/anime_rate/top10", Controller.GetAnimeRankingHandler)
	r.PUT("/api/v1/do_rate", Controller.DoRateHandler)
	r.GET("/api/v1/comment/:id", Controller.GetCommentHandler)
	r.POST("/api/v1/comment/:id", Controller.DoCommentHandler)
	return r
}
