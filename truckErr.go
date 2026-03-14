package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}
type ElectricTruck struct {
	id           string
	cargo        int
	batteryLevel float64
}

var (
	ErrNotImplemented = errors.New("Not implemented")
	ErrTruckNotFound  = errors.New("Truck not found")
)

func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo -= 1
	return nil
}
func (t *ElectricTruck) LoadCargo() error {
	t.cargo += 2
	t.batteryLevel -= 1
	return nil
}

func (t *ElectricTruck) UnloadCargo() error {
	t.cargo -= 1
	t.batteryLevel -= 1
	return nil
}

func processTruck(t Truck) error {
	fmt.Printf("Processing truck :%+v\n", t)
	err := t.LoadCargo()
	if err != nil {
		return fmt.Errorf("Error loading cargo : %w", err)
	}

	err = t.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo : %w", err)
	}
	return nil
}

func main() {
	nt := NormalTruck{id: "1", cargo: 10}
	et := ElectricTruck{id: "2", cargo: 20, batteryLevel: 100}
	err := processTruck(&nt)
	if err != nil {
		log.Printf("Error processing truck: %v", err)
	}
	err = processTruck(&et)
	if err != nil {
		log.Printf("Error processing truck: %v", err)
	}
	log.Printf("processing complete \n")
	log.Printf("Normal Truck: %+v\n", nt)
	log.Printf("Electric Truck: %+v\n", et)

	//map of strings
	person := make(map[string]any, 0)
	person["name"] = "John"
	person["age"] = 30

	age, exists := person["age"]
	if exists {
		log.Printf("Age: %v\n", age)
	} else {
		log.Printf("Age not found\n")
	}

	width, exists := person["width"]
	if !exists {
		log.Fatalf("Width not found\n")
	}
	log.Printf("Width: %v\n", width)
	log.Printf("Person: %+v\n", person)
}
