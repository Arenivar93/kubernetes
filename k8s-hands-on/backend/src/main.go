package main

import (
	"net/http"
	"encoding/json"
	"os"
	"time"
)

type HandsOn struct {
	Time 		time.Time	`json:"time"`
	Hostname 	string 		`json:"hostname"`
}

func ServerHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}

	resp := HandsOn{
		Time:		time.Now(),
		Hostname:	os.Getenv("HOSTNAME"),
	}
	jsonResp, err := json.Marshal(&resp)
	if err != nil {
		w.Write([]byte("Error"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Acces-Control-Allow-Origin", "*")
	//resp := fmt.Sprintf("La hora es %v y hostname es %v", time.Now(), os.Getenv("HOSTNAME"))

    w.Write(jsonResp)
}

func main() {
    http.HandleFunc("/", ServerHTTP)
    http.ListenAndServe(":9090", nil)
}
