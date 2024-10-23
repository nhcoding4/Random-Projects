#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

void init_window(Config *config)
{
    SetConfigFlags(FLAG_MSAA_4X_HINT);
    SetConfigFlags(FLAG_WINDOW_RESIZABLE);
    SetConfigFlags(FLAG_WINDOW_HIGHDPI);

    InitWindow(config->width, config->height, config->title.data());
    SetTargetFPS(config->target_fps);
}

// --------------------------------------------------------------------------------------------------------------------

void draw_fps_counter(Config *config)
{
    DrawText(config->fps_text.data(), 0, 0, 40, GREEN);
}

// --------------------------------------------------------------------------------------------------------------------

void resize_window(std::vector<Particle> *particles, Config *config)
{
    if (IsWindowResized())
    {
        config->width = GetScreenWidth();
        config->height = GetScreenHeight();

        set_position(particles, config);
    }
}

// --------------------------------------------------------------------------------------------------------------------

void update_fps_counter(Config *config)
{
    char buffer[5]{};
    sprintf(buffer, "%d", GetFPS());
    config->fps_text = buffer;
}

// --------------------------------------------------------------------------------------------------------------------
