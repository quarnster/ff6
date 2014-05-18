package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"hash/crc32"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	//	"reflect"
	"regexp"
)

// something5[1]  23  38
// unknown22d[0] 232 245
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

	Inventory struct {
		SortOrder [288]Item
		Count     [288]uint8
	}

	SpellMasteries [13]SpellMastery
	Header         struct {
		Checksum uint32
		Padding  [12]byte
	}
	SpellMastery [64]uint8
	Save         struct {
		Header       //`json:"-"`
		Characters   [16]Character
		Unknown2     [16]byte //`json:"-"`
		Gil          uint24
		Time         Time
		Steps        uint24
		SpellMastery SpellMasteries
		Allzeroes    [333]byte   //`json:"-"`
		Unknown222   [51]byte    //`json:"-"`
		LoreMastery  LoreMastery //`json:"-"`
		Beastiary    [40]uint8
		DanceMastery uint8     //
		Unknown22    [352]byte //`json:"-"`
		Something2   uint8     //`json:"-"`
		Unknown22a   [182]byte //`json:"-"`
		Something3   uint8     //`json:"-"`
		Unknown22b   [51]byte  //`json:"-"`
		Something4   [5]byte   //`json:"-"`
		Unknown22c   [88]byte  //`json:"-"`
		Something5   [2]byte   //`json:"-"`
		Unknown22d   [17]byte  //`json:"-"`
		Inventory    Inventory
		Unknown4     [1460]uint8 //`json:"-"`
		Names        [16]CharacterName
		Unknown5     [736]uint8  //`json:"-"`
		Allzeroes2   [15212]byte //`json:"-"`
	}

	FileHeader struct {
		Unknown  [36]byte
		Version  uint32
		Unknown2 [24]byte
		Header
	}
	String16 [16]byte
	SaveFile struct {
		FF6        String16   //`json:"-"`
		Mobile     String16   //`json:"-"`
		FileHeader FileHeader //`json:"-"`
		Saves      [5]Save
		HaltData   struct {
			Header
			Data [25600]byte
		} //`json:"-"`
		Unknown [16512]byte //`json:"-"`
	}
	Character struct {
		Id         CharId
		Portrait   Portrait
		Unknown    [6]byte
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

const (
	FileChecksumOffset = 32 + 36 + 4 + 24
	HeaderSize         = 16
	SaveEntrySize      = 21008
	FirstSaveOffset    = 112
)

func (inv Inventory) MarshalJSON() ([]byte, error) {
	d := make(map[string]interface{})
	d["SortingOrder"] = &inv.SortOrder
	for i, k := range inv.SortOrder {
		d[iname[k]] = inv.Count[i]
	}
	return json.Marshal(d)
}

func (sm SpellMasteries) MarshalJSON() ([]byte, error) {
	d := make(map[string]*SpellMastery)
	for i := range sm {
		d[cname[CharId(i)]] = &sm[i]
	}
	return json.Marshal(d)
}

func (sm SpellMastery) MarshalJSON() ([]byte, error) {
	d := make(map[string]uint8)
	for k, v := range spellname {
		val := sm[k]
		if val > 100 {
			val = 100
		}
		d[v] = val
	}
	return json.Marshal(d)
}

func (u uint24) Uint() uint32 {
	var j uint32
	for i, v := range u.Data {
		j |= uint32(v) << uint32(i*8)
	}
	return j
}

func (u uint24) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.Uint())
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

func updateChecksums(data []byte) {
	for i := 0; i < 5; i++ {
		saveData := data[FirstSaveOffset+i*SaveEntrySize : FirstSaveOffset+(i+1)*SaveEntrySize]
		binary.LittleEndian.PutUint32(saveData, crc32.ChecksumIEEE(saveData[HeaderSize:]))
	}
	copy(data[FileChecksumOffset:], make([]byte, 4))
	binary.LittleEndian.PutUint32(data[FileChecksumOffset:], crc32.ChecksumIEEE(data[HeaderSize:]))
}

func process(in, out, out2 string) {
	f, err := os.Open(in)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()
	var s SaveFile
	log.Println(binary.Read(f, binary.LittleEndian, &s))
	d, err := json.MarshalIndent(s, "", "\t")
	log.Println(err)
	ioutil.WriteFile(out, d, 0644)
	//	s.Saves[3].Characters[Strago].TotHp = 9999
	//	s.Saves[3].Time.Minutes++

	buf := bytes.NewBuffer(nil)
	log.Println(binary.Write(buf, binary.LittleEndian, s))
	data2 := buf.Bytes()
	updateChecksums(data2)
	log.Println(ioutil.WriteFile(out2, data2, 0644))
}

func main() {
	// log.Println(reflect.TypeOf(Save{}).Size())
	// log.Println(unsafe.Sizeof(Save{}), reflect.ValueOf(Save{}).Type().Size())
	// log.Println(unsafe.Sizeof(Header{}))
	// log.Println(unsafe.Sizeof(SaveFile{}))
	// return
	re := regexp.MustCompile(`(\d+).*?(\d+)`)
	s, _ := filepath.Glob("files*.sav.*")
	for _, p := range s {
		match := re.FindStringSubmatch(p)[1:]
		out := fmt.Sprintf("save%s_%s.json", match[0], match[1])
		out2 := "hack/" + p
		log.Println(p, out, out2)
		process(p, out, out2)
		//		break
	}
}
