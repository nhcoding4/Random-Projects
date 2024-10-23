This is the version I attempted to make on my own before giving in and using a guide for the Python version.

It mostly works with the blocks shifting down, rotation, keeping blocks within bounds but there is an error that randomly crashes the program.
I'm probably attempting to insert into an array out of bounds but after working on this for nearly a week with little progress I decided to look at a tutorial found:
https://youtu.be/wVYKG_ch4yM.

This version differs from the python version as everything is done on the grid. There are no block objects being imposed upon the grid, just changing of cell values on the grid.

The rotation and bounds checking is way less elegant than the python version as there are way more checks going on to prevent errors from blocks moving through themselves. The python version
has a way simpler method and just moves the block back if an invalid move has happened where as this version attempts to scan a move into the future before making the move.

When remaking this I need to:

-Simplify the out of bounds checking.
-Maybe encapsulate more manipulation of grid elements to remove potential errors. Once I have a method that works I can just use that instead of reaching into the structs and editing things over and over.
-Maybe split block and grid logic into their own objects. I still think this can be done just by manipulating the grid itself similar to game of life.
