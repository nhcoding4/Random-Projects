package main

import (
	"fmt"
)

// --------------------------------------------------------------------------------------------------------------------

func main() {

	// Demo

	// Create a linked list.
	linked_list := Linked_List()

	// Add some values
	for i := 0; i < 10; i++ {
		linked_list.Add_Value(i + 1)
	}
	fmt.Println("Adding some values to the list:")
	linked_list.Print_List()

	// Editing the values in the list
	for i := 0; i < linked_list.length; i++ {
		linked_list.Edit_Value(i, (i+1)*10)
	}
	fmt.Println("\nList post edit:")
	linked_list.Print_List()

	// Remove and add 5 values from the list to a slice.
	five_values := []int{}
	for i := 0; i < 5; i++ {
		five_values = append(five_values, linked_list.Get_First_Value())
	}
	fmt.Println("\nFive Values from the start of the list.")
	for i := 0; i < len(five_values); i++ {
		fmt.Println(five_values[i])
	}
	fmt.Println("\nList post value removal:")
	linked_list.Print_List()

	// Delete a value at X
	linked_list.Delete(2)
	fmt.Println("\nList post delete at index 2")
	linked_list.Print_List()

	// Get a (copy of a) value from X index.
	value, err := linked_list.Get_Index_Value(0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("\nValue at index 0:", value)
	}
	linked_list.Print_List()

}

// --------------------------------------------------------------------------------------------------------------------

func Linked_List() *Node {
	// Returns root node.
	return &Node{value: 0, next: nil, root: true, length: 0}
}

// --------------------------------------------------------------------------------------------------------------------

type Node struct {
	value  int
	next   *Node
	root   bool
	length int
}

// --------------------------------------------------------

func (n *Node) Add_Value(value int) {
	// Create the new node.
	new_node := &Node{value: value, next: nil, root: false}

	// Add the node to the end of the list.
	last_node := n
	for {
		if last_node.next == nil {
			last_node.next = new_node
			break
		}
		last_node = last_node.next
	}
	n.length++
}

// --------------------------------------------------------

func (n *Node) Edit_Value(index, new int) {
	// Find and edits a node at X index of the linked list.

	// Attempt to get the value and index X.
	found_index, err := n.Get_Index(index)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Edit the value at index X if there are no errors.
	found_index.value = new
}

// --------------------------------------------------------

func (n *Node) Get_First_Value() int {
	// Takes and removes the first in value from the linked list.

	value := n.next.value
	n.next = n.next.next
	n.length--
	return value
}

// --------------------------------------------------------

func (n *Node) Get_Index(index int) (*Node, error) {
	// Add 1 to the index to skip the root value
	index++

	// Check for invalid inputs.
	if index > n.length || index <= 0 {
		return nil, fmt.Errorf("error: index out of range - min 0, max %v", n.length-1)
	}

	// Go to the requested index and change the value.
	current := n
	for index != 0 {
		current = current.next
		index--
	}
	return current, nil
}

// --------------------------------------------------------

func (n *Node) Get_Index_Value(index int) (int, error) {
	// Returns the value at index X of the node.

	// Attempt to find and get the node at X index.
	found_node, err := n.Get_Index(index)
	if err != nil {
		return 0, err
	}

	return found_node.value, nil
}

// --------------------------------------------------------

func (n *Node) Print_List() {
	// Prints all of the nodes contained in the list. Start -> finish.

	// Track the current node.
	i := 0
	current := n

	for {
		// Ignore the root value.
		if !current.root {
			fmt.Println(i, current.value)
		}
		// Stop the process if there are no more nodes.
		if current.next == nil {
			return
		}
		// Move onto the next node.
		current = current.next
		i++
	}
}

// --------------------------------------------------------

// Delete value

func (n *Node) Delete(index int) error {
	// Attempts to delete a node at index X

	index++
	// Check for invalid inputs.
	if index > n.length || index <= 0 {
		return fmt.Errorf("error: index out of range - min 0, max %v", n.length-1)
	}

	// Go to the requested index and change the value.
	var previous *Node
	current := n
	for index != 0 {
		previous = current
		current = current.next
		index--
	}

	// Remove the node from the list.
	previous.next = current.next
	n.length--

	return nil
}

// --------------------------------------------------------
