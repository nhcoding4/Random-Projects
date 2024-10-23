
// --------------------------------------------------------------------------------------------------------------------

class Layer {
    context: CanvasRenderingContext2D | null
    gameFrame: number
    gameSpeed: number
    height: number
    image: HTMLImageElement
    speed: number
    speedModifier: number
    width: number
    x2: number
    x: number
    y: number

    // --------------------------------------------------------

    // Object setup.

    constructor(
        image: HTMLImageElement,
        speedModifier: number,
        gameSpeed: number,
        context: CanvasRenderingContext2D | null,
        height: number,
        width: number,
    ) {
        this.context = context
        this.gameFrame = 0
        this.height = height
        this.image = image
        this.speed = gameSpeed * speedModifier
        this.speedModifier = speedModifier
        this.width = width
        this.x2 = width
        this.x = 0
        this.y = 0
    }

    // --------------------------------------------------------

    // Draws images to the canvas.

    draw() {
        this.context?.drawImage(this.image, this.x, this.y, this.width, this.height)
        this.context?.drawImage(this.image, this.x2, this.y, this.width, this.height)
    }

    // --------------------------------------------------------

    // Change the x position of the layer

    update() {
        const moveBackground = ((xValue: number, offset: number) => {
            if (xValue < -this.width) {
                xValue = this.width - this.speed + offset
            } else {
                xValue -= this.speed
            }
            return xValue
        })

        this.x = Math.floor(moveBackground(this.x, this.x2))
        this.x2 = Math.floor(moveBackground(this.x2, this.x))
    }

    // --------------------------------------------------------        

    // Updates the speed of the layer

    updateSpeed(newSpeed: number) {
        this.speed = newSpeed * this.speedModifier
    }

    // --------------------------------------------------------

}

// --------------------------------------------------------------------------------------------------------------------

class Backgrounds {
    canvas: HTMLCanvasElement
    canvasHeight: number
    canvasWidth: number
    context: CanvasRenderingContext2D | null

    layers: Layer[]

    gamespeed: number
    x: number
    x2: number

    paragraph: HTMLParagraphElement
    slider: HTMLInputElement

    // --------------------------------------------------------

    // Object setup.

    constructor() {
        document.body.style.background = "black"

        this.canvasHeight = 700
        this.canvasWidth = 800
        this.gamespeed = 5
        this.x = 0
        this.x2 = 2400

        this.createElements()
        this.loadImages()
    }

    // --------------------------------------------------------

    // Draw objects.

    animate() {
        this.context?.clearRect(0, 0, this.canvasWidth, this.canvasHeight)
        this.gamespeed = parseInt(this.slider.value)
        this.paragraph.textContent = `Game Speed: ${this.gamespeed}`

        for (let layer of this.layers) {
            layer.updateSpeed(this.gamespeed)
            layer.update()
            layer.draw()
        }

        requestAnimationFrame(() => {
            this.animate()
        })
    }

    // --------------------------------------------------------

    // Create object to draw on. 

    createCanvas(): HTMLCanvasElement {
        const canvas = document.createElement("canvas")
        canvas.setAttribute("id", "canvas")

        canvas.height = this.canvasHeight
        canvas.width = this.canvasWidth

        canvas.style.position = "relative"
        canvas.style.width = `${this.canvasWidth}px`
        canvas.style.height = `${this.canvasHeight}px`

        document.body.appendChild(canvas)
        this.context = canvas.getContext("2d")

        this.canvas = canvas

        return canvas
    }

    // --------------------------------------------------------

    // Creates a container for all the other objects.

    createContainer(): HTMLDivElement {
        const container = document.createElement("div")

        container.style.position = "absolute"
        container.style.width = `${this.canvasWidth}px`
        container.style.transform = "translate(-50%, -50%)"
        container.style.top = "50%"
        container.style.left = "50%"
        container.style.border = "3px solid white"
        container.style.fontSize = "25px"

        return container
    }

    // --------------------------------------------------------

    // Creates the canvas and slider elements.

    createElements() {
        const container = this.createContainer()
        document.body.appendChild(container)

        const canvas = this.createCanvas()
        container.appendChild(canvas)

        const paragraph = document.createElement("p")
        paragraph.textContent = `Game Speed: ${this.gamespeed}`
        paragraph.style.color = "white"
        container.appendChild(paragraph)
        this.paragraph = paragraph

        const span = document.createElement("span")
        span.setAttribute("id", "showGameSpeed")
        paragraph.appendChild(span)

        const slider = this.createSlider()
        slider.style.width = "100%"

        container.appendChild(slider)
    }

    // --------------------------------------------------------

    // Creates a slider element.

    createSlider(): HTMLInputElement {
        const input = document.createElement("input")
        input.setAttribute("type", "range")
        input.setAttribute("min", "0")
        input.setAttribute("max", "20")
        input.setAttribute("value", "5")
        input.setAttribute("class", "slider")
        input.setAttribute("id", "slider")

        this.slider = input
        this.slider.value = `${this.gamespeed}`

        return input
    }

    // --------------------------------------------------------

    // Loads images from files and creates layer objects out of them.

    loadImages(): Layer[] {
        const base = "./backgroundLayers/"
        const images: HTMLImageElement[] = []

        for (let i = 0; i < 5; i++) {
            const backgroundLayer = new Image()
            backgroundLayer.src = `${base}layer-${i + 1}.png`
            images.push(backgroundLayer)
        }

        const layers: Layer[] = []
        let speedModifier = 0.1

        for (let image of images) {
            const newLayer = new Layer(image, speedModifier, this.gamespeed, this.context, this.canvasHeight, this.x2)
            layers.push(newLayer)
            speedModifier += 0.1
        }

        return this.layers = layers
    }

    // --------------------------------------------------------

    // Main loop.

    run() {
        this.animate()
    }

    // --------------------------------------------------------

}

// --------------------------------------------------------------------------------------------------------------------

window.addEventListener("load", () => {
    const background = new Backgrounds()
    background.run()
})

