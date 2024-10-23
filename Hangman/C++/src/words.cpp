#include <iostream>
#include <fstream>
#include <vector>

using namespace std;

string word_bank()
{
    // Vector to store each word.
    vector<string> all_words{};

    // Open the file.
    ifstream words_file("./words.txt");

    // Buffer to store the current word
    string current_word{};

    // Read words and add to vector.
    if (words_file.is_open())
    {
        while (getline(words_file, current_word))
        {
            all_words.push_back(current_word);
        }
        words_file.close();
    }
    // Return a set word so the program doesn't crash if the file cannot be opened.
    else
    {
        cout << "There was an error opening the word bank file - test word returned.\n";
        return "apples";
    }

    // Select and return a random word.
    auto random_index{rand() % all_words.size()};
    return all_words[random_index];
}