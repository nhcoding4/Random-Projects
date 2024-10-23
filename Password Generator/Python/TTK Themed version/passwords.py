import tkinter.messagebox
from tkinter import ttk
from ttkthemes import ThemedTk

import PIL.Image
from PIL import ImageTk

from data import *

# ---------------------------------------------------------------------------------------------------------------------

# Constants.

GUI_FONT: tuple[str, int] = ("Arial", 20)
IMAGE_LOCATION: str = os.path.join(os.getcwd(), "key_clear.png")
Y_PAD: int = 5
X_PAD: int = 1


# ---------------------------------------------------------------------------------------------------------------------


def main():
    err: str = init()
    gui(err)


# ---------------------------------------------------------------------------------------------------------------------


def gui(error_status: str) -> None:
    """Main Window of the GUI."""

    # Root Window.

    window_size: tuple[int, int] = (400, 250)
    root = ThemedTk(theme="arc", background=True)
    root.geometry(f"{window_size[0]}x{window_size[1]}")
    root.minsize(window_size[0], window_size[1])
    root.title("Password Manager")

    if not error_status == "":
        tkinter.messagebox.showerror(
            message="Cannot generate required folders and files for saving passwords. Some functionality will not work."
        )

    # -----------------------------------------------------

    # Loads, resizes and displays the title image.

    frame_size: tuple[int, int] = (
        int((window_size[0] / 100) * 60),
        int((window_size[1] / 100) * 40),
    )
    title_frame = ttk.Frame(master=root, width=frame_size[0], height=frame_size[1])
    title_frame.pack(padx=X_PAD, pady=Y_PAD)

    data = PIL.Image.open(IMAGE_LOCATION)
    resized_image = data.resize((int(data.width * 0.4), int(data.height * 0.4)))
    img = ImageTk.PhotoImage(resized_image)
    image_label = ttk.Label(master=title_frame, image=img)
    image_label.pack()

    # ------------------------------------------------------------

    # Texbox, buttons.

    mid_size: tuple[int, int] = (
        int((window_size[0] / 100) * 75),
        int((window_size[1] / 100) * 40),
    )
    mid_frame = ttk.Frame(master=root, width=mid_size[0], height=mid_size[1])
    mid_frame.pack(padx=X_PAD, pady=Y_PAD)

    textbox = tkinter.Text(
        mid_frame, width=int((mid_size[0] / 100) * 75), height=1, font=GUI_FONT
    )
    textbox.pack()

    button_frame = ttk.Frame(
        master=mid_frame, width=mid_size[0], height=mid_size[1] / 2
    )
    button_frame.pack()

    clipboard_button = ttk.Button(
        master=button_frame,
        text="Copy",
        command=lambda: copy_password(textbox.get("1.0", END)),
    )
    clipboard_button.grid(row=0, column=1, pady=Y_PAD, padx=X_PAD)

    save_button = ttk.Button(
        master=button_frame,
        text="Save",
        command=lambda: add_password_frame(root, textbox.get("1.0", END)),
    )
    save_button.grid(row=0, column=0)

    generate_button = ttk.Button(
        master=button_frame, text="Generate", command=lambda: generate_password(textbox)
    )
    generate_button.grid(row=0, column=2, pady=Y_PAD, padx=X_PAD)

    view_edit_button = ttk.Button(
        master=button_frame, text="Edit/View", command=lambda: view_edit_frame(root)
    )
    view_edit_button.grid(row=0, column=3, pady=Y_PAD, padx=X_PAD)

    # ------------------------------------------------------------

    root.mainloop()


# ---------------------------------------------------------------------------------------------------------------------


