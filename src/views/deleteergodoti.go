package views

import (
	"net/http"
	"storage"
	"strconv"
)

func DeleteErgodoti(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(r.PostFormValue("idergodoti"), 10, 64)
	var db = storage.GetDb()
	stmt, _ := db.Prepare("DELETE FROM ergodotes WHERE id = ?;")
	res, _ := stmt.Exec(id)
	stmt, _ = db.Prepare("DELETE FROM xrewseis where xrewstis = ?;")
	res, _ = stmt.Exec(id)
	_ = res
	http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
}
