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
	"strings"
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

	if err := http.ListenAndServe("192.168.1.84:3000", mux); err != nil {
		print("Error")
		log.Fatal(err)
	}

}
func AddbulkLoc(w http.ResponseWriter, request *http.Request) {
	var contentArray []interface{}
	var loc Location
	if request.Body == nil {
		//utils.ErrorWithJSON(w, "Please send a request body", http.StatusNotFound)
		fmt.Println("ERRER 1")
		return
	}
	data, _ := ioutil.ReadAll(request.Body)
	asString := string(data)
	locations := strings.Split(asString,"#");
	for _, location := range locations {
		err := json.Unmarshal([]byte(location), &loc)
		if err != nil {
			panic(err)
		}
		contentArray =  append(contentArray, &loc)
	}
	mgoSession, _ := mgo.Dial(host)
	defer mgoSession.Close()
	collection := mgoSession.DB(db).C(collection)
	bulk := collection.Bulk()
	bulk.Insert(contentArray...)
	_, err := bulk.Run()
	if err != nil {
		fmt.Println("ERROR! Bulk Insert. Datos a Insertar:",contentArray," tamaño: ", len(contentArray))
		panic(err)
	}
}


