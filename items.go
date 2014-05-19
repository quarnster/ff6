// Copyright 2014
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
)

type (
	Item      uint8
	Inventory struct {
		SortOrder [288]Item
		Count     [288]uint8
	}
)

const (

	// Dirk (0)
	Dagger Item = iota
	MythrilKnife
	MainGauche
	AirKnife
	ThiefsKnife
	AssassinsDagger
	ManEater
	SwordBreaker
	Gladius
	ValiantKnife
	// Sword (0xA)
	MythrilSword
	GreatSword
	RuneBlade
	Flametounge
	Icebrand
	ThunderBlade
	BastardSword
	Stoneblade
	BloodSword
	Enhancer
	CrystalSword
	Falchion
	SoulSabre
	Organyx
	Excalibur
	Zantetsuken
	Lightbringer
	Ragnarok
	UltimaWeapon
	// Lance (0x1d)
	MythrilSpear
	Trident
	HeavyLance
	Partisian
	HolyLance
	GoldLance
	RadiantLance
	Impartisian
	//  Dirk (0x25)
	Kunai
	Kodachi
	Sakura
	Sasuke
	Ichigeki
	Kagenui
	// Knife (0x2B)
	Ashura
	Kotetsu
	Kikuichimonji
	Kazekiri
	Murasame
	Masamune
	Murakuma
	Mutsunokami
	// Rod (0x33)
	HealingRod
	MythrilRod
	FlameRod
	IceRod
	ThunderRod
	PoisonRod
	HolyRod
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
	FumaShuriken
	Pinwheel
	// Special (0x44)
	ChainFlail
	MoonringBlade
	MorningStar
	Boomerang
	RisingSun
	Hawkeye
	BoneClub
	Sniper
	WingEdge
	// Gambler (0x4d)
	Cards
	Darts
	DeathTarot
	ViperDarts
	Dice
	FixedDice
	// Claw (0x53)
	MetalKnuckle
	MythrilClaw
	KaiserKnuckle
	VenomClaws
	BurningFist
	DragonClaws
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
	FlameScroll
	WaterScroll
	LightningScroll
	InvisibilityScroll
	ShadowScroll
	// Relics (0xB0)
	SilverSpectacles
	StarPendant
	PeaceRing
	Amulet
	WhiteCape
	JeweledRing
	FairyRing
	BarrierRing
	MythrilGlove
	ProtectRing
	HermesSandals
	ReflectRing
	AngelWings
	AngelRing
	KnightsCode
	DragonBoots
	ZephyrCloak
	PrincessRing
	CursedCing
	Earrings
	GigasGlove
	BlizzardRing
	BerserkerRing
	ThiefsBracer
	GuardBracelet
	HerosRing
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
	PrayerBeads
	BlackBelt
	HejisJitte
	FakeMoustache
	SoulOfTamasa
	DragonHorn
	MeritAward
	MementoRing
	SafetyBit
	LichRing
	MolulusCharm
	WardBangle
	MiracleShoes
	AlarmEarring
	GaleHairpin
	SniperEye
	GrowthEgg
	Tintinabulum
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
	Dagger:             "Dagger",
	MythrilKnife:       "MythrilKnife",
	MainGauche:         "MainGauche",
	AirKnife:           "AirKnife",
	ThiefsKnife:        "ThiefsKnife",
	AssassinsDagger:    "AssassinsDagger",
	ManEater:           "ManEater",
	SwordBreaker:       "SwordBreaker",
	Gladius:            "Gladius",
	ValiantKnife:       "ValiantKnife",
	MythrilSword:       "MythrilSword",
	GreatSword:         "GreatSword",
	RuneBlade:          "RuneBlade",
	Flametounge:        "Flametounge",
	Icebrand:           "Icebrand",
	ThunderBlade:       "ThunderBlade",
	BastardSword:       "BastardSword",
	Stoneblade:         "Stoneblade",
	BloodSword:         "BloodSword",
	Enhancer:           "Enhancer",
	CrystalSword:       "CrystalSword",
	Falchion:           "Falchion",
	SoulSabre:          "SoulSabre",
	Organyx:            "Organyx",
	Excalibur:          "Excalibur",
	Zantetsuken:        "Zantetsuken",
	Lightbringer:       "Lightbringer",
	Ragnarok:           "Ragnarok",
	UltimaWeapon:       "UltimaWeapon",
	MythrilSpear:       "MythrilSpear",
	Trident:            "Trident",
	HeavyLance:         "HeavyLance",
	Partisian:          "Partisian",
	HolyLance:          "HolyLance",
	GoldLance:          "GoldLance",
	RadiantLance:       "RadiantLance",
	Impartisian:        "Impartisian",
	Kunai:              "Kunai",
	Kodachi:            "Kodachi",
	Sakura:             "Sakura",
	Sasuke:             "Sasuke",
	Ichigeki:           "Ichigeki",
	Kagenui:            "Kagenui",
	Ashura:             "Ashura",
	Kotetsu:            "Kotetsu",
	Kikuichimonji:      "Kikuichimonji",
	Kazekiri:           "Kazekiri",
	Murasame:           "Murasame",
	Masamune:           "Masamune",
	Murakuma:           "Murakuma",
	Mutsunokami:        "Mutsunokami",
	HealingRod:         "HealingRod",
	MythrilRod:         "MythrilRod",
	FlameRod:           "FlameRod",
	IceRod:             "IceRod",
	ThunderRod:         "ThunderRod",
	PoisonRod:          "PoisonRod",
	HolyRod:            "HolyRod",
	GravityRod:         "GravityRod",
	Punisher:           "Punisher",
	MagusRod:           "MagusRod",
	ChocoboBrush:       "ChocoboBrush",
	DaVinciBrush:       "DaVinciBrush",
	MagicalBrush:       "MagicalBrush",
	RainbowBrush:       "RainbowBrush",
	Shuriken:           "Shuriken",
	FumaShuriken:       "FumaShuriken",
	Pinwheel:           "Pinwheel",
	ChainFlail:         "ChainFlail",
	MoonringBlade:      "MoonringBlade",
	MorningStar:        "MorningStar",
	Boomerang:          "Boomerang",
	RisingSun:          "RisingSun",
	Hawkeye:            "Hawkeye",
	BoneClub:           "BoneClub",
	Sniper:             "Sniper",
	WingEdge:           "WingEdge",
	Cards:              "Cards",
	Darts:              "Darts",
	DeathTarot:         "DeathTarot",
	ViperDarts:         "ViperDarts",
	Dice:               "Dice",
	FixedDice:          "FixedDice",
	MetalKnuckle:       "MetalKnuckle",
	MythrilClaw:        "MythrilClaw",
	KaiserKnuckle:      "KaiserKnuckle",
	VenomClaws:         "VenomClaws",
	BurningFist:        "BurningFist",
	DragonClaws:        "DragonClaws",
	TigerFang:          "TigerFang",
	Buckler:            "Buckler",
	HeavyShield:        "HeavyShield",
	MythrilShield:      "MythrilShield",
	GoldShield:         "GoldShield",
	AegisShield:        "AegisShield",
	DiamondShield:      "DiamondShield",
	FlameShield:        "FlameShield",
	IceShield:          "IceShield",
	ThunderShield:      "ThunderShield",
	CrystalShield:      "CrystalShield",
	GenjiShield:        "GenjiShield",
	TortoiseShield:     "TortoiseShield",
	CursedShield:       "CursedShield",
	PaladinShield:      "PaladinShield",
	ForceShield:        "ForceShield",
	LeatherHat:         "LeatherHat",
	HairBand:           "HairBand",
	PlumedHat:          "PlumedHat",
	Beret:              "Beret",
	MagusHat:           "MagusHat",
	Bandana:            "Bandana",
	IronHelmet:         "IronHelmet",
	Coronet:            "Coronet",
	BardHat:            "BardHat",
	GreenBeret:         "GreenBeret",
	HeadBand:           "HeadBand",
	MithrilHelm:        "MithrilHelm",
	Tiara:              "Tiara",
	GoldHelmet:         "GoldHelmet",
	TigerMask:          "TigerMask",
	RedCap:             "RedCap",
	MysteryVeil:        "MysteryVeil",
	Circlet:            "Circlet",
	RoyalCrown:         "RoyalCrown",
	DiamondHelm:        "DiamondHelm",
	DarkHood:           "DarkHood",
	CrystalHelm:        "CrystalHelm",
	OathViel:           "OathViel",
	CatHood:            "CatHood",
	GenjiHelm:          "GenjiHelm",
	Thornlet:           "Thornlet",
	Titanium:           "Titanium",
	LeatherArmor:       "LeatherArmor",
	CottonRobe:         "CottonRobe",
	KungFuSuit:         "KungFuSuit",
	IronArmor:          "IronArmor",
	SilkRobe:           "SilkRobe",
	MithrilVest:        "MithrilVest",
	NinjaGear:          "NinjaGear",
	WhiteDress:         "WhiteDress",
	MithrilMail:        "MithrilMail",
	GaiaGear:           "GaiaGear",
	MirageDress:        "MirageDress",
	GoldArmor:          "GoldArmor",
	PowerSash:          "PowerSash",
	LightRobe:          "LightRobe",
	DiamondVest:        "DiamondVest",
	RedJacket:          "RedJacket",
	ForceArmor:         "ForceArmor",
	DiamondArmor:       "DiamondArmor",
	BlackGarb:          "BlackGarb",
	TaoRobe:            "TaoRobe",
	CrystalMail:        "CrystalMail",
	CzarinaGown:        "CzarinaGown",
	GenjiArmor:         "GenjiArmor",
	ReedCloak:          "ReedCloak",
	MinervaBustier:     "MinervaBustier",
	TabbySuit:          "TabbySuit",
	ChocoboSuit:        "ChocoboSuit",
	MoogleSuit:         "MoogleSuit",
	NutkinSuit:         "NutkinSuit",
	BehemothSuit:       "BehemothSuit",
	SnowMuffler:        "SnowMuffler",
	NoiseBlaster:       "NoiseBlaster",
	BioBlaster:         "BioBlaster",
	Flash:              "Flash",
	ChainSaw:           "ChainSaw",
	Debilitator:        "Debilitator",
	Drill:              "Drill",
	AirAnchor:          "AirAnchor",
	AutoCrossbow:       "AutoCrossbow",
	FlameScroll:        "FlameScroll",
	WaterScroll:        "WaterScroll",
	LightningScroll:    "LightningScroll",
	InvisibilityScroll: "InvisibilityScroll",
	ShadowScroll:       "ShadowScroll",
	SilverSpectacles:   "SilverSpectacles",
	StarPendant:        "StarPendant",
	PeaceRing:          "PeaceRing",
	Amulet:             "Amulet",
	WhiteCape:          "WhiteCape",
	JeweledRing:        "JeweledRing",
	FairyRing:          "FairyRing",
	BarrierRing:        "BarrierRing",
	MythrilGlove:       "MythrilGlove",
	ProtectRing:        "ProtectRing",
	HermesSandals:      "HermesSandals",
	ReflectRing:        "ReflectRing",
	AngelWings:         "AngelWings",
	AngelRing:          "AngelRing",
	KnightsCode:        "KnightsCode",
	DragonBoots:        "DragonBoots",
	ZephyrCloak:        "ZephyrCloak",
	PrincessRing:       "PrincessRing",
	CursedCing:         "CursedCing",
	Earrings:           "Earrings",
	GigasGlove:         "GigasGlove",
	BlizzardRing:       "BlizzardRing",
	BerserkerRing:      "BerserkerRing",
	ThiefsBracer:       "ThiefsBracer",
	GuardBracelet:      "GuardBracelet",
	HerosRing:          "HerosRing",
	Ribbon:             "Ribbon",
	MuscleBelt:         "MuscleBelt",
	CrystalOrb:         "CrystalOrb",
	GoldHairpin:        "GoldHairpin",
	Celestriad:         "Celestriad",
	BrigandsGlove:      "BrigandsGlove",
	Gauntlet:           "Gauntlet",
	GenjiGlove:         "GenjiGlove",
	HyperWrist:         "HyperWrist",
	MastersScroll:      "MastersScroll",
	PrayerBeads:        "PrayerBeads",
	BlackBelt:          "BlackBelt",
	HejisJitte:         "HejisJitte",
	FakeMoustache:      "FakeMoustache",
	SoulOfTamasa:       "SoulOfTamasa",
	DragonHorn:         "DragonHorn",
	MeritAward:         "MeritAward",
	MementoRing:        "MementoRing",
	SafetyBit:          "SafetyBit",
	LichRing:           "LichRing",
	MolulusCharm:       "MolulusCharm",
	WardBangle:         "WardBangle",
	MiracleShoes:       "MiracleShoes",
	AlarmEarring:       "AlarmEarring",
	GaleHairpin:        "GaleHairpin",
	SniperEye:          "SniperEye",
	GrowthEgg:          "GrowthEgg",
	Tintinabulum:       "Tintinabulum",
	SprintShoes:        "SprintShoes",
	RenameCard:         "RenameCard",
	Potion:             "Potion",
	HiPotion:           "HiPotion",
	XPotion:            "XPotion",
	Ether:              "Ether",
	HiEther:            "HiEther",
	XEther:             "XEther",
	Elixir:             "Elixir",
	Megalixir:          "Megalixir",
	PhoenixDown:        "PhoenixDown",
	Revivify:           "Revivify",
	Antidote:           "Antidote",
	Eyedrop:            "Eyedrop",
	Soft:               "Soft",
	Remedy:             "Remedy",
	SleepingBag:        "SleepingBag",
	Tent:               "Tent",
	GreenCherry:        "GreenCherry",
	Magicite:           "Magicite",
	SuperBall:          "SuperBall",
	EchoScreen:         "EchoScreen",
	SmokeBomb:          "SmokeBomb",
	TeleportStone:      "TeleportStone",
	DriedMeat:          "DriedMeat",
	Nothing:            "Nothing",
}

func (i Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(iname[i])
}

func (i *Item) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	for k, v := range iname {
		if v == s {
			*i = k
			return nil
		}
	}
	return fmt.Errorf("\"%s\" isn't a known Item", s)
}

func (inv *Inventory) UnmarshalJSON(data []byte) error {
	d := make(map[string]json.RawMessage)
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	if err := json.Unmarshal(d["SortingOrder"], &inv.SortOrder); err != nil {
		return err
	}
	for i, k := range inv.SortOrder {
		if err := json.Unmarshal(d[iname[k]], &inv.Count[i]); err != nil {
			return err
		}
	}
	for i := 256; i < len(inv.Count); i++ {
		inv.Count[i] = 0
		inv.SortOrder[i] = 0xff
	}
	return nil
}

func (inv Inventory) MarshalJSON() ([]byte, error) {
	d := make(map[string]interface{})
	d["SortingOrder"] = &inv.SortOrder
	for i, k := range inv.SortOrder {
		d[iname[k]] = inv.Count[i]
	}
	return json.Marshal(d)
}
