package views

import (
	"models"
	"net/http"
	"storage"
	"strconv"
)

func RetrieveIdioktiti(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/idioktites", http.StatusMovedPermanently)
	}
	var db = storage.GetDb()
	res, _ := db.Query("SELECT * FROM idioktites WHERE id = ?", id)
	var idioktitis models.Idioktitis
	if res.Next() {
		res.Scan(
			&idioktitis.ID,
			&idioktitis.Name,
			&idioktitis.Surname,
			&idioktitis.Phone)
	}
	res.Close()
	var t = storage.GetTemplate("idioktitis.html")
	t.ExecuteTemplate(w, "base", idioktitis)
}
