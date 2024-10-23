mod ball;
mod parameters;

use crate::ball::*;
use crate::parameters::*;
use raylib::prelude::*;

fn main() {
    // --------------------------------------------------------------------------------------------
    // Init
    // --------------------------------------------------------------------------------------------

    let mut parameters = Parameters::new();

    let (mut rl, thread) = init()
        .msaa_4x()
        .vsync()
        .resizable()
        .size(parameters.width as i32, parameters.height as i32)
        .title(&parameters.title)
        .build();

    // --------------------------------------------------------------------------------------------
    // Asset Creation
    // --------------------------------------------------------------------------------------------

    let mut balls: Vec<Ball> = Vec::new();

    for _ in 0..parameters.total_balls {
        balls.push(Ball::new(
            &parameters.time_step,
            &parameters.scaled_width,
            &parameters.scaled_height,
            &parameters.gravity,
        ));
    }

    while !rl.window_should_close() {
        // --------------------------------------------------------------------------------------------
        // Update
        // --------------------------------------------------------------------------------------------

        parameters.update(&rl);

        for ball in balls.iter_mut() {
            ball.update(
                &parameters.scaled_width,
                &parameters.scaled_height,
                &parameters.scale,
                &parameters.height,
            );
        }

        // --------------------------------------------------------------------------------------------
        // Draw
        // --------------------------------------------------------------------------------------------

        let mut rl_draw = rl.begin_drawing(&thread);
        rl_draw.clear_background(Color::BLACK);

        for ball in balls.iter() {
            ball.draw(&mut rl_draw);
        }

        parameters.draw(&mut rl_draw);

        // --------------------------------------------------------------------------------------------
    }
}
