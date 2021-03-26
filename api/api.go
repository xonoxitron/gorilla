package api

import (
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/xonoxitron/gorilla/utils"
)

func Request(uri string) []byte {
	res, err := http.Get(uri)
	utils.ErrorCheck(err)
	resBody, err := ioutil.ReadAll(res.Body)
	utils.ErrorCheck(err)
	return resBody
}

func GetKrakenPublicAssets() gjson.Result {
	return gjson.Parse(string(Request("https://api.kraken.com/0/public/Assets"))).Get("result")
}

func GetKrakenPublicAssetPairs() gjson.Result {
	return gjson.Parse(string(Request("https://api.kraken.com/0/public/AssetPairs"))).Get("result")
}

func GetKrakenPublicTicker(pair string) gjson.Result {
	return gjson.Parse(string(Request("https://api.kraken.com/0/public/Ticker?pair=" + pair))).Get("result")
}
