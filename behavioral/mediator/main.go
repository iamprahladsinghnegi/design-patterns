package main

import (
	mediator "github.com/iamprahladsinghnegi/design-patterns/behavioral/mediator/mediator"
)

func main() {
	stationManager := mediator.NewStationManger()

	passengerTrain := &mediator.PassengerTrain{
		Mediator: stationManager,
	}
	freightTrain := &mediator.FreightTrain{
		Mediator: stationManager,
	}

	passengerTrain.Arrive()
	freightTrain.Arrive()
	passengerTrain.Depart()
}
