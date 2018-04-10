package views

import (
	"models"
	"net/http"
	"storage"
)

/*
CreateErgodoti dimiourgia ergodoti
*/
func CreateIdioktiti(w http.ResponseWriter, r *http.Request) {
	var idioktitis models.Idioktitis
	idioktitis.Name = r.PostFormValue("name")
	idioktitis.Surname = r.PostFormValue("surname")
	idioktitis.Phone = r.PostFormValue("phone")
	var db = storage.GetDb()
	stmt, _ := db.Prepare("INSERT INTO idioktites (name, surname, phone)values(?,?,?)")
	res, _ := stmt.Exec(idioktitis.Name, idioktitis.Surname, idioktitis.Phone)
	_ = res
	http.Redirect(w, r, "/idioktites", http.StatusMovedPermanently)
}
