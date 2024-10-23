from __future__ import annotations

from typing import List, Type


# ----------------------------------------------------------------------------------------------------------------------

class Node:

    def __init__(self, value: any, is_root: bool = False, next_node: Node = None) -> None:
        self.value: any = value
        self.is_root: bool = is_root
        self.next_node: Node = next_node

    # -----------------------------------------------------

    def get_value(self) -> any:
        """ Returns the value held in the node. """

        if not self.is_root:
            return self.value

    # -----------------------------------------------------

    def edit_value(self, search_value: any, new_value: any) -> bool:
        """ Checks the value held in the node against a passed value. If they match then it replaces the held value
         with the new value. """

        if self.value == search_value:
            self.value = new_value
            return True

        return False


# ----------------------------------------------------------------------------------------------------------------------
def main() -> None:
    root: Node = Node(value=None, is_root=True)

    some_values: List[int] = [10, 20, 30, 40, 50, 60]
    add_nodes(root, some_values)

    edit_node(root, 30, 35)
    edit_node(root, 101, 102)
    print_list(root)
    print()

    delete_node(root, 100)
    delete_node(root, 20)
    print()
    print_list(root)
    print()

    print(f"Value at index 10: {get_value_index(root, 10)}")
    print(f"Value at index 2: {get_value_index(root, 2)}")


# ----------------------------------------------------------------------------------------------------------------------


def add_nodes(root_node: Node, values: List[any]) -> None:
    """ Takes a list of values and adds a node to the list for every value. """

    for value in values:
        new_node: Node = Node(value=value)

        if root_node.next_node is not None:
            new_node.next_node = root_node.next_node

        root_node.next_node = new_node


# ----------------------------------------------------------------------------------------------------------------------

def delete_node(root_node: Node, to_delete: any) -> None:
    """ Searches for and removes a node from the list (if found). """

    current_node: Node = root_node
    previous_node: Node = root_node

    while True:
        if current_node.value == to_delete and not current_node.is_root:
            if current_node.next_node is None:
                del current_node
            else:
                previous_node.next_node = current_node.next_node
                del current_node
            print(f"Node with value {to_delete} has been removed from the list.")
            return

        if current_node.next_node is None:
            print(f"No node with value {to_delete} has been found.")
            return

        previous_node = current_node
        current_node = current_node.next_node


# ----------------------------------------------------------------------------------------------------------------------

def edit_node(root_node: Node, to_edit: any, new_value: any) -> None:
    """ Searches for a node holding X value and changes it to a new value (if found). """

    current_node: Node = root_node

    while True:
        edited: bool = current_node.edit_value(to_edit, new_value)

        if edited:
            print(f"Node containing: {to_edit} has been changed to {new_value}.")
            return

        if current_node.next_node is None:
            print(f"No node with: {to_edit} has been found. No data has been changed.")
            return

        current_node = current_node.next_node


# ----------------------------------------------------------------------------------------------------------------------


def get_value_index(root_node: Node, index: int) -> any:
    """ Attempts to get a value at a certain index of the node and returns the value to the user. """

    node_values: List[any] = list_values(root_node)
    try:
        return node_values[index]
    except IndexError:
        print(f"Error, index out of range. Please enter a value between 0 and {len(node_values) - 1}")
        return None


# ----------------------------------------------------------------------------------------------------------------------


def list_values(root_node: Node) -> List[any]:
    """ Creates a list of values which contains the value of each node. """

    values: List[any] = []

    if root_node is None:
        print("No nodes in the list.")
        return values

    current_node: Node = root_node

    while True:

        if not current_node.is_root:
            values.append(current_node.get_value())

        if current_node.next_node is None:
            return values

        current_node = current_node.next_node


# ----------------------------------------------------------------------------------------------------------------------


def print_list(root_node: Node) -> None:
    """ Prints all the values stored in the list. """

    values: List[any] = list_values(root_node)
    if len(values) == 0:
        print("No values are stored in the list. Add some using the add nodes function before printing")
        return

    for i, value in enumerate(values):
        print(f"Node: {i + 1}: {value}")


# ----------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()
