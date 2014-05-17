package main

import (
	//	"bytes"
	//	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cheggaaa/pb"
	"github.com/quarnster/util/encoding/binary"
	"hash/crc32"
	//	"io/ioutil"
	"log"
	"os"
	"unsafe"
)

// http://www.gamefaqs.com/snes/554041-final-fantasy-iii/faqs/11544

// version 166 168 170
// file checksum:
// 113 172 147   /\   59       -25 (230)
//  79 197  64   /\  118      -133 (122)
//  23  34  37   /\   11         3
// 137 128 201   /\   -9(246)   73

// slot 1 checksum
//  43  250 185  /\   207  -65
//  201 178 253  /\   -23   75
//  193 139  21  /\   -54 -118
//  155   7 217  /\  -148  210

// Minutes 			37 37 40  /\   0  3
// Seconds 			36 45 56  /\   9  11
// Something2 		 1  1 65  /\   0  64
// Something5[0]	 9 18 96  /\   9  78
// sum              65 65 65
//

// Save slot 1
//       0x44 a4 a6 a8 aa ac
//       0x60 [a2 9d 32 00] [71 4f 17 89] [ac c5 22 80] [93 40 25 c9] [ff 80 9b 16]
// +0    0x70 [c1 a7 16 6b] [2b c9 c1 9b] [fa b2 8b 07] [b9 fd 15 d9] [9d 4a dd b4]
// +628  0x2e4 0x25 0x25 0x25 0x28 0x29
// +629  0x2e5 0x1f 0x24 0x2d 0x38 0x0b
// +2246 0x936 0x41 0x01 0x01 0x41 0x41
// +2429 0x9ed 0xf4 0xee 0xee 0xee 0xee
// +2481 0xa21 [0xf4 0xf4 0xf4 0xf4 0xf4] [0xee 0xee 0xee 0xee 0xee] [0xee 0xee 0xee 0xee 0xee] [0xee 0xee 0xee 0xee 0xee] [0xee 0xee 0xee 0xee 0xee]  [0xee 0xee 0xee 0xee 0xee]
// +2574 0xa7e 0x68 0x09 0x12 0x60 0x34

// Save slot 4, 0.sav 5, 6 and 7
//      0x00034 0x01                0x04                0x04
//      0x00044 0xac                0xb0                0xb2
//      0x00060 [ff 80 9b 16]       [5d ac b5 b1]       [ad 87 75 21]
//+0    0x0f6a0	[1e c0 f9 1b]       [e5 44 2d 95]       [fe bc 59 2c]
//+629  0x0f915	1b                  1f                  35
//+2429 0x1001d da                  d0                  d0
//+2481 0x10051 [51 89 7f 90 7f]    [d0 d0 d0 d0 d0]    [d0 d0 d0 d0 d0]
//+2544 0x10090 0d                  09                  09
//+2552 0x10098 0d                  09                  09
//+5228 0x1009e 0d                  09                  09
//+2574 0x100ae [62 4f]             [f8 50]             [0e 51]

type (
	uint24 uint32

	CharacterName struct {
		Name string `length:"27"`
	}
	Time struct {
		Hours   uint8
		Minutes uint8
		Seconds uint8
	}

	Header struct {
		Checksum uint32
		Padding  [12]byte
	}
	SpellMastery [64]uint8
	Save         struct {
		Header
		Terra, Locke, Cyan, Shadow, Edgar, Sabin, Celes, Strago, Relm, Setzer, Mog, Gau, Gogo, Umaro, Leo, Kefka Character
		Unknown2                                                                                                 [16]byte `json:"-"`
		Gil                                                                                                      uint24
		Time                                                                                                     Time
		Steps                                                                                                    uint24
		SpellMastery                                                                                             [13]SpellMastery `json:"-"`
		Unknown222                                                                                               [384]byte        `json:"-"`
		LoreMastery                                                                                              LoreMastery
		Beastiary                                                                                                [40]uint8 `json:"-"`
		DanceMastery                                                                                             uint8
		Unknown22                                                                                                [352]byte   `json:"-"`
		Something2                                                                                               uint8       `json:"-"`
		Unknown22a                                                                                               [182]byte   `json:"-"`
		Something3                                                                                               uint8       `json:"-"`
		Unknown22b                                                                                               [51]byte    `json:"-"`
		Something4                                                                                               [5]byte     `json:"-"`
		Unknown22c                                                                                               [88]byte    `json:"-"`
		Something5                                                                                               [2]byte     `json:"-"`
		Unknown22d                                                                                               [16]byte    `json:"-"`
		InventorySortOrder                                                                                       [288]Item   `json:"-"`
		InventoryCount                                                                                           [288]uint8  `json:"-"`
		Unknown4                                                                                                 [1460]uint8 `json:"-"`
		Names                                                                                                    [16]CharacterName
		Unknown5                                                                                                 [15868 + 80]uint8 `json:"-"`
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

	FileHeader struct {
		Unknown  [36]byte
		Version  uint8
		Unknown2 [27]byte
		Header
	}
	SaveFile struct {
		Title      string `length:"32"`
		FileHeader FileHeader
		Saves      [5]Save
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

const (
	HeaderSize      = 16
	SaveEntrySize   = 21008
	FirstSaveOffset = 112
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
func reverse(data []byte) {
	j := len(data) - 1
	for i := 0; i < len(data)/2; i++ {
		data[i], data[j] = data[j], data[i]
		j--
	}
}
func sum(checksum uint32, data []byte) {
	fmt.Printf("%x: %x\n", checksum, crc32.ChecksumIEEE(data))
}
func process(fn int) {
	f, err := os.Open(fmt.Sprintf("/Users/quarnster/Downloads/rs/filesRockDat0.sav.%d", fn))
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
	log.Println(string(d))
	log.Println(unsafe.Offsetof(s.Saves))
	var headers []byte
	for i, save := range s.Saves {
		log.Println(br.Seek(FirstSaveOffset+int64(i*SaveEntrySize), 0))
		data, _ := br.Read(SaveEntrySize)
		sum(save.Checksum, data[16:])
		headers = append(headers, data...)
	}
	log.Println(br.Seek(0, 1))
	br.Seek(96, 0)
	//	pre, _ := br.Read(96)
	data, _ := br.Read(147184)
	log.Println(br.Seek(0, 1))
	log.Printf("%+#v", data[:4])
	sum(s.FileHeader.Checksum, data[16:])
	sum(s.FileHeader.Checksum, headers)
	//_ = pre
	// sum(s.FileHeader.Checksum, append(pre, data[16:]...))
	// sum(s.FileHeader.Checksum, append(pre, headers...))
	// sum(s.FileHeader.Checksum, append(pre[32:], data[16:]...))
	// sum(s.FileHeader.Checksum, append(pre[32:], headers...))

	bar := pb.StartNew(len(data[4:]))
	defer bar.Finish()
	for i := range data[4:] {
		bar.Increment()
		if v := crc32.ChecksumIEEE(data[4 : 4+i]); v == s.FileHeader.Checksum {
			fmt.Printf("%d bytes: %x: %x\n", i, v, s.FileHeader.Checksum)
		}
	}
}
func main() {
	for i := 1; i < 5; i++ {
		process(i)
		break
	}
}
