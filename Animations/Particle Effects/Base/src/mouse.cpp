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

void push_particles(std::vector<Particle> *particles, Mouse *mouse)
{
    if (mouse->active)
    {
        for (auto &particle : *particles)
        {
            auto dx{particle.position.x - mouse->position.x};
            auto dy{particle.position.y - mouse->position.y};
            auto distance{std::hypot(dx, dy)};

            if (distance <= mouse->radius)
            {
                auto power{mouse->radius / distance};
                auto angle{std::atan2(dy, dx)};
                particle.push.x = std::cos(angle) * power;
                particle.push.y = std::sin(angle) * power;
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