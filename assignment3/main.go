package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/alok87/goutils/pkg/random"
)

type Weather struct {
	Status Status `json:"status"`
}

type Status struct {
	Water       int `json:"water"`
	Wind        int `json:"wind"`
	WaterStatus string
	WindStatus  string
}

var PORT = ":8000"

func main() {
	data := &Weather{}
	http.HandleFunc("/", data.display)
	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("asset"))))

	http.ListenAndServe(PORT, nil)
}

func (data *Weather) display(w http.ResponseWriter, r *http.Request) {
	openFile, err := os.Open("weather.json")

	if err != nil {
		fmt.Println("Error Open File")
		return
	}

	jsonData, err := ioutil.ReadAll(openFile)
	if err != nil {
		fmt.Println("Error Parse File")
		return
	}

	if err := json.Unmarshal(jsonData, &data); err != nil {
		fmt.Println("Error unmarshal", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	waterNow := random.RangeInt(1, 10, 1)
	data.Status.Water = waterNow[0]

	windNow := random.RangeInt(1, 20, 1)
	data.Status.Wind = windNow[0]

	marshalJson, _ := json.Marshal(data)
	err = ioutil.WriteFile("weather.json", marshalJson, 0644)

	if data.Status.Water < 5 {
		data.Status.WaterStatus = "Aman"
	} else if data.Status.Water < 9 {
		data.Status.WaterStatus = "Siaga"
	} else {
		data.Status.WaterStatus = "Bahaya"
	}

	if data.Status.Wind < 6 {
		data.Status.WindStatus = "Aman"
	} else if data.Status.Water < 16 {
		data.Status.WindStatus = "Siaga"
	} else {
		data.Status.WindStatus = "Bahaya"
	}

	tpl, err := template.ParseFiles("template.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tpl.Execute(w, data.Status)
}
