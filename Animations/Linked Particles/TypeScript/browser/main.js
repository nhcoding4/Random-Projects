// src/particles.ts
class Particle {
  context;
  width;
  height;
  mouse;
  x;
  y;
  radius = Math.floor(Math.random() * 15 + 1);
  dx = Math.random() * 2 - 1;
  dy = Math.random() * 2 - 1;
  pushForceX = 0;
  pushForceY = 0;
  friction;
  constructor(context, width, height, mouse) {
    this.context = context;
    this.width = width;
    this.height = height;
    this.mouse = mouse;
    this.setPosition();
    this.friction = this.calculateFriction();
  }
  bounds() {
    if (this.x < 0 + this.radius) {
      this.x = this.radius;
      this.dx *= -1;
    }
    if (this.x > this.width - this.radius) {
      this.x = this.width - this.radius;
      this.dx *= -1;
    }
    if (this.y < 0 + this.radius) {
      this.y = this.radius;
      this.dy *= -1;
    }
    if (this.y > this.height - this.radius) {
      this.y = this.height - this.radius;
      this.dy *= -1;
    }
  }
  calculateFriction() {
    if (this.radius <= 7) {
      return 0.99;
    } else if (this.radius <= 10) {
      return 0.975;
    } else {
      return 0.95;
    }
  }
  draw() {
    if (this.context != null) {
      this.context.beginPath();
      this.context.arc(this.x, this.y, this.radius, 0, 2 * Math.PI, false);
      this.context.fill();
      this.context.stroke();
    }
  }
  getPosition() {
    return [this.x, this.y];
  }
  getRadius() {
    return this.radius;
  }
  move() {
    this.x += this.dx + this.pushForceX;
    this.y += this.dy + this.pushForceY;
  }
  pushWithMouse() {
    if (this.mouse.getPressedStatus()) {
      let mousePositions = this.mouse.getPosition();
      let dx = this.x - mousePositions[0];
      let dy = this.y - mousePositions[1];
      let distance = Math.hypot(dx, dy);
      if (distance < this.mouse.getRadius()) {
        let force = this.mouse.getRadius() / distance;
        let angle = Math.atan2(dy, dx);
        this.pushForceX = Math.cos(angle) * force;
        this.pushForceY = Math.sin(angle) * force;
      }
    }
  }
  resizeCanvas(width, height) {
    this.width = width;
    this.height = height;
    this.setPosition();
  }
  setPosition() {
    this.x = Math.random() * this.width - this.radius * 2 + this.radius * 2;
    this.y = Math.random() * this.height - this.radius * 2 + this.radius * 2;
  }
  update() {
    this.pushWithMouse();
    this.move();
    this.bounds();
    this.updateForce();
  }
  updateForce() {
    this.pushForceX *= this.friction;
    this.pushForceY *= this.friction;
  }
}

// src/effect.ts
class Effect {
  context;
  width;
  height;
  totalParticles;
  connectionDistance;
  lineWidth;
  mouse;
  particles = [];
  constructor(context, width, height, totalParticles, connectionDistance, lineWidth, mouse) {
    this.context = context;
    this.width = width;
    this.height = height;
    this.totalParticles = totalParticles;
    this.connectionDistance = connectionDistance;
    this.lineWidth = lineWidth;
    this.mouse = mouse;
    this.createParticles();
  }
  draw() {
    this.connectParticles();
    for (let particle of this.particles) {
      particle.draw();
    }
  }
  createParticles() {
    for (let i = 0;i < this.totalParticles; i++) {
      let newParticle = new Particle(this.context, this.width, this.height, this.mouse);
      this.particles.push(newParticle);
    }
  }
  connectParticles() {
    for (let i = 0;i < this.totalParticles; i++) {
      let iParticlePositions = this.particles[i].getPosition();
      for (let j = i;j < this.totalParticles; j++) {
        if (i == j) {
          continue;
        }
        let jParticlePositions = this.particles[j].getPosition();
        let dx = iParticlePositions[0] - jParticlePositions[0];
        let dy = iParticlePositions[1] - jParticlePositions[1];
        let distance = Math.hypot(dx, dy);
        if (distance <= this.connectionDistance && this.context != null) {
          let opacity = 1 - distance / this.connectionDistance;
          this.context.save();
          this.context.beginPath();
          this.context.moveTo(iParticlePositions[0], iParticlePositions[1]);
          this.context.lineTo(jParticlePositions[0], jParticlePositions[1]);
          this.context.lineWidth = this.lineWidth;
          this.context.globalAlpha = opacity;
          this.context.strokeStyle = "white";
          this.context.stroke();
          this.context.restore();
        }
      }
    }
  }
  resize(width, height) {
    this.width = width;
    this.height = height;
    for (let particle of this.particles) {
      particle.resizeCanvas(this.width, this.height);
    }
  }
  update() {
    for (let particle of this.particles) {
      particle.update();
    }
  }
}

