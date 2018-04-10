package views

import (
	"log"
	"models"
	"net/http"
	"storage"
	"strconv"
)

func UpdateIdioktiti(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ShowUpdateFormidioktiti(w, r)
		return
	}
	id, err := strconv.ParseInt(r.PostFormValue("ididioktiti"), 10, 64)
	log.Println(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/idioktites", http.StatusMovedPermanently)
		return
	}
	var db = storage.GetDb()
	stmt, _ := db.Prepare("UPDATE idioktites set phone = ? where id = ?")
	result, err := stmt.Exec(r.PostFormValue("phone"), id)
	_ = result
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/idioktitis?id="+r.PostFormValue("ididioktiti"), http.StatusMovedPermanently)
}

func ShowUpdateFormidioktiti(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/idioktites", http.StatusMovedPermanently)
	}
	var db = storage.GetDb()
	res, _ := db.Query("SELECT * FROM idioktites WHERE id = ?", r.URL.Query().Get("id"))
	var idioktitis models.Idioktitis
	if res.Next() {
		err := res.Scan(
			&idioktitis.ID,
			&idioktitis.Name,
			&idioktitis.Surname,
			&idioktitis.Phone)
		if err != nil {
			log.Println(err)
		}
	}
	res.Close()
	var t = storage.GetTemplate("updateidioktiti.html")
	t.ExecuteTemplate(w, "base", idioktitis)
}
