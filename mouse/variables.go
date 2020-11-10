package mouse

// Trying out some global variables for real.
// Because these are constrained to this package,
// I don't see any benefit coupling them to the
// overarching struct, despite it making sense
// to do so. Let's see how this goes.

// This cluster of variables
// help with (de)selecting walls.
var startPoint = -1
var endPoint = -1
var hasClicked = false
var firstClickedSprite = -1
