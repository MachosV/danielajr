package views

import (
	"models"
	"net/http"
	"storage"
	"strconv"
)

func RetrieveErgodoti(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
	}
	var db = storage.GetDb()
	res, _ := db.Query("SELECT * FROM ergodotes WHERE id = ?", id)
	var ergodotis models.Ergodotis
	if res.Next() {
		res.Scan(
			&ergodotis.ID,
			&ergodotis.Name,
			&ergodotis.Surname,
			&ergodotis.Phone,
			&ergodotis.Poso)
	}
	res.Close()
	var t = storage.GetTemplate("ergodotis.html")
	t.ExecuteTemplate(w, "base", ergodotis)
}
