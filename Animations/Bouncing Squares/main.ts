
// --------------------------------------------------------------------------------------------------------------------

class Box {
    colour: string
    context: CanvasRenderingContext2D | null
    canvasWidth: number
    canvasHeight: number
    movementX: number
    movementY: number
    height: number
    width: number
    x: number
    y: number

    constructor(
        x: number,
        y: number,
        height: number,
        width: number,
        canvasWidth: number,
        canvasHeight: number,
        context: CanvasRenderingContext2D | null,
        colour: string
    ) {
        this.x = x
        this.y = y
        this.height = height
        this.width = width
        this.canvasHeight = canvasHeight
        this.canvasWidth = canvasWidth
        this.context = context
        this.colour = colour
        this.movementX = this.setSpeed()
        this.movementY = this.setSpeed()
        this.randomizeDirection()
    }

    // -----------------------------------------------------

    // Draw the asset on the canvas.

    draw() {
        if (this.context != null) {
            this.context.fillStyle = this.colour
        }
        this.context?.fillRect(this.x, this.y, this.width, this.height)
    }

    // -----------------------------------------------------

    // Set the starting direction for the object.

    randomizeDirection() {
        const randomize = ((direction: number): number => {
            const choice = Math.floor(Math.random() * 2)
            return (choice == 0) ? direction *= - 1 : direction
        })

        this.movementX = randomize(this.movementX)
        this.movementY = randomize(this.movementY)
    }

    // -----------------------------------------------------

    // Set the speed of the object.

    setSpeed() {
        let i: number
        i = Math.floor(Math.random() * 2 + 1)

        return i
    }

    // -----------------------------------------------------

    // Updates the position to be drawn on the next frame and collision detection.

    update() {
        const objectXEdge = this.x + this.width
        const objectYEdge = this.y + this.height

        if (this.x <= 0 || objectXEdge >= this.canvasWidth) {
            if (this.movementX > 0) {
                this.movementX = this.setSpeed()
                this.movementX *= -1
            } else {
                this.movementX = this.setSpeed()
            }

            if (this.x < 0) {
                this.x = 0 + this.width
            }
            if (this.x > this.canvasWidth) {
                this.x = this.canvasWidth - this.width
            }
        }

        if (this.y <= 0 || objectYEdge >= this.canvasHeight) {
            if (this.movementY > 0) {
                this.movementY = this.setSpeed()
                this.movementY *= -1
            } else {
                this.movementY = this.setSpeed()
            }

            if (this.y < 0) {
                this.y = 0 + this.height
            }
            if (this.y > this.canvasHeight) {
                this.y = this.canvasHeight - this.height
            }
        }

        this.x += this.movementX
        this.y += this.movementY
    }

    // -----------------------------------------------------

    reverse() {
        this.movementY *= -1
        this.movementX *= -1

        this.x += this.movementX
        this.y += this.movementY
    }

}

// --------------------------------------------------------------------------------------------------------------------

class Game {
    canvas: HTMLCanvasElement
    context: CanvasRenderingContext2D | null
    boxes: Box[]
    totalBoxes: number
    height: number
    width: number

    constructor() {
        this.boxes = []
        this.height = 1000
        this.width = 500
        this.totalBoxes = 20

        this.createCanvas()
        this.createBox()
    }

    // -----------------------------------------------------

    // Animate assets.

    animate() {
        this.context?.clearRect(0, 0, this.width, this.height)

        for (let object of this.boxes) {
            object.draw()
            object.update()
            this.checkBoxCollision()
        }

        requestAnimationFrame(() => {
            this.animate()
        })
    }

    // -----------------------------------------------------

    // Deals with collision with other boxes.

    checkBoxCollision() {
        const length = this.boxes.length

        for (let i = 0; i < length; i++) {
            let currentBox = this.boxes[i]

            for (let j = 0; j < length; j++) {
                if (i == j) {
                    continue
                }
                let otherBox = this.boxes[j]
                if (Math.abs(currentBox.x - otherBox.x) <= 10) {
                    if (Math.abs(currentBox.y - otherBox.y) <= 10) {
                        currentBox.reverse()
                        otherBox.reverse()
                    }
                }
            }
        }
    }

    // -----------------------------------------------------

    // Canvas properties.

    createCanvas(): HTMLCanvasElement {
        const canvas = document.createElement("canvas")
        canvas.setAttribute("id", "canvas")

        canvas.height = this.height
        canvas.width = this.width

        canvas.style.border = "3px solid black"
        canvas.style.position = "absolute"
        canvas.style.top = "50%"
        canvas.style.left = "50%"
        canvas.style.transform = "translate(-50%, -50%)"
        canvas.style.height = `${this.height}px`
        canvas.style.width = `${this.width}px`

        this.context = canvas.getContext("2d")
        document.body.appendChild(canvas)

        return this.canvas = canvas
    }

    // -----------------------------------------------------

    // Create an an enemy object.

    createBox() {
        const xSize = 20
        const ySize = 20
        const xPos = Math.floor(Math.random() * this.width - xSize + 1)
        const yPos = Math.floor(Math.random() * this.height - ySize + 1)

        const colours = ["red", "blue", "green", "yellow", "purple", "pink", "black"]
        const colourSelection = colours[Math.floor(Math.random() * colours.length)]

        const newBox = new Box(xPos, yPos, ySize, xSize, this.width, this.height, this.context, colourSelection)
        this.boxes.push(newBox)
    }

    // -----------------------------------------------------

    // Sets up the game. Starts the recursive animation function.

    run() {
        for (let i = 0; i < this.totalBoxes; i++) {
            this.createBox()
        }
        this.animate()
    }

    // -----------------------------------------------------
}

// --------------------------------------------------------------------------------------------------------------------

window.addEventListener("load", () => {
    const game = new Game()
    game.run()
})
