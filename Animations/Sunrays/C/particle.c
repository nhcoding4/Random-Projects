#include "def.h"

// --------------------------------------------------------------------------------------------------------------------

void collision(Particle *particle, Config *config)
{
    if (particle->x < particle->radius)
    {
        particle->x = particle->radius;
        particle->movementX *= -1;
    }

    if (particle->x > config->width - particle->radius)
    {
        particle->x = config->width - particle->radius;
        particle->movementX *= -1;
    }

    if (particle->y < particle->radius)
    {
        particle->y = particle->radius;
        particle->movementY *= -1;
    }

    if (particle->y > config->height - particle->radius)
    {
        particle->y = config->height - particle->radius;
        particle->movementY *= -1;
    }
}

// --------------------------------------------------------------------------------------------------------------------

void pushParticle(Particle *particle, Mouse *mouse)
{
    int32_t dx = particle->x - mouse->x;
    int32_t dy = particle->y - mouse->y;
    double distance = hypot(dx, dy);

    if (distance < mouse->radius)
    {
        double power = mouse->radius / distance;
        double angle = atan2(dy, dx);
        particle->pushX = cos(angle) * power;
        particle->pushY = sin(angle) * power;
    }
}

// --------------------------------------------------------------------------------------------------------------------

void moveParticle(Particle *particle)
{
    particle->x += particle->movementX + round(particle->pushX);
    particle->y += particle->movementY + round(particle->pushY);

    if (particle->pushX != 0)
    {
        particle->pushX *= particle->friction;
    }
    if (particle->pushY != 0)
    {
        particle->pushY *= particle->friction;
    }
}

// --------------------------------------------------------------------------------------------------------------------
