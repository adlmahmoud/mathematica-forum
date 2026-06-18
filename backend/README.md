# Mathematica Forum - Backend Go

## 📋 Vue d'ensemble
Ce projet est un backend RESTful en Go pour un forum de discussions mathématiques.

### Architecture
- **Framework**: Gorilla Mux (routing HTTP)
- **Base de données**: MySQL
- **Authentification**: JWT
- **Structuration**: Architecture en couches (Repository → Service → Controller)

---

## 🚀 Installation & Démarrage

### 1. Prérequis
- Go 1.18+
- MySQL 8.0+
- Golang JWT package: `github.com/golang-jwt/jwt/v5`
- Gorilla Mux: `github.com/gorilla/mux`
- MySQL driver: `github.com/go-sql-driver/mysql`
- Godotenv: `github.com/joho/godotenv`

### 2. Installation des dépendances
```bash
cd backend
go mod download
```

### 3. Configuration DB
Créer la base de données et les tables:
```sql
SOURCE migration/migrations.sql;
```

### 4. Variables d'environnement
Le fichier `.env` à la racine du backend doit contenir:
```
DB_NAME=mathematica_forum
DB_USER=root
DB_PWD=
DB_HOST=127.0.0.1
DB_PORT=3306
JWT_SECRET=your_super_secret_jwt_key_change_in_production
```

### 5. Lancer le serveur
```bash
go run main.go
```

Le serveur écoute sur `http://localhost:8080`

---

## 📡 API Endpoints

### Authentication
- `POST /api/auth/login` - Connexion (retourne JWT token)

### Utilisateurs
- `GET /api/utilisateurs` - Lister tous les utilisateurs
- `GET /api/utilisateurs/{id}` - Récupérer un utilisateur
- `POST /api/utilisateurs` - Créer un utilisateur
- `PUT /api/utilisateurs/{id}` - Modifier un utilisateur (🔐 Authentifié)
- `DELETE /api/utilisateurs/{id}` - Supprimer un utilisateur (🔐 Authentifié)

### Fils de discussion (Threads)
- `GET /api/fils` - Lister tous les fils
- `GET /api/fils/{id}` - Récupérer un fil
- `POST /api/fils` - Créer un fil (🔐 Authentifié)
- `PUT /api/fils/{id}` - Modifier un fil (🔐 Authentifié)
- `DELETE /api/fils/{id}` - Supprimer un fil (🔐 Authentifié)

### Messages (Réponses)
- `GET /api/messages` - Lister tous les messages
- `GET /api/messages/{id}` - Récupérer un message
- `GET /api/fils/{filId}/messages` - Lister les messages d'un fil
- `POST /api/messages` - Créer un message (🔐 Authentifié)
- `PUT /api/messages/{id}` - Modifier un message (🔐 Authentifié)
- `DELETE /api/messages/{id}` - Supprimer un message (🔐 Authentifié)

### 🔐 Routes protégées
Nécessitent le header:
```
Authorization: Bearer <JWT_TOKEN>
```

---

## 📂 Structure du projet
```
backend/
├── app/                    # Initialisation application
├── config/                 # Configuration (DB, ENV)
├── controllers/            # Contrôleurs HTTP
├── middleware/             # Middlewares (auth, etc.)
├── models/                 # Modèles de données
├── repositories/           # Accès aux données
├── routes/                 # Définition des routes
├── services/               # Logique métier
├── migration/              # Migrations SQL
├── .env                    # Variables d'environnement
├── main.go                 # Point d'entrée
└── go.mod / go.sum         # Dépendances Go
```

---

## 🔐 Authentification JWT

### Exemple de login
```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"user@example.com","password":"Password123!"}'
```

**Réponse:**
```json
{
  "token": "eyJhbGc...",
  "user": {
    "id": 1,
    "username": "john_doe",
    "email": "user@example.com",
    "is_admin": false
  }
}
```

### Utiliser le token
```bash
curl -X POST http://localhost:8080/api/messages \
  -H "Authorization: Bearer eyJhbGc..." \
  -H "Content-Type: application/json" \
  -d '{"contenu":"Mon message","id_utilisateur":1,"id_fil":1}'
```

---

## 💡 Points forts du projet
✅ Architecture en couches bien organisée  
✅ Authentification JWT sécurisée  
✅ Middleware d'autorisation  
✅ Gestion complète des erreurs  
✅ SQL préparé (prévention injection)  
✅ Routage avec Gorilla Mux  
✅ Pagination et limitation des requêtes  
✅ Support des associations (Fil ↔ Messages)  

---

## 🛠️ Dépannage

### "Erreur connection base de donnees"
- Vérifier que MySQL est en cours d'exécution
- Vérifier les variables .env (DB_HOST, DB_PORT, DB_USER)

### "Token invalide"
- Vérifier que le header `Authorization` est au format `Bearer <token>`
- Vérifier que la clé JWT_SECRET est correcte

### "Table doesn't exist"
- Exécuter le fichier `migration/migrations.sql`

---

## 📝 Notes pour la soutenance
- Architecture MVC: Repository → Service → Controller
- Authentification: JWT avec expiration 24h
- Base de données: Schema normalisé avec clés étrangères
- Middleware: Protection des routes sensibles
- Validation: Côté service (métier) et contrôleur
