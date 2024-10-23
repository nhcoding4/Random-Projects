use rand::prelude::*;
use raylib::prelude::*;

#[derive(Default)]
pub struct Ball {
    pub radius: f32,
    pub time_step: f32,
    pub scaled_x: f32,
    pub scaled_y: f32,
    pub position: Vector2,
    pub velocity: Vector2,
    pub gravity: Vector2,
    pub color: Color,
}

impl Ball {
    // --------------------------------------------------------------------------------------------
    // Public Methods
    // --------------------------------------------------------------------------------------------

    pub fn new(
        time_step: &f32,
        scaled_width: &f32,
        scaled_height: &f32,
        gravity: &Vector2,
    ) -> Ball {
        let mut rng = rand::thread_rng();

        Self {
            radius: rng.gen_range(10.0..20.0),
            time_step: *time_step,
            scaled_x: Default::default(),
            scaled_y: Default::default(),
            position: Vector2 {
                x: rng.gen_range(0.0..*scaled_width),
                y: rng.gen_range(0.0..*scaled_height),
            },
            velocity: Vector2 {
                x: rng.gen_range(-10.0..10.0),
                y: 0.0,
            },
            gravity: *gravity,
            color: Color::ORANGE,
        }
    }

    // --------------------------------------------------------------------------------------------

    pub fn draw(&self, rl: &mut RaylibDrawHandle) {
        rl.draw_circle(
            f32::round(self.scaled_x) as i32,
            f32::round(self.scaled_y) as i32,
            self.radius + 2.0,
            Color::BLACK,
        );
        rl.draw_circle(
            f32::round(self.scaled_x) as i32,
            f32::round(self.scaled_y) as i32,
            self.radius,
            Color::WHITE,
        );
        rl.draw_circle(
            f32::round(self.scaled_x) as i32,
            f32::round(self.scaled_y) as i32,
            self.radius,
            self.color,
        );
    }

    // --------------------------------------------------------------------------------------------

    pub fn update(
        &mut self,
        scaled_width: &f32,
        scaled_height: &f32,
        scale_factor: &f32,
        window_height: &f32,
    ) {
        self.update_position();
        self.check_bounds(scaled_width, scaled_height, scale_factor);
        self.update_scaled_position(scale_factor, window_height);
        self.update_color(scale_factor, window_height);
    }

    // --------------------------------------------------------------------------------------------
    // Private Methods
    // --------------------------------------------------------------------------------------------

    fn check_bounds(&mut self, scaled_width: &f32, scaled_height: &f32, scale_factor: &f32) {
        let scaled_radius = self.radius / scale_factor;

        if self.position.x - scaled_radius < 0.0 {
            self.position.x = scaled_radius;
            self.velocity.x *= -1.0;
        }
        if self.position.x + scaled_radius > *scaled_width {
            self.position.x = scaled_width - scaled_radius;
            self.velocity.x *= -1.0;
        }
        if self.position.y - scaled_radius < 0.0 {
            self.position.y = scaled_radius;
            self.velocity.y *= -1.0;
        }
        if self.position.y + scaled_radius > *scaled_height {
            self.position.y = scaled_height - scaled_radius;
        }
    }

    // --------------------------------------------------------------------------------------------

    fn update_color(&mut self, scale_factor: &f32, window_height: &f32) {
        let scaled_position = self.position.y * scale_factor;
        self.color.a = f32::round(255.0 * (1.0 - (scaled_position / window_height))) as u8;
    }

    // --------------------------------------------------------------------------------------------

    fn update_scaled_position(&mut self, scale_factor: &f32, window_height: &f32) {
        self.scaled_x = self.position.x * scale_factor;
        self.scaled_y = window_height - self.position.y * scale_factor;
    }

    // --------------------------------------------------------------------------------------------

    fn update_position(&mut self) {
        self.velocity.x += self.gravity.x * self.time_step;
        self.velocity.y += self.gravity.y * self.time_step;

        self.position.x += self.velocity.x * self.time_step;
        self.position.y += self.velocity.y * self.time_step;
    }
    // --------------------------------------------------------------------------------------------
}
