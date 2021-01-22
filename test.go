    package main

import (
	"log"
	"net/http"
   "encoding/json"
	"github.com/gorilla/mux"
	"os/exec"
	"io/ioutil"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-movie", getMovie)
	log.Fatal(http.ListenAndServe(":8080", router))
}

type JSONOutput struct {
	Movie         string `json:"movie"`
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var local JSONOutput
	json.Unmarshal(reqBody, &local)
	cmd := exec.Command("ruby", "obmdb.rb", local.Movie)
	output, _ := cmd.Output()
    json.NewEncoder(w).Encode(string(output))
	}


	
	

