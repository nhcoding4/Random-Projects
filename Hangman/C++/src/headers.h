#include <iostream>
#include <vector>
using namespace std;

// --------------------------------------------------------------------------------------------------------------------

// Structure definitions.
typedef struct
{
    string word;
    string partial_word;
    string art;
    vector<string> guesses;
    int lives;

} game;

typedef struct
{
    int wins;
    int losses;
} scoreboard;

// --------------------------------------------------------------------------------------------------------------------

// Function definitions.
game game_round();
int user_input_int(const string prompt, const int limit);
string art(const int option);
string user_input(const string prompt);
string word_bank();
void print_score(const scoreboard *scores);

// --------------------------------------------------------------------------------------------------------------------