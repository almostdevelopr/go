package main

import "fmt"

// entry point
func main() {
	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(10),
		SaturatedFattyAcids: SaturatedFattyAcids(2),
		Sodium:              SodiumMilligram(500),
		Fruits:              FruitsPercent(60),
		Fibre:               FibreGram(4),
		Protein:             ProteinGram(2),
	}, Food)

	fmt.Printf("Nutrictional Score: %d\n", ns.Value)
	fmt.Printf("Nutrictional Score Grade: %s\n", ns.GetNutriScore())
}
