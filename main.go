// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"github.com/robmerrell/comandante"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// http://www.gamefaqs.com/snes/554041-final-fantasy-iii/faqs/11544 for basics,
// then loads of trial and error.

type (
	uint24 struct {
		Data [3]uint8
	}

	CharacterName [27]byte
	Time          struct {
		Hours   uint8
		Minutes uint8
		Seconds uint8
	}

	Header struct {
		Checksum uint32
		Padding  [12]byte
	}
	Body struct {
		Characters   [16]Character
		Unknown2     [16]byte //`json:"-"`
		Gil          uint24
		Time         Time
		Steps        uint24
		SpellMastery SpellMasteries
		Allzeroes    [332]byte `json:"-"`
		Something    uint8
		Unknown222   [49]byte //`json:"-"`
		Swdtech      uint8
		BlitzMastery BlitzMastery
		LoreMastery  LoreMastery //`json:"-"`
		Beastiary    [41]uint8
		DanceMastery uint8     //
		Unknown22    [352]byte //`json:"-"`
		Something2   uint8     //`json:"-"`
		Unknown22a   [182]byte //`json:"-"`
		Something3   uint8     //`json:"-"`
		Unknown22b   [51]byte  //`json:"-"`
		Something4   [5]byte   //`json:"-"`
		Unknown22c   [88]byte  //`json:"-"`
		Checksum     uint16    //`json:"-"`
	}
	Iv struct {
		EspersOwned EspersOwned
		Unknown     [12]uint8
		Inventory   Inventory
		Unknown4    [430]uint8
		IvChecksum  uint16
	}

	Body2 struct {
		Unknown41 [1028]uint8 //`json:"-"`
		Names     [16]CharacterName
		Unknown5  [736]uint8 //`json:"-"`
	}

	Save struct {
		Header     `json:"-"`
		Body       Body
		Iv         Iv
		Body2      Body2
		Allzeroes2 [15212]byte `json:"-"`
	}

	FileHeader struct {
		Unknown  [36]byte
		Version  uint32
		Unknown2 [24]byte //`json:"-"`
		Header   `json:"-"`
	}
	String16 [16]byte
	HaltData struct {
		Header `json:"-"`
		Data   [25600]byte
	}
	SaveFile struct {
		FF6        String16   //`json:"-"`
		Mobile     String16   //`json:"-"`
		FileHeader FileHeader //`json:"-"`
		Saves      [5]Save
		HaltData   HaltData
		Unknown    [16512]byte //`json:"-"`
	}
	Character struct {
		Id         CharId
		Portrait   Portrait
		Unknown    [6]uint8
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
		RightHand  Equipable
		LeftHand   Equipable
		Helmet     Equipable
		Body       Equipable
		Relic1     Equipable
		Relic2     Equipable
	}
	Writer interface {
		Write(io.Writer) error
	}
)

const (
	FileChecksumOffset = 32 + 36 + 4 + 24
	HeaderSize         = 16
	SaveEntrySize      = 21008
	FirstSaveOffset    = 112
)

func csData(v interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	binary.Write(buf, binary.LittleEndian, v)
	data := buf.Bytes()
	// Don't include the actual checksum value in the
	// calculation of it ;)
	off := len(data) - 2
	var sum uint16
	for _, v := range data[:off] {
		sum += uint16(v)
	}

	binary.LittleEndian.PutUint16(data[off:], sum)

	return data
}

func (i *Iv) Encode() []byte {
	data := csData(*i)
	off := len(data) - 2
	curr := binary.LittleEndian.Uint16(data[off:])
	if curr != 0 {
		// TODO(.): why do I need to do this?
		binary.LittleEndian.PutUint16(data[off:], curr+1)
	}

	return data
}

func (b *Body) Encode() []byte {
	return csData(*b)
}

