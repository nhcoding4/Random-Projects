#include "../headers/def.hpp"

// --------------------------------------------------------------------------------------------------------------------

std::vector<Connector> find_connections(std::vector<Particle> *particles, Config *config)
{
    std::vector<Connector> connections{};

    for (auto i = 0; i < config->total_particles; i++)
    {
        for (auto j = i; j < config->total_particles; j++)
        {
            if (i == j)
            {
                continue;
            }

            auto dx{(*particles)[i].position.x - (*particles)[j].position.x};
            auto dy{(*particles)[i].position.y - (*particles)[j].position.y};
            auto distance{std::hypot(dx, dy)};

            if (distance <= config->connection_distance)
            {
                Connector new_connector{
                    .start{(*particles)[i].position},
                    .end{(*particles)[j].position},
                    .color{config->connector_color},
                    .highlight{WHITE},
                };

                new_connector.color.a = std::round(255 * (1 - (distance / config->connection_distance)));
                new_connector.highlight.a = new_connector.color.a;

                connections.push_back(new_connector);
            }
        }
    }

    return connections;
}

// --------------------------------------------------------------------------------------------------------------------

void draw_connections(std::vector<Connector> *connections, Config *config)
{
    for (auto &connector : *connections)
    {
        DrawLineEx(
            connector.start,
            connector.end,
            config->connection_thickness,
            connector.highlight);

        DrawLineEx(
            connector.start,
            connector.end,
            config->connection_thickness,
            connector.color);
    }
}

// --------------------------------------------------------------------------------------------------------------------
