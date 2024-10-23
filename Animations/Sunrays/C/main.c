#include "def.h"

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    // ----------------------------------------------------------------------------------------------------------------

    // Game Data.

    // ----------------------------------------------------------------------------------------------------------------

    Config config = {
        width : 1000,
        height : 1000,
        targetFPS : 144,
        title : "Sun Rays",
    };

    Mouse mouse = {
        x : config.width / 2,
        y : config.height / 2,
        radius : 250,
    };
    Vector2 mouseVect = {mouse.x, mouse.y};
    mouse.mouseVector2 = mouseVect;

    EffectManager effect = {
        totalParticles : 1000,
        connectionDistance : 100,
        rayDivisor:  3,
        lineThickness :
            1.5,
        lines : {capacity : 1000, 0},
        colors : (Color *)malloc(sizeof(Color) * effect.totalParticles + 1),
    };

    fpsData fpsCounter = {
        currentFps : 0,
    };

    // ----------------------------------------------------------------------------------------------------------------

    // Setup.

    // ----------------------------------------------------------------------------------------------------------------

    SetConfigFlags(FLAG_WINDOW_RESIZABLE);
    InitWindow(config.width, config.height, config.title);
    SetTargetFPS(config.targetFPS);

    // Set effect data.
    allocateLines(&effect);
    createParticles(&effect, &config);

    // ----------------------------------------------------------------------------------------------------------------

    // Mainloop.

    // ----------------------------------------------------------------------------------------------------------------

    while (!WindowShouldClose())
    {
        // ------------------------------------------------------------------------------------------------------------

        // Update.
        resizeScreen(&config, &effect);
        allocateLines(&effect);
        updateEffect(&effect, &mouse, &config);
        updateCounter(&fpsCounter);

        // ------------------------------------------------------------------------------------------------------------

        // Drawing.

        BeginDrawing();
        ClearBackground(BLACK);
        draw(&effect, &mouse);
        DrawText(fpsCounter.fpsString, 0, 0, 40, GREEN);
        EndDrawing();

        deallocateLines(&effect);
    }
    CloseWindow();
    free(effect.particles);
    free(effect.colors);
}

// --------------------------------------------------------------------------------------------------------------------
