package main

func main() {
	println("Old Macdonald Farm")
	cows := []Cow{Cow{Animal{Name: "Cow 1"}}, Cow{Animal{Name: "Cow 2"}}, Cow{Animal{Name: "Cow 3"}}, Cow{Animal{Name: "Cow 4"}}}
	pigs := []Pig{Pig{Animal{Name: "Pig 1"}}, Pig{Animal{Name: "Pig 2"}}, Pig{Animal{Name: "Pig 3"}}}
	chicks := []Chick{Chick{Animal{Name: "Chick 1"}}, Chick{Animal{Name: "Chick 2"}}, Chick{Animal{Name: "Chick 3"}}, Chick{Animal{Name: "Chick 4"}}, Chick{Animal{Name: "Chick 5"}}}

	job := OldMacdonaldJob{}
	job.NormalDayOnAFarm(cows, pigs, chicks)
}
