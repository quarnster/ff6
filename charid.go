package main

import "encoding/json"

type (
	CharId   uint8
	Portrait CharId
)

const (
	Terra CharId = iota
	Locke
	Cyan
	Shadow
	Edgar
	Sabin
	Celes
	Strago
	Relm
	Setzer
	Mog
	Gau
	Gogo
	Umaro
	Vicks_Wedge
	Imp
	Leo
	Banon
	EsperTerra    // (Terra)
	MerchantLocke // (Soldier)
	Ghost
	Kefka          //(Terra)
	EmperorGestahl // (Terra)
	SoldierLocke   // [Village Elder] (Terra)
	Empty1         // [Young Man] (Terra) *cannot fight during battle*
	Empty2         // [Interceptor] (Terra) *cannot fight during battle*
	Empty3         // [Opera Celes] (Celes) *cannot fight during battle*
	Empty4         //  [Scholar] (Glitch) *will hang the game when fighting*
)

var cname = map[CharId]string{
	Terra:          "Terra",
	Locke:          "Locke",
	Cyan:           "Cyan",
	Shadow:         "Shadow",
	Edgar:          "Edgar",
	Sabin:          "Sabin",
	Celes:          "Celes",
	Strago:         "Strago",
	Relm:           "Relm",
	Setzer:         "Setzer",
	Mog:            "Mog",
	Gau:            "Gau",
	Gogo:           "Gogo",
	Umaro:          "Umaro",
	Vicks_Wedge:    "Vicks_Wedge",
	Imp:            "Imp",
	Leo:            "Leo",
	Banon:          "Banon",
	EsperTerra:     "EsperTerra",
	MerchantLocke:  "MerchantLocke",
	Ghost:          "Ghost",
	Kefka:          "Kefka",
	EmperorGestahl: "EmperorGestahl",
	SoldierLocke:   "SoldierLocke",
	Empty1:         "Empty1",
	Empty2:         "Empty2",
	Empty3:         "Empty3",
	Empty4:         "Empty4",
}

func (c CharId) MarshalJSON() ([]byte, error) {
	return json.Marshal(cname[c])
}

func (p Portrait) MarshalJSON() ([]byte, error) {
	return CharId(p).MarshalJSON()
}
