package main

import (
	"fmt"
	"log"
	"net/http"
   "encoding/json"
	"github.com/gorilla/mux"
	"os/exec"
	"io/ioutil"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to this API")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/get-movie", getMovie)
	log.Fatal(http.ListenAndServe(":8080", router))
}

type JSONOutput struct {
	Movie         string `json:"movie"`
	Error         string `json:"error"`
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	reqBody,err:= ioutil.ReadAll(r.Body)
	var local JSONOutput
	json.Unmarshal(reqBody, &local)
	cmd := exec.Command("ruby", "obmdb.rb", local.Movie)
	output,err := cmd.Output()
	if ((err != nil) || (local.Movie == "")) {
		fmt.Fprintf(w, "Kindly input a json movie title")
	}
	json.NewEncoder(w).Encode(string(output))
	
}
