package main

func main() {
	stationManager := NewStationManager()

	passengerTrain := &PassengerTrain{
		stationManager,
	}

	freightTrain := &FreightTrain{
		stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
