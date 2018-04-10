package views

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"storage"
)

type DataXrewsewn struct {
	Xrewseis  []models.Xrewsi
	Ergodotes []models.Ergodotis
}

/*
ListXrewseis lista xrewsewn, 8a tis bgazei oles
*/
func ListXrewseis(w http.ResponseWriter, r *http.Request) {
	var db = storage.GetDb()
	var dataxrewsewn DataXrewsewn
	results, err := db.Query("SELECT id,name,surname from ergodotes")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var ergodotis models.Ergodotis
	for results.Next() {
		err = results.Scan(
			&ergodotis.ID,
			&ergodotis.Name,
			&ergodotis.Surname)
		dataxrewsewn.Ergodotes = append(dataxrewsewn.Ergodotes, ergodotis)
	}
	results.Close()
	results, err = db.Query("SELECT " +
		"xrewseis.ID," +
		"xrewseis.poso," +
		"xrewseis.status," +
		"xrewseis.workdate," +
		"xrewseis.dateregistered," +
		"xrewseis.aitiologia," +
		"ergodotes.name," +
		"ergodotes.surname " +
		"FROM xrewseis JOIN ergodotes ON " +
		"xrewseis.xrewstis = ergodotes.ID ORDER BY workdate")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var xrewsi models.Xrewsi
	for results.Next() {
		err = results.Scan(
			&xrewsi.ID,
			&xrewsi.Poso,
			&xrewsi.Status,
			&xrewsi.WorkDate,
			&xrewsi.DateRegistered,
			&xrewsi.Aitiologia,
			&xrewsi.Name,
			&xrewsi.Surname)
		dataxrewsewn.Xrewseis = append(dataxrewsewn.Xrewseis, xrewsi)
	}
	//fmt.Fprint(w, xrewseis.Xrewseis)
	results.Close()
	var t = storage.GetTemplate("xrewseis.html")
	t.ExecuteTemplate(w, "base", dataxrewsewn)
}
