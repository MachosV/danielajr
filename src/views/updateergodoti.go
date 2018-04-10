package views

import (
	"log"
	"models"
	"net/http"
	"storage"
	"strconv"
)

func UpdateErgodoti(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ShowUpdateFormErgodoti(w, r)
		return
	}
	id, err := strconv.ParseInt(r.PostFormValue("idergodoti"), 10, 64)
	log.Println(id)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
		return
	}
	var db = storage.GetDb()
	stmt, _ := db.Prepare("UPDATE ergodotes set phone = ? where id = ?")
	result, err := stmt.Exec(r.PostFormValue("phone"), id)
	_ = result
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/ergodotis?id="+r.PostFormValue("idergodoti"), http.StatusMovedPermanently)
}

func ShowUpdateFormErgodoti(w http.ResponseWriter, r *http.Request) {
	_, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/ergodotes", http.StatusMovedPermanently)
	}
	var db = storage.GetDb()
	res, _ := db.Query("SELECT * FROM ergodotes WHERE id = ?", r.URL.Query().Get("id"))
	var ergodotis models.Ergodotis
	if res.Next() {
		err := res.Scan(
			&ergodotis.ID,
			&ergodotis.Name,
			&ergodotis.Surname,
			&ergodotis.Phone,
			&ergodotis.Poso)
		if err != nil {
			log.Println(err)
		}
	}
	res.Close()
	var t = storage.GetTemplate("updateergodoti.html")
	t.ExecuteTemplate(w, "base", ergodotis)
}
