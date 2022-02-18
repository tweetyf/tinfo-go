class Sprite {
    constructor(name, initLeft, initBottom, width, height) {
        this.name = name
        this.bottom = initBottom
        this.left = initLeft
        this.width = width
        this.height = height
        this.visualDiv = document.createElement('div')
        this.visualDiv.classList.add(name)
        this.speedX = 0
        this.speedY = 0
    }

    update(newLeft, newBottom) {
        this.bottom = newBottom
        this.left = newLeft
    }

    move(leftpx, bottompx) {
        this.bottom += bottompx
        this.left += leftpx
    }

    render(viewportX, viewportY) {
        let x = this.left - viewportX
        let y = this.bottom - viewportY
        this.visualDiv.style.left = x + 'px'
        this.visualDiv.style.bottom = y + 'px'
    }

    destory() {
        this.visualDiv.classList.remove(this.name)
    }
}

class Scene {
    constructor(width, height) {
        this.GRAVITY = 0.98
        this.isGameOver = false
        this.isOnPlatform = true
        this.finalScore = 0
        this.SceneWidth = width
        this.SceneHeight = height
        // init objects
        this.platforms = []
        this.platformCount = 6
        this.createPlatforms()
        this.newDoodler = new Sprite("doodler", this.platforms[3].left, this.platforms[3].bottom, 87, 85)
        // Init view port
        this.viewPort = new Viewport(400, 600)
        this.render()
        this.viewPort.addObj(this.newDoodler.visualDiv)
        this.platforms.forEach(plt => {
            this.viewPort.addObj(plt.visualDiv)
        });
    }

    createPlatforms() {
        for (let i = 0; i < this.platformCount; i++) {
            let platGap = 100
            let x = getRndInteger(0, (this.SceneWidth - 100))
            let y = platGap * i + 0
            let newPlatform = new Sprite("platform", x, y, 85, 15)
            this.platforms.push(newPlatform)
        }
    }

    update() {
        // viewport
        this.viewPort.move(0, 1)
        // doodler: gravity
        let vt = this.newDoodler.speedY - this.GRAVITY * 1
        let y = vt * 1
        this.newDoodler.speedY = vt
        // doodler: boundry check
        if (this.newDoodler.left <= 0) {
            this.newDoodler.left = 0
        } else if (this.newDoodler.left >= this.SceneWidth - 85) {
            this.newDoodler.left = this.SceneWidth - 85
        }
        if (this.newDoodler.bottom >= this.viewPort.viewportY + 525) {
            this.newDoodler.bottom = this.viewPort.viewportY + 525
            this.newDoodler.speedY = 0
        }
        // doodler: check platform collision
        if (this.newDoodler.speedY <= 0) {
            this.platforms.forEach(plt => {
                if (this.newDoodler.left >= plt.left - 0.4 * this.newDoodler.left &&
                    this.newDoodler.left <= plt.left + plt.width &&
                    this.newDoodler.bottom >= plt.bottom &&
                    this.newDoodler.bottom <= plt.bottom + plt.height) {
                    // doodler can stand on platform now.
                    this.newDoodler.speedX = 0
                    this.newDoodler.speedY = 0
                    this.isOnPlatform = true
                }
            });
        } else {
            this.isOnPlatform = false
        }
        // move doodler
        this.newDoodler.move(this.newDoodler.speedX, this.newDoodler.speedY)
        //platform
        if (this.newDoodler.bottom < this.viewPort.viewportY) {
            this.gameOver()
        }
        this.platforms.forEach(plt => {
            if (plt.bottom < this.viewPort.viewportY) {
                plt.destory()
                this.platforms.shift()
                let x = getRndInteger(0, (this.SceneWidth - 100))
                let y = this.viewPort.viewportY + this.viewPort.height
                let newPlatform = new Sprite("platform", x, y, 85, 15)
                this.viewPort.addObj(newPlatform.visualDiv)
                this.platforms.push(newPlatform)
                this.finalScore += 1
                this.viewPort.infoDiv.innerHTML = '' + this.finalScore
            }
        });
    }

    updateInput(e) {
        console.log(e.key)
        if (e.key === 'ArrowLeft') {
            this.newDoodler.speedX += -4
        } else if (e.key === 'ArrowRight') {
            this.newDoodler.speedX += 4
        } else if (e.key === 'ArrowUp') {
            if (this.isOnPlatform) { this.newDoodler.speedY = 20 }
        } else if (e.key === 'Enter') {
            // reset the game.
            this.viewPort.infoDiv.innerHTML = '' + this.finalScore
            this.isGameOver = false
            this.newDoodler.update(this.platforms[3].left, this.platforms[3].bottom)
            this.newDoodler.speedX = 0
            this.newDoodler.speedY = 0
            this.finalScore = 0
        }
    }

    render() {
        let x = this.viewPort.viewportX
        let y = this.viewPort.viewportY
        this.newDoodler.render(x, y)
        this.platforms.forEach(plt => {
            plt.render(x, y)
        });

    }

    gameOver() {
        this.isGameOver = true
        //this.viewPort.visualDiv.removeChild(this.viewPort.visualDiv.firstChild)
        this.viewPort.infoDiv.innerHTML = "Game Over! <br>Press Enter to restart<br>Your Final Score: " + this.finalScore
        console.log("game over!")
    }

}

class Viewport {
    constructor(width, height) {
        this.width = width
        this.height = height
        this.viewportX = 0
        this.viewportY = 0
        this.visualDiv = document.querySelector('.grid')
        this.infoDiv = document.querySelector('.info')

    }

    addObj(objNode) {
        this.visualDiv.appendChild(objNode)
    }

    move(leftpx, bottompx) {
        this.viewportX += leftpx
        this.viewportY += bottompx
    }
}

function main() {

    let worldScene = new Scene(400, 600)

    function control(e) {
        worldScene.updateInput(e)
    }

    function updateScene() {
        if (!worldScene.isGameOver) {
            worldScene.update()
            worldScene.render()
        }
    }

    if (!worldScene.isGameOver) {
        setInterval(updateScene, 10);
        document.addEventListener('keydown', control)
    }

}

function getRndInteger(min, max) {
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

document.addEventListener("DOMContentLoaded", () => {
    main()
})