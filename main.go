package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	port = ":8080"
	host = "http://localhost"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

func group(w http.ResponseWriter, r *http.Request) {
	// Create a new template and parse the letter into it.
	t, _ := template.ParseFiles("templates/group.html")
	// Make an HTTP GET request to the JSON file
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists/" + "1")

	defer response.Body.Close()

	// Read the response body
	body, _ := ioutil.ReadAll(response.Body)

	// Unmarshal the byte slice into a Data struct.
	var user Artist
	json.Unmarshal(body, &user)

	var buffer bytes.Buffer
	t.Execute(&buffer, user)

	w.Write(buffer.Bytes())

	// fmt.Println(user.Name)
	// fmt.Println(user.Members)

}

func artistHandler(w http.ResponseWriter, r *http.Request) {

	// Create a new template and parse the letter into it.
	t, _ := template.ParseFiles("templates/artist.html")

	// Add that letter to the body of our email
	response, _ := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	var artists []Artist
	json.Unmarshal(body, &artists)
	t.Execute(w, artists)
	// fmt.Println(artists)

}

func main() {

	http.HandleFunc("/group", group)
	http.HandleFunc("/artist", artistHandler)

	// log.Println("Server started at", "\033[1;32m"+host+port+"/hello", "\033[0m")
	log.Println("Server started at", host+port+"/group")
	log.Println("Server started at", host+port+"/artist")

	http.ListenAndServe(port, nil)

}
