#include "def.h"

// --------------------------------------------------------------------------------------------------------------------

void calculateColors(EffectManager *effect)
{
    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        Color color = ColorFromHSV((float)effect->particles[i].x + effect->particles[i].y, 1.0, 1.0);
        effect->colors[i] = color;
    }
}

// --------------------------------------------------------------------------------------------------------------------

void createParticles(EffectManager *effect, Config *config)
{
    Particle *particleArray = (Particle *)malloc(sizeof(Particle) * effect->totalParticles + 1);
    effect->particles = particleArray;

    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        Particle newParticle = {
            radius : GetRandomValue(5, 15),
            friction : 0.95,
            pushX : 0,
            pushY : 0,
        };
        effect->particles[i] = newParticle;
    }

    setPosition(effect, config);
}

// --------------------------------------------------------------------------------------------------------------------

void draw(EffectManager *effect, Mouse *mouse)
{
    Color color = {255, 255, 255, 255 * 0.4};
    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        if (i % effect->rayDivisor == 0)
        {
            Vector2 start = {effect->particles[i].x, effect->particles[i].y};
            DrawLineEx(start, mouse->mouseVector2, 1.0, color);
        }
    }

    for (int32_t i = 0; i < effect->lines.total; i++)
    {
        DrawLineEx(effect->lines.start[i], effect->lines.end[i], effect->lineThickness, effect->lines.color[i]);
    }

    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        DrawCircle(effect->particles[i].x, effect->particles[i].y, (float)effect->particles[i].radius, effect->colors[i]);
    }
}

// --------------------------------------------------------------------------------------------------------------------

void findConnections(EffectManager *effect)
{
    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        for (int32_t j = i; j < effect->totalParticles; j++)
        {
            if (i == j)
            {
                continue;
            }

            int32_t dx = effect->particles[i].x - effect->particles[j].x;
            int32_t dy = effect->particles[i].y - effect->particles[j].y;
            double distance = hypot(dx, dy);

            if (distance < effect->connectionDistance)
            {
                Vector2 start = {effect->particles[i].x,
                                 effect->particles[i].y};
                Vector2 end = {effect->particles[j].x,
                               effect->particles[j].y};

                double opacity = 1 - (distance / effect->connectionDistance);
                Color color = {
                    255,
                    255,
                    255,
                    255 * opacity,
                };

                effect->lines.start[effect->lines.total] = start;
                effect->lines.end[effect->lines.total] = end;
                effect->lines.color[effect->lines.total] = color;
                effect->lines.total++;

                checkAllocationLines(effect);
            }
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

void setPosition(EffectManager *effect, Config *config)
{
    for (int i = 0; i < effect->totalParticles; i++)
    {
        effect->particles[i].x = GetRandomValue(effect->particles[i].radius, config->width - effect->particles[i].radius);
        effect->particles[i].y = GetRandomValue(effect->particles[i].radius, config->height - effect->particles[i].radius);
        effect->particles[i].movementX = GetRandomValue(-2, 2);
        effect->particles[i].movementY = GetRandomValue(-2, 2);
    }
}

// --------------------------------------------------------------------------------------------------------------------

void updateEffect(EffectManager *effect, Mouse *mouse, Config *config)
{

    for (int32_t i = 0; i < effect->totalParticles; i++)
    {
        if (IsMouseButtonDown(MOUSE_BUTTON_LEFT))
        {
            updateMouse(mouse);
            pushParticle(&effect->particles[i], mouse);
        }

        moveParticle(&effect->particles[i]);
        collision(&effect->particles[i], config);
    }

    findConnections(effect);
    calculateColors(effect);
}

// --------------------------------------------------------------------------------------------------------------------
