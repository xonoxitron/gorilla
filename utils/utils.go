package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/tidwall/gjson"
	"github.com/xonoxitron/gorilla/config"
	"github.com/xonoxitron/gorilla/parser"
)

var crlf = "\r\n"

func Log(message string) {
	log.Println("[gorilla] " + message)
}

func ExtractAssets(payload gjson.Result) []string {
	var assets []string

	payload.ForEach(func(key, value gjson.Result) bool {
		assets = append(assets, key.String())

		return true
	})

	return assets
}

func ExtractAssetPairs(payload gjson.Result, currency string) []string {
	var assetPairs []string

	payload.ForEach(func(key, value gjson.Result) bool {
		assetPair := key.String()

		if strings.HasSuffix(assetPair, currency) {
			assetPairs = append(assetPairs, assetPair)
		}

		return true
	})

	return assetPairs
}

func ExtractTickerData(payload gjson.Result) []string {
	var tickerData []string

	payload.ForEach(func(key, value gjson.Result) bool {
		var data = key.String() + "="

		value.ForEach(func(key, value gjson.Result) bool {
			data += key.String() + ":" + strings.NewReplacer("[", "", "]", "", "\"", "").Replace(value.String()) + ";"
			return true
		})

		tickerData = append(tickerData, data)

		return true
	})

	return tickerData
}

func ToMap(slice []string) map[string]string {
	m := make(map[string]string)

	for _, v := range slice {
		m[v] = v
	}

	return m
}

func GetDifference(oldAssets []string, newAssets []string) []string {
	var difference []string
	oldAssetsMap := ToMap(oldAssets)
	newAssetsMap := ToMap(newAssets)

	for _, asset := range newAssets {
		if _, ok := oldAssetsMap[asset]; !ok {
			difference = append(difference, asset)
		}
	}

	for _, asset := range oldAssets {
		if _, ok := newAssetsMap[asset]; !ok {
			difference = append(difference, asset)
		}
	}

	return difference
}

func EvaluateAssets(oldAssets []string, newAssets []string) string {
	response := ""

	if oldAssets[0] != "" {
		assetsDifference := GetDifference(oldAssets, newAssets)
		oldAssetsCount := len(oldAssets)
		newAssetsCount := len(newAssets)

		if len(assetsDifference) != 0 {
			assets := strings.Join(assetsDifference, crlf)
			response += "⚠️[Alert]" + crlf

			if oldAssetsCount > newAssetsCount {
				response += "Asset/s not more supported:" + crlf + assets
			} else if newAssetsCount > oldAssetsCount {
				response += "New asset/s supported:" + crlf + assets
			}

			response += crlf
		}
	}

	return response
}

func EvaluateTickers(oldTickersData []string, newTickersData []string) string {
	response := ""

	if oldTickersData[0] != "" && len(oldTickersData) == len(newTickersData) {
		oldTickers := parser.GetTickers(oldTickersData)
		newTickers := parser.GetTickers(newTickersData)
		currency := config.Get().Currency
		info := ""

		for t := 0; t < len(newTickersData); t++ {
			oldTicker := oldTickers[t]
			newTicker := newTickers[t]

			if oldTicker.Name != newTicker.Name {
				panic("tickers not aligned")
			}

			if newTicker.Low[0] > oldTicker.Low[0] {
				info += "⬇️ " + newTicker.Name + " reached a new LOW: " + fmt.Sprintf("%f", newTicker.High[0]) + currency + crlf
			}

			if newTicker.High[0] > oldTicker.High[0] {
				info += "⬆️ " + newTicker.Name + " reached a new HIGH: " + fmt.Sprintf("%f", newTicker.High[0]) + currency + crlf
			}
		}

		if info != "" {
			response += "ℹ️[Info]" + crlf + info
		}

	}
	return response
}

func ErrorCheck(e error) {
	if e != nil {
		panic(e.Error())
	}
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}
