package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"

	_ "github.com/go-sql-driver/mysql"
)

type TagRepository struct {
	db *sql.DB
}

func InitTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db}
}

func (r *TagRepository) CreateTag(tag models.Tag) (int, error) {
	query := "INSERT INTO `tag`(`nom_tag`) VALUES (?);"

	sqlResult, sqlErr := r.db.Exec(query, tag.Nom)
	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création tag - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création tag - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *TagRepository) ReadAll() ([]models.Tag, error) {
	var listTags []models.Tag
	sqlResult, sqlErr := r.db.Query("SELECT `id_tag`, `nom_tag` FROM `tag`;")
	if sqlErr != nil {
		return listTags, fmt.Errorf("Erreur récupération tags - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var tag models.Tag
		errScan := sqlResult.Scan(&tag.ID, &tag.Nom)
		if errScan != nil {
			return nil, errScan
		}
		listTags = append(listTags, tag)
	}

	return listTags, nil
}

func (r *TagRepository) ReadById(id int) (models.Tag, error) {
	var tag models.Tag
	query := "SELECT `id_tag`, `nom_tag` FROM `tag` WHERE `id_tag` = ?;"

	sqlErr := r.db.QueryRow(query, id).Scan(&tag.ID, &tag.Nom)
	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Tag{}, nil
		}
		return models.Tag{}, fmt.Errorf("Erreur récupération tag - Erreur : \n\t %s", sqlErr.Error())
	}

	return tag, nil
}

func (r *TagRepository) DeleteTagById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM `tag` WHERE `id_tag`=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression tag - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur suppression tag - Aucun tag supprimé")
	}

	return nil
}

func (r *TagRepository) GetFilsByTag(tagID int) ([]int, error) {
	var filIDs []int
	sqlResult, sqlErr := r.db.Query("SELECT `id_fil` FROM `fil_tag` WHERE `id_tag` = ?;", tagID)
	if sqlErr != nil {
		return filIDs, fmt.Errorf("Erreur récupération fils - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var filID int
		errScan := sqlResult.Scan(&filID)
		if errScan != nil {
			return nil, errScan
		}
		filIDs = append(filIDs, filID)
	}

	return filIDs, nil
}

func (r *TagRepository) AddTagToFil(filID int, tagID int) error {
	query := "INSERT INTO `fil_tag`(`id_fil`, `id_tag`) VALUES (?,?);"

	_, sqlErr := r.db.Exec(query, filID, tagID)
	if sqlErr != nil {
		return fmt.Errorf("Erreur ajout tag au fil - Erreur : \n\t %s", sqlErr.Error())
	}

	return nil
}

func (r *TagRepository) RemoveTagFromFil(filID int, tagID int) error {
	_, sqlErr := r.db.Exec("DELETE FROM `fil_tag` WHERE `id_fil`=? AND `id_tag`=?;", filID, tagID)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression tag du fil - Erreur : \n\t %s", sqlErr.Error())
	}

	return nil
}
