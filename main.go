package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GeertJohan/go.rice"
	"github.com/abijr/GameOfLife/game"
)

func main() {
	log.Println("Running webapp at http://localhost:8080")
	http.Handle("/", http.FileServer(rice.MustFindBox("static/dist").HTTPBox()))
	http.HandleFunc("/play", play)
	http.ListenAndServe(":8080", nil)
}

func play(w http.ResponseWriter, r *http.Request) {
	var d [][]bool
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)

	err := dec.Decode(&d)
	if err != nil {
		log.Println("Err: ", err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", "localhost")

	g := game.New(d)
	done := make(chan struct{})
	results := g.Play(done)
	s := <-results

	data, err := json.Marshal(s.World())
	fmt.Fprint(w, string(data))
}
