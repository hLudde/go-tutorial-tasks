package main

import (
	"fmt"
	"net/http"
	"os"
)

type car struct {
	posX          float64
	posY          float64
	steeringWheel int16
	topSpeedKmh   float64
}

func (c car) hardTurn() int16 {
	return -c.steeringWheel
}

var aCar = createCar(100.4, 250.3, 225.0, 5430)

func createCar(posX, posY, topSpeedKmh float64, steeringWheel int16) car {
	return car{posX, posY, steeringWheel, topSpeedKmh}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<p>values of aCar with type car: posX=%f posY=%f topSpeedKmh=%f steeringWheel=%d</p> <a href=\"/hardTurn\">Do a hard turn</a> <a href=\"/stop\">Stop server</a>", aCar.posX, aCar.posY, aCar.topSpeedKmh, aCar.steeringWheel)
}

func turnHandler(w http.ResponseWriter, r *http.Request) {
	aCar.steeringWheel = aCar.hardTurn()
	indexHandler(w, r)
}

func stopServer(w http.ResponseWriter, r *http.Request) {
	os.Exit(0)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/stop", stopServer)
	http.HandleFunc("/hardTurn", turnHandler)
	http.ListenAndServe(":80", nil)
}
