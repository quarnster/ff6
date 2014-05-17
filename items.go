package main

import "encoding/json"

type Item uint8

const (

	// Dirk (0)
	Dirk Item = iota
	MithrilKnife
	Guardian
	AirLancet
	ThiefKnife
	AssassinsKnife
	ManEater
	SwordBreaker
	Gladius
	ValiantKnife
	// Sword (0xA)
	MithrilBlade
	RegalCutlass
	RuneEdge
	FlameSword
	BlizzardSword
	ThunderSword
	Epee
	BreakBlade
	Drainer
	Enhancer
	Crystal
	Falchion
	SoulSabre
	OgreNix
	Excalibur
	Zantetsuken
	Illumina
	Ragnarok
	UltimaWeapon
	// Lance (0x1d)
	MithrilPike
	Trident
	StoutSpear
	Partisian
	HolyLance
	GoldLance
	RadiantLance
	Impartisian
	//  Dirk (0x25)
	Imperial
	Kodachi
	Blossom
	Hardened
	Ichigeki
	Kagenui
	// Knife (0x2B)
	Ashura
	Kotetsu
	Forged
	Tempest
	Murasame
	Aura
	Strato
	SkyRender
	// Rod (0x33)
	HealRod
	MithrilRod
	FireRod
	IceRod
	ThunderRod
	PoisonRod
	PearlRod
	GravityRod
	Punisher
	MagusRod
	// Brush (0x3d)
	ChocoboBrush
	DaVinciBrush
	MagicalBrush
	RainbowBrush
	// Stars (0x41)
	Shuriken
	NinjaStar
	TackStar
	// Special (0x44)
	Flail
	FullMoon
	MorningStar
	Boomerang
	RisingSun
	HawkEye
	BoneClub
	Sniper
	WingEdge
	// Gambler (0x4d)
	Cards
	Darts
	DeathTarot
	Trump
	Dice
	FixedDice
	// Claw (0x53)
	MetalKnuckle
	MithrilClaw
	Kaiser
	PoisonClaw
	FireKnuckle
	DragonClaw
	TigerFang
	// Shields (0x5A)
	Buckler
	HeavyShield
	MythrilShield
	GoldShield
	AegisShield
	DiamondShield
	FlameShield
	IceShield
	ThunderShield
	CrystalShield
	GenjiShield
	TortoiseShield
	CursedShield
	PaladinShield
	ForceShield
	// Helmet (0x69)
	LeatherHat
	HairBand
	PlumedHat
	Beret
	MagusHat
	Bandana
	IronHelmet
	Coronet
	BardHat
	GreenBeret
	HeadBand
	MithrilHelm
	Tiara
	GoldHelmet
	TigerMask
	RedCap
	MysteryVeil
	Circlet
	RoyalCrown
	DiamondHelm
	DarkHood
	CrystalHelm
	OathViel
	CatHood
	GenjiHelm
	Thornlet
	Titanium
	// Armor (0x84)
	LeatherArmor
	CottonRobe
	KungFuSuit
	IronArmor
	SilkRobe
	MithrilVest
	NinjaGear
	WhiteDress
	MithrilMail
	GaiaGear
	MirageDress
	GoldArmor
	PowerSash
	LightRobe
	DiamondVest
	RedJacket
	ForceArmor
	DiamondArmor
	BlackGarb
	TaoRobe
	CrystalMail
	CzarinaGown
	GenjiArmor
	ReedCloak
	MinervaBustier
	TabbySuit
	ChocoboSuit
	MoogleSuit
	NutkinSuit
	BehemothSuit
	SnowMuffler
	// Tool (0xa3)
	NoiseBlaster
	BioBlaster
	Flash
	ChainSaw
	Debilitator
	Drill
	AirAnchor
	AutoCrossbow
	// Skean (0xAB)
	FireSkean
	WaterEdge
	BoltEdge
	InvizEdge
	ShadowEdge
	// Relics (0xB0)
	Goggles
	StarPendant
	PeaceRing
	Amulet
	WhiteCape
	JewelRing
	FairRing
	BarrierRing
	MithrilGlove
	GuardRing
	HermesSandals
	WallRing
	CherubDown
	CureRing
	TrueKnight
	DragonBoots
	ZephyrCape
	CzarinaRing
	CursedCing
	Earrings
	AtlasArmlet
	BlizzardRing
	RageRing
	SneakRing
	PodBracelet
	HeroRing
	Ribbon
	MuscleBelt
	CrystalOrb
	GoldHairpin
	Celestriad
	BrigandsGlove
	Gauntlet
	GenjiGlove
	HyperWrist
	MastersScroll
	Beads
	BlackBelt
	CoinToss
	FakeMustache
	SoulOfTamasa
	DragonHorn
	MeritAward
	MementoRing
	SafetyBit
	RelicRing
	MolulusCharm
	CharmBangle
	MarvelShoes
	AlarmEarring
	GaleHairpin
	SniperSight
	GrowthEgg
	TintinaBar
	SprintShoes
	// Misc (0xE7)
	RenameCard
	Potion
	HiPotion
	XPotion
	Ether
	HiEther
	XEther
	Elixir
	Megalixir
	PhoenixDown
	Revivify
	Antidote
	Eyedrop
	Soft
	Remedy
	SleepingBag
	Tent
	GreenCherry
	Magicite
	SuperBall
	EchoScreen
	SmokeBomb
	TeleportStone
	DriedMeat
	Nothing
)

