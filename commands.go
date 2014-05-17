package main

import "encoding/json"

type Command uint8

const (
	Fight Command = iota
	ItemCmd
	Magic
	Morph
	Revert
	Steal
	Capture
	Swdtech
	Throw
	Tools
	BlitzCmd
	Runic
	LoreCmd
	Sketch
	Control
	Slot
	Rage
	Leap
	Mimic
	Dance
	Row
	Def
	Jump
	XMagic
	GP_Rain
	Summon
	Health
	Shock
	Possess
	MagiTek
)

var cmdName = map[Command]string{
	Fight:    "Fight",
	ItemCmd:  "Item",
	Magic:    "Magic",
	Morph:    "Morph",
	Revert:   "Revert",
	Steal:    "Steal",
	Capture:  "Capture",
	Swdtech:  "Swdtech",
	Throw:    "Throw",
	Tools:    "Tools",
	BlitzCmd: "Blitz",
	Runic:    "Runic",
	LoreCmd:  "Lore",
	Sketch:   "Sketch",
	Control:  "Control",
	Slot:     "Slot",
	Rage:     "Rage",
	Leap:     "Leap",
	Mimic:    "Mimic",
	Dance:    "Dance",
	Row:      "Row",
	Def:      "Def",
	Jump:     "Jump",
	XMagic:   "XMagic",
	GP_Rain:  "GP_Rain",
	Summon:   "Summon",
	Health:   "Health",
	Shock:    "Shock",
	Possess:  "Possess",
	MagiTek:  "MagiTek",
}

func (c Command) MarshalJSON() ([]byte, error) {
	return json.Marshal(cmdName[c])
}
