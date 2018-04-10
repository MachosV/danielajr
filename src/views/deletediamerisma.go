package views

import (
	"log"
	"net/http"
	"storage"
	"strconv"
)

func DeleteDiamerisma(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.PostFormValue("iddiamerisma"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/diamerismata", http.StatusMovedPermanently)
		return
	}
	var db = storage.GetDb()
	stmt, _ := db.Prepare("DELETE FROM diamerismata where id = ?")
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/diamerismata", http.StatusMovedPermanently)
}
