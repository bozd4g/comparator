package products

import (
	"errors"
	"fmt"
	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gocolly/colly/v2"
	"sort"
	"strings"
)

func New(config configs.Service) Service {
	return service{collector: colly.NewCollector(colly.AllowURLRevisit(), colly.UserAgent(SAMPLE_USER_AGENT)), configService: config}
}

func (s service) GetAll(name string) ([]Dto, error) {
	configs, err := s.configService.GetAll()
	if err != nil {
		return nil, err
	}

	var dtos []Dto
	for _, config := range configs {
		configDtos := s.collectDataFromSite(name, config)
		dtos = append(dtos, configDtos...)
	}

	return dtos, nil
}

func (s service) GetAllByCategory(name, category string) ([]Dto, error) {
	config, err := s.configService.GetByName(category)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("An error occurred while retrieving the config!")
	}

	dtos := s.collectDataFromSite(name, *config)
	return dtos, nil
}

func (s service) collectDataFromSite(name string, config configs.Dto) []Dto {
	var dtos []Dto
	for _, site := range config.Sites {
		dto := Dto{Name: site.Name}

		sort.Slice(site.Steps, func(i, j int) bool {
			return site.Steps[i].Id > site.Steps[j].Id
		})

		s.collector.OnError(func(response *colly.Response, err error) {
			fmt.Println(err)
		})

		for _, step := range site.Steps {
			if step.Action == configs.SEARCH {
				err := s.collector.Visit(fmt.Sprintf("%s%s", site.Address, fmt.Sprintf(step.Selector, name)))
				if err != nil {
					fmt.Println(err)
				}
			} else if step.Action == configs.LINK {
				s.collector.OnHTML(step.Selector, func(e *colly.HTMLElement) {
					href := e.Attr("href")
					if strings.Contains(href, site.Address) {
						dto.Link = href
					} else {
						dto.Link = fmt.Sprintf("%s%s", site.Address, href)
					}
				})
			} else if step.Action == configs.VALUE {
				s.collector.OnHTML(step.Selector, func(e *colly.HTMLElement) {
					dto.Price = s.clearCurrencies(e.Text)
				})
			}
		}

		dtos = append(dtos, dto)
	}

	return dtos
}

func (s service) clearCurrencies(price string) string {
	currencies := []string{"TL", "TRY", "EUR", "USD"}
	price = strings.Replace(price, ".", ",", 1)

	for _, currency := range currencies {
		price = strings.TrimLeft(price, currency)
		price = strings.TrimRight(price, currency)
		price = strings.TrimSpace(price)
	}

	return price
}
