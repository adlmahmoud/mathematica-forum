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

func (r *TagRepository) AddTagToFil(filID int, tagID int) error {
	query := "INSERT INTO fil_tag (id_fil, id_tag) VALUES (?, ?);"

	_, err := r.db.Exec(query, filID, tagID)
	if err != nil {
		return fmt.Errorf("Erreur lors de l'association du tag : %s", err.Error())
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
