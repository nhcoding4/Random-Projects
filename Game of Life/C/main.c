#include <raylib.h>
#include <stdlib.h>
#include <stdio.h>

typedef struct
{
    int width, height, target_fps;
    float cell_size;
    char *title;
    Color background;
} Config;

typedef struct
{
    int x, y;
} Offset;

typedef struct
{
    int rows, columns;
    Offset offsets[8];
    int **grid;
} Grid;

void free_grid(Grid *grid);
void populate_grid(Grid *grid, Config *config);

int main()
{
    Config config = {
        .width = 1900,
        .height = 1000,
        .target_fps = 144,
        .cell_size = 3,
        .title = "Game of Life C",
        .background = BLACK,
    };

    Grid grid = {
        .rows = 0,
        .columns = 0,
        .offsets = {
            {-1, 0},
            {1, 0},
            {0, -1},
            {0, 1},
            {-1, -1},
            {-1, 1},
            {1, -1},
            {1, 1},
        },
    };
    populate_grid(&grid, &config);

    SetConfigFlags(FLAG_WINDOW_RESIZABLE);
    SetConfigFlags(FLAG_MSAA_4X_HINT);
    SetConfigFlags(FLAG_VSYNC_HINT);
    SetConfigFlags(FLAG_WINDOW_HIGHDPI);

    InitWindow(config.width, config.height, config.title);
    SetTargetFPS(config.target_fps);

    while (!WindowShouldClose())
    {
        // Resize Window
        if (IsWindowResized())
        {
            config.width = GetScreenWidth();
            config.height = GetScreenHeight();
            populate_grid(&grid, &config);
        }

        // Update grid
        int **new_grid = malloc(sizeof(int *) * grid.rows);

        for (int i = 0; i < grid.rows; i++)
        {
            int *new_row = malloc(sizeof(int) * grid.columns);

            for (int j = 0; j < grid.columns; j++)
            {
                int total_live_neighbours = 0;

                for (int x = 0; x < 8; x++)
                {
                    int row = (i + grid.offsets[x].y + grid.rows) % grid.rows;
                    int column = (j + grid.offsets[x].x + grid.columns) % grid.columns;

                    total_live_neighbours += grid.grid[row][column];
                }

                if (total_live_neighbours == 3)
                {
                    new_row[j] = 1;
                }
                else if (total_live_neighbours == 2)
                {
                    new_row[j] = grid.grid[i][j];
                }
                else
                {
                    new_row[j] = 0;
                }
            }
            new_grid[i] = new_row;
        }

        free_grid(&grid);
        grid.grid = new_grid;

        // Update FPS counter
        int fps = GetFPS();
        char buffer[5];
        sprintf(buffer, "%d", fps);

        BeginDrawing();
        ClearBackground(config.background);

        // Draw fps counter
        DrawText(buffer, 0, 0, 40, GREEN);

        // Draw Grid
        for (int i = 0; i < grid.rows; i++)
        {
            for (int j = 0; j < grid.columns; j++)
            {
                Color color = config.background;
                if (grid.grid[i][j] == 1)
                {
                    color = WHITE;
                }

                DrawRectangle(
                    config.cell_size * j + 1,
                    config.cell_size * i + 1,
                    config.cell_size - 1,
                    config.cell_size - 1,
                    color);
            }
        }
        EndDrawing();
    }
    CloseWindow();
    free_grid(&grid);
}

void free_grid(Grid *grid)
{
    for (int i = 0; i < grid->rows; i++)
    {
        free(grid->grid[i]);
    }

    free(grid->grid);
}

void populate_grid(Grid *grid, Config *config)
{
    if (grid->rows != 0)
    {
        free_grid(grid);
    }

    grid->rows = config->height / config->cell_size;
    grid->columns = config->width / config->cell_size;

    int **new_grid = malloc(sizeof(int *) * grid->rows);
    if (new_grid == NULL)
    {
        printf("Error allocating memory for grid\n");
        exit(1);
    }

    for (int i = 0; i < grid->rows; i++)
    {
        int *new_row = malloc(sizeof(int) * grid->columns);
        if (new_row == NULL)
        {
            printf("Error allocating memory for row\n");
            exit(1);
        }

        for (int j = 0; j < grid->columns; j++)
        {
            if (GetRandomValue(0, 10) < 2)
            {
                new_row[j] = 1;
            }
            else
            {
                new_row[j] = 0;
            }
        }
        new_grid[i] = new_row;
    }

    grid->grid = new_grid;
}