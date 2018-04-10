package views

import (
	"models"
	"net/http"
	"storage"
)

/*
CreateErgodoti dimiourgia ergodoti
*/
func CreateErgodoti(w http.ResponseWriter, r *http.Request) {
	var ergodotis models.Ergodotis
	ergodotis.Name = r.PostFormValue("name")
	ergodotis.Surname = r.PostFormValue("surname")
	ergodotis.Phone = r.PostFormValue("phone")
	ergodotis.Poso = 0
	var db = storage.GetDb()
	stmt, _ := db.Prepare("INSERT INTO ergodotes (name, surname, phone, poso)values(?,?,?,?)")
	res, _ := stmt.Exec(ergodotis.Name, ergodotis.Surname, ergodotis.Phone, ergodotis.Poso)
	_ = res
	http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
}
