#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

Mouse init_mouse(Config *config)
{
    Mouse new_mouse{
        .position{
            .x{static_cast<float>(config->width) / 2},
            .y{static_cast<float>(config->height) / 2},
        },
        .radius{config->mouse_radius},
        .active{false},
    };

    return new_mouse;
}
// --------------------------------------------------------------------------------------------------------------------

void inflate_particles(Mouse *mouse, std::vector<Particle> *particles)
{
    if (mouse->active)
    {
        for (auto &particle : *particles)
        {
            auto dx = particle.position.x - mouse->position.x;
            auto dy = particle.position.y - mouse->position.y;
            auto distance = std::hypot(dx, dy);

            if (distance <= mouse->radius && particle.radius < 50)
            {
                particle.radius += 0.5;
            }
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

void update_mouse(Mouse *mouse)
{
    if (IsMouseButtonDown(MOUSE_BUTTON_LEFT))
    {
        mouse->position = GetMousePosition();
        mouse->active = true;
    }
    else
    {
        mouse->active = false;
    }
}

// --------------------------------------------------------------------------------------------------------------------