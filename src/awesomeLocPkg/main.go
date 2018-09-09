package main

import (
	"fmt"
	"log"
	"net/http"


	"goji.io"
	"goji.io/pat"
	"encoding/json"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"bytes"
)


func AddLoc(w http.ResponseWriter, request *http.Request) {
	fmt.Println("recibi un llamado\n")

	var loc Location
	if request.Body == nil {
		//utils.ErrorWithJSON(w, "Please send a request body", http.StatusNotFound)
		fmt.Println("ERRER 1")
		return
	}
	err := json.NewDecoder(request.Body).Decode(&loc)
	if err != nil {
		fmt.Println("ERRER 2")
		//utils.ErrorWithJSON(w, err.Error(), http.StatusNotFound)
		return
	}

	fmt.Println(loc)
/*

	buf, bodyErr := ioutil.ReadAll(request.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}

	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	request.Body = rdr2

	*/
	print("Webservice iniciado")
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()

	collection := mgoSession.DB(db).C(collection)

	//query_bson := bson.M{"userID":"1","lat":-30, "lon":-31.55, "timestamp":123456789, "accuracy":10, "altitude":100, "speed":50}
	err = collection.Insert(loc)
	if err != nil {
		panic(err)
	}

}


func main() {
	print("Webservice iniciado")

	mux := goji.NewMux()
	//routes
	mux.HandleFunc(pat.Post("/addloc"), AddLoc)
	mux.HandleFunc(pat.Post("/addbulkloc"), AddbulkLoc)

	if err := http.ListenAndServe("192.168.1.31:3000", mux); err != nil {
		print("Error")
		log.Fatal(err)
	}
/*
	c := c.New(cors.Options{
		AllowedOrigins:[]string{utils.Config.Server.Frontend},
		AllowedHeaders:[]string{"X-Requested-With", "Content-Type", "Authorization"},
		AllowedMethods:[]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowCredentials:true,
	})

	http.ListenAndServe(":8080",  c.Handler(mux))
*/
}
func AddbulkLoc(w http.ResponseWriter, request *http.Request) {
	fmt.Println("recibi un Bulk llamado\n")
/*
	var loc Location
	if request.Body == nil {
		fmt.Println("ERROR 1")
		return
	}

	fmt.Println(loc)
	*/
	buf, bodyErr := ioutil.ReadAll(request.Body)
	if bodyErr != nil {
		log.Print("bodyErr ", bodyErr.Error())
		http.Error(w, bodyErr.Error(), http.StatusInternalServerError)
		return
	}
	rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
	rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))
	log.Printf("BODY: %q", rdr1)
	request.Body = rdr2

}


