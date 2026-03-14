package main

import (
	"errors"
	"fmt"
	"log"
)

type Truck struct {
	id string
}

var (
	ErrTruckProcessingFailed = errors.New("Truck processing failed")
	ErrNotImplemented        = errors.New("Not implemented")
	ErrTruckNotFound         = errors.New("Truck not found")
	ErrTruckNotFoundwithID   = func(id string) error {
		return fmt.Errorf("Truck with ID %s not found", id)
	}
)

// Load simulates loading the truck and returns nil for success.
func (t Truck) LoadCargo() error {
	// You can add logic here, for now just return nil to indicate success.
	fmt.Printf("Loading Cargo ->%s\n", t.id)
	if t.id == "Truck2" {
		return ErrTruckNotFoundwithID(t.id)
	}
	return nil

}

func processTruck(t Truck) error {
	fmt.Printf("Processing Truck ->%s\n", t.id)
	if err := t.LoadCargo(); err != nil {
		return fmt.Errorf("Error loading cargo for truck %s: %w", t.id, err)
	}
	return nil
}

func main() {
	trucks := []Truck{
		{id: "Truck1"},
		{id: "Truck2"},
		{id: "Truck3"},
	}
	for _, truck := range trucks {
		fmt.Printf("Truck has arrived :->%s\n", truck.id)
		err := processTruck(truck)
		switch err {
		case ErrTruckProcessingFailed:
			log.Printf("Failed to process truck %s: %v\n", truck.id, err)
		case ErrTruckNotFound:
			log.Printf("Truck %s not found\n", truck.id)
		case ErrNotImplemented:
			log.Printf("Processing of truck %s is not implemented yet\n", truck.id)
		default:
			if err != nil {
				log.Printf("Error processing truck %s: %v\n", truck.id, err)
			}
		}
	}
}
