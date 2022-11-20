package cineparser

import (
	"fmt"
	"go_cineparse/pkg/models"
	"log"
	"strings"

	"github.com/anaskhan96/soup"
)

type Service struct {
	url string
}

func NewService(url string) *Service {
	return &Service{url: url}
}

func (s *Service) Echo() {
	fmt.Println("cineparser service")
}

// ----------------------------------
func (s *Service) GetDataShit() ([]*models.CineRow, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ParseURL - error")
		}
	}()

	dataRows := make([]*models.CineRow, 0)
	pageStr, err := soup.Get("https://vhost.fastserv.com/~cinemix/samPHPweb/web/playing.php?buster=06090525471")
	if err != nil {
		log.Println(err)
		return dataRows, err
	}
	//fmt.Println(pageStr)
	doc := soup.HTMLParse(pageStr)

	listTR := doc.Find("div", "id", "recently_played").Find("tbody").FindAll("tr")
	fmt.Println("len=", len(listTR)) //fmt.Println(strings.TrimSpace(div1[0].HTML()))

	for _, elem := range listTR {
		val, err := ParseRow(elem)
		if err == nil {
			dataRows = append(dataRows, val)
		}
	}
	//
	return dataRows, nil
}

func ParseRow(elem soup.Root) (*models.CineRow, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ParseRow - error")
		}
	}()
	listTD := elem.FindAll("td")
	tmp := &models.CineRow{
		AuthorTitle: strings.TrimSpace(listTD[1].Text()),
		Album:       strings.TrimSpace(listTD[2].Text()),
		Duration:    strings.TrimSpace(listTD[3].Text()),
	}

	return tmp, nil
}
