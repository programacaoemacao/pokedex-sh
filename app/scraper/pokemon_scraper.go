package scraper

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
	"github.com/programacaoemacao/pokedex-sh/app/model"
)

type pokeScraper struct {
	collector *colly.Collector
	cookies   []*http.Cookie
}

func NewPokeScraper() *pokeScraper {
	return &pokeScraper{
		collector: colly.NewCollector(),
	}
}

func (s *pokeScraper) mountURLToScrape(id int) string {
	pokemonIDStr := strconv.Itoa(id)
	const totalOfChars = 4
	for len(pokemonIDStr) < totalOfChars {
		pokemonIDStr = "0" + pokemonIDStr
	}

	const baseURL = "https://sg.portal-pokemon.com/play/pokedex"
	url := fmt.Sprintf("%s/%s", baseURL, pokemonIDStr)
	return url
}

func (s *pokeScraper) getRequestHeaders() http.Header {
	headers := http.Header{
		"accept":                    []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7"},
		"accept-language":           []string{"pt-BR,pt;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6"},
		"cache-control":             []string{"max-age=0"},
		"sec-ch-ua":                 []string{"\"Microsoft Edge\";v=\"119\", \"Chromium\";v=\"119\", \"Not?A_Brand\";v=\"24\""},
		"sec-ch-ua-mobile":          []string{"?0"},
		"sec-ch-ua-platform":        []string{"\"Windows\""},
		"sec-fetch-dest":            []string{"document"},
		"sec-fetch-mode":            []string{"navigate"},
		"sec-fetch-site":            []string{"none"},
		"sec-fetch-user":            []string{"?1"},
		"upgrade-insecure-requests": []string{"1"},
		"user-agent":                []string{"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0"},
	}
	return headers
}

func (s *pokeScraper) scrape(url string) (*model.Pokemon, error) {
	c := s.collector.Clone()

	var pokemon *model.Pokemon = &model.Pokemon{}

	c.OnXML("//p[@class='pokemon-slider__main-no size-28']", func(x *colly.XMLElement) {
		id, err := strconv.Atoi(x.Text)
		if err == nil {
			pokemon.ID = id
		}
	})
	c.OnXML("//p[@class='pokemon-slider__main-name size-35']", func(x *colly.XMLElement) {
		pokemon.Name = strings.TrimSpace(x.Text)
	})
	c.OnXML("//span[@class='pokemon-info__title size-14'][text()='Height']/following-sibling::span", func(x *colly.XMLElement) {
		pokemon.Height = strings.TrimSpace(x.Text)
	})
	c.OnXML("//span[@class='pokemon-info__title size-14'][text()='Weight']/following-sibling::span", func(x *colly.XMLElement) {
		pokemon.Weight = strings.TrimSpace(x.Text)
	})
	c.OnXML("//div[contains(@class,'pokemon-info__category')]//span[2]/span", func(x *colly.XMLElement) {
		pokemon.Category = strings.TrimSpace(x.Text)
	})
	c.OnXML("//div[@class='pokemon-info__abilities']//span[contains(@class,'pokemon-info__value')]", func(x *colly.XMLElement) {
		pokemon.Abilities = append(pokemon.Abilities, strings.TrimSpace(x.Text))
	})
	c.OnXML("//span[@class='pokemon-info__value pokemon-info__value--body size-14']/span", func(x *colly.XMLElement) {
		pokemon.AbilitiesInfo = append(pokemon.AbilitiesInfo, strings.TrimSpace(x.Text))
	})
	c.OnXML("//div[contains(@class,'pokemon-type__type')]//span", func(x *colly.XMLElement) {
		pokemon.Types = append(pokemon.Types, strings.TrimSpace(x.Text))
	})
	c.OnXML("//div[contains(@class,'pokemon-weakness__btn')]//span", func(x *colly.XMLElement) {
		pokemon.Weakness = append(pokemon.Weakness, strings.TrimSpace(x.Text))
	})

	// The image on this website has a white shadow, so it's preferable to get the content from another place
	// c.OnXML("//img[@class='pokemon-img__front']", func(x *colly.XMLElement) {
	// 	pokemon.ImageSrc = "https://sg.portal-pokemon.com" + x.Attr("src")
	// })

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Request(http.MethodGet, url, nil, colly.NewContext(), s.getRequestHeaders())
	if err != nil {
		return nil, err
	}

	pokemon.ImageSrc = fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/official-artwork/%d.png", pokemon.ID)

	return pokemon, nil
}

func (s *pokeScraper) ScrapePokemonInfo(id int) (*model.Pokemon, error) {
	url := s.mountURLToScrape(id)

	pokemon, err := s.scrape(url)
	if err != nil {
		return nil, err
	}

	return pokemon, nil
}
