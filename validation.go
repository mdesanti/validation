package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	HTTP_PORT = "8008" //591, 8000, 8008, 8080
	CONN_HOST = "localhost"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = HTTP_PORT
	}

	log.WithFields(log.Fields{
		"host": CONN_HOST,
		"port": port,
	}).Info("SYNC_GLOBAL listening for SYNC_LOCAL")

	fmt.Println("[HTTP] VALIDATION listening for SYNC_GLOBAL on " + CONN_HOST + ":" + port)

	router := mux.NewRouter()
	router.HandleFunc("/validate/tag", tag).Methods("POST")
	router.HandleFunc("/validate/card", card).Methods("POST")
	router.HandleFunc("/validate/contingency", contingency).Methods("POST")
	http.ListenAndServe(CONN_HOST+":"+port, router)
}

func tag(w http.ResponseWriter, r *http.Request) {

	fmt.Println("tag()")

	// decrypt
	// check token

	validateTag() // + get vehicle id + vehicle GPS boolean

	validateStation() // + get station Lat/Lon

	getVehicleLocation() // if vehicle has GPS

	validateVehicleAtStation() // if vehicle has GPS

}

func card(w http.ResponseWriter, r *http.Request) {

	fmt.Println("card()")

	// decrypt
	// check token

	validateCard() // + get vehicle id + vehicle GPS boolean

	validateStation() // + get station Lat/Lon

	getVehicleLocation() // if vehicle has GPS

	validateVehicleAtStation() // if vehicle has GPS

}

func contingency(w http.ResponseWriter, r *http.Request) {

	fmt.Println("contingency()")

}

func validateTag() {

	fmt.Println("validateTag()")

}

func validateCard() {

	fmt.Println("validateCard()")

}

func validateStation() {

	fmt.Println("validateStation()")

}

func getVehicleLocation() {

	fmt.Println("getVehicleLocation()")

}

func validateVehicleAtStation() {

	fmt.Println("validateVehicleAtStation()")

}
