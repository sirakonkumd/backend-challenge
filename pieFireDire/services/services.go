package services

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2/log"
)

type PieFireDireResponse struct {
	Resp map[string]int `json:"beef"`
}

func PieFireDireService() (PieFireDireResponse, error) {
	var res PieFireDireResponse
	resMeat, err := callGetMeat()
	if err != nil {
		log.Error("call 3rd error ", err)
		return res, err
	}

	var mapBeef = make(map[string]int)
	findTxt := regexp.MustCompile(`\s+|\.|,`)
	sptxt := findTxt.Split(strings.ToLower(string(resMeat)), -1)
	for _, beef := range sptxt {
		if beef != "" {
			if _, ok := mapBeef[beef]; !ok {
				mapBeef[beef] = 1
			} else {
				mapBeef[beef]++
			}
		}
	}
	// fmt.Println("=>   ", len(mapBeef))
	res.Resp = mapBeef
	return res, nil
}
