package main

type GrassProvider interface {
	getSomeGrass()
}

type GrassProviderS struct {
	NameOfGrass string
}

type GrainProvider interface {
	getSomeGrains()
}

type TransportCompany interface {
	orderDelivery()
}

type TransportCompanyS struct {
	creditCardNumber string
}

func (transportCompany TransportCompanyS) orderDelivery() {
	println("Withdraw from " + transportCompany.creditCardNumber)
}

type GrainProviderS struct {
	transportCompany TransportCompany
}

func (provider GrassProviderS) getSomeGrass() {
	println("Grass Provider ships some grass")
}

func (provider GrainProviderS) getSomeGrains() {
	provider.transportCompany.orderDelivery()
	println("Grain Provider ships some grains")
}

type FoodStation interface {
	feed(animal Animal)
}

type FoodStationS struct {
	grassProvider GrassProvider
	grainProvider GrainProvider
}

func (fs FoodStationS) feed(animal Animal) {
	fs.grassProvider.getSomeGrass()
	fs.grainProvider.getSomeGrains()
	println("Feeding " + animal.Name)
}

type OldMacdonaldJob struct {
	foodStation FoodStation
}

func (oldMacdonaldJob OldMacdonaldJob) NormalDayOnAFarm(cows []Cow, pigs []Pig, chicks []Chick) {
	for _, cow := range cows {
		oldMacdonaldJob.foodStation.feed(cow.Animal)
	}
	for _, pig := range pigs {
		oldMacdonaldJob.foodStation.feed(pig.Animal)
	}
	for _, chick := range chicks {
		oldMacdonaldJob.foodStation.feed(chick.Animal)
	}
}
