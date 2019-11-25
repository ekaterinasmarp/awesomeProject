package main

func main() {
	println("Old Macdonald Farm")
	cows := generateCows()
	pigs := generatePigs()
	chicks := generateChicks()

	job := OldMacdonaldJob{foodStation: FoodStationS{
		grassProvider: GrassProviderS{NameOfGrass: "Green Grass"},
		grainProvider: GrainProviderS{transportCompany: TransportCompanyS{creditCardNumber: "XXXX XXXX XXXX 1234"}}}}
	job.NormalDayOnAFarm(cows, pigs, chicks)
}

func generateChicks() []Chick {
	return []Chick{Chick{Animal{Name: "Chick 1"}}, Chick{Animal{Name: "Chick 2"}}, Chick{Animal{Name: "Chick 3"}}, Chick{Animal{Name: "Chick 4"}}, Chick{Animal{Name: "Chick 5"}}}
}

func generatePigs() []Pig {
	return []Pig{Pig{Animal{Name: "Pig 1"}}, Pig{Animal{Name: "Pig 2"}}, Pig{Animal{Name: "Pig 3"}}}
}

func generateCows() []Cow {
	return []Cow{Cow{Animal{Name: "Cow 1"}}, Cow{Animal{Name: "Cow 2"}}, Cow{Animal{Name: "Cow 3"}}, Cow{Animal{Name: "Cow 4"}}}
}
