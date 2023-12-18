package model

import "strconv"

const (
	LastPokemonID      = 1010
	formattingMaxChars = 4
)

type Pokemon struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Height        string   `json:"height"`
	Weight        string   `json:"weight"`
	Category      string   `json:"category"`
	Abilities     []string `json:"abilities"`
	AbilitiesInfo []string `json:"abilites_info"`
	Types         []string `json:"types"`
	Weakness      []string `json:"weakness"`
	ImageSrc      string   `json:"image_src"`
}

func (p *Pokemon) GetFormattedID() string {
	stringID := strconv.Itoa(p.ID)
	for len(stringID) < formattingMaxChars {
		stringID = "0" + stringID
	}
	return stringID
}
