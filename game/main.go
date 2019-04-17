package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

var rnd = rand.New(rand.NewSource(99))

func main() {
	http.HandleFunc("/game", game)

	log.Println("Listening on port 8888")
	http.ListenAndServe(":8888", nil)
}

// Very inefficient code to choose 2 random hexagons.
func game(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RequestURI)

	var info Info
	if err := readJSON("/count", r.Header, &info); err != nil {
		http.Error(w, "Unable to read count", 500)
		return
	}
	count := info.Count

	randomGuess := rnd.Intn(count)
	var guess Hexagon
	if err := readJSON(fmt.Sprintf("/hexagon/%d", randomGuess), r.Header, &guess); err != nil {
		http.Error(w, "Unable to read guess", 500)
		return
	}

	tries := 0
	randomChoice := randomGuess
	choice := guess
	for randomChoice == randomGuess || guess.Image == choice.Image {
		if tries == 20 {
			http.Error(w, "Couldn't find a choice different than the guess", 500)
			return
		}
		tries++

		randomChoice = rnd.Intn(count)
		if err := readJSON(fmt.Sprintf("/hexagon/%d", randomChoice), r.Header, &choice); err != nil {
			http.Error(w, "Unable to read choice", 500)
			return
		}
	}

	game := Game{
		Guess:  guess,
		Choice: choice,
		Flavor: 1 + rnd.Intn(2),
	}

	writeJSON(w, game)
}

func readJSON(url string, header http.Header, v interface{}) error {
	req, err := http.NewRequest("GET", "http://ref:8888"+url, nil)
	if err != nil {
		return err
	}

	copyHeader(req.Header, header)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.New("Unable to read response")
	}

	err = json.Unmarshal(buf, v)
	if err != nil {
		return errors.New("Unable to parse response")
	}

	return nil
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
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
