
// Global Variables

const width = 600
const height = 600
const canvas = createCanvas(width, height)
const ctx = canvas.getContext("2d")

const spriteWidth = 575
const spriteHeight = 523

// --------------------------------------------------------------------------------------------------------------------

function main() {
    document.body.appendChild(animationSelector())

    const image = loadImage("./shadow_dog.png")
    const frameX = 0
    const frameY = 0
    const gameFrame = 0
    const staggerFrame = 5
    const animationData = generateAnimations()
    const animation = "idle"
    animate(image, frameX, frameY, gameFrame, staggerFrame, animationData, animation)
}

// --------------------------------------------------------------------------------------------------------------------

// Animates an object - Recursive.

function animate(image, frameX, frameY, gameFrame, staggerFrame, animationData, animationState) {
    // Allows the animation to change state depending on user selection    
    const dropDown = document.getElementById("animations")
    dropDown.addEventListener("change", function (e) {
        animationState = e.target.value
    })

    // Get the position of the current frame on the sprite sheet
    let currentLength = animationData[animationState].loc.length
    let position = Math.floor(gameFrame / staggerFrame) % currentLength

    frameX = spriteWidth * position
    frameY = animationData[animationState].loc[position].y

    // Draw sprite.
    ctx.clearRect(0, 0, width, height)
    ctx.drawImage(image, frameX, frameY, spriteWidth, spriteHeight, 0, 0, spriteWidth, spriteHeight)
    
    gameFrame++
    if (gameFrame > currentLength * 10) {
        gameFrame = 1
    }

    requestAnimationFrame(() => {
        animate(image, frameX, frameY, gameFrame, staggerFrame, animationData, animationState)
    })
}

// --------------------------------------------------------------------------------------------------------------------

// Allows changing of animation from browser using a dropdown.

function animationSelector() {
    const root = document.createElement("div")
    root.setAttribute("class", "controls")
    root.style.position = "absolute"
    root.style.zIndex = "10"
    root.style.top = "50px"
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

    return root
}

// --------------------------------------------------------------------------------------------------------------------

// Creates a canvas.

function createCanvas(width, height) {
    const canvas = document.createElement("canvas")
    canvas.setAttribute("id", "canvas1")
    document.body.appendChild(canvas)

    canvas.width = width
    canvas.height = height

    canvas.style.width = `${width}px`
    canvas.style.height = `${height}px`
    canvas.style.border = "5px solid black"
    canvas.style.position = "absolute"
    canvas.style.top = "50%"
    canvas.style.left = "50%"
    canvas.style.transform = "translate(-50%, -50%)"

    return canvas
}

// --------------------------------------------------------------------------------------------------------------------

// Generate sprite animations.

function generateAnimations() {
    const spirteAnimations = []
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

    // Iterate over each animation state and get animation data.
    animationStates.forEach((state, index) => {
        let frames = {
            loc: [],
        }
        for (let i = 0; i < state.frames; i++) {
            let positionX = i * spriteWidth
            let positionY = index * spriteHeight
            frames.loc.push({ x: positionX, y: positionY })
        }
        spirteAnimations[state.name] = frames
    })

    return spirteAnimations
}

// --------------------------------------------------------------------------------------------------------------------

// Load an image.

function loadImage(path) {
    const image = new Image()
    image.src = `${path}`

    return image
}

// --------------------------------------------------------------------------------------------------------------------

main()
