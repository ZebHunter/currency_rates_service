package repository

import (
	"encoding/xml"
	"fmt"
	"github.com/corpix/uarand"
	"golang.org/x/net/html/charset"
	"hw1/internal/pkg/model"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Rep struct {
}

func NewRep() *Rep {
	return &Rep{}
}

func (rp *Rep) FetchData(date string) (model.CBRes, error) {
	baseURL := "https://www.cbr.ru/scripts/XML_daily.asp"
	params := url.Values{}
	if date == "" {
		params.Add("date_req", time.Now().Format("02.01.2006"))
	} else {
		parsedDate, err := time.Parse("2006-01-02", date)
		if err != nil {
			fmt.Println(err)
			return model.CBRes{}, err
		}
		params.Add("date_req", parsedDate.Format("02.01.2006"))
	}

	urlParams := fmt.Sprintf("%s?%s", baseURL, params.Encode())
	fmt.Println("Sending request to:", urlParams)
	req, err := http.NewRequest("GET", urlParams, nil)
	if err != nil {
		return model.CBRes{}, err
	}

	req.Header.Add("User-Agent", uarand.GetRandom())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return model.CBRes{}, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	decoder := xml.NewDecoder(resp.Body)
	var result model.CBRes
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&result)

	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}
