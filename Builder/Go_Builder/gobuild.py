# Unlike C++, go has a good, easy to use build system so this will only be for project setup.
import os
import os.path
import sys

# ---------------------------------------------------------------------------------------------------------------------


def main():
    """  This program starts a Golang project """
    #
    if len(sys.argv) != 2:
        print("Error: Please pass a name for this project (no spaces) to the command.")
        return

    # -----------------------------------------------------

    # Create folder, main file and create the go mod file.
    try:
        path: str = os.path.join(os.getcwd(), "Golang")
        sub_path: str = os.path.join(path, "src")
        if not os.path.exists(path):
            os.mkdir(path)
            os.mkdir(sub_path)

        main_path: str = os.path.join(sub_path, "main.go")
        if not os.path.exists(main_path):
            with open(main_path, "w") as file:
                file.writelines(main_file_contents())

        if not os.path.exists(os.path.join(sub_path, "go.mod")):
            command: str = f"go mod init {sys.argv[1]}"
            os.chdir(sub_path)
            os.system(command)

        print(f"Project {sys.argv[1]} successfully started.")

    except Exception as e:
        print("There was an error creating the project structure: ", e)


# ---------------------------------------------------------------------------------------------------------------------


def main_file_contents() -> str:
    """ Contains the initial main file layout. """

    contents: str = """
package main

func main(){
}

"""
    return contents

# ---------------------------------------------------------------------------------------------------------------------



if __name__ == "__main__":
    main()
