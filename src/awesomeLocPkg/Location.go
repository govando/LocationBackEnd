package main

import "gopkg.in/mgo.v2/bson"

type Location struct {
	ID			bson.ObjectId `bson:"_id,omitempty"`
	UserID		string	`json:"userID"`
	Lat			float64	`json:"lat"`
	Lon 		float64	`json:"lon"`
	Timestamp	float64	`json:"timestamp"`
	Accuracy 	float32	`json:"accuracy"`
	Altitude	float64	`json:"altitude"`
	Speed		float32	`json:"speed"`
}



