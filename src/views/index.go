package views

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"storage"
)

type Results struct {
	Ergodotes    []models.Ergodotis
	Idioktites   []models.Idioktitis
	Diamerismata []models.Diamerisma
}

func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		DisplaySearchForm(w, r)
		return
	} else if r.Method == http.MethodPost {
		DisplayResults(w, r)
		return
	} else {
		DisplaySearchForm(w, r)
	}
}

func DisplaySearchForm(w http.ResponseWriter, r *http.Request) {
	var t = storage.GetTemplate("index.html")
	t.ExecuteTemplate(w, "base", nil)
}

func DisplayResults(w http.ResponseWriter, r *http.Request) {
	anazitisi := r.PostFormValue("anazitisi")
	var apotelesmata Results
	var query = "SELECT " +
		"diamerismata.ID," +
		"diamerismata.name," +
		"diamerismata.dieythinsi," +
		"diamerismata.perigrafi," +
		"idioktites.name AS onomata," +
		"idioktites.surname," +
		"idioktites.phone " +
		"FROM diamerismata " +
		"JOIN idioktites ON " +
		"diamerismata.idioktitis = idioktites.id " +
		"WHERE dieythinsi LIKE '%" + anazitisi + "%' OR diamerismata.name LIKE '%" + anazitisi + "%' ORDER BY diamerismata.name COLLATE NOCASE ASC"
	var db = storage.GetDb()
	results, err := db.Query(query)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var diamerisma models.Diamerisma
	var idioktitis models.Idioktitis
	diamerisma.IdioktitisObj = idioktitis
	for results.Next() {
		results.Scan(
			&diamerisma.ID,
			&diamerisma.Name,
			&diamerisma.Dieythinsi,
			&diamerisma.Perigrafi,
			&diamerisma.IdioktitisObj.Name,
			&diamerisma.IdioktitisObj.Surname,
			&diamerisma.IdioktitisObj.Phone)
		apotelesmata.Diamerismata = append(apotelesmata.Diamerismata, diamerisma)
	}
	results.Close()
	query = "SELECT * from idioktites WHERE surname LIKE '%" + anazitisi + "%'"
	results, err = db.Query(query)
	if err != nil {
		log.Println(err)
		log.Println("query idioktites")
	}
	for results.Next() {
		results.Scan(
			&idioktitis.ID,
			&idioktitis.Name,
			&idioktitis.Surname,
			&idioktitis.Phone)
		apotelesmata.Idioktites = append(apotelesmata.Idioktites, idioktitis)
	}
	query = "SELECT * FROM ergodotes WHERE surname LIKE '%" + anazitisi + "%'"
	results, err = db.Query(query)
	if err != nil {
		log.Println(err)
		log.Println("query ergodotes")
	}
	var ergodotis models.Ergodotis
	for results.Next() {
		results.Scan(
			&ergodotis.ID,
			&ergodotis.Name,
			&ergodotis.Surname,
			&ergodotis.Phone,
			&ergodotis.Poso)
		apotelesmata.Ergodotes = append(apotelesmata.Ergodotes, ergodotis)
	}
	var t = storage.GetTemplate("index.html")
	t.ExecuteTemplate(w, "base", apotelesmata)
}
