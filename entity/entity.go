// Package entity describes all
// base structs that might need
// to be shared between packages
// which would normally create
// cyclic dependencies.
package entity

// Entity defines the position and sprite
// of an in-game object
type Entity struct {
	Idx    int
	Sprite int
}
