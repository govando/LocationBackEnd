package main

type location struct {
	_id			int		`json:"id"`
	userID		string	`json:"userID"`
	lat			float64	`json:"lat"`
	lon 		float64	`json:"lon"`
	timestamp	float32	`json:"timestamp"`
	accuracy 	float32	`json:"accuracy"`
	altitude	float64	`json:"altitude"`
	speed		float32	`json:"speed"`
}



