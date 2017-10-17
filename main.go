package main

import (
	"encoding/json"
	"fmt"
	"os"
	//	"github.com/nats-io/go-nats"
)

// Encapsulating a range with a unit of measurement
type Range struct {
	Min   float32 `json:"Min"`
	Max   float32 `json:"Max"`
	Units string  `json:"Units"`
}

// Combines data with
type Datum struct {
	Data  float32 `json:"Data"`
	Units string  `json:"Units"`
}

type Ranger interface {
	InRange(r Range, d Datum) (b bool, e error)
}

// Cycle contains timing information of a crop
type Cycle struct {
	CropName string `json:"CropName"`
	Times    Range  `json:"Times"`
	Spacing  Range  `json:"Spacing"`
	Length   Range  `json:"Length"`
}
type GrowthType int

// Constants defining the separate stages of crop growth
const (
	SpawnRun GrowthType = iota
	PrimordiaFormation
	FruitingBodyDevelopment
)

// GrowthStage holds environmental variables for a stage of a plants growth
type Parameters struct {
	CropName         string     `json:"CropName"`
	CropStage        GrowthType `json:"CropStage"`
	Temperature      Range      `json:"Temperature"` // will track Centigrade and convert if needed
	RelativeHumidity Range      `json:"RelativeHumidity"`
	Duration         Range      `json:"DurationDays"`  //typically in days
	CarbonDioxide    Range      `json:"CarbonDioxide"` //CO2 ppm
	AirExchange      Range      `json:"AirExchage"`
	LightRequirement Range      `json:"Light"` // lux
}

func NewRange(a float32, b float32, s string) *Range {
	r := Range{a, b, s}
	return &r
}

func SaveParams(p *Parameters, s string) {
	f, err := os.Create(s)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(128)
	} else {
		m, err := json.MarshalIndent(p, "", "  ")
		if err != nil {
			fmt.Println("Error: ", err)
		}
		w, err := f.Write(m)
		if w < 0 {
			fmt.Println("Error: ", err)
		}
	}
}

func main() {
	// if a range is exceeded send an alert using nats tagged with the data
	// and the
	//	nc, _ := nats.Connect(nats.DefaultURL)
	//	c, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer c.Close()
	p := &Parameters{}
	SaveParams(p, "Pleurotus_ostreatus.json")
	fmt.Println("hello hypha")
}
