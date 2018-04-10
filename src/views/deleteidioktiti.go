package views

import (
	"net/http"
	"storage"
	"strconv"
)

func DeleteIdioktiti(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PostFormValue("ididioktiti"), 10, 64)
	var db = storage.GetDb()
	stmt, _ := db.Prepare("DELETE FROM idioktites WHERE id = ?;")
	res, _ := stmt.Exec(id)
	stmt, _ = db.Prepare("DELETE FROM diamerismata where idioktitis = ?;")
	res, _ = stmt.Exec(id)
	_ = res
	http.Redirect(w, r, "/idioktites", http.StatusMovedPermanently)
}
