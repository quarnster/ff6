package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/quarnster/util/encoding/binary"
	"log"
	"os"
	"unsafe"
)

// 01 3f 99
// 51 47 4b
// http://www.gamefaqs.com/snes/554041-final-fantasy-iii/faqs/11544

type (
	uint24 uint32

	CharacterName struct {
		Name string `length:"27"`
	}

	SaveFile struct {
		Title                                                                                                    string `length:"32"`
		Unknown                                                                                                  [96]byte
		Terra, Locke, Cyan, Shadow, Edgar, Sabin, Celes, Strago, Relm, Setzer, Mog, Gau, Gogo, Umaro, Leo, Kefka Character
		Unknown2                                                                                                 [16]byte `json:"-"`
		Gil                                                                                                      uint24
		Something                                                                                                uint24
		Steps                                                                                                    uint24
		Unknown22                                                                                                [1947]byte `json:"-"`
		Unknowna                                                                                                 [12]byte
		InventorySortOrder                                                                                       [256]Item         `json:"-"`
		Unknown3                                                                                                 [32]byte          `json:"-"`
		InventoryCount                                                                                           [256]uint8        `json:"-"`
		Unknown4                                                                                                 [1492]uint8       `json:"-"`
		Names                                                                                                    [16]CharacterName `json:"-"`
		Unknown5                                                                                                 [15979]uint8      `json:"-"`
	}
	Character struct {
		Id         CharId
		Portrait   Portrait
		Name       [6]byte
		Level      uint8
		CurrHp     uint16
		TotHp      uint16
		CurrMp     uint16
		TotMp      uint16
		Experience uint32
		Unknown2   byte
		Command    [4]Command
		Strength   uint8
		Speed      uint8
		Stamina    uint8
		Magic      uint8
		Esper      Esper
		RightHand  Item
		LeftHand   Item
		Helmet     Item
		Body       Item
		Relic1     Item
		Relic2     Item
	}
)

func (u *uint24) Read(br *binary.BinaryReader) error {
	d, err := br.Read(3)
	if err != nil {
		return err
	}

	for i, v := range d {
		*u |= uint24(v) << uint24(i*8)
	}
	log.Println(*u)
	return nil
}

func main() {
	f, err := os.Open("/Users/quarnster/Downloads/rs/filesRockDat0.sav")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	br := binary.BinaryReader{
		Reader:    f,
		Endianess: binary.LittleEndian,
	}
	var s SaveFile
	log.Println(br.ReadInterface(&s))
	d, _ := json.MarshalIndent(s, "", "\t")
	_ = d
	log.Println("\n", string(d))
	_ = hex.Dump
	//	log.Println("\n", hex.Dump(s.Unknown22[:]))
	log.Println(unsafe.Sizeof(s))
}
