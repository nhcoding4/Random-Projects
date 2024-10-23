#include "def.h"

// --------------------------------------------------------------------------------------------------------------------

void allocateLines(EffectManager *effect)
{
    effect->lines.capacity = 1000;
    effect->lines.total = 0;

    effect->lines.start = (Vector2 *)malloc(sizeof(Vector2) * effect->lines.capacity + 1);
    effect->lines.end = (Vector2 *)malloc(sizeof(Vector2) * effect->lines.capacity + 1);
    effect->lines.color = (Color *)malloc(sizeof(Color) * effect->lines.capacity + 1);
}

// --------------------------------------------------------------------------------------------------------------------

void checkAllocationLines(EffectManager *effect)
{
    if (effect->lines.total == effect->lines.capacity - 1)
    {
        effect->lines.capacity += 1000;
        effect->lines.start = (Vector2 *)realloc(effect->lines.start, sizeof(Vector2) * effect->lines.capacity + 1);
        effect->lines.end = (Vector2 *)realloc(effect->lines.end, sizeof(Vector2) * effect->lines.capacity + 1);
        effect->lines.color = (Color *)realloc(effect->lines.color, sizeof(Color) * effect->lines.capacity + 1);
    }
}

// --------------------------------------------------------------------------------------------------------------------

void deallocateLines(EffectManager *effect)
{
    free(effect->lines.start);
    free(effect->lines.end);
    free(effect->lines.color);
}

// --------------------------------------------------------------------------------------------------------------------

void resizeScreen(Config *config, EffectManager *effect)
{
    if (IsWindowResized())
    {
        config->width = GetScreenWidth();
        config->height = GetScreenHeight();

        setPosition(effect, config);
    }
}

// --------------------------------------------------------------------------------------------------------------------
