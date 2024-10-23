#include "./headers.h"
#include <iostream>
#include <vector>

using namespace std;

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    scoreboard score_keeper{};

    while (1)
    {
        // Clears the screen.
        cout << "\033[2J\033[1;1H";

        auto new_round{game_round()};

        while (1)
        {
            // Display the current game stage:
            if (new_round.art.length() > 0)
            {
                cout << new_round.art << endl;
            }

            // Display the amount of current charcaters the user has guessed and lives.
            cout << endl
                 << new_round.partial_word << endl
                 << "Lives: " << new_round.lives << endl;

            // If there have been any incorrect guesses, display them.
            if (new_round.guesses.size() > 0)
            {
                cout << "Incorrect guesses: ";
                for (auto i = 0; i < new_round.guesses.size(); i++)
                {
                    cout << new_round.guesses[i] << " ";
                }
                cout << endl;
            }

            // Get a guess from the user and compare it to the current word.
            auto guess{user_input("Guess a letter: ")};
            auto found = false;

            for (auto i = 0; i < new_round.word.length(); i++)
            {
                // If the user guessed correctly, change the _ word in the correct position.
                if (guess[0] == new_round.word[i])
                {
                    new_round.partial_word[i] = guess[0];
                    found = true;
                }
            }
            // Incorrect guess, take away a life and add it to the incorrect guesses vector.
            if (!found)
            {
                new_round.lives--;
                new_round.guesses.push_back(guess);
                new_round.art = art(7 - (new_round.lives + 1));
            }

            // Game victory or loss check.
            if (new_round.lives == 0)
            {
                cout << "No lives remaining. YOU LOSE!\n"
                     << "The word was: " << new_round.word << endl;

                score_keeper.losses++;
                cout << new_round.art << endl;
                print_score(&score_keeper);
                break;
            }
            if (new_round.word == new_round.partial_word)
            {
                cout << "You have correctly guessed the word. YOU WIN!\n";
                score_keeper.wins++;
                print_score(&score_keeper);
                break;
            }

            cout << "\033[2J\033[1;1H";
        }

        // Give option to continue
        auto choice{user_input_int("Press 1 to play another round, 2 to Quit.", 2)};
        if (choice == 2)
        {
            cout << "\033[2J\033[1;1H";
            exit(1);
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Creates a struct that contains all the information the game needs for a new round.
game game_round()
{
    // Select a random word.
    auto selected_word{word_bank()};

    // Create a string of underscores equal to the length of the selected word.
    string underscore_string{};
    for (auto i = 0; i < selected_word.length(); i++)
    {
        underscore_string += "_";
    }

    // Create a new round structure.
    game new_round{};
    new_round.word = selected_word;
    new_round.partial_word = underscore_string;
    new_round.lives = 7;

    return new_round;
}

// --------------------------------------------------------------------------------------------------------------------

// Prompts and takes user input.
string user_input(const string prompt)
{
    while (1)
    {
        cout << prompt << endl;
        string input{};
        cin >> input;

        if (input.length() == 1)
        {
            return input;
        }

        cout << "Invalid input. Please enter a single character\n\n";
    }
}

// --------------------------------------------------------------------------------------------------------------------

// Gets an integer from the user.
int user_input_int(const string prompt, const int limit)
{
    while (1)
    {
        cout << prompt << endl;
        int input{};
        try
        {
            cin >> input;
            if (input >= 1 && input <= limit)
            {
                return input;
            }
        }
        catch (exception &e)
        {
            cout << "Invalid input. ";
        }
        cout << "Please enter a number between 1 and " << limit << endl;
    }
}
// --------------------------------------------------------------------------------------------------------------------

// Prints the score.
void print_score(const scoreboard *scores)
{
    auto wins{scores->wins};
    auto losses{scores->losses};

    cout << "Wins: " << wins << " Losses: " << losses << endl;
}

// --------------------------------------------------------------------------------------------------------------------