def add_password_frame(root, password: str) -> None:
    """Password popup frame used for saving password data to a json file."""

    if len(password) <= 1:
        tkinter.messagebox.showerror(message="No password has been generated.")
        return

    # -----------------------------------------------------

    # Popup Window when user clicks save from the main window.

    pop_up_size: tuple[int, int] = (300, 110)
    pop_up = tkinter.Toplevel(root)
    pop_up.geometry(f"{pop_up_size[0]}x{pop_up_size[1]}")
    pop_up.title("Save")

    pop_up_frame = ttk.Frame(master=pop_up, width=pop_up_size[0], height=pop_up_size[1])
    pop_up_frame.pack(padx=X_PAD, pady=Y_PAD)

    pop_up_label = ttk.Label(master=pop_up_frame, text="Enter A Name:")
    pop_up_label.pack()

    text_box = tkinter.Text(
        master=pop_up_frame,
        width=int((pop_up_size[0] / 100) * 80),
        height=1,
        font=GUI_FONT,
    )
    text_box.pack(pady=Y_PAD)

    # -----------------------------------------------------

    # Subframe containing buttons organized in a grid.

    button_frame = ttk.Frame(
        master=pop_up_frame, width=pop_up_size[0], height=int(pop_up_size[1] / 2)
    )
    button_frame.pack()

    close_button = ttk.Button(master=button_frame, text="Close", command=pop_up.destroy)
    close_button.grid(row=0, column=0, padx=X_PAD)

    save_button = ttk.Button(
        master=button_frame,
        text="Save",
        command=lambda: save_json(text_box.get("1.0", END), password, pop_up),
    )
    save_button.grid(row=0, column=1, padx=X_PAD)


# ---------------------------------------------------------------------------------------------------------------------


def view_edit_frame(root) -> None:
    """Spawns the view and edit frame which allows the user to view or edit all the passwords saved."""

    # View/Edit window root.

    json_data: json = load_json()

    view_frame_size: tuple[int, int] = (600, 300)
    view_frame = tkinter.Toplevel(root)
    view_frame.geometry = f"{view_frame_size[0]}x{view_frame_size[1]}"
    view_frame.title("Edit/View")

    # ------------------------------------------------------------

    # Main frame containing all other frames and labels for the info grid.

    main_frame = ttk.Frame(
        master=view_frame, width=view_frame_size[0], height=view_frame_size[1]
    )
    main_frame.pack(padx=X_PAD, pady=Y_PAD)

    label_grid = ttk.Frame(master=main_frame, width=view_frame_size[0], height=1)
    label_grid.pack()

    label_name = ttk.Label(master=label_grid, text="Name", font=("Arial", 12))
    label_name.grid(row=0, column=0, padx=50)
    label_value = ttk.Label(master=label_grid, text="Value", font=("Arial", 12))
    label_value.grid(row=0, column=4, padx=50)

    # ------------------------------------------------------------

    # Display the saved data and allow the user to view, edit and save entries.

    data_frame = tkinter.Canvas(
        master=main_frame,
        width=int(view_frame_size[0] * 0.90),
        height=int(view_frame_size[1] * 0.80),
    )
    data_frame.pack(pady=Y_PAD)

    row: int = 0
    key_elements: List[ttk.Entry] = []
    value_elements: List[ttk.Entry] = []

    for key, value in json_data.items():
        new_name = ttk.Entry(master=data_frame)
        new_name.insert(INSERT, key)
        new_name.grid(row=row, column=0)

        new_value = ttk.Entry(master=data_frame)
        new_value.insert(INSERT, value)
        new_value.grid(row=row, column=1)

        key_elements.append(new_name)
        value_elements.append(new_value)
        row += 1

    scrollbar = ttk.Scrollbar(
        master=data_frame, orient="vertical", command=data_frame.yview
    )
    data_frame.bind(
        "<MouseWheel>",
        lambda event: data_frame.yview_scroll(-int(event.delta / 60), "units"),
    )
    data_frame.configure(yscrollcommand=scrollbar.set)
    scrollbar.place(relx=1, rely=0, relheight=1, anchor="ne")

    # ------------------------------------------------------------

    # Container holding buttons.

    button_frame = ttk.Frame(
        master=main_frame,
        width=int((view_frame_size[0] / 100) * 90),
        height=int((view_frame_size[1] / 100) * 20),
    )
    button_frame.pack(padx=X_PAD, pady=Y_PAD)

    exit_button = ttk.Button(
        master=button_frame, text="Exit", command=view_frame.destroy
    )
    exit_button.grid(row=0, column=0)

    save_button = ttk.Button(
        master=button_frame,
        text="Save",
        command=lambda: edit_json(key_elements, value_elements),
    )
    save_button.grid(row=0, column=1)


# ---------------------------------------------------------------------------------------------------------------------


if __name__ == "__main__":
    main()

# ---------------------------------------------------------------------------------------------------------------------
