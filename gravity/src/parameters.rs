use raylib::prelude::*;

#[derive(Default)]
pub struct Parameters {
    pub width: f32,
    pub height: f32,
    pub min_width: f32,
    pub scale: f32,
    pub scaled_width: f32,
    pub scaled_height: f32,
    pub time_step: f32,
    pub total_balls: i32,
    pub title: String,
    pub fps: String,
    pub gravity: Vector2,
}

impl Parameters {
    // --------------------------------------------------------------------------------------------
    // Public Methods.
    // --------------------------------------------------------------------------------------------

    pub fn new() -> Parameters {
        let mut new_parameter = Self {
            width: 1000.0,
            height: 1000.0,
            min_width: 20.0,
            scale: Default::default(),
            scaled_width: Default::default(),
            scaled_height: Default::default(),
            time_step: 1.0 / 144.0,
            total_balls: 10,
            title: "Gravity".to_string(),
            fps: "".to_string(),
            gravity: Vector2 { x: 0.0, y: -15.0 },
        };
        new_parameter.calculate_scale();

        new_parameter
    }

    // --------------------------------------------------------------------------------------------

    pub fn calculate_scale(&mut self) {
        let calculate_min = |a: f32, b: f32| -> f32 {
            if a > b {
                return a;
            }
            b
        };
        self.scale = calculate_min(self.width, self.height) / self.min_width;
        self.scaled_width = self.width / self.scale;
        self.scaled_height = self.height / self.scale;
    }

    // --------------------------------------------------------------------------------------------

    pub fn draw(&self, rl: &mut RaylibDrawHandle) {
        rl.draw_text(&self.fps, 0, 0, 40, Color::GREEN);
    }

    // --------------------------------------------------------------------------------------------

    pub fn update(&mut self, rl: &RaylibHandle) {
        self.update_window_size(rl);
        self.update_fps(rl);
    }

    // --------------------------------------------------------------------------------------------
    // Private Methods
    // --------------------------------------------------------------------------------------------

    fn update_fps(&mut self, rl: &RaylibHandle) {
        self.fps = rl.get_fps().to_string();
    }

    // --------------------------------------------------------------------------------------------

    fn update_window_size(&mut self, rl: &RaylibHandle) {
        if rl.is_window_resized() {
            self.width = rl.get_screen_width() as f32;
            self.height = rl.get_screen_height() as f32;
            self.calculate_scale();
        }
    }

    // --------------------------------------------------------------------------------------------
}
