#include <iostream>
#include <map>
#include <bits/stdc++.h>

using namespace std;

int main()
{

    // Take names for each person from the user and combine those names.
    cout << "Enter a name for person 1\n";
    string name1{};
    cin >> name1;

    cout << "Enter a name for person 2\n";
    string name2{};
    cin >> name2;

    auto completed{name1 + "loves" + name2};

    // Total the frequency of each character found in the completed phrase.
    map<char, int> characters{};

    for (auto i = 0; i < completed.size(); i++)
    {
        // Convert each character to lower-case.
        auto c = (char)tolower(completed[i]);

        // Skip non a-z characters provide.
        if (c < 'a' || c > 'z')
        {
            continue;
        }

        // Add the character to the map count.
        if (characters.find(c) == characters.end())
        {
            characters[c] = 1;
        }
        else
        {
            characters[c]++;
        }
    }

    vector<int> values{};
    for (auto const &[key, val] : characters)
    {
        values.push_back(val);
    }

    // Pair off all the values together until we get 2 values remaining.
    while (values.size() > 2)
    {
        auto i = 0;
        while (i < values.size() && values.size() > 2)
        {
            values[i] += values[values.size() - 1];
            values.erase(values.begin() + values.size() - 1);
            i++;
        }
    }

    auto total = (values[0] * 10) + values[1] / 10;
    cout << total << "% match!!!\n";
}
