package products

import (
	"errors"
	"fmt"
	"github.com/bozd4g/comparator/internal/services/configs"
	"github.com/gocolly/colly/v2"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

func New(config configs.Service) Service {
	return service{
		collector:     colly.NewCollector(colly.AllowURLRevisit(), colly.UserAgent(SAMPLE_USER_AGENT)),
		configService: config,
	}
}

func (s service) GetAll(name string) ([]Dto, error) {
	configs, err := s.configService.GetAll()
	if err != nil {
		return nil, err
	}

	var dtos = make([]Dto, 0)
	for _, config := range configs {
		configDtos := s.collectDataFromSite([]string{name}, config)
		for name, dto := range configDtos {
			dtos = append(dtos, Dto{
				Name:    name,
				Product: dto[0],
			})
		}
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

	var dtos = make([]Dto, 0)
	grouppedDtos := s.collectDataFromSite([]string{name}, *config)
	for name, dto := range grouppedDtos {
		if len(dto) == 0 {
			continue
		}

		dtos = append(dtos, Dto{Name: name, Product: dto[0]})
	}

	return dtos, nil
}

func (s service) GetAllMultipleByCategory(names []string, category string) ([]MultipleDto, error) {
	config, err := s.configService.GetByName(category)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, errors.New("An error occurred while retrieving the config!")
	}

	var multipleDtos = make([]MultipleDto, 0)
	grouppedDtos := s.collectDataFromSite(names, *config)
	for name, dtos := range grouppedDtos {
		var sum float64 = 0
		for _, v := range dtos {
			sum += v.Price
		}
		multipleDtos = append(multipleDtos, MultipleDto{
			Name:     name,
			Products: dtos,
			Total:    sum,
		})
	}

	return multipleDtos, err
}

func (s service) collectDataFromSite(names []string, config configs.Dto) map[string][]ProductDto {
	var dtos = make(map[string][]ProductDto)
	for _, name := range names {
		for _, site := range config.Sites {
			sort.Slice(site.Steps, func(i, j int) bool {
				return site.Steps[i].Id > site.Steps[j].Id
			})

			var dto ProductDto
			for _, step := range site.Steps {
				encodedName := url.PathEscape(name)

				if step.Action == configs.SEARCH {
					site := fmt.Sprintf("%s%s", site.Address, fmt.Sprintf(step.Selector, encodedName))
					err := s.collector.Visit(site)
					if err != nil {
						break
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

			if dto.Price != 0 {
				dtos[site.Name] = append(dtos[site.Name], dto)
			}
		}
	}
	return dtos
}

func (s service) clearCurrencies(price string) float64 {
	currencies := []string{"TL", "TRY", "EUR", "USD"}

	for _, currency := range currencies {
		price = strings.TrimLeft(price, currency)
		price = strings.TrimRight(price, currency)
		price = strings.TrimSpace(price)
	}

	if !strings.Contains(price, ".") {
		price = strings.ReplaceAll(price, ",", ".")
	}

	priceAsFloat, err := strconv.ParseFloat(strings.TrimSpace(price), 64)
	if err != nil {
		return 0
	}

	return priceAsFloat
}
