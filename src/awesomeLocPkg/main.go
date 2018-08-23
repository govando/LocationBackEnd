package main

import (
"fmt"
"log"
"net/http"

"github.com/gorilla/mux"
)


func AddLoc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")

	/*
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	*/

}


func main() {
	print("sadsa")
	r := mux.NewRouter()

	r.HandleFunc("/addloc", AddLoc).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}


