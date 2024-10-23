#include "def.h"

void updateCounter(fpsData *fps)
{
    fps->currentFps = GetFPS();
    sprintf(fps->fpsString, "%d", fps->currentFps);
}