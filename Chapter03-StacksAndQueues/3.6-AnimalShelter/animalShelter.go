package animal_shelter

/*
3.6 Animal Shelter
An animal shelter, which holds only dogs and cats, operates on a strictly "first in, first out" basis.
People must adopt either the "oldest" (based on arrival time) of all animals at the shelter, or they can select whether they would prefer a dog or a cat (and will receive the oldest animal of that type).
They cannot select which specific animal they would like. Create the data structures to maintain this system and implement operations such as enqueue, dequeueAny, dequeueDog, and dequeueCat.
You may use the built-in Linked List data structure.
*/

import (
	"container/list"
	"errors"
	"fmt"
)

// Animal is an interface that defines common methods for all animals.
type Animal interface {
	GetID() int
	SetID(id int)
}

// BaseAnimal implements common functionality for all animals.
type BaseAnimal struct {
	ID int
}

// GetID returns the ID of the animal.
func (a *BaseAnimal) GetID() int {
	return a.ID
}

// SetID sets the ID of the animal.
func (a *BaseAnimal) SetID(id int) {
	a.ID = id
}

// Dog represents a dog in the shelter, embedding BaseAnimal.
type Dog struct {
	BaseAnimal
}

// Cat represents a cat in the shelter, embedding BaseAnimal.
type Cat struct {
	BaseAnimal
}

// AnimalShelter maintains the shelter state with queues for all animals and for each type.
type AnimalShelter struct {
	allAnimals   *list.List // Queue for all animals
	dogQueue     *list.List // Queue for dogs
	catQueue     *list.List // Queue for cats
	arrivalOrder int        // Used to assign unique ID to each animal
}

// NewAnimalShelter creates a new AnimalShelter instance.
func NewAnimalShelter() *AnimalShelter {
	return &AnimalShelter{
		allAnimals:   list.New(),
		dogQueue:     list.New(),
		catQueue:     list.New(),
		arrivalOrder: 0,
	}
}

// enqueue adds an animal to the shelter.
func (shelter *AnimalShelter) enqueue(animal Animal) {
	shelter.arrivalOrder++
	animal.SetID(shelter.arrivalOrder)
	shelter.allAnimals.PushBack(animal)
	switch animal.(type) {
	case *Dog:
		shelter.dogQueue.PushBack(animal)
	case *Cat:
		shelter.catQueue.PushBack(animal)
	default:
		fmt.Println("Unsupported animal type")
	}
}

// dequeueAny removes and returns the oldest animal from the shelter.
func (shelter *AnimalShelter) dequeueAny() (Animal, error) {
	if shelter.allAnimals.Len() == 0 {
		return nil, errors.New("no animals in the shelter")
	}
	element := shelter.allAnimals.Front()
	animal := element.Value.(Animal)
	shelter.allAnimals.Remove(element)
	switch animal.(type) {
	case *Dog:
		shelter.removeAnimalFromQueue(shelter.dogQueue, animal)
	case *Cat:
		shelter.removeAnimalFromQueue(shelter.catQueue, animal)
	}
	return animal, nil
}

// dequeueDog removes and returns the oldest dog from the shelter.
func (shelter *AnimalShelter) dequeueDog() (*Dog, error) {
	if shelter.dogQueue.Len() == 0 {
		return nil, errors.New("no dogs in the shelter")
	}
	element := shelter.dogQueue.Front()
	animal := element.Value.(*Dog)
	shelter.dogQueue.Remove(element)
	shelter.removeAnimalFromQueue(shelter.allAnimals, animal)
	return animal, nil
}

// dequeueCat removes and returns the oldest cat from the shelter.
func (shelter *AnimalShelter) dequeueCat() (*Cat, error) {
	if shelter.catQueue.Len() == 0 {
		return nil, errors.New("no cats in the shelter")
	}
	element := shelter.catQueue.Front()
	animal := element.Value.(*Cat)
	shelter.catQueue.Remove(element)
	shelter.removeAnimalFromQueue(shelter.allAnimals, animal)
	return animal, nil
}

// removeAnimalFromQueue removes a specific animal from the provided queue.
func (shelter *AnimalShelter) removeAnimalFromQueue(queue *list.List, animal Animal) {
	for element := queue.Front(); element != nil; element = element.Next() {
		if e, ok := element.Value.(Animal); ok && e.GetID() == animal.GetID() {
			queue.Remove(element)
			break
		}
	}
}
