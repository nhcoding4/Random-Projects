class Node {
    constructor(value, nextNode) {
        this.value = value
        this.nextNode = nextNode
    }
    root = false
    totalNodes = 0

    setRoot() {
        this.root = !this.root
    }

    editValue(newValue) {
        this.value = newValue
    }

    increaseCount() {
        if (this.root) {
            this.totalNodes++
        }
    }

    printNode() {
        console.log(`Value: ${this.value} | Nextnode: ${this.nextNode}`)
    }

    isRoot() {
        return this.root
    }
}

// --------------------------------------------------------------------------------------------------------------------

function main() {
    let root = new Node("LIST ROOT", null)
    root.setRoot()

    addNode(root, "Apples")
    addNode(root, "Turnips")
    addNode(root, "Cherry")

    console.log()
    printList(root)

    editNode(root, "Apples", "Oranges")
    editNode(root, "Backwards", "Forwards")

    console.log()
    printList(root)

    deleteNode(root, "Turnips")
    deleteNode(root, "Backwards")

    console.log()
    printList(root)

    console.log()
    let value = getValue(root, 1)
    console.log(value)

    console.log(dumpValues(root))
}

// --------------------------------------------------------------------------------------------------------------------

// Adds a node to the list.

function addNode(listRoot, nodeValue) {
    listRoot.increaseCount()
    let next = null

    if (listRoot.nextNode != null) {
        next = listRoot.nextNode
    }
    const newNode = new Node(nodeValue, next)
    listRoot.nextNode = newNode
    console.log(`Node holding ${nodeValue} added to list.`)
}

// --------------------------------------------------------------------------------------------------------------------

// Removes a node from the list.

function deleteNode(node, nodeValue) {
    let previousNode = null
    let currentNode = node

    while (1) {
        if (currentNode.value == nodeValue) {
            if (currentNode.listRoot) {
                console.log("Cannot delete the root node")
                return
            }
            previousNode.nextNode = currentNode.nextNode
            return
        }
        if (currentNode.nextNode == null) {
            console.log(`Node with value ${nodeValue} has not been found in the list`)
            return
        }
        previousNode = currentNode
        currentNode = currentNode.nextNode
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Returns an array of all the values held in the list.

function dumpValues(root) {
    const listValues = []
    let currentNode = root.nextNode
    while (1) {
        if (currentNode == null) {
            return listValues
        }
        listValues.push(currentNode.value)
        currentNode = currentNode.nextNode
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Searches for an edits a node in the list.

function editNode(node, nodeValue, newValue) {
    if (node.value == nodeValue) {
        node.editValue(newValue)
        console.log(`Changed ${nodeValue} to ${newValue}.`)
        return
    }
    if (node.nextNode == null) {
        console.log(`Unable to find ${nodeValue} in the linked list.`)
        return
    }
    editNode(node.nextNode, nodeValue, newValue)
}

// --------------------------------------------------------------------------------------------------------------------

// Returns value from the list from an index.

function getValue(listRoot, index) {
    if ((listRoot.totalNodes - 1 < index) || (index < 1)) {
        return "index out of range"
    }
    let currentNode = listRoot
    while (1) {
        index--
        if (index == -1) {
            console.log("Cannot return value of root node")
            return
        }
        currentNode = currentNode.nextNode
        if (index == 0) {
            return currentNode.value
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Prints the entire list.

function printList(node) {
    node.printNode()
    if (node.nextNode == null) {
        return
    }
    printList(node.nextNode)
}

// --------------------------------------------------------------------------------------------------------------------

main()
