package repositories

import (
	"database/sql"
	"fmt"
	"mathematica-forum/models"

	_ "github.com/go-sql-driver/mysql"
)

type ReactionRepository struct {
	db *sql.DB
}

func InitReactionRepository(db *sql.DB) *ReactionRepository {
	return &ReactionRepository{db}
}

func (r *ReactionRepository) CreateReaction(reaction models.Reaction) (int, error) {
	query := "INSERT INTO `reaction`(`type_reaction`, `id_utilisateur`, `id_message`) VALUES (?,?,?);"

	sqlResult, sqlErr := r.db.Exec(query,
		reaction.Type,
		reaction.UtilisateurID,
		reaction.MessageID,
	)

	if sqlErr != nil {
		return -1, fmt.Errorf("Erreur création réaction - Erreur : \n\t %s", sqlErr.Error())
	}

	id, idErr := sqlResult.LastInsertId()
	if idErr != nil {
		return -1, fmt.Errorf("Erreur création réaction - Erreur récupération identifiant : \n\t %s", idErr.Error())
	}

	return int(id), nil
}

func (r *ReactionRepository) ReadByMessageID(messageID int) ([]models.Reaction, error) {
	var listReactions []models.Reaction
	sqlResult, sqlErr := r.db.Query("SELECT `id_reaction`, `type_reaction`, `id_utilisateur`, `id_message` FROM `reaction` WHERE `id_message` = ?;", messageID)
	if sqlErr != nil {
		return listReactions, fmt.Errorf("Erreur récupération réactions - Erreur : \n\t %s", sqlErr.Error())
	}

	defer sqlResult.Close()

	for sqlResult.Next() {
		var reaction models.Reaction
		errScan := sqlResult.Scan(&reaction.ID, &reaction.Type, &reaction.UtilisateurID, &reaction.MessageID)
		if errScan != nil {
			return nil, errScan
		}
		listReactions = append(listReactions, reaction)
	}

	return listReactions, nil
}

func (r *ReactionRepository) ReadByUserAndMessage(userID int, messageID int) (models.Reaction, error) {
	var reaction models.Reaction
	query := "SELECT `id_reaction`, `type_reaction`, `id_utilisateur`, `id_message` FROM `reaction` WHERE `id_utilisateur` = ? AND `id_message` = ?;"

	sqlErr := r.db.QueryRow(query, userID, messageID).
		Scan(&reaction.ID, &reaction.Type, &reaction.UtilisateurID, &reaction.MessageID)

	if sqlErr != nil {
		if sqlErr == sql.ErrNoRows {
			return models.Reaction{}, nil
		}
		return models.Reaction{}, fmt.Errorf("Erreur récupération réaction - Erreur : \n\t %s", sqlErr.Error())
	}

	return reaction, nil
}

func (r *ReactionRepository) DeleteReactionById(id int) error {
	sqlResult, sqlErr := r.db.Exec("DELETE FROM `reaction` WHERE `id_reaction`=?;", id)
	if sqlErr != nil {
		return fmt.Errorf("Erreur suppression réaction - Erreur : \n\t %s", sqlErr.Error())
	}

	if nbrRow, _ := sqlResult.RowsAffected(); nbrRow <= 0 {
		return fmt.Errorf("Erreur suppression réaction - Aucune réaction supprimée")
	}

	return nil
}

func (r *ReactionRepository) CountLikesByMessageID(messageID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM `reaction` WHERE `id_message` = ? AND `type_reaction` = 'like';"

	sqlErr := r.db.QueryRow(query, messageID).Scan(&count)
	if sqlErr != nil {
		return 0, fmt.Errorf("Erreur comptage likes - Erreur : \n\t %s", sqlErr.Error())
	}

	return count, nil
}

func (r *ReactionRepository) CountDislikesByMessageID(messageID int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM `reaction` WHERE `id_message` = ? AND `type_reaction` = 'dislike';"

	sqlErr := r.db.QueryRow(query, messageID).Scan(&count)
	if sqlErr != nil {
		return 0, fmt.Errorf("Erreur comptage dislikes - Erreur : \n\t %s", sqlErr.Error())
	}

	return count, nil
}