// src/mouse.ts
class Mouse {
  radius;
  x = 0;
  y = 0;
  pressed = false;
  constructor(radius) {
    this.radius = radius;
    window.addEventListener("mousedown", (e) => {
      this.pressed = true;
      this.x = e.x;
      this.y = e.y;
    });
    window.addEventListener("mouseup", (e) => {
      this.pressed = false;
      this.x = e.x;
      this.y = e.y;
    });
    window.addEventListener("mousemove", (e) => {
      if (this.getPressedStatus()) {
        this.x = e.x;
        this.y = e.y;
      }
    });
  }
  getPosition() {
    return [this.x, this.y];
  }
  getPressedStatus() {
    return this.pressed;
  }
  getRadius() {
    return this.radius;
  }
}

// src/main.ts
class Simulation {
  width = window.innerWidth;
  height = window.innerHeight;
  totalParticles = 750;
  canvas;
  context;
  mouse;
  effect;
  constructor() {
    this.canvas = this.makeCanvas();
    this.context = this.getContext();
    this.setContextProperties();
    this.mouse = this.makeMouse();
    this.effect = this.makeEffect();
    window.addEventListener("resize", (_) => {
      let width = window.innerWidth;
      let height = window.innerHeight;
      if (Math.abs(width - this.width) != 0 || Math.abs(height - this.height) != 0) {
        this.width = width;
        this.height = height;
        this.canvas.width = this.width;
        this.canvas.height = this.height;
        this.effect.resize(this.width, this.height);
        this.setContextProperties();
      }
    });
  }
  getContext() {
    return this.canvas.getContext("2d");
  }
  makeCanvas() {
    let newCanvas = document.createElement("canvas");
    newCanvas.setAttribute("id", "canvas1");
    newCanvas.width = this.width;
    newCanvas.height = this.height;
    newCanvas.style.background = "black";
    newCanvas.style.position = "abosolute";
    newCanvas.style.left = "0";
    newCanvas.style.top = "0";
    document.body.appendChild(newCanvas);
    return newCanvas;
  }
  makeEffect() {
    let connectionDistance = 100;
    let lineWidth = 1.5;
    return new Effect(this.context, this.width, this.height, this.totalParticles, connectionDistance, lineWidth, this.mouse);
  }
  makeMouse() {
    let radius = 300;
    return new Mouse(radius);
  }
  run() {
    if (this.context != null) {
      this.context.clearRect(0, 0, this.width, this.height);
    }
    this.effect.update();
    this.effect.draw();
    requestAnimationFrame(() => {
      this.run();
    });
  }
  setContextProperties() {
    if (this.context != null) {
      let gradient = this.context.createLinearGradient(0, 0, this.width, this.height);
      gradient.addColorStop(0, "cyan");
      gradient.addColorStop(0.5, "purple");
      gradient.addColorStop(1, "orangered");
      this.context.fillStyle = gradient;
      this.context.strokeStyle = "black";
    }
  }
}
var simulation = new Simulation;
simulation.run();
