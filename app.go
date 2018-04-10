package main

import (
	"log"
	"net/http"
	_ "storage"
	"views"
)

func main() {
	log.Println("Starting server..")
	http.HandleFunc("/idioktitis", views.RetrieveIdioktiti)
	http.HandleFunc("/ergodotis", views.RetrieveErgodoti)
	http.HandleFunc("/diamerisma", views.RetrieveDiamerisma)
	http.HandleFunc("/deleteidioktiti", views.DeleteIdioktiti)
	http.HandleFunc("/deleteergodoti", views.DeleteErgodoti)
	http.HandleFunc("/deletediamerisma", views.DeleteDiamerisma)
	http.HandleFunc("/deletexrewsi", views.DeleteXrewsi)
	http.HandleFunc("/createxrewsi", views.CreateXrewsi)
	http.HandleFunc("/createergodoti", views.CreateErgodoti)
	http.HandleFunc("/createidioktiti", views.CreateIdioktiti)
	http.HandleFunc("/creatediamerisma", views.CreateDiamerisma)
	http.HandleFunc("/idioktites", views.ListIdioktites)
	http.HandleFunc("/diamerismata", views.ListDiamerismata)
	http.HandleFunc("/ergodotes", views.ListErgodotes)
	http.HandleFunc("/xrewseis", views.ListXrewseis)
	http.HandleFunc("/updateidioktiti", views.UpdateIdioktiti)
	http.HandleFunc("/updateergodoti", views.UpdateErgodoti)
	http.HandleFunc("/updateergodotiposo", views.UpdateErgodotiPoso)
	http.HandleFunc("/", views.Index)
	http.ListenAndServe(":8000", nil)
}
