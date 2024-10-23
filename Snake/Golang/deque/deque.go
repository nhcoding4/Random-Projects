package Deque

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// ---------------------------------------------------------------------------------------------------------------------

// Deque data structure. This example is specialised for the snake game.

// ---------------------------------------------------------------------------------------------------------------------

type Deque struct {
	length int
	Elems  []rl.Vector2
}

// ---------------------------------------------------------------------------------------------------------------------

// Adds data to the back of the deque.

func (d *Deque) Append(elems ...rl.Vector2) {
	d.Elems = append(d.Elems, elems...)
	d.updateLength()
}

// ---------------------------------------------------------------------------------------------------------------------

// Clears all the elements held in the deque.

func (d *Deque) Clear() {
	var newDeque []rl.Vector2
	d.Elems = newDeque
	d.updateLength()
}

// --------------------------------------------------------------------------------------------------------------------

// Get a value of X index

func (d *Deque) GetValue(index int) (rl.Vector2, error) {
	if index > d.length-1 {
		return rl.Vector2{}, fmt.Errorf("deque GetValue: %v is out of range %v", index, d.length-1)
	}
	return d.Elems[index], nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Adds data to the start of the queue.

func (d *Deque) Prepend(elems ...rl.Vector2) {
	var newDeque []rl.Vector2
	newDeque = append(newDeque, elems...)
	newDeque = append(newDeque, d.Elems...)

	d.Elems = newDeque
	d.updateLength()
}

// ---------------------------------------------------------------------------------------------------------------------

// Remove and return items from the back of the deque.

func (d *Deque) RemoveBack(items int) ([]rl.Vector2, error) {
	if items > d.length-1 {
		return nil, fmt.Errorf("deque GetBack: %v is out of range %v", items, d.length-1)
	}
	if items == 0 {
		return nil, fmt.Errorf("deque GetBack: %v is empty", items)
	}

	sliceIndex := d.length - items
	slicedItems := d.Elems[sliceIndex:]
	d.Elems = d.Elems[:sliceIndex]
	d.updateLength()

	return slicedItems, nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Remove and return items from the front of the deque.

func (d *Deque) RemoveFront(items int) ([]rl.Vector2, error) {
	if items > d.length-1 {
		return nil, fmt.Errorf("deque GetFront: %v is out of range %v", items, d.length-1)
	}
	slicedItems := d.Elems[:items]
	d.Elems = d.Elems[items:]
	d.updateLength()

	return slicedItems, nil
}

// ---------------------------------------------------------------------------------------------------------------------

// Returns the total amount of elements held in the deque.

func (d *Deque) Size() int {
	return d.length
}

// --------------------------------------------------------------------------------------------------------------------

// Prints the contents of the deque.

func (d *Deque) String() string {
	return fmt.Sprintf("%v | %v total elements", d.Elems, d.Size())
}

// ---------------------------------------------------------------------------------------------------------------------

// Updates the length variable.

func (d *Deque) updateLength() {
	d.length = len(d.Elems)
}

// ---------------------------------------------------------------------------------------------------------------------
