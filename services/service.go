package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/juliofernandolepore/webserver/models"
)

var dbconn *sqlx.DB

func SetDB(db *sqlx.DB) {
	dbconn = db
}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var posts = models.GetPosts()

	sqlStmt := `SELECT * FROM posts`
	rows, err := dbconn.Queryx(sqlStmt)

	if err == nil {
		var tempPost = models.GetPost()

		for rows.Next() {
			err = rows.StructScan(&tempPost)
			posts = append(posts, tempPost)
		}

		switch err {
		case sql.ErrNoRows:
			{
				log.Println("no rows returned.")
				http.Error(w, err.Error(), http.StatusNoContent)
			}
		case nil:
			json.NewEncoder(w).Encode(&posts)

		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	} else {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
