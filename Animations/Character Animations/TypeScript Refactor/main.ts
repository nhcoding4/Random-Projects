interface framePosition {
    x: number
    y: number
}

interface frameData {
    loc: framePosition[]
}

// --------------------------------------------------------------------------------------------------------------------

class Game {
    height = 600
    spriteHeight = 523
    spriteWidth = 575
    staggerFrame = 5
    width = 600

    animationDropdown: HTMLDivElement
    canvas: HTMLCanvasElement
    context: CanvasRenderingContext2D | null
    defaultAnimation = "idle"
    imageData: HTMLImageElement
    spirteAnimations: frameData[] = []

    // --------------------------------------------------------------------------------------------------------------------

    animate(gameFrame: number, animationState: string) {
        const dropDown = document.getElementById("animations")
        dropDown?.addEventListener("change",  (e) => {
            animationState = (e.target as HTMLInputElement).value
        })

        let currentLength = this.spirteAnimations[animationState].loc.length
        let position = Math.floor(gameFrame / this.staggerFrame) % currentLength

        let frameX = this.spriteWidth * position
        let frameY = this.spirteAnimations[animationState].loc[position].y

        this.context?.clearRect(0, 0, this.width, this.height)
        this.context?.drawImage(this.imageData, frameX, frameY, this.spriteWidth, this.spriteHeight, 0, 0, this.spriteWidth, this.spriteHeight)

        gameFrame++
        if (gameFrame > currentLength * 10) {
            gameFrame = 1
        }

        requestAnimationFrame(() => {
            this.animate(gameFrame, animationState)
        })
    }

    // --------------------------------------------------------------------------------------------------------------------

    // Creates a box that allows animation selection

    animationSelector(): HTMLDivElement {
        const root = document.createElement("div")
        root.setAttribute("class", "controls")
        root.style.position = "absolute"
        root.style.zIndex = "10"
        root.style.top = "10px"
        root.style.left = "50%"
        root.style.transform = "translateX(-50%)"

        const label = document.createElement("label")
        label.setAttribute("for", "animations")
        label.textContent = "Choose Animation:"
        label.style.fontSize = "25px"

        const select = document.createElement("select")
        select.setAttribute("id", "animations")
        select.setAttribute("name", "animations")
        select.style.fontSize = "25px"

        const options = ["idle", "jump", "fall", "run", "dizzy", "sit", "roll", "bite", "ko", "getHit"]

        for (let i = 0; i < options.length; i++) {
            const newOption = document.createElement("option")
            const value = options[i]
            newOption.setAttribute("value", value)
            newOption.textContent = value
            newOption.style.fontSize = "25px"

            select.appendChild(newOption)
        }

        root.appendChild(label)
        root.appendChild(select)

        document.body.appendChild(root)

        return this.animationDropdown = root
    }

    // --------------------------------------------------------------------------------------------------------------------

    // Create a base canvas.

    createCanvas(): HTMLCanvasElement {
        const canvas = document.createElement("canvas")
        canvas.setAttribute("id", "canvas")

        canvas.width = this.width
        canvas.height = this.height

        canvas.style.width = `${this.width}px`
        canvas.style.height = `${this.height}px`
        canvas.style.border = "3px solid black"
        canvas.style.position = "absolute"
        canvas.style.top = "50%"
        canvas.style.left = "50%"
        canvas.style.transform = "translate(-50%, -50%)"

        document.body.appendChild(canvas)

        return this.canvas = canvas
    }

    // --------------------------------------------------------------------------------------------------------------------

    // Load sprite sheet

    loadImage(): HTMLImageElement {
        const image = new Image()
        const imageLocation = "./shadow_dog.png"
        image.src = imageLocation

        return this.imageData = image
    }

    // --------------------------------------------------------------------------------------------------------------------

    // Generate sprite sheet animation locations

    generateAnimations() {
        const animationStates = [
            {
                name: "idle",
                frames: 7,
            },
            {
                name: "jump",
                frames: 7,
            },
            {
                name: "fall",
                frames: 7,
            },
            {
                name: "run",
                frames: 9,
            },
            {
                name: "dizzy",
                frames: 11,
            },
            {
                name: "sit",
                frames: 5,
            },
            {
                name: "roll",
                frames: 7,
            },
            {
                name: "bite",
                frames: 7,
            },
            {
                name: "ko",
                frames: 12,
            },
            {
                name: "getHit",
                frames: 4,
            },
        ]

        animationStates.forEach((state, index) => {
            let frames: frameData = {
                loc: []
            }
            for (let i = 0; i < state.frames; i++) {
                let positionX = i * this.spriteWidth
                let positionY = index * this.spriteHeight
                let newFrame: framePosition = { x: positionX, y: positionY }

                frames.loc.push(newFrame)
            }
            this.spirteAnimations[state.name] = frames
        })
    }


    // --------------------------------------------------------------------------------------------------------------------

    run() {
        this.loadImage()
        this.createCanvas()
        this.context = this.canvas.getContext("2d")
        this.animationSelector()
        this.generateAnimations()
        this.animate(0, this.defaultAnimation)
    }
}

// --------------------------------------------------------------------------------------------------------------------

function main() {
    const game = new Game()
    game.run()
}

main()
