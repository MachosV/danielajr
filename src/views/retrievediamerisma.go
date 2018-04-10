package views

import (
	"models"
	"net/http"
	"storage"
	"strconv"
)

func RetrieveDiamerisma(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Redirect(w, r, "/diamerismata", http.StatusMovedPermanently)
		return
	}
	db := storage.GetDb()
	var query = "SELECT " +
		"diamerismata.ID," +
		"diamerismata.name," +
		"diamerismata.dieythinsi," +
		"diamerismata.perigrafi," +
		"idioktites.name," +
		"idioktites.surname," +
		"idioktites.phone " +
		"FROM diamerismata " +
		"JOIN idioktites ON " +
		"diamerismata.idioktitis = idioktites.id WHERE diamerismata.id = ?"
	res, _ := db.Query(query, id)
	var diamerisma models.Diamerisma
	if res.Next() {
		var idioktitis models.Idioktitis
		diamerisma.IdioktitisObj = idioktitis
		res.Scan(
			&diamerisma.ID,
			&diamerisma.Name,
			&diamerisma.Dieythinsi,
			&diamerisma.Perigrafi,
			&diamerisma.IdioktitisObj.Name,
			&diamerisma.IdioktitisObj.Surname,
			&diamerisma.IdioktitisObj.Phone)
	}
	res.Close()
	var t = storage.GetTemplate("diamerisma.html")
	t.ExecuteTemplate(w, "base", diamerisma)
}
