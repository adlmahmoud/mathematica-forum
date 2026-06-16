package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type UtilisateurRepository struct {
	db *sql.DB
}

func InitUtilisateurRepository(db *sql.DB) *UtilisateurRepository {
	return &UtilisateurRepository{db}
}

func (r *UtilisateurRepository) CreateUtilisateur(user models.Utilisateur) (int, error) {
	query := "INSERT INTO `UTILISATEUR`(`nom_utilisateur`, `email`, `mot_de_passe_hash`, `date_inscription`) VALUES (?,?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		user.Username,
		user.Email,
		user.PasswordHash,
		time.Now().Format("2006-01-02 15:04:05"),
	)

	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur ajout utilisateur - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur ajout utilisateur - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *UtilisateurRepository) ReadAll() ([]models.Utilisateur, error) {
	var listUsers []models.Utilisateur
	sqlResult, sqlErr := r.db.Query("SELECT `id_utilisateur`, `nom_utilisateur`, `email`, `mot_de_passe_hash`, `date_inscription` FROM `UTILISATEUR`;")
	if sqlErr != nil {
		return listUsers, fmt.Errorf("Erreur récupération utilisateurs - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var user models.Utilisateur
		errScan := sqlResult.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.DateInscription)
		if errScan != nil {
			return nil, errScan
		}
		listUsers = append(listUsers, user)
	}

	return listUsers, nil
}

func (r *UtilisateurRepository) ReadById(id int) (models.Utilisateur, error) {
	var user models.Utilisateur
	query := "SELECT `id_utilisateur`, `nom_utilisateur`, `email`, `mot_de_passe_hash`, `date_inscription` FROM `UTILISATEUR` WHERE `id_utilisateur` = ?;"

	sqlErr := r.db.QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.DateInscription)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Utilisateur{}, nil
		}
		return models.Utilisateur{}, fmt.Errorf("Erreur récupération utilisateur - Erreur : \n\t %s", sqlErr.Error())
	}

	return user, nil
}

func (r *UtilisateurRepository) UpdateUtilisateurById(user models.Utilisateur) error {
	query := "UPDATE `UTILISATEUR` SET `nom_utilisateur`=?, `email`=?, `mot_de_passe_hash`=? WHERE `id_utilisateur`=?;"

	sqlResult, sqlErr := r.db.Exec(query,
		user.Username,
		user.Email,
		user.PasswordHash,
		user.ID)

	if sqlErr != nil {
		return fmt.Errorf("Erreur modification utilisateur - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur modification utilisateur - Aucune ligne modifiée")
	}

	return nil
}

func (r *UtilisateurRepository) DeleteUtilisateurById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM `UTILISATEUR` WHERE `id_utilisateur`=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression utilisateur - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur suppression utilisateur - Aucun utilisateur supprimé")
	}

	return nil
}
