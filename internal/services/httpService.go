package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type AgeModel struct {
	Count int
	Name  string
	Age   int
}

type SexModel struct {
	Count       int
	Name        string
	Gender      string
	Probability float64
}

type nationCountryModel struct {
	Country_id  string
	Probability float64
}

type NationModel struct {
	Count   int
	Name    string
	Country []nationCountryModel
}

func FetchAge(url string) AgeModel {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	}

	req.Header.Add("Content-type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	data := new(AgeModel)
	unMarshalErr := json.Unmarshal(body, &data)
	if unMarshalErr != nil {
		log.Fatalf("Failed to unmarshal response body: %v", unMarshalErr)
	}
	defer resp.Body.Close()
	return *data
}

func FetchSex(url string) SexModel {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	}

	req.Header.Add("Content-type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	data := new(SexModel)
	unMarshalErr := json.Unmarshal(body, &data)
	if unMarshalErr != nil {
		log.Fatalf("Failed to unmarshal response body: %v", unMarshalErr)
	}
	defer resp.Body.Close()
	return *data
}

func FetchNation(url string) NationModel {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	}

	req.Header.Add("Content-type", "application/json; charset=utf-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}
	data := new(NationModel)
	unMarshalErr := json.Unmarshal(body, &data)

	if unMarshalErr != nil {
		log.Fatalf("Failed to unmarshal response body: %v", unMarshalErr)
	}
	defer resp.Body.Close()
	return *data
}
