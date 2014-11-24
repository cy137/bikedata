package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Result struct {
	Stations   []Station `xml:"station"`
	LastUpdate string    `xml:"lastUpdate"`
}

func main() {
	url := "http://www.capitalbikeshare.com/data/stations/bikeStations.xml"

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	result := Result{}

	err = xml.Unmarshal(body, &result)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(result)

}
