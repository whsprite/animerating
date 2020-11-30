package mysql

import (
	"Goweb_cli/modules"
	"database/sql"
)

func GetAnime() (animes []modules.Anime, err error) {
	animes = make([]modules.Anime, 0)
	sqlStr := "select * from anime"
	err = db.Select(&animes, sqlStr)
	return animes, err
}
func GetAnimeRanking() (animes []modules.Anime, err error) {
	animes = make([]modules.Anime, 0)
	sqlStr := "select * from anime order by rating desc limit 10"
	err = db.Select(&animes, sqlStr)
	return animes, err
}
func RateAnime(rp *modules.RateParam) (err error) {
	oldRp := new(modules.RateParam)
	getSqlStr := "select rating,rate_num from anime where id = ?"
	err = db.Get(oldRp, getSqlStr, rp.Id)
	if err != nil {
		return err
	}
	//获取旧的数据 计算出新的数据
	newNum := float32(oldRp.RateNum)

	newRate := ((oldRp.Rate * newNum) + rp.Rate)
	newRate = newRate / (newNum + 1)
	updateSqlStr := "update anime set rating = ? where id = ?"
	_, err = db.Exec(updateSqlStr, newRate, rp.Id)
	if err != nil {
		return err
	}
	updateSqlStr = "update anime set rate_num = ? where id = ?"
	_, err = db.Exec(updateSqlStr, oldRp.RateNum+1, rp.Id)
	if err != nil {
		return err
	}
	return
}
func GetComment(animeID string) (comments []*modules.Comment, err error) {
	sqlStr := "select * from comment where anime_id = ?"
	if err = db.Select(&comments, sqlStr, animeID); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrorInvalidParam
		}
	}
	return comments, err
}
func DoComment(animeID string, rp *modules.CommentParam) error {
	sqlStr := "insert into comment (comment_string,rate,anime_id) values (?,?,?)"
	_, err := db.Exec(sqlStr, rp.Comment, rp.Rate, animeID)
	return err
}
