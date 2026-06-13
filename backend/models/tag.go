package models

type Tag struct {
	ID  int    `db:"id"`
	Nom string `db:"nom"`
}

type FilTag struct {
	FilID int `db:"fil_id"`
	TagID int `db:"tag_id"`
}
