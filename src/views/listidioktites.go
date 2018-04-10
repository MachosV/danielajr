package views

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"storage"
)

type Idioktites struct {
	Idioktites []models.Idioktitis
}

func ListIdioktites(w http.ResponseWriter, r *http.Request) {
	var db = storage.GetDb()
	results, err := db.Query("SELECT * FROM idioktites ORDER BY surname COLLATE NOCASE ASC")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var idioktitis models.Idioktitis
	var idioktites Idioktites
	for results.Next() {
		err = results.Scan(
			&idioktitis.ID,
			&idioktitis.Name,
			&idioktitis.Surname,
			&idioktitis.Phone)
		idioktites.Idioktites = append(idioktites.Idioktites, idioktitis)
	}
	results.Close()
	var t = storage.GetTemplate("idioktites.html")
	t.ExecuteTemplate(w, "base", idioktites)
}
