package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Request struct {
	Param string `json: "param"`
}

type ResponseParam struct {
	Status int
	Desc   bool
}

func main() {
	const port = ":8080"
	router := mux.NewRouter()

	router.HandleFunc("/is-palindrom", isPalindrom).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Starting at port", port)

	http.ListenAndServe(port, nil)
}

func isPalindrom(w http.ResponseWriter, r *http.Request) {
	var req Request
	var res ResponseParam

	c, errRead := ioutil.ReadAll(r.Body)
	if errRead != nil {
		panic(errRead)
	}

	defer r.Body.Close()

	errUnmarshal := json.Unmarshal([]byte(c), &req)

	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	palindrom := palindrom(req.Param)

	res = ResponseParam{
		Status: 200,
		Desc:   palindrom,
	}

	json.NewEncoder(w).Encode(res)
}

func palindrom(s string) bool {
	s = strings.ToLower(s)
	arr := strings.Split(s, "")
	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}
	return true
}
