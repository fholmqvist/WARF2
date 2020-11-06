package mouse

// Trying out some global variables for real.
// Because these are constrained to this package,
// I don't see any benefit coupling them to the
// overarching struct, despite that this makes sense.
// Let's see how this goes.

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1

// Remembering last frame
// in order to reset selected
// tiles without having to
// redraw the entire screen.
var previousStartPoint = -1
var previousEndPoint = -1
