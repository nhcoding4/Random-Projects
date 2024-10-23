#include <iostream>
#include <vector>

using namespace std;

// --------------------------------------------------------------------------------------------------------------------

// Function signatures.

int take_input_int(string prompt);
vector<long long> calculate_sequence(int total);

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    auto numbers{take_input_int("How many Fionacci numbers would you like to calculate?\n")};
    auto sequence{calculate_sequence(numbers)};
    auto counter{0};

    for (auto &i : sequence)
    {
        cout << i << " ";
        counter++;
        if (counter == 10)
        {
            cout << endl;
            counter = 0;
        }
    }

    cout << endl;
}

// --------------------------------------------------------------------------------------------------------------------

int take_input_int(string prompt)
{
    while (1)
    {
        cout << prompt;
        try
        {
            int input{};
            cin >> input;

            if (input < 1)
            {
                cout << "Please enter a positive integer greater than 1\n";
                continue;
            }

            return input;
        }
        catch (exception &e)
        {
            cout << "Invalid input. Please enter a number greater than 1\n";
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

vector<long long> calculate_sequence(int total)
{
    vector<long long> sequence{1, 1};

    for (auto i = 0; i < total; i++)
    {
        auto new_number = sequence[i] + sequence[i + 1];
        sequence.push_back(new_number);
    }

    return sequence;
}

// --------------------------------------------------------------------------------------------------------------------
