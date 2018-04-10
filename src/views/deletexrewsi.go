package views

import (
	"log"
	"net/http"
	"storage"
	"strconv"
)

/*
DeleteXrewsi diagrafei xrewseis
*/
func DeleteXrewsi(w http.ResponseWriter, r *http.Request) {
	var ids = r.PostFormValue("id_xrewsis")
	id, err := strconv.Atoi(ids)
	_ = id
	if err != nil {
		http.Redirect(w, r, "/xrewseis", http.StatusMovedPermanently)
	}
	var db = storage.GetDb()
	stmt, err := db.Prepare("DELETE FROM xrewseis where xrewseis.id = ?")
	_, err = stmt.Exec(ids)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/xrewseis", http.StatusMovedPermanently)
}
