#include <iostream>
#include <vector>

using namespace std;

// Prints art for the game.
string art(const int option)
{
      vector<string> ascii_art{
          R"(
  +---+
  |   |
      |
      |
      |
      |
=========)",
          R"(
  +---+
  |   |
  O   |
      |
      |
      |
=========)",
          R"(
  +---+
  |   |
  O   |
  |   |
      |
      |
=========)",
          R"(
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========)",
          R"(
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========)",
          R"(
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========)",
          R"(  
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========)"};

      return ascii_art[option];
}

