package animal_shelter

import (
	"fmt"
	"testing"
)

func TestAnimalShelter(t *testing.T) {
	shelter := NewAnimalShelter()

	// Step 1: Enqueue dogs and cats
	fmt.Println("Step 1: Enqueue dogs and cats")
	shelter.enqueue(&Dog{})
	shelter.enqueue(&Cat{})
	shelter.enqueue(&Dog{})

	// Step 2: Dequeue Any (should return the oldest animal, which is the first Dog)
	fmt.Println("Step 2: Dequeue Any")
	animal, err := shelter.dequeueAny()
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}
	if _, ok := animal.(*Dog); !ok {
		t.Errorf("Expected Dog, got %T", animal)
	}
	fmt.Printf("Dequeued Any: %+v\n", animal)

	// Step 3: Dequeue Dog (should return the second Dog, the one added last)
	fmt.Println("Step 3: Dequeue Dog")
	dog, err := shelter.dequeueDog()
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}
	if dog == nil {
		t.Errorf("Expected Dog, got nil")
	}
	fmt.Printf("Dequeued Dog: %+v\n", dog)

	// Step 4: Dequeue Cat (should return the only Cat)
	fmt.Println("Step 4: Dequeue Cat")
	cat, err := shelter.dequeueCat()
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}
	if cat == nil {
		t.Errorf("Expected Cat, got nil")
	}
	fmt.Printf("Dequeued Cat: %+v\n", cat)

	// Step 5: Dequeue Any (should return an empty error because no animals left)
	fmt.Println("Step 5: Dequeue Any")
	_, err = shelter.dequeueAny()
	if err == nil {
		t.Error("Expected error, got nil")
	}
	fmt.Printf("Dequeued Any Error: %v\n", err)
}
