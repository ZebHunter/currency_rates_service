package services

import (
	"errors"
	"fmt"
	"hw1/configs"
	"hw1/internal/pkg/model"
	"strconv"
	"strings"
)

var (
	errorInfo     = errors.New("impossible to get info")
	errorCurrency = errors.New("impossible to get currency")
)

type RepImpl interface {
	FetchData(data string) (model.CBRes, error)
}

type Service struct {
	rp RepImpl
}

func NewService(rp RepImpl) *Service {
	return &Service{rp}
}

func (s *Service) GetInfo() (model.InfoRes, error) {
	jsonInfo := model.InfoRes{
		Version: configs.GetEnv("VERSION"),
		Service: "currency",
		Author:  "m.rogachev",
	}

	return jsonInfo, nil
}

func (s *Service) GetCurrency(data string, name string) (model.CurrencyRes, error) {
	currency, err := s.rp.FetchData(data)
	if err != nil {
		return model.CurrencyRes{}, err
	}
	cur := model.CurrencyRes{Service: "currency", Data: map[string]float64{}}

	for _, v := range currency.Currencies {
		if v.CharCode == name || name == "" {
			v.Exchange = strings.Replace(v.Exchange, ",", ".", 1)
			cur.Data[v.CharCode], _ = strconv.ParseFloat(v.Exchange, 64)
			if name != "" {
				break
			}
		}
	}
	fmt.Println(cur)
	return cur, nil
}
