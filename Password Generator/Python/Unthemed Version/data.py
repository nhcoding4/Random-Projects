import json
import os
import random
import tkinter
from random import shuffle
from tkinter import *
from tkinter import messagebox
from typing import List

import pyperclip

# ---------------------------------------------------------------------------------------------------------------------

# Constants
PATH_FOLDER: str = os.path.join(os.getcwd(), "passwords")
PATH_JSON: str = os.path.join(PATH_FOLDER, "passwords.json")
# ---------------------------------------------------------------------------------------------------------------------


def copy_password(password: str) -> None:
    """Copies text contained inside the textbox to the system's clipboard or displays an error."""

    if len(password) > 1:
        pyperclip.copy(password)
        tkinter.messagebox.showinfo(
            message="Password has been copied to your clipboard."
        )
    else:
        tkinter.messagebox.showerror(message="No password has been generated.")


# ---------------------------------------------------------------------------------------------------------------------


def edit_json(
    key_elements: List[tkinter.Entry], value_elements: List[tkinter.Entry]
) -> None:
    """Attempts to create a new json file from edited data."""

    data: dict = {}
    for i, _ in enumerate(key_elements):
        if key_elements[i].get() in data.keys():
            tkinter.messagebox.showerror(
                title="Error",
                message=f"Repeat entries. Please rename {key_elements[i].get()}.",
            )
            return

        data[key_elements[i].get()] = value_elements[i].get()

    with open(PATH_JSON, "w") as file:
        json.dump(data, file, ensure_ascii=False, indent=4)
        tkinter.messagebox.showinfo(
            title="Success", message="Data has been successfully edited."
        )


# ---------------------------------------------------------------------------------------------------------------------


def generate_password(textbox: tkinter.Text) -> None:
    """Randomly selects characters, numbers and symbols to form a password and injects it into the main textbox."""

    # Password generation.

    alphabetical: List[str] = [
        (
            chr(random.randint(65, 90))
            if random.randint(0, 1) == 0
            else chr(random.randint(97, 122))
        )
        for _ in range(6)
    ]

    numerical: List[str] = [
        chr(random.randint(48, 57)) for _ in range(random.randint(2, 4))
    ]

    symbols: List[str] = [
        chr(random.randint(60, 64)) for _ in range(random.randint(2, 4))
    ]

    characters: List[str] = alphabetical + numerical + symbols
    shuffle(characters)
    password: str = "".join(characters)

    # -----------------------------------------------------

    # Inject the password into the textbox.

    textbox.delete("1.0", END)
    textbox.insert(INSERT, chars=password)


# ---------------------------------------------------------------------------------------------------------------------


def init() -> str:
    """Initialization checks. If we don't have a passwords folder or file then we create it on application launch."""

    try:

        if not os.path.exists(PATH_FOLDER):
            os.mkdir(PATH_FOLDER)

        if not os.path.exists(PATH_JSON):
            with open(PATH_JSON, "w") as file:
                file_init = {}
                json.dump(file_init, file)

        return ""

    except Exception as exception:
        return str(exception)


# ---------------------------------------------------------------------------------------------------------------------


def load_json() -> json:
    """Attempts to load the json file at the expected storage path. If it fails then attempt to create a new one."""

    with open(PATH_JSON, "r", encoding="utf-8") as file:
        try:
            return json.load(file)
        except Exception as e:
            tkinter.messagebox.showerror(
                title="Error",
                message=str(f"Error loading file:{e}\nAttempting to recreate."),
            )
            init()


# ---------------------------------------------------------------------------------------------------------------------


def save_json(name: str, password: str, pop_up: tkinter.Toplevel) -> None:
    """Appends a name:password pair to a Json file."""

    if len(name) <= 1:
        tkinter.messagebox.showerror(title="Error", message="No name has been entered.")
        return

    name_new = name.replace("\n", "")
    password_new = password.replace("\n", "")

    # -----------------------------------------------------

    # Append data to Json and close popup.

    json_data = load_json()

    if name_new in json_data.keys():
        tkinter.messagebox.showerror(
            title="Error", message="Name already exists. Please enter a different name."
        )
        return

    json_data[name_new] = password_new

    with open(PATH_JSON, "w") as json_file:
        try:
            json.dump(json_data, json_file, ensure_ascii=False, indent=4)
            tkinter.messagebox.showinfo(
                title="Saved", message=f"Data has successfully been saved."
            )
            pop_up.destroy()
        except Exception as e:
            tkinter.messagebox.showerror(title="Error", message=str(e))


# ---------------------------------------------------------------------------------------------------------------------
