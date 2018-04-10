package views

import (
	"log"
	"models"
	"net/http"
	"storage"
	"strconv"
	"time"
)

/*
CreateXrewsi creates a bil
*/
func CreateXrewsi(w http.ResponseWriter, r *http.Request) {
	var xrewsi models.Xrewsi
	poso, err := strconv.ParseFloat(r.PostFormValue("poso"), 64)
	if err != nil {
		http.Redirect(w, r, "/xrewseis", http.StatusMovedPermanently)
		return
	}
	xrewstis, err := strconv.ParseInt(r.PostFormValue("xrewstis"), 10, 64)
	xrewsi.Poso = poso
	xrewsi.Xrewstis = int(xrewstis)
	xrewsi.Status = false
	xrewsi.DateRegistered = time.Now().Format("Monday 02/01/2006 15:04:05")
	xrewsi.WorkDate = r.PostFormValue("workdate")
	xrewsi.Aitiologia = r.PostFormValue("aitiologia")
	var db = storage.GetDb()
	stmt, _ := db.Prepare("INSERT INTO xrewseis (poso, status, xrewstis, workdate, dateregistered, aitiologia) values(?,?,?,?,?,?)")
	res, _ := stmt.Exec(xrewsi.Poso, xrewsi.Status, xrewsi.Xrewstis, xrewsi.WorkDate, xrewsi.DateRegistered, xrewsi.Aitiologia)
	_ = res
	stmt, _ = db.Prepare("UPDATE ergodotes SET poso = poso + ? where id = ?")
	res, _ = stmt.Exec(poso, xrewsi.Xrewstis)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/xrewseis", http.StatusMovedPermanently)
}
