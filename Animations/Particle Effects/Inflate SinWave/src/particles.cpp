#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

void create_particle(std::vector<Particle> *particles, Mouse *mouse, Config *config)
{
    Particle new_particle{
        .radius{static_cast<float>(GetRandomValue(20, 40))},
        .position{
            .x{static_cast<float>(GetRandomValue(std::round(mouse->position.x - mouse->radius), std::round(mouse->position.x + mouse->radius)))},
            .y{static_cast<float>(GetRandomValue(std::round(mouse->position.y - mouse->radius), std::round(mouse->position.y + mouse->radius)))},
        },
        .color{250, 0, 100, 255},
    };
    new_particle.starting_radius = new_particle.radius;
    new_particle.color.r = static_cast<unsigned char>(static_cast<float>(255 * new_particle.position.x / static_cast<float>(config->width)));

    particles->push_back(new_particle);
}

// --------------------------------------------------------------------------------------------------------------------

void draw_particles(std::vector<Particle> *particles)
{
    for (auto &particle : *particles)
    {
        DrawCircle(
            std::round(particle.position.x),
            std::round(particle.position.y),
            particle.radius + 2,
            BLACK);

        DrawCircle(
            std::round(particle.position.x),
            std::round(particle.position.y),
            particle.radius,
            WHITE);

        DrawCircle(
            std::round(particle.position.x),
            std::round(particle.position.y),
            particle.radius,
            particle.color);

        DrawCircle(
            std::round(particle.position.x - particle.radius * 0.2),
            std::round(particle.position.y - particle.radius * 0.3),
            particle.radius * 0.6,
            WHITE);
    }
}

// --------------------------------------------------------------------------------------------------------------------

void update_particles(std::vector<Particle> *particles)
{
    for (auto i = 0; i < particles->size(); i++)
    {
        (*particles)[i].radius -= 0.1;

        if ((*particles)[i].radius <= 0)
        {
            (*particles).erase((*particles).begin() + i);
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------
