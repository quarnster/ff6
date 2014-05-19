package main

import (
	"encoding/json"
	"fmt"
)

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
	NoEsper Esper = 255
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
	NoEsper:       "",
}

func (e Esper) MarshalJSON() ([]byte, error) {
	return json.Marshal(ename[e])
}

func (e *Esper) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range ename {
		if v == s {
			*e = k
			return nil
		}
	}
	return fmt.Errorf("\"%s\" isn't a known Esper", s)
}
