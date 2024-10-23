#include <raylib.h>
#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <math.h>
#include <stdio.h>

// --------------------------------------------------------------------------------------------------------------------

typedef struct
{
    int32_t width, height, targetFPS;
    char *title;
} Config;

typedef struct
{
    int32_t x, y, radius;
    Vector2 mouseVector2;
} Mouse;

typedef struct
{
    int32_t x, y, radius, movementX, movementY;
    double pushX, pushY, friction;
    Color color;
} Particle;

typedef struct
{
    Vector2 *start;
    Vector2 *end;
    Color *color;
    int32_t capacity, total;
} Lines;

typedef struct
{
    int32_t totalParticles, connectionDistance, rayDivisor;
    double lineThickness;
    Particle *particles;
    Lines lines;
    Color *colors;
} EffectManager;

typedef struct
{
    int32_t currentFps;
    char fpsString[5];
} fpsData;

// --------------------------------------------------------------------------------------------------------------------

void allocateLines(EffectManager *effect);
void checkAllocationLines(EffectManager *effect);
void deallocateLines(EffectManager *effect);

void calculateColors(EffectManager *effect);
void createParticles(EffectManager *effect, Config *config);
void draw(EffectManager *effect, Mouse *mouse);
void setPosition(EffectManager *effect, Config *config);
void updateEffect(EffectManager *effect, Mouse *mouse, Config *config);

void collision(Particle *particle, Config *config);
void moveParticle(Particle *particle);
void pushParticle(Particle *particle, Mouse *mouse);

void updateMouse(Mouse *mouse);

void updateCounter(fpsData *fps);

void resizeScreen(Config *config, EffectManager *effect);
