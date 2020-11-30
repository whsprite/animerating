package modules

type Anime struct {
	Id       int     `db:"id"`
	Name     string  `db:"name"`
	Url      string  `db:"url"`
	Category string  `db:"category"`
	Year     int     `db:"year"`
	Rating   float32 `db:"rating"`
	RateNum  int     `db:"rate_num"`
}

type Comment struct {
	ID            int    `db:"id"`
	AnimeID       int    `db:"anime_id"`
	CommentString string `db:"comment_string"`
	Rate          int    `db:"rate"`
}
