package parser

import (
	"strconv"
	"strings"

	"github.com/xonoxitron/gorilla/models"
)

func ToFloat32(rawNumber string) float32 {
	var n float64
	var err error

	n, err = strconv.ParseFloat(rawNumber, 32)

	if err != nil {
		panic("cannot convert to float32")
	}

	return float32(n)
}

func ToFloat32Array(rawData string) []float32 {
	rawNumbers := strings.Split(rawData, ",")
	var numbers []float32

	for _, rawNumber := range rawNumbers {
		numbers = append(numbers, ToFloat32(rawNumber))
	}

	return numbers
}

func ToTicker(rawTickerData string) models.Ticker {
	var ticker models.Ticker
	chunks := strings.Split(rawTickerData, "=")
	ticker.Name = chunks[0]
	data := strings.Split(strings.TrimSpace(chunks[1]), ";")

	for _, d := range data {
		if d != "" {
			splits := strings.Split(d, ":")
			switch splits[0] {
			case "a":
				ticker.Ask = ToFloat32Array(splits[1])
			case "b":
				ticker.Bid = ToFloat32Array(splits[1])
			case "c":
				ticker.LastTradeClosed = ToFloat32Array(splits[1])
			case "v":
				ticker.Volume = ToFloat32Array(splits[1])
			case "p":
				ticker.VolumeWeightedAveragePrice = ToFloat32Array(splits[1])
			case "t":
				ticker.NumberOfTrades = ToFloat32Array(splits[1])
			case "l":
				ticker.Low = ToFloat32Array(splits[1])
			case "h":
				ticker.High = ToFloat32Array(splits[1])
			case "o":
				ticker.OpeningPrice = ToFloat32(splits[1])
			}
		}
	}

	return ticker
}

func GetTickers(rawTickersData []string) []models.Ticker {
	var tickers []models.Ticker

	for _, rawTickerData := range rawTickersData {
		tickers = append(tickers, ToTicker(rawTickerData))
	}

	return tickers
}
