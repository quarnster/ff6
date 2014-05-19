package main

import "encoding/json"

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

func (l *LoreMastery) UnmarshalJSON(data []byte) error {
	d := make(map[string]bool)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}
	var i Lore
	for k, v := range lorenames {
		if d[v] {
			i |= k
		}
	}
	(*uint24)(l).Set(uint32(i))
	return nil
}

func (l LoreMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]bool)
	for k, v := range lorenames {
		d[v] = (uint24(l).Uint() & uint32(k)) != 0
	}
	return json.Marshal(d)
}

var lorenames = map[Lore]string{
	Condemned:          "Condemned",
	Roulette:           "Roulette",
	Tsunami:            "Tsunami",
	AquaBreath:         "Aqua Breath",
	Aero:               "Aero",
	OneThousandNeedles: "1000 Needles",
	BigGuard:           "BigGuard",
	RevengeBlast:       "RevengeBlast",
	PearlWind:          "PearlWind",
	L5Death:            "L.5 Death",
	L4Flare:            "L.4 Flare",
	L3Confuse:          "L.3 Confuse",
	ReflectLore:        "Reflect",
	LPearl:             "L.? Pearl",
	StepMine:           "StepMine",
	ForceField:         "ForceField",
	Dischord:           "Dischord",
	SourMouth:          "SourMouth",
	PepUp:              "PepUp",
	Rippler:            "Rippler",
	StoneLore:          "Stone",
	Quasar:             "Quasar",
	GrandTrain:         "GrandTrain",
	SelfDestruct:       "Self Destruct",
}

//var a = Roulette + Tsunami + AquaBreath + Aero + OneThousandNeedles + RevengeBlast + L5Death + ForceField + Dischord + Stone + SelfDestruct

//01111101
//128+32+16+8+4+2
