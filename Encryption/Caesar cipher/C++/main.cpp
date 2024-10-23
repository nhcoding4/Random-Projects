#include <iostream>

using namespace std;

// --------------------------------------------------------------------------------------------------------------------

int integer_input(string prompt, int limit);
string string_input(string prompt);
string encode_decode(string input, int offset, int option);
void menu_wrapper(string prompt, int option);

// --------------------------------------------------------------------------------------------------------------------

int main()
{
    while (1)
    {
        // Display options.
        cout << "CAESAR'S CIPHER\n------------\n"
             << "\n1)Encode\n2)Decode\n3)Exit\n\n";

        // ----------

        // Take user choice.
        constexpr auto menu_options{3};
        auto selection{integer_input("Select an option:", menu_options)};

        // ----------

        switch (selection)
        {
        case 1:
        {
            // Encode a message.
            constexpr auto encode_prompt{"Enter a message to encode"};
            constexpr auto encode_option{1};
            menu_wrapper(encode_prompt, encode_option);
            break;
        }

            // ----------

        case 2:
        {
            // Decode a message.
            constexpr auto decode_prompt{"Enter a message to decode"};
            constexpr auto decode_option{2};
            menu_wrapper(decode_prompt, decode_option);
            break;
        }

            // ----------

        case 3:
        {
            // Exit the program.
            cout << "\033[2J\033[1;1H";
            exit(1);
        }
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

string string_input(string prompt)
// Takes a string from the user.
{
    // Prompt the user for an input and store said input
    cout << prompt << endl;
    string input{};
    getline(cin, input);

    // ----------

    // Check for empty strings, keep prompting until something has been entered.
    while (input.empty())
    {
        cout << "Enter a non empty string.\n";
        cout << prompt;
        getline(cin, input);
    }

    // ----------

    return input;
}

// --------------------------------------------------------------------------------------------------------------------

int integer_input(string prompt, int limit)
// Takes an interger from the user.
{
    string input;
    while (1)
    {
        // Take input from the user.
        cout << prompt << endl;
        getline(cin, input);

        // ----------

        // Attempt to convert input to integer between 1 and the defined limit.
        try
        {
            auto number{stoi(input)};

            if (number < 1 || number > limit)
            {
                // The user has entered a number too high or low.
                cout << "Invalid input, please enter a number between 1 and " << limit << endl;
            }
            else
            {
                return number;
            }
        }
        catch (exception &err)
        {
            // A non integer has been provided by the user.
            cout << "Invalid input. That input is not a valid integer.\n";
        }
    }
}

// --------------------------------------------------------------------------------------------------------------------

string encode_decode(string input, int offset, int option)
// Encodes or decodes a string based upon the option fed to the function. 1 for encode, 2 for decode
{
    string converted{};
    for (int i = 0; i < input.length(); i++)
    {
        // Current character of the string we are working with.
        auto character = input[i];

        // Add non A-Z characters back into the converted string as they are.
        if (character < 'A' || character > 'Z')
        {
            if (character < 'a' || character > 'z')
            {
                converted += character;
                continue;
            }
        }

        // ----------

        // Set limits. If we move outside of these character limits when encoding/decoding we will move the counter back
        // to the limit.
        auto lower{'a'};
        auto upper{'z'};
        if (character >= 'A' && character <= 'Z')
        {
            lower = 'A';
            upper = 'Z';
        }

        // ----------

        // Encode or decode the string depending on the option fed into the function.
        switch (option)
        {
        case 1:
            // Encode a character.
            for (int j = 0; j < offset; j++)
            {
                character++;
                if (character > upper)
                {
                    character = lower;
                }
            }
            break;

            // ----------

        case 2:
            // Decode a character.
            for (int j = 0; j < offset; j++)
            {
                character--;
                if (character < lower)
                {
                    character = upper;
                }
            }
            break;

            // ----------
        }

        // Add the encoded/decoded character to the string.
        converted += character;
    }

    // ----------

    return converted;
}

// --------------------------------------------------------------------------------------------------------------------

void menu_wrapper(string prompt, int option)
// Wraps a bunch of repeat code into a function, used for code reduction.
{
    // Take the message to encode from the user.
    auto user_string{string_input(prompt)};

    // ----------

    // Take the offset from the user.
    constexpr auto integer_prompt{"Enter an offset"};
    constexpr auto max_offset{__INT_MAX__};
    auto offset{integer_input(integer_prompt, max_offset)};

    // ----------

    // Encode or decode the string
    auto changed_message{encode_decode(user_string, offset, option)};
    cout << changed_message << endl;

    // ----------
    // Pause the program until the user wants to continue
    cout << "Press any key to continue...\n";
    system("read");

    // ----------

    // Clear the screen
    cout << "\033[2J\033[1;1H";
}

// --------------------------------------------------------------------------------------------------------------------