func (s *Save) Encode() []byte {
	buf := bytes.NewBuffer(nil)
	buf.Write(s.Body.Encode())
	buf.Write(s.Iv.Encode())
	binary.Write(buf, binary.LittleEndian, s.Body2)
	buf.Write(s.Allzeroes2[:])
	s.Checksum = crc32.ChecksumIEEE(buf.Bytes())
	buf2 := bytes.NewBuffer(nil)
	binary.Write(buf2, binary.LittleEndian, s.Header)
	buf2.Write(buf.Bytes())
	return buf2.Bytes()
}

func (s *SaveFile) Encode() []byte {
	s.FileHeader.Checksum = 0
	buf := bytes.NewBuffer(nil)
	buf.Write(s.Mobile[:])
	binary.Write(buf, binary.LittleEndian, s.FileHeader)

	for i := range s.Saves {
		buf.Write(s.Saves[i].Encode())
	}
	s.HaltData.Checksum = crc32.ChecksumIEEE(s.HaltData.Data[:])
	binary.Write(buf, binary.LittleEndian, s.HaltData)
	binary.Write(buf, binary.LittleEndian, s.Unknown)

	data := buf.Bytes()

	s.FileHeader.Checksum = crc32.ChecksumIEEE(data)
	binary.LittleEndian.PutUint32(data[80:], s.FileHeader.Checksum)
	buf2 := bytes.NewBuffer(nil)
	buf2.Write(s.FF6[:])
	buf2.Write(data)
	return buf2.Bytes()
}

func (u *uint24) Set(j uint32) {
	for i := range u.Data {
		u.Data[i] = uint8((j >> (uint(i) * 8)) & 0xff)
	}
}
func (u uint24) Uint() uint32 {
	var j uint32
	for i, v := range u.Data {
		j |= uint32(v) << uint32(i*8)
	}
	return j
}

func (u *uint24) UnmarshalJSON(data []byte) error {
	var j uint32
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	u.Set(j)
	return nil
}

func (u uint24) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Uint())
}

func (n *CharacterName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	copy(n[:], []byte(s))
	return nil
}

func (n CharacterName) MarshalJSON() ([]byte, error) {
	bytes := n[:]
	i := len(bytes) - 1
	for i > 0 && bytes[i] == 0 {
		i--
	}
	return json.Marshal(string(bytes[:i+1]))
}

func (n String16) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(n[:]))
}

func (n *String16) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	copy(n[:], []byte(s))
	return nil
}

var (
	savfile  = "filesRockDat0.sav"
	jsonfile = "filesRockDat0.json"
)

func tojson() error {
	f, err := os.Open(savfile)
	if err != nil {
		return err
	}
	defer f.Close()
	var s SaveFile
	if err := binary.Read(f, binary.LittleEndian, &s); err != nil {
		return err
	}

	d, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(jsonfile, d, 0644)
}

func fromjson() error {
	data, err := ioutil.ReadFile(jsonfile)
	if err != nil {
		return err
	}
	var s SaveFile
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	_, err = os.Stat(savfile)
	if err == nil {
		_, err = os.Stat(savfile + ".bak")
		if err != nil {
			os.Rename(savfile, savfile+".bak")
		}
	}
	return ioutil.WriteFile(savfile, s.Encode(), 0644)
}

func main() {
	bin := comandante.New("ff6", "Manage Final Fantasy VI mobile save files")
	bin.IncludeHelp()
	tj := comandante.NewCommand("tojson", "Convert .sav save file to .json", tojson)
	tj.FlagInit = func(fs *flag.FlagSet) {
		fs.StringVar(&savfile, "i", savfile, "Input .sav file")
		fs.StringVar(&jsonfile, "o", jsonfile, "Output .json file")
	}
	bin.RegisterCommand(tj)
	fj := comandante.NewCommand("fromjson", "Convert .json save file to .sav", fromjson)
	fj.FlagInit = func(fs *flag.FlagSet) {
		fs.StringVar(&jsonfile, "i", jsonfile, "Input .json file")
		fs.StringVar(&savfile, "o", savfile, "Output .sav file")
	}
	bin.RegisterCommand(fj)

	flag.Parse()

	if err := bin.Run(); err != nil {
		log.Fatalln(err)
	}
}
