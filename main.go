package main

import (
	"fmt"
)

// Encapsulating a range with a unit of measurement
type Range struct {
	Min  int
	Max  int
	Unit string
}

type GrowthType int

/* Constants defining the separate stages of crop growth
 */
const (
	SpawnRun GrowthType = iota
	PrimordiaFormation
	FruitingBodyDevelopment
	CroppingCycle
)

/* GrowthStage holds environmental variables for a stage of a plants growth
 */
type GrowthStage struct {
	CropName         string
	Stage            GrowthType
	Temperature      Range // will track Centigrade and convert if needed
	RelativeHumidity Range
	Duration         Range //typically in days
	CarbonDioxide    Range //CO2 ppm
	AirExchange      Range
	LightRequirement Range // lux
	CroppingCycle    Range
}

func main() {
	fmt.Println("hello hypha")
}
