package main

import (
	CoR "github.com/iamprahladsinghnegi/design-patterns/behavioral/chain-of-responsibility/chain-of-responsibility"
)

func main() {

	cashier := &CoR.Cashier{}

	//Set next for medical department
	medical := &CoR.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &CoR.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &CoR.Reception{}
	reception.SetNext(doctor)

	patient := &CoR.Patient{Name: "abc"}
	//Patient visiting
	reception.Execute(patient)
}
