import os
import os.path
import sys
from typing import List


def main() -> None:

    # Check for stupid cmdline input
    if not 2 <= len(sys.argv):
        print("Stupid input detected. Aborting.")
        return

    # Invoke options.
    if sys.argv[1] == "init":
        init()
    elif sys.argv[1] == "build":
        build()
    else:
        print(f"{sys.argv[1]} not recognized.")


# ---------------------------------------------------------------------------------------------------------------------


def build() -> None:
    """
    Attempts to build a C++ project using source files found in a /src/ directory and with options from cmdline and
    a textfile called options.txt.
    """

    # Get all file names in the src directory and save them to a list.

    # Get all C++ source files in the 'src' directory.
    try:
        files: List[str] = os.listdir("./src/")
        file_names: List[str] = []
        for file in files:
            if ".cpp" or ".h" in file:
                file_names.append(f"./src/{file} ")

    # No src directory detected
    except Exception as e:
        print("error: ", e)
        return

    # No files in the src directory.
    if len(file_names) == 0:
        print("No files found")
        return

    # ---------------------------------------------------------------

    # Build a valid command and pass it to the operating system to execute.

    # Make a command:
    command: str = f"g++ -o {sys.argv[2]}"
    files_used: str = "".join(f" {file}" for file in file_names)

    # Add in any extra command line arguments passed to the program.
    if len(sys.argv) > 3:
        options: str = "".join(f"{ sys.argv[i]}" for i in range(3, len(sys.argv) - 1))
    else:
        options: str = ""

    command += files_used + options

    # ---------------------------------------------------------------

    # Look for persistent options in a textfile.
    path: str = os.path.join(os.getcwd(), "options.txt")

    if not os.path.exists(path):
        pass

    else:
        with open("options.txt", "r") as file:
            text_options: List[str] = file.readlines()
            additional_options: str = "".join(f" {arg}" for arg in text_options)
            command += additional_options

    # ---------------------------------------------------------------

    # Run command:
    try:
        os.system(command)
        print("built")
    except Exception as e:
        print("error: ", e)


# ---------------------------------------------------------------------------------------------------------------------


def init() -> None:
    """Creates the file structure and main source file for a new C++ project."""

    contents: str = """#include <iostream>

using namespace std;

int main()
{

}

"""
    # Check if the path exists, if not make it.

    path: str = os.path.join(os.getcwd(), "src")
    if not os.path.exists(path):
        os.mkdir(path)

    # Create a file.
    main_path: str = os.path.join(path, "main.cpp")
    if not os.path.exists(main_path):
        with open(main_path, "w") as file:
            file.writelines(contents)
        print("Project has been started.")
    else:
        print("main.cpp already exists.")


# ---------------------------------------------------------------------------------------------------------------------

if __name__ == "__main__":
    main()
