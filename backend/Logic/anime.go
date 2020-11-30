package Logic

import (
	"Goweb_cli/dao/mysql"
	"Goweb_cli/modules"
)

func GetComment(animeID string) ([]*modules.Comment, error) {

	return mysql.GetComment(animeID)
}
func DoComment(animeID string, rp *modules.CommentParam) error {

	return mysql.DoComment(animeID, rp)
}
