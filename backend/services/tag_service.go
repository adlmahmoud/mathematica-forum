package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"
)

type TagRepository struct {
	db *sql.DB
}

func InitTagRepository(db *sql.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) CreateTag(tag models.Tag) (int, error) {
	query := "INSERT INTO tag (nom_tag) VALUES (?);"

	sqlResult, err := r.db.Exec(query, tag.Nom)
	if err != nil {
		return 0, err
	}

	id, err := sqlResult.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *TagRepository) ReadAll() ([]models.Tag, error) {
	var tags []models.Tag
	query := "SELECT id_tag, nom_tag FROM tag;"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag models.Tag
		if err := rows.Scan(&tag.ID, &tag.Nom); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (r *TagRepository) ReadById(id int) (models.Tag, error) {
	var tag models.Tag
	query := "SELECT id_tag, nom_tag FROM tag WHERE id_tag = ?;"

	err := r.db.QueryRow(query, id).Scan(&tag.ID, &tag.Nom)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Tag{}, nil
		}
		return models.Tag{}, err
	}

	return tag, nil
}

func (r *TagRepository) DeleteTagById(id int) error {
	query := "DELETE FROM tag WHERE id_tag = ?;"

	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Tag introuvable")
	}

	return nil
}

func (r *TagRepository) AddTagToFil(filID int, tagID int) error {
	query := "INSERT INTO fil_tag (id_fil, id_tag) VALUES (?, ?);"

	_, err := r.db.Exec(query, filID, tagID)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'association du tag : %s", err.Error())
	}

	return nil
}

func (r *TagRepository) RemoveTagFromFil(filID int, tagID int) error {
	query := "DELETE FROM fil_tag WHERE id_fil = ? AND id_tag = ?;"

	_, err := r.db.Exec(query, filID, tagID)
	if err != nil {
		return fmt.Errorf("Erreur lors de la suppression de l'association : %s", err.Error())
	}

	return nil
}

func (r *TagRepository) GetFilsByTag(tagID int) ([]models.Fil, error) {
	var listFils []models.Fil
	query := `
		SELECT f.id_fil, f.titre, f.statut, f.date_creation, f.id_utilisateur
		FROM fil_discussion f
		INNER JOIN fil_tag ft ON f.id_fil = ft.id_fil
		WHERE ft.id_tag = ? AND f.statut != 'archivé'
		ORDER BY f.date_creation DESC;`

	rows, err := r.db.Query(query, tagID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var fil models.Fil
		if err := rows.Scan(&fil.ID, &fil.Titre, &fil.Statut, &fil.DateCreation, &fil.UtilisateurID); err != nil {
			return nil, err
		}
		listFils = append(listFils, fil)
	}

	return listFils, nil
}