var iname = map[Item]string{
	Dirk:           "Dirk",
	MithrilKnife:   "MithrilKnife",
	Guardian:       "Guardian",
	AirLancet:      "AirLancet",
	ThiefKnife:     "ThiefKnife",
	AssassinsKnife: "AssassinsKnife",
	ManEater:       "ManEater",
	SwordBreaker:   "SwordBreaker",
	Gladius:        "Gladius",
	ValiantKnife:   "ValiantKnife",
	MithrilBlade:   "MithrilBlade",
	RegalCutlass:   "RegalCutlass",
	RuneEdge:       "RuneEdge",
	FlameSword:     "FlameSword",
	BlizzardSword:  "BlizzardSword",
	ThunderSword:   "ThunderSword",
	Epee:           "Epee",
	BreakBlade:     "BreakBlade",
	Drainer:        "Drainer",
	Enhancer:       "Enhancer",
	Crystal:        "Crystal",
	Falchion:       "Falchion",
	SoulSabre:      "SoulSabre",
	OgreNix:        "OgreNix",
	Excalibur:      "Excalibur",
	Zantetsuken:    "Zantetsuken",
	Illumina:       "Illumina",
	Ragnarok:       "Ragnarok",
	UltimaWeapon:   "UltimaWeapon",
	MithrilPike:    "MithrilPike",
	Trident:        "Trident",
	StoutSpear:     "StoutSpear",
	Partisian:      "Partisian",
	HolyLance:      "HolyLance",
	GoldLance:      "GoldLance",
	RadiantLance:   "RadiantLance",
	Impartisian:    "Impartisian",
	Imperial:       "Imperial",
	Kodachi:        "Kodachi",
	Blossom:        "Blossom",
	Hardened:       "Hardened",
	Ichigeki:       "Ichigeki",
	Kagenui:        "Kagenui",
	Ashura:         "Ashura",
	Kotetsu:        "Kotetsu",
	Forged:         "Forged",
	Tempest:        "Tempest",
	Murasame:       "Murasame",
	Aura:           "Aura",
	Strato:         "Strato",
	SkyRender:      "SkyRender",
	HealRod:        "HealRod",
	MithrilRod:     "MithrilRod",
	FireRod:        "FireRod",
	IceRod:         "IceRod",
	ThunderRod:     "ThunderRod",
	PoisonRod:      "PoisonRod",
	PearlRod:       "PearlRod",
	GravityRod:     "GravityRod",
	Punisher:       "Punisher",
	MagusRod:       "MagusRod",
	ChocoboBrush:   "ChocoboBrush",
	DaVinciBrush:   "DaVinciBrush",
	MagicalBrush:   "MagicalBrush",
	RainbowBrush:   "RainbowBrush",
	Shuriken:       "Shuriken",
	NinjaStar:      "NinjaStar",
	TackStar:       "TackStar",
	Flail:          "Flail",
	FullMoon:       "FullMoon",
	MorningStar:    "MorningStar",
	Boomerang:      "Boomerang",
	RisingSun:      "RisingSun",
	HawkEye:        "HawkEye",
	BoneClub:       "BoneClub",
	Sniper:         "Sniper",
	WingEdge:       "WingEdge",
	Cards:          "Cards",
	Darts:          "Darts",
	DeathTarot:     "DeathTarot",
	Trump:          "Trump",
	Dice:           "Dice",
	FixedDice:      "FixedDice",
	MetalKnuckle:   "MetalKnuckle",
	MithrilClaw:    "MithrilClaw",
	Kaiser:         "Kaiser",
	PoisonClaw:     "PoisonClaw",
	FireKnuckle:    "FireKnuckle",
	DragonClaw:     "DragonClaw",
	TigerFang:      "TigerFang",
	Buckler:        "Buckler",
	HeavyShield:    "HeavyShield",
	MythrilShield:  "MythrilShield",
	GoldShield:     "GoldShield",
	AegisShield:    "AegisShield",
	DiamondShield:  "DiamondShield",
	FlameShield:    "FlameShield",
	IceShield:      "IceShield",
	ThunderShield:  "ThunderShield",
	CrystalShield:  "CrystalShield",
	GenjiShield:    "GenjiShield",
	TortoiseShield: "TortoiseShield",
	CursedShield:   "CursedShield",
	PaladinShield:  "PaladinShield",
	ForceShield:    "ForceShield",
	LeatherHat:     "LeatherHat",
	HairBand:       "HairBand",
	PlumedHat:      "PlumedHat",
	Beret:          "Beret",
	MagusHat:       "MagusHat",
	Bandana:        "Bandana",
	IronHelmet:     "IronHelmet",
	Coronet:        "Coronet",
	BardHat:        "BardHat",
	GreenBeret:     "GreenBeret",
	HeadBand:       "HeadBand",
	MithrilHelm:    "MithrilHelm",
	Tiara:          "Tiara",
	GoldHelmet:     "GoldHelmet",
	TigerMask:      "TigerMask",
	RedCap:         "RedCap",
	MysteryVeil:    "MysteryVeil",
	Circlet:        "Circlet",
	RoyalCrown:     "RoyalCrown",
	DiamondHelm:    "DiamondHelm",
	DarkHood:       "DarkHood",
	CrystalHelm:    "CrystalHelm",
	OathViel:       "OathViel",
	CatHood:        "CatHood",
	GenjiHelm:      "GenjiHelm",
	Thornlet:       "Thornlet",
	Titanium:       "Titanium",
	LeatherArmor:   "LeatherArmor",
	CottonRobe:     "CottonRobe",
	KungFuSuit:     "KungFuSuit",
	IronArmor:      "IronArmor",
	SilkRobe:       "SilkRobe",
	MithrilVest:    "MithrilVest",
	NinjaGear:      "NinjaGear",
	WhiteDress:     "WhiteDress",
	MithrilMail:    "MithrilMail",
	GaiaGear:       "GaiaGear",
	MirageDress:    "MirageDress",
	GoldArmor:      "GoldArmor",
	PowerSash:      "PowerSash",
	LightRobe:      "LightRobe",
	DiamondVest:    "DiamondVest",
	RedJacket:      "RedJacket",
	ForceArmor:     "ForceArmor",
	DiamondArmor:   "DiamondArmor",
	BlackGarb:      "BlackGarb",
	TaoRobe:        "TaoRobe",
	CrystalMail:    "CrystalMail",
	CzarinaGown:    "CzarinaGown",
	GenjiArmor:     "GenjiArmor",
	ReedCloak:      "ReedCloak",
	MinervaBustier: "MinervaBustier",
	TabbySuit:      "TabbySuit",
	ChocoboSuit:    "ChocoboSuit",
	MoogleSuit:     "MoogleSuit",
	NutkinSuit:     "NutkinSuit",
	BehemothSuit:   "BehemothSuit",
	SnowMuffler:    "SnowMuffler",
	NoiseBlaster:   "NoiseBlaster",
	BioBlaster:     "BioBlaster",
	Flash:          "Flash",
	ChainSaw:       "ChainSaw",
	Debilitator:    "Debilitator",
	Drill:          "Drill",
	AirAnchor:      "AirAnchor",
	AutoCrossbow:   "AutoCrossbow",
	FireSkean:      "FireSkean",
	WaterEdge:      "WaterEdge",
	BoltEdge:       "BoltEdge",
	InvizEdge:      "InvizEdge",
	ShadowEdge:     "ShadowEdge",
	Goggles:        "Goggles",
	StarPendant:    "StarPendant",
	PeaceRing:      "PeaceRing",
	Amulet:         "Amulet",
	WhiteCape:      "WhiteCape",
	JewelRing:      "JewelRing",
	FairRing:       "FairRing",
	BarrierRing:    "BarrierRing",
	MithrilGlove:   "MithrilGlove",
	GuardRing:      "GuardRing",
	HermesSandals:  "HermesSandals",
	WallRing:       "WallRing",
	CherubDown:     "CherubDown",
	CureRing:       "CureRing",
	TrueKnight:     "TrueKnight",
	DragonBoots:    "DragonBoots",
	ZephyrCape:     "ZephyrCape",
	CzarinaRing:    "CzarinaRing",
	CursedCing:     "CursedCing",
	Earrings:       "Earrings",
	AtlasArmlet:    "AtlasArmlet",
	BlizzardRing:   "BlizzardRing",
	RageRing:       "RageRing",
	SneakRing:      "SneakRing",
	PodBracelet:    "PodBracelet",
	HeroRing:       "HeroRing",
	Ribbon:         "Ribbon",
	MuscleBelt:     "MuscleBelt",
	CrystalOrb:     "CrystalOrb",
	GoldHairpin:    "GoldHairpin",
	Celestriad:     "Celestriad",
	BrigandsGlove:  "BrigandsGlove",
	Gauntlet:       "Gauntlet",
	GenjiGlove:     "GenjiGlove",
	HyperWrist:     "HyperWrist",
	MastersScroll:  "MastersScroll",
	Beads:          "Beads",
	BlackBelt:      "BlackBelt",
	CoinToss:       "CoinToss",
	FakeMustache:   "FakeMustache",
	SoulOfTamasa:   "SoulOfTamasa",
	DragonHorn:     "DragonHorn",
	MeritAward:     "MeritAward",
	MementoRing:    "MementoRing",
	SafetyBit:      "SafetyBit",
	RelicRing:      "RelicRing",
	MolulusCharm:   "MolulusCharm",
	CharmBangle:    "CharmBangle",
	MarvelShoes:    "MarvelShoes",
	AlarmEarring:   "AlarmEarring",
	GaleHairpin:    "GaleHairpin",
	SniperSight:    "SniperSight",
	GrowthEgg:      "GrowthEgg",
	TintinaBar:     "TintinaBar",
	SprintShoes:    "SprintShoes",
	RenameCard:     "RenameCard",
	Potion:         "Potion",
	HiPotion:       "HiPotion",
	XPotion:        "XPotion",
	Ether:          "Ether",
	HiEther:        "HiEther",
	XEther:         "XEther",
	Elixir:         "Elixir",
	Megalixir:      "Megalixir",
	PhoenixDown:    "PhoenixDown",
	Revivify:       "Revivify",
	Antidote:       "Antidote",
	Eyedrop:        "Eyedrop",
	Soft:           "Soft",
	Remedy:         "Remedy",
	SleepingBag:    "SleepingBag",
	Tent:           "Tent",
	GreenCherry:    "GreenCherry",
	Magicite:       "Magicite",
	SuperBall:      "SuperBall",
	EchoScreen:     "EchoScreen",
	SmokeBomb:      "SmokeBomb",
	TeleportStone:  "TeleportStone",
	DriedMeat:      "DriedMeat",
}

func (i Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(iname[i])
}
