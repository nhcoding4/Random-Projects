#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    // ----------------------------------------------------------------------------------------------------------------

    // Init Resources

    Config config{
        .width{750},
        .height{750},
        .target_fps{144},
        .total_particles{1000},
        .mouse_radius{250.0},
        .title{"Base C++"},
        .particle_color{ORANGE},
    };
    init_window(&config);

    std::vector<Particle> particles{};
    create_particle(&particles, &config);

    Mouse mouse = init_mouse(&config);

    while (!WindowShouldClose())
    {
        // ------------------------------------------------------------------------------------------------------------

        // Update

        resize_window(&particles, &config);
        update_mouse(&mouse);
        push_particles(&particles, &mouse);
        update_particles(&particles, &config);
        update_fps_counter(&config);

        // ------------------------------------------------------------------------------------------------------------

        // Draw

        BeginDrawing();

        ClearBackground(BLACK);
        DrawRectangleGradientV(0, 0, config.width, config.height, DARKBLUE, BLACK);

        draw_fps_counter(&config);

        draw_particles(&particles);

        EndDrawing();

        // ------------------------------------------------------------------------------------------------------------
    }

    // Free resources

    CloseWindow();
}

// --------------------------------------------------------------------------------------------------------------------
