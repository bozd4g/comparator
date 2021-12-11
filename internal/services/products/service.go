package products

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/bozd4g/comparator/internal/services/configs"
)

func New(configService configs.Servicer) Service {
	return Service{
		configService: configService,
	}
}

func (s Service) GetAll(name string) ([]Dto, error) {
	configs, err := s.configService.GetAll()
	if err != nil {
		return nil, err
	}

	var dtos = make([]Dto, 0)
	for _, config := range configs {
		configDtos := s.collectDataFromSite(name, config)
		for name, dto := range configDtos {
			dtos = append(dtos, Dto{
				Name:    name,
				Product: dto[0],
			})
		}
	}

	return dtos, nil
}

func (s Service) GetAllByConfig(name, config string) ([]Dto, error) {
	c, err := s.configService.GetByName(config)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New("An error occurred while retrieving the config!")
	}

	var dtos = make([]Dto, 0)
	grouppedDtos := s.collectDataFromSite(name, *c)
	for name, dto := range grouppedDtos {
		if len(dto) == 0 {
			continue
		}

		dtos = append(dtos, Dto{Name: name, Product: dto[0]})
	}

	return dtos, nil
}

func (s Service) GetAllMultipleByConfig(names []string, config string) ([]MultipleDto, error) {
	c, err := s.configService.GetByName(config)
	if err != nil {
		return nil, err
	}

	if c == nil {
		return nil, errors.New("An error occurred while retrieving the config!")
	}

	var multipleDtos = make([]MultipleDto, 0)
	var grouppedDtos = make(map[string][]ProductDto)
	for _, name := range names {
		dtos := s.collectDataFromSite(name, *c)
		for site, dto := range dtos {
			grouppedDtos[site] = append(grouppedDtos[site], dto...)
		}
	}

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

func (s Service) collectDataFromSite(name string, config configs.Dto) map[string][]ProductDto {
	var dtos = make(map[string][]ProductDto)
	encodedName := url.QueryEscape(name)

	for _, site := range config.Sites {
		url := fmt.Sprintf("%s%s", site.Address, fmt.Sprintf(site.Search, encodedName))
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()
		if res.StatusCode != 200 {
			fmt.Println(fmt.Sprintf("Unexpected status code(%d) for %s from %s", res.StatusCode, name, url))
			continue
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			fmt.Println(fmt.Sprintf("An error occured while reading the body of response! Site: %s, Error: %+v", url, err))
			continue
		}

		var dto ProductDto

		link := doc.Find(site.Link)
		if len(link.Nodes) != 0 {
			dto.Link = link.First().AttrOr("href", "")
		}

		value := doc.Find(site.Value)
		if len(value.Nodes) != 0 {
			dto.Price = s.clearCurrencies(value.First().Text())
		}

		if dto.Price != 0 {
			dtos[site.Name] = append(dtos[site.Name], dto)
		}
	}

	return dtos
}

func (s Service) clearCurrencies(price string) float64 {
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
