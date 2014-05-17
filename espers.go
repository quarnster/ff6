package main

import "encoding/json"

type Esper uint8

const (
	Ramuh Esper = iota
	Ifrit
	Shiva
	Siren
	Midgardsormr
	Maduin
	Catoblepas
	Bismark
	CaitSith
	Quetzalli
	Valigarmanda
	Odin
	Raiden
	Bahamut
	Alexandr
	Crusader
	RagnarokEsper
	Kirin
	ZoonaSeeker
	Carbunkl
	Phantom
	Seraph
	Golem
	Unicorn
	Fenrir
	Lakshmi
	Phoenix
	Leviathan
)

var ename = map[Esper]string{
	Ramuh:         "Ramuh",
	Ifrit:         "Ifrit",
	Shiva:         "Shiva",
	Siren:         "Siren",
	Midgardsormr:  "Midgardsormr",
	Maduin:        "Maduin",
	Catoblepas:    "Catoblepas",
	Bismark:       "Bismark",
	CaitSith:      "CaitSith",
	Quetzalli:     "Quetzalli",
	Valigarmanda:  "Valigarmanda",
	Odin:          "Odin",
	Raiden:        "Raiden",
	Bahamut:       "Bahamut",
	Alexandr:      "Alexandr",
	Crusader:      "Crusader",
	RagnarokEsper: "RagnarokEsper",
	Kirin:         "Kirin",
	ZoonaSeeker:   "ZoonaSeeker",
	Carbunkl:      "Carbunkl",
	Phantom:       "Phantom",
	Seraph:        "Seraph",
	Golem:         "Golem",
	Unicorn:       "Unicorn",
	Fenrir:        "Fenrir",
	Lakshmi:       "Lakshmi",
	Phoenix:       "Phoenix",
	Leviathan:     "Leviathan",
}

func (e Esper) MarshalJSON() ([]byte, error) {
	return json.Marshal(ename[e])
}
