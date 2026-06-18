package models

type Tag struct {
	ID  int    `db:"id_tag"`
	Nom string `db:"nom_tag"`
}

type FilTag struct {
	FilID int `db:"id_fil"`
	TagID int `db:"id_tag"`
}
