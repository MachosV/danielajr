package views

import (
	"log"
	"models"
	"net/http"
	"storage"
	"strconv"
)

func CreateDiamerisma(w http.ResponseWriter, r *http.Request) {
	var diamerisma models.Diamerisma
	idioktitis, err := strconv.ParseInt(r.PostFormValue("idioktitis"), 10, 64)
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/diamerismata", http.StatusMovedPermanently)
		return
	}
	diamerisma.Name = r.PostFormValue("name")
	diamerisma.Idioktitis = int(idioktitis)
	diamerisma.Dieythinsi = r.PostFormValue("dieythinsi")
	diamerisma.Perigrafi = r.PostFormValue("perigrafi")
	var db = storage.GetDb()
	stmt, err := db.Prepare("INSERT INTO diamerismata (name,idioktitis,dieythinsi,perigrafi) VALUES(?,?,?,?);")
	res, err := stmt.Exec(diamerisma.Name, diamerisma.Idioktitis, diamerisma.Dieythinsi, diamerisma.Perigrafi)
	if err != nil {
		log.Println(err)
	}
	_ = res
	stmt.Close()
	http.Redirect(w, r, "/diamerismata", http.StatusMovedPermanently)
}
