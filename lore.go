package main

//import "encoding/json"

type (
	Lore        uint32
	LoreMastery uint24
)

const (
	Condemned Lore = 1 << iota
	Roulette
	Tsunami
	AquaBreath
	Aero
	OneThousandNeedles
	BigGuard
	RevengeBlast
	PearlWind
	L5Death
	L4Flare
	L3Confuse
	ReflectLore
	LPearl
	StepMine
	ForceField
	Dischord
	SourMouth
	PepUp
	Rippler
	StoneLore
	Quasar
	GrandTrain
	SelfDestruct
)

// func (l LoreMastery) MarshalJSON() ([]byte, error) {
// 	str := ""
// 	for i := uint(0); i < 24; i++ {
// 		if l&(1<<i) != 0 {
// 			str += "+" + lorenames[Lore(1<<i)]
// 		}
// 	}
// 	return json.Marshal(str)
// }

var lorenames = map[Lore]string{
	Condemned:          "Condemned",
	Roulette:           "Roulette",
	Tsunami:            "Tsunami",
	AquaBreath:         "AquaBreath",
	Aero:               "Aero",
	OneThousandNeedles: "OneThousandNeedles",
	BigGuard:           "BigGuard",
	RevengeBlast:       "RevengeBlast",
	PearlWind:          "PearlWind",
	L5Death:            "L5Death",
	L4Flare:            "L4Flare",
	L3Confuse:          "L3Confuse",
	ReflectLore:        "ReflectLore",
	LPearl:             "LPearl",
	StepMine:           "StepMine",
	ForceField:         "ForceField",
	Dischord:           "Dischord",
	SourMouth:          "SourMouth",
	PepUp:              "PepUp",
	Rippler:            "Rippler",
	StoneLore:          "StoneLore",
	Quasar:             "Quasar",
	GrandTrain:         "GrandTrain",
	SelfDestruct:       "SelfDestruct",
}

//var a = Roulette + Tsunami + AquaBreath + Aero + OneThousandNeedles + RevengeBlast + L5Death + ForceField + Dischord + Stone + SelfDestruct

//01111101
//128+32+16+8+4+2
