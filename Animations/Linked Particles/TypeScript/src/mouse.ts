export class Mouse {
    private x = 0
    private y = 0
    private pressed = false

    constructor(
        private radius: number
    ) {
        window.addEventListener("mousedown", e => {
            this.pressed = true
            this.x = e.x
            this.y = e.y
        })

        window.addEventListener("mouseup", e => {
            this.pressed = false
            this.x = e.x
            this.y = e.y
        })

        window.addEventListener("mousemove", e => {
            if (this.getPressedStatus()) {
                this.x = e.x
                this.y = e.y
            }
        })
    }

    // --------------------------------------------------------------------------------------------

    // Returns the position of the mouse cursor.

    public getPosition(): number[] {
        return [this.x, this.y]
    }

    // --------------------------------------------------------------------------------------------

    // Returns the left click status of the cursor. 

    public getPressedStatus(): boolean {
        return this.pressed
    }

    // --------------------------------------------------------------------------------------------

    // Returns the radius set around the mouse cursor. 

    public getRadius(): number {
        return this.radius
    }

    // --------------------------------------------------------------------------------------------

}

