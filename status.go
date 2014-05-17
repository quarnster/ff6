package main

type Status uint8

const (
	Darkness Status = 1 << iota
	Zombie
	Poison
	Magitek
	Invisible
	Imp
	Stone
	Wounded
)
