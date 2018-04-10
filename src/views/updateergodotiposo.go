package views

import (
	"log"
	"net/http"
	"storage"
	"strconv"
)

func UpdateErgodotiPoso(w http.ResponseWriter, r *http.Request) {
	var ypoloipo float64
	poso, err := strconv.ParseFloat(r.PostFormValue("poso"), 64)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
		return
	}
	id, err := strconv.ParseInt(r.PostFormValue("idergodoti"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
		return
	}
	var db = storage.GetDb()
	res, err := db.Query("SELECT poso from ergodotes where id = ?", id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
		return
	}
	if res.Next() {
		res.Scan(&ypoloipo)
	}
	res.Close()
	if ypoloipo <= poso {
		poso = ypoloipo
	}
	stmt, _ := db.Prepare("UPDATE ergodotes set poso = poso - ? where id = ?")
	result, err := stmt.Exec(poso, id)
	_ = result
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/ergodotis?id="+r.PostFormValue("idergodoti"), http.StatusMovedPermanently)
}
