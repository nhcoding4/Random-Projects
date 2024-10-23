import pyray as pr
import os
from ctypes import *


class Sounds:
    def __init__(self) -> None:
        self.init_audio()
        self.music: pr.Music = pr.load_music_stream(os.getcwd() + "/Sounds/music.mp3")
        self.rotate_sound: pr.Sound = pr.load_sound(os.getcwd() + "/Sounds/rotate.mp3")
        self.clear_sound: pr.Sound = pr.load_sound(os.getcwd() + "/Sounds/clear.mp3")

        self.play_background_music()

    # -------------------------------------------------------------------------

    def init_audio(self) -> None:
        """Sets up the audio device. Must be ran before any sound functions are used."""

        pr.init_audio_device()

    # -------------------------------------------------------------------------

    def play_background_music(self) -> None:
        """Plays background music."""

        pr.play_music_stream(self.music)

    # -------------------------------------------------------------------------

    def play_clear_sound(self) -> None:
        """Plays the rotate sound."""

        pr.play_sound(self.clear_sound)

    # -------------------------------------------------------------------------

    def play_rotate_sound(self) -> None:
        """Plays the rotate sound."""

        pr.play_sound(self.rotate_sound)

    # -------------------------------------------------------------------------

    def close_audio_device(self) -> None:
        """Closes the auto device. Run before closing the program."""

        pr.unload_music_stream(self.music)
        pr.unload_sound(self.rotate_sound)
        pr.unload_sound(self.clear_sound)
        pr.close_audio_device

    # -------------------------------------------------------------------------
