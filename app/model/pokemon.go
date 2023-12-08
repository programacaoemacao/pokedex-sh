package model

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
