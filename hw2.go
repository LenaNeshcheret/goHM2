package main

import (
	"fmt"
)

func main() {
	// Initialize the animals
	animals := []Animal{
		{"Elephant", "Dambo", 3, 2},
		{"Lion", "Leo", 2, 1.51},
		{"Monkey", "McScruff", 0.53, 0.52},
		{"Giraffe", "Dottie", 5.01, 1.8},
		{"Tiger", "Teo", 2.57, 1.23},
		{"Panda", "Bo-bo", 1.85, 1.23},
		{"Zebra", "Ziggy", 2.3, 1.4},
		{"Kangaroo", "Kenny", 1.9, 0.8},
		{"Crocodile", "Snappy", 4.5, 80},
		{"Penguin", "Waddles", 0.7, 5},
		{"Hippo", "Bubbles", 4.2, 2.5},
		{"Gorilla", "Gigi", 1.95, 160},
		{"Parrot", "Polly", 0.3, 0.4},
		{"Ostrich", "Ozzy", 2.1, 10},
		{"Rhino", "Rex", 1.8, 2.3},
		{"Turtle", "Sheldon", 0.2, 0.25},
		{"Polar Bear", "Snowy", 1.6, 350},
		{"Koala", "Kobi", 0.6, 12},
		{"Cheetah", "Chet", 1.1, 60},
		{"Sloth", "Sid", 0.5, 8},
		{"Giraffe", "Stretch", 5.3, 2},
		{"Penguin", "Pingu", 0.65, 4},
		{"Panda", "Pandy", 1.75, 1.1},
		{"Elephant", "Jumbo", 3.5, 3.2},
	}

	// Initialize the zookeeper
	zookeeper := Zookeeper{Name: "John"}

	// Initialize the cages
	cage1 := Cage{
		Name:        "Cage 1",
		AnimalsInfo: []string{"Dambo", "Dottie", "Ziggy", "Kenny", "Rex", "Ozzy", "Stretch"},
		Animals: []Animal{
			{"Elephant", "Dambo", 3, 2},
			{"Giraffe", "Dottie", 5.01, 1.8},
			{"Zebra", "Ziggy", 2.3, 1.4},
		},
	}
	cage2 := Cage{
		Name:        "Cage 2",
		AnimalsInfo: []string{"Ziggy", "Waddles", "Pingu"},
		Animals: []Animal{
			{"Zebra", "Ziggy", 2.3, 1.4},
			{"Penguin", "Waddles", 0.7, 5},
			{"Penguin", "Pingu", 0.65, 4},
		},
	}
	cage3 := Cage{
		Name:        "Cage 3",
		AnimalsInfo: []string{"Bo-bo", "Pandy", "Jumbo"},
		Animals: []Animal{
			{"Panda", "Bo-bo", 1.85, 1.23},
			{"Panda", "Pandy", 1.75, 1.1},
		},
	}
	cages := []Cage{cage1, cage2, cage3}

	// Initialize the zoo
	zoo := Zoo{Zookeeper: zookeeper, Cages: cages, Animals: animals}

	// Check for escaped animals
	checkEscapedAnimals(zoo)
}

type Zookeeper struct {
	Name string
}

type Animal struct {
	Species string
	Name    string
	Height  float64
	Weight  float64
}

type Cage struct {
	Name        string   // Name of the cage
	AnimalsInfo []string // Names of animals that should be in the cage
	Animals     []Animal // Actual animals in the cage
}

type Zoo struct {
	Zookeeper Zookeeper
	Cages     []Cage
	Animals   []Animal
}

// Function to check for escaped animals and print their details
func checkEscapedAnimals(zoo Zoo) {
	for _, cage := range zoo.Cages {
		escaped := len(cage.AnimalsInfo) - len(cage.Animals)
		if escaped == 0 {
			fmt.Printf("All animals are in %s.\n", cage.Name)
		} else {
			fmt.Printf("Number of animals who have escaped from %s: %d\n", cage.Name, escaped)
			fmt.Println("Escaped animals:")
			for _, animalInfo := range cage.AnimalsInfo {
				found := false
				for _, animal := range cage.Animals {
					if animal.Name == animalInfo {
						found = true
						break
					}
				}
				if !found {
					escapedAnimal := getAnimalByName(animalInfo, zoo.Animals)
					fmt.Printf("Species: %s, Name: %s, Height: %.2f, Weight: %.2f\n",
						escapedAnimal.Species, escapedAnimal.Name, escapedAnimal.Height, escapedAnimal.Weight)
				}
			}
		}
	}
}

// Helper function to find an animal by name
func getAnimalByName(name string, animals []Animal) Animal {
	for _, animal := range animals {
		if animal.Name == name {
			return animal
		}
	}
	return Animal{}
}
