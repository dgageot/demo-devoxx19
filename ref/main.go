package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/count", count)
	r.HandleFunc("/hexagon/{index:[0-9]+}", hexagon)

	log.Println("Listening on port 8888")
	http.ListenAndServe(":8888", r)
}

func count(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	hexagons, err := readHexagons()
	if err != nil {
		http.Error(w, "Unable to read hexagons", 500)
		return
	}

	writeJSON(w, Info{
		Count: len(hexagons),
	})
}

func hexagon(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	vars := mux.Vars(r)
	index, _ := strconv.Atoi(vars["index"])

	hexagons, err := readHexagons()
	if err != nil {
		http.Error(w, "Unable to read hexagons", 500)
		return
	}

	writeJSON(w, hexagons[index])
}

func readHexagons() ([]Hexagon, error) {
	db, err := sql.Open("postgres", "postgres://hexagons:notsosecret@db/hexagons?sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT image, category, name, description, url FROM hexagons")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var hexagons []Hexagon
	for rows.Next() {
		var h Hexagon
		if err := rows.Scan(&h.Image, &h.Category, &h.Name, &h.Description, &h.URL); err != nil {
			return nil, err
		}

		hexagons = append(hexagons, h)
	}

	return hexagons, nil
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	buf, err := json.Marshal(v)
	if err != nil {
		http.Error(w, "Unable to marshal value", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(buf)
}
