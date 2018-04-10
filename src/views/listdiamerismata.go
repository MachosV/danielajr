package views

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"storage"
)

type Diamerismata struct {
	Diamerismata []models.Diamerisma
	Idioktites   []models.Idioktitis
}

func ListDiamerismata(w http.ResponseWriter, r *http.Request) {
	var query = "SELECT " +
		"diamerismata.ID," +
		"diamerismata.name," +
		"diamerismata.dieythinsi," +
		"diamerismata.perigrafi," +
		"idioktites.name," +
		"idioktites.surname," +
		"idioktites.phone " +
		"FROM diamerismata " +
		"JOIN idioktites ON " +
		"diamerismata.idioktitis = idioktites.id ORDER BY diamerismata.name COLLATE NOCASE ASC"
	var db = storage.GetDb()
	results, err := db.Query("SELECT id, name, surname from idioktites")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var idioktitis models.Idioktitis
	var diamerismata Diamerismata
	for results.Next() {
		err = results.Scan(
			&idioktitis.ID,
			&idioktitis.Name,
			&idioktitis.Surname)
		diamerismata.Idioktites = append(diamerismata.Idioktites, idioktitis)
	}
	results.Close()
	results, err = db.Query(query)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var diamerisma models.Diamerisma
	diamerisma.IdioktitisObj = idioktitis
	for results.Next() {
		err = results.Scan(
			&diamerisma.ID,
			&diamerisma.Name,
			&diamerisma.Dieythinsi,
			&diamerisma.Perigrafi,
			&diamerisma.IdioktitisObj.Name,
			&diamerisma.IdioktitisObj.Surname,
			&diamerisma.IdioktitisObj.Phone)
		diamerismata.Diamerismata = append(diamerismata.Diamerismata, diamerisma)
	}
	results.Close()
	var t = storage.GetTemplate("diamerismata.html")
	t.ExecuteTemplate(w, "base", diamerismata)
}
