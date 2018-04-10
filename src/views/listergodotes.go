package views

import (
	"fmt"
	"log"
	"models"
	"net/http"
	"storage"
)

type Ergodotes struct {
	Ergodotes []models.Ergodotis
}

func ListErgodotes(w http.ResponseWriter, r *http.Request) {
	var db = storage.GetDb()
	results, err := db.Query("SELECT * FROM ergodotes ORDER BY surname COLLATE NOCASE ASC")
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, "Database error.")
		return
	}
	var ergodotis models.Ergodotis
	var ergodotes Ergodotes
	for results.Next() {
		err = results.Scan(
			&ergodotis.ID,
			&ergodotis.Name,
			&ergodotis.Surname,
			&ergodotis.Phone,
			&ergodotis.Poso)
		ergodotes.Ergodotes = append(ergodotes.Ergodotes, ergodotis)
	}
	results.Close()
	var t = storage.GetTemplate("ergodotes.html")
	t.ExecuteTemplate(w, "base", ergodotes)
}
