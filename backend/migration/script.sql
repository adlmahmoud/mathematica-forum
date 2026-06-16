DROP DATABASE IF EXISTS mathematica_forum;
CREATE DATABASE mathematica_forum
    DEFAULT CHARACTER SET utf8mb4
    DEFAULT COLLATE utf8mb4_unicode_ci;

USE mathematica_forum;

CREATE TABLE niveau (
    id_niveau INT AUTO_INCREMENT PRIMARY KEY,
    niveau_d_etude VARCHAR(100) NOT NULL
) ENGINE=InnoDB;

CREATE TABLE tag (
    id_tag INT AUTO_INCREMENT PRIMARY KEY,
    nom_tag VARCHAR(50) NOT NULL UNIQUE
) ENGINE=InnoDB;

CREATE TABLE categorie (
    id_categorie INT AUTO_INCREMENT PRIMARY KEY,
    nom_categorie VARCHAR(100) NOT NULL UNIQUE
) ENGINE=InnoDB;

CREATE TABLE utilisateur (
    id_utilisateur INT AUTO_INCREMENT PRIMARY KEY,
    nom_utilisateur VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    mot_de_passe_hash VARCHAR(255) NOT NULL,
    sel VARCHAR(255),
    is_admin BOOLEAN DEFAULT FALSE,
    is_banni BOOLEAN DEFAULT FALSE,
    date_inscription TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_niveau INT,

    CONSTRAINT fk_utilisateur_niveau
        FOREIGN KEY (id_niveau)
        REFERENCES niveau(id_niveau)
        ON DELETE SET NULL
) ENGINE=InnoDB;

CREATE TABLE fil_discussion (
    id_fil INT AUTO_INCREMENT PRIMARY KEY,
    titre VARCHAR(255) NOT NULL,
    statut ENUM('ouvert', 'fermé', 'archivé') DEFAULT 'ouvert',
    date_creation TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_utilisateur INT NOT NULL,

    CONSTRAINT fk_fil_discussion_utilisateur
        FOREIGN KEY (id_utilisateur)
        REFERENCES utilisateur(id_utilisateur)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE token_jwt (
    id_token INT AUTO_INCREMENT PRIMARY KEY,
    token VARCHAR(512) NOT NULL,
    date_expiration DATETIME NOT NULL,
    id_utilisateur INT NOT NULL,

    CONSTRAINT fk_token_jwt_utilisateur
        FOREIGN KEY (id_utilisateur)
        REFERENCES utilisateur(id_utilisateur)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE message (
    id_message INT AUTO_INCREMENT PRIMARY KEY,
    contenu TEXT NOT NULL,
    date_envoi TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    id_utilisateur INT NOT NULL,
    id_fil INT NOT NULL,
    id_message_parent INT,

    CONSTRAINT fk_message_utilisateur
        FOREIGN KEY (id_utilisateur)
        REFERENCES utilisateur(id_utilisateur)
        ON DELETE CASCADE,

    CONSTRAINT fk_message_fil_discussion
        FOREIGN KEY (id_fil)
        REFERENCES fil_discussion(id_fil)
        ON DELETE CASCADE,

    CONSTRAINT fk_message_message_parent
        FOREIGN KEY (id_message_parent)
        REFERENCES message(id_message)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE fil_tag (
    id_fil INT NOT NULL,
    id_tag INT NOT NULL,
    PRIMARY KEY (id_fil, id_tag),

    CONSTRAINT fk_fil_tag_fil
        FOREIGN KEY (id_fil)
        REFERENCES fil_discussion(id_fil)
        ON DELETE CASCADE,

    CONSTRAINT fk_fil_tag_tag
        FOREIGN KEY (id_tag)
        REFERENCES tag(id_tag)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE fil_categorie (
    id_fil INT NOT NULL,
    id_categorie INT NOT NULL,
    PRIMARY KEY (id_fil, id_categorie),

    CONSTRAINT fk_fil_categorie_fil
        FOREIGN KEY (id_fil)
        REFERENCES fil_discussion(id_fil)
        ON DELETE CASCADE,

    CONSTRAINT fk_fil_categorie_categorie
        FOREIGN KEY (id_categorie)
        REFERENCES categorie(id_categorie)
        ON DELETE CASCADE
) ENGINE=InnoDB;

CREATE TABLE reaction (
    id_reaction INT AUTO_INCREMENT PRIMARY KEY,
    type_reaction ENUM('like', 'dislike') NOT NULL,
    id_utilisateur INT NOT NULL,
    id_message INT NOT NULL,

    UNIQUE KEY unique_user_message (id_utilisateur, id_message),

    CONSTRAINT fk_reaction_utilisateur
        FOREIGN KEY (id_utilisateur)
        REFERENCES utilisateur(id_utilisateur)
        ON DELETE CASCADE,

    CONSTRAINT fk_reaction_message
        FOREIGN KEY (id_message)
        REFERENCES message(id_message)
        ON DELETE CASCADE
) ENGINE=InnoDB;