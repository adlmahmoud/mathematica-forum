package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type FilRepository struct {
	db *sql.DB
}

func InitFilRepository(db *sql.DB) *FilRepository {
	return &FilRepository{db}
}

func (r *FilRepository) CreateFil(fil models.Fil) (int, error) {
	query := "INSERT INTO `fil_discussion`(`titre`, `statut`, `date_creation`, `id_utilisateur`) VALUES (?,?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		fil.Titre,
		fil.Statut,
		time.Now().Format("2006-01-02 15:04:05"),
		fil.UtilisateurID,
	)

	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création fil - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création fil - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *FilRepository) ReadAll() ([]models.Fil, error) {
	var listFils []models.Fil
	sqlResult, sqlErr := r.db.Query("SELECT `id_fil`, `titre`, `statut`, `date_creation`, `id_utilisateur` FROM `fil_discussion`;")
	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fils - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var fil models.Fil
		errScan := sqlResult.Scan(&fil.ID, &fil.Titre, &fil.Statut, &fil.DateCreation, &fil.UtilisateurID)
		if errScan != nil {
			return nil, errScan
		}
		listFils = append(listFils, fil)
	}

	return listFils, nil
}

func (r *FilRepository) ReadById(id int) (models.Fil, error) {
	var fil models.Fil
	query := "SELECT `id_fil`, `titre`, `statut`, `date_creation`, `id_utilisateur` FROM `fil_discussion` WHERE `id_fil` = ?;"

	sqlErr := r.db.QueryRow(query, id).
		Scan(&fil.ID, &fil.Titre, &fil.Statut, &fil.DateCreation, &fil.UtilisateurID)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Fil{}, nil
		}
		return models.Fil{}, fmt.Errorf("Erreur récupération fil - Erreur : \n\t %s", sqlErr.Error())
	}

	return fil, nil
}

func (r *FilRepository) UpdateFilById(fil models.Fil) error {
	query := "UPDATE `fil_discussion` SET `titre`=?, `statut`=? WHERE `id_fil`=?;"

	sqlResult, sqlErr := r.db.Exec(query,
		fil.Titre,
		fil.Statut,
		fil.ID)

	if sqlErr != nil {
		return fmt.Errorf("Erreur modification fil - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur modification fil - Aucune ligne modifiée")
	}

	return nil
}

func (r *FilRepository) DeleteFilById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM `fil_discussion` WHERE `id_fil`=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression fil - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur suppression fil - Aucun fil supprimé")
	}

	return nil
}

func (r *FilRepository) ReadAllPaginated(page int, limit int) ([]models.Fil, error) {
	var listFils []models.Fil
	offset := (page - 1) * limit

	sqlResult, sqlErr := r.db.Query("SELECT `id_fil`, `titre`, `statut`, `date_creation`, `id_utilisateur` FROM `fil_discussion` WHERE `statut` != 'archivé' ORDER BY `date_creation` DESC LIMIT ? OFFSET ?;", limit, offset)
	if sqlErr != nil {
		return listFils, fmt.Errorf("Erreur récupération fils - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var fil models.Fil
		errScan := sqlResult.Scan(&fil.ID, &fil.Titre, &fil.Statut, &fil.DateCreation, &fil.UtilisateurID)
		if errScan != nil {
			return nil, errScan
		}
		listFils = append(listFils, fil)
	}

	return listFils, nil
}
