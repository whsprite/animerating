package Controller

import (
	"Goweb_cli/Logic"
	"Goweb_cli/dao/mysql"
	"Goweb_cli/modules"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func GetAnimeHandler(c *gin.Context) {
	animes := make([]modules.Anime, 0)
	var err error
	animes, err = mysql.GetAnime()
	if err != nil {
		zap.L().Error("mysql.GetAnime() failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
		return
	}
	modules.ResponseSuccess(c, animes)
	return
}

func GetAnimeRankingHandler(c *gin.Context) {
	animes := make([]modules.Anime, 0)
	var err error
	animes, err = mysql.GetAnimeRanking()
	if err != nil {
		zap.L().Error("mysql.GetAnime() failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
		return
	}
	modules.ResponseSuccess(c, animes)
	return
}

func DoRateHandler(c *gin.Context) {
	rp := new(modules.RateParam)
	err := c.BindJSON(rp)
	if err != nil {
		zap.L().Error("c.BindJSON failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
	}
	err = mysql.RateAnime(rp)
	if err != nil {
		zap.L().Error("mysql.RateAnime failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
	}
	modules.ResponseSuccess(c, nil)
}

func GetCommentHandler(c *gin.Context) {
	animeID := c.Param("id")
	comments, err := Logic.GetComment(animeID)
	if err != nil {
		zap.L().Error("Logic.GetComment failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
		return
	}
	modules.ResponseSuccess(c, comments)
	return
}

func DoCommentHandler(c *gin.Context) {
	animeID := c.Param("id")
	cp := new(modules.CommentParam)
	err := c.BindJSON(cp)
	if err != nil {
		zap.L().Error("c.BindJSON failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
	}
	err = Logic.DoComment(animeID, cp)
	if err != nil {
		zap.L().Error("c.BindJSON failed", zap.Error(err))
		modules.ResponseError(c, modules.CodeServerBusy)
	}
	modules.ResponseSuccess(c, nil)
}
