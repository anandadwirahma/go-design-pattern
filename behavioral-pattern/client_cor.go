package main

import "go-design-pattern/behavioral-pattern/cor"

func CallChainOfResponsibility() {

	cashier := &cor.Cashier{}

	//Set next for medical department
	medical := &cor.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &cor.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := &cor.Reception{}
	reception.SetNext(doctor)

	patient := &cor.Patient{Name: "Ananda"}
	// patient visiting
	reception.Execute(patient)
}
