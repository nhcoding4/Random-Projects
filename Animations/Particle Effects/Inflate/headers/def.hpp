#ifndef DEFINITIONS
#define DEFININTIONS

// --------------------------------------------------------------------------------------------------------------------

#include <raylib.h>
#include <string>
#include <vector>
#include <bits/stdc++.h>
#include <cmath>

// --------------------------------------------------------------------------------------------------------------------

typedef struct
{
    int width, height, target_fps, total_particles;
    float mouse_radius;
    std::string title, fps_text;
    Color particle_color;
} Config;

typedef struct
{
    float radius, starting_radius;
    Vector2 position, movement;
    Color color;
} Particle;

typedef struct
{
    Vector2 position;
    float radius;
    bool active;
} Mouse;

// --------------------------------------------------------------------------------------------------------------------

void init_window(Config *config);
void draw_fps_counter(Config *config);
void create_particle(std::vector<Particle> *particles, Config *config);
void resize_window(std::vector<Particle> *particles, Config *config);
void update_fps_counter(Config *config);

void draw_particles(std::vector<Particle> *particles);
void set_position(std::vector<Particle> *particles, Config *config);
void update_particles(std::vector<Particle> *particles, Config *config);

Mouse init_mouse(Config *config);
void inflate_particles(Mouse *mouse, std::vector<Particle> *particles);
void update_mouse(Mouse *mouse);

// --------------------------------------------------------------------------------------------------------------------

#endif