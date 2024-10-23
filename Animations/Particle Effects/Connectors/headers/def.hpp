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
    float connection_distance, connection_thickness, mouse_radius;
    std::string title, fps_text;
    Color particle_color, connector_color;
} Config;

typedef struct
{
    float radius, friction;
    Vector2 position, movement, push;
    Color color;
} Particle;

typedef struct
{
    Vector2 start, end;
    Color color, highlight;
} Connector;

typedef struct
{
    Vector2 position;
    float radius;
    bool active;
} Mouse;

// --------------------------------------------------------------------------------------------------------------------

std::vector<Connector> find_connections(std::vector<Particle> *particles, Config *config);
void draw_connections(std::vector<Connector> *connections, Config *config);

void init_window(Config *config);
void draw_fps_counter(Config *config);
void create_particle(std::vector<Particle> *particles, Config *config);
void resize_window(std::vector<Particle> *particles, Config *config);
void update_fps_counter(Config *config);

void draw_particles(std::vector<Particle> *particles);
void set_position(std::vector<Particle> *particles, Config *config);
void update_particles(std::vector<Particle> *particles, Config *config);

Mouse init_mouse(Config *config);
void push_particles(std::vector<Particle> *particles, Mouse *mouse);
void update_mouse(Mouse *mouse);

// --------------------------------------------------------------------------------------------------------------------

#endif