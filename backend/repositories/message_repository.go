package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MessageRepository struct {
	db *sql.DB
}

func InitMessageRepository(db *sql.DB) *MessageRepository {
	return &MessageRepository{db}
}

func (r *MessageRepository) CreateMessage(message models.Message) (int, error) {
	query := "INSERT INTO `message`(`contenu`, `date_envoi`, `id_utilisateur`, `id_fil`) VALUES (?,?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		message.Contenu,
		time.Now().Format("2006-01-02 15:04:05"),
		message.UtilisateurID,
		message.FilID,
	)

	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création message - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création message - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *MessageRepository) ReadAll() ([]models.Message, error) {
	var listMessages []models.Message
	sqlResult, sqlErr := r.db.Query("SELECT `id_message`, `contenu`, `date_envoi`, `id_utilisateur`, `id_fil` FROM `message` ORDER BY `date_envoi` DESC LIMIT 100;")
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		errScan := sqlResult.Scan(&message.ID, &message.Contenu, &message.DateCreation, &message.UtilisateurID, &message.FilID)
		if errScan != nil {
			return nil, errScan
		}
		listMessages = append(listMessages, message)
	}

	return listMessages, nil
}

func (r *MessageRepository) ReadById(id int) (models.Message, error) {
	var message models.Message
	query := "SELECT `id_message`, `contenu`, `date_envoi`, `id_utilisateur`, `id_fil` FROM `message` WHERE `id_message` = ?;"

	sqlErr := r.db.QueryRow(query, id).
		Scan(&message.ID, &message.Contenu, &message.DateCreation, &message.UtilisateurID, &message.FilID)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Message{}, nil
		}
		return models.Message{}, fmt.Errorf("Erreur récupération message - Erreur : \n\t %s", sqlErr.Error())
	}

	return message, nil
}

func (r *MessageRepository) ReadByFilId(filId int) ([]models.Message, error) {
	var listMessages []models.Message
	sqlResult, sqlErr := r.db.Query("SELECT `id_message`, `contenu`, `date_envoi`, `id_utilisateur`, `id_fil` FROM `message` WHERE `id_fil` = ? ORDER BY `date_envoi` ASC;", filId)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages du fil - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		errScan := sqlResult.Scan(&message.ID, &message.Contenu, &message.DateCreation, &message.UtilisateurID, &message.FilID)
		if errScan != nil {
			return nil, errScan
		}
		listMessages = append(listMessages, message)
	}

	return listMessages, nil
}

func (r *MessageRepository) UpdateMessageById(message models.Message) error {
	query := "UPDATE `message` SET `contenu`=? WHERE `id_message`=?;"

	sqlResult, sqlErr := r.db.Exec(query,
		message.Contenu,
		message.ID)

	if sqlErr != nil {
		return fmt.Errorf("Erreur modification message - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur modification message - Aucune ligne modifiée")
	}

	return nil
}

func (r *MessageRepository) DeleteMessageById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM `message` WHERE `id_message`=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression message - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur suppression message - Aucun message supprimé")
	}

	return nil
}

func (r *MessageRepository) ReadByFilIdPaginated(filId int, page int, limit int) ([]models.Message, error) {
	var listMessages []models.Message
	offset := (page - 1) * limit

	sqlResult, sqlErr := r.db.Query("SELECT `id_message`, `contenu`, `date_envoi`, `id_utilisateur`, `id_fil` FROM `message` WHERE `id_fil` = ? ORDER BY `date_envoi` DESC LIMIT ? OFFSET ?;", filId, limit, offset)
	if sqlErr != nil {
		return listMessages, fmt.Errorf("Erreur récupération messages du fil - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		errScan := sqlResult.Scan(&message.ID, &message.Contenu, &message.DateCreation, &message.UtilisateurID, &message.FilID)
		if errScan != nil {
			return nil, errScan
		}
		listMessages = append(listMessages, message)
	}

	return listMessages, nil
}

func (r *MessageRepository) ReadByFilIdWithScores(filId int) (map[int]models.Message, error) {
	messageMap := make(map[int]models.Message)

	sqlResult, sqlErr := r.db.Query("SELECT `id_message`, `contenu`, `date_envoi`, `id_utilisateur`, `id_fil` FROM `message` WHERE `id_fil` = ?;", filId)
	if sqlErr != nil {
		return messageMap, fmt.Errorf("Erreur récupération messages - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var message models.Message
		errScan := sqlResult.Scan(&message.ID, &message.Contenu, &message.DateCreation, &message.UtilisateurID, &message.FilID)
		if errScan != nil {
			return nil, errScan
		}
		messageMap[message.ID] = message
	}

	return messageMap, nil
}
