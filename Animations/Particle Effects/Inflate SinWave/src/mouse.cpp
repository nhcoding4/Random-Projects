#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

Mouse init_mouse(Config *config)
{
    Mouse new_mouse{
        .position{
            .x{static_cast<float>(config->width / 2)},
            .y{static_cast<float>(config->height / 2)},
        },
        .radius{config->mouse_radius},
        .angle{0},
        .angle_speed{static_cast<float>(0.5) + static_cast<float>(rand()) / static_cast<float>(RAND_MAX)},
        .x_div{540.0},
        .y_div{720.0},
    };

    return new_mouse;
}
// --------------------------------------------------------------------------------------------------------------------

void mouse_effect(Mouse *mouse, std::vector<Particle> *particles, Config *config)
{
    if (IsMouseButtonDown(MOUSE_BUTTON_LEFT))
    {
        mouse->position = GetMousePosition();
    }
    else
    {
        mouse->angle += mouse->angle_speed;

        auto cast_width{static_cast<float>(config->width)};
        auto cast_height{static_cast<float>(config->height)};
        auto cast_pi{static_cast<float>(PI)};

        mouse->position.x = ((cast_width - mouse->radius) / 2.0 * std::sin(mouse->angle * cast_pi / mouse->x_div)) + (cast_width / 2.0 - mouse->radius);
        mouse->position.y = ((cast_height - mouse->radius) / 2.0 * std::cos(mouse->angle * cast_pi / mouse->y_div)) + (cast_height / 2.0 - mouse->radius);

        if (mouse->position.x - mouse->radius * 2 < 0)
        {
            mouse->position.x = mouse->radius * 2;
        }
        if (mouse->position.x + mouse->radius * 2 > static_cast<float>(config->width))
        {
            mouse->position.x = static_cast<float>(config->width) - (mouse->radius * 2);
        }
        if (mouse->position.y - mouse->radius * 2 < 0)
        {
            mouse->position.y = mouse->radius * 2;
        }
        if (mouse->position.y + mouse->radius * 2 > static_cast<float>(config->height))
        {
            mouse->position.y = static_cast<float>(config->height) - (mouse->radius * 2);
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------
