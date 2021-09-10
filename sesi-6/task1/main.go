package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type structLang struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}

type Relation struct {
	InfluenceBy []string `json:"influenced-by"`
	Influences  []string `json:"influences"`
}

type RequestParam struct {
	Param string `json:"param"`
}
type ResponseParam struct {
	Status int
	Desc   bool
}

type ResponseLang struct {
	Status int
	Desc   string
	Detail structLang
}

const port = ":8080"

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/lang", langHandler).Methods("POST")
	http.Handle("/", router)
	fmt.Println("running at http://localhost", port)
	http.ListenAndServe(port, nil)
}

func langHandler(w http.ResponseWriter, r *http.Request) {
	var response ResponseLang
	var request structLang
	defer r.Body.Close()
	c, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(c, &request)

	if err != nil {
		panic(err)
	}
	response.Desc = "success"
	response.Status = 200
	response.Detail = request

	json.NewEncoder(w).Encode(response)
}
