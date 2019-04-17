package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const maxRetry = 30

func main() {
	http.HandleFunc("/game", proxy("http://game:8888/game"))
	http.HandleFunc("/title", proxy("http://title:8888/title"))

	log.Println("Listening on port 8888")
	http.ListenAndServe(":8888", nil)
}

func proxy(url string) func(http.ResponseWriter, *http.Request) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		for retry := 1; retry <= maxRetry; retry++ {
			log.Println(r.RequestURI)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				http.Error(w, "Unable to create request", 500)
				return
			}

			copyHeader(req.Header, r.Header)
			resp, err := client.Do(req)
			if err != nil {
				if retry == maxRetry {
					http.Error(w, "unable to connect to upstream", 500)
					return
				}

				log.Println("Error:", r.RequestURI, err, "Retrying in 1s")
				time.Sleep(500 * time.Millisecond)
				continue
			}
			defer resp.Body.Close() // TODO fix this

			if resp.StatusCode >= 400 {
				if retry == maxRetry {
					http.Error(w, resp.Status, resp.StatusCode)
					return
				}

				log.Println("Error:", r.RequestURI, err, "Retrying in 1s")
				time.Sleep(500 * time.Millisecond)
				continue
			}

			buf, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				http.Error(w, "Unable to read hexagons", 500)
				return
			}

			copyHeader(w.Header(), resp.Header)
			w.Write(buf)
			return
		}
	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}
