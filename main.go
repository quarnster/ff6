package main

import (
	"encoding/hex"
	"encoding/json"
	//	"fmt"
	"github.com/quarnster/util/encoding/binary"
	//	"io/ioutil"
	"log"
	"os"
	//	"unsafe"
)

// 01 3f 99
// 51 47 4b
// http://www.gamefaqs.com/snes/554041-final-fantasy-iii/faqs/11544

type (
	uint24 uint32

	CharacterName struct {
		Name string `length:"27"`
	}

	SpellMastery [64]uint8
	Save         struct {
		Unknown                                                                                                  [96]byte
		Terra, Locke, Cyan, Shadow, Edgar, Sabin, Celes, Strago, Relm, Setzer, Mog, Gau, Gogo, Umaro, Leo, Kefka Character
		Unknown2                                                                                                 [16]byte `json:"-"`
		Gil                                                                                                      uint24
		Something                                                                                                uint24
		Steps                                                                                                    uint24
		SpellMastery                                                                                             [13]SpellMastery
		Unknown222                                                                                               [384]byte
		LoreMastery                                                                                              LoreMastery
		Beastiary                                                                                                [40]uint8
		DanceMastery                                                                                             uint8
		Unknown22                                                                                                [698]byte         `json:"-"`
		InventorySortOrder                                                                                       [288]Item         `json:"-"`
		InventoryCount                                                                                           [288]uint8        `json:"-"`
		Unknown4                                                                                                 [1460]uint8       `json:"-"`
		Names                                                                                                    [16]CharacterName `json:"-"`
		Unknown5                                                                                                 [15868]uint8      `json:"-"`
	}
	// Locke meteor=26% = 0x1a
	//       flood = 62% = 0x3e
	// Edgar Drain 55% 0x37
	//	     graviga 53% = 0x35
	//	     tornado 53% = 0x35
	//       esuna 84% = 0x54
	//		reraise 53% = 0x35
	// Sabin reraise 85%  0x55
	// sabin drain 8% = 0x08
	//	     death 50% = 0x32
	//       meltdown = 0x0a
	SaveFile struct {
		Title string `length:"32"`
		Saves [5]Save
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
		Status     Status
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
	//log.Println(br.ReadInterface(&s))

	d, _ := json.MarshalIndent(s, "", "\t")
	_ = d
	log.Println("\n", string(d))
	// d, _ = json.MarshalIndent(s.Saves[1].InventorySortOrder, "", "\t")
	// log.Println("\n", string(d))
	for _, save := range s.Saves {
		// 	//	ioutil.WriteFile(fmt.Sprintf("save%d", i), save.Unknown5[:], 0644)
		// 	log.Printf("%s: %+v", cname[Terra], save.Spells[Terra])
		// 	for i, v := range save.Spells[Sabin] {
		// 		log.Printf("%s: %d", spellname[Spell(i)], v)
		// 	}
		// 	log.Printf("%s: %+v", cname[Locke], save.Spells[Locke])
		// 	// for _, m := range save.Spells {
		// 	// 	log.Printf("%+v", m)
		log.Println("\n", hex.Dump(save.Unknown2[:]))
		// log.Println("\n", hex.Dump(save.Unknown222[:]))
		// log.Println("\n", hex.Dump(save.Unknown22[:]))
		// log.Println("\n", hex.Dump(save.InventoryCount[:]))
		// log.Println("\n", hex.Dump(save.Unknown4[:]))
		// log.Println("\n", hex.Dump(save.Unknown5[:]))
		//		log.Printf("%+v", save.Names)
		// 	// }
	}
	// log.Println(br.Seek(4096*5, 1))
	// data, _ := br.Read(4096)
	// log.Println("\n", hex.Dump(data))
	// log.Println(br.Seek(0, 2))
	_ = hex.Dump
}
