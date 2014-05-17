package main

type Spell uint8

const (
	Fire Spell = iota
	Blizzard
	Thunder
	PoisonSpell
	Drain
	Fira
	Blizzara
	Thundara
	Bio
	Firaga
	Blizzaga
	Thundaga
	Break
	Death
	Holy
	Flare
	Gravity
	Graviga
	Banish
	Meteor
	Ultima
	Quake
	Tornado
	Meltdown
	Libra
	Slow
	Rasp
	Silence
	Protect
	Sleep
	Confuse
	Haste
	Stop
	Berserk
	Float
	ImpSpell
	Reflect
	Shell
	Vanish
	Hastega
	Slowga
	Osmose
	Teleport
	Quick
	Dispel
	Cure
	Cura
	Curaga
	Raise
	Arise
	Poisona
	Esuna
	Regen
	Reraise
	Flood
)

var spellname = map[Spell]string{
	Fire:        "Fire",
	Blizzard:    "Blizzard",
	Thunder:     "Thunder",
	PoisonSpell: "Poison",
	Drain:       "Drain",
	Fira:        "Fira",
	Blizzara:    "Blizzara",
	Thundara:    "Thundara",
	Bio:         "Bio",
	Firaga:      "Firaga",
	Blizzaga:    "Blizzaga",
	Thundaga:    "Thundaga",
	Break:       "Break",
	Death:       "Death",
	Holy:        "Holy",
	Flare:       "Flare",
	Gravity:     "Gravity",
	Graviga:     "Graviga",
	Banish:      "Banish",
	Meteor:      "Meteor",
	Ultima:      "Ultima",
	Quake:       "Quake",
	Tornado:     "Tornado",
	Meltdown:    "Meltdown",
	Libra:       "Libra",
	Slow:        "Slow",
	Rasp:        "Rasp",
	Silence:     "Silence",
	Protect:     "Protect",
	Sleep:       "Sleep",
	Confuse:     "Confuse",
	Haste:       "Haste",
	Stop:        "Stop",
	Berserk:     "Berserk",
	Float:       "Float",
	ImpSpell:    "Imp",
	Reflect:     "Reflect",
	Shell:       "Shell",
	Vanish:      "Vanish",
	Hastega:     "Hastega",
	Slowga:      "Slowga",
	Osmose:      "Osmose",
	Teleport:    "Teleport",
	Quick:       "Quick",
	Dispel:      "Dispel",
	Cure:        "Cure",
	Cura:        "Cura",
	Curaga:      "Curaga",
	Raise:       "Raise",
	Arise:       "Arise",
	Poisona:     "Poisona",
	Esuna:       "Esuna",
	Regen:       "Regen",
	Reraise:     "Reraise",
	Flood:       "Flood",
}
