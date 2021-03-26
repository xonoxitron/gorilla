package scheduler

import (
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/xonoxitron/gorilla/api"
	"github.com/xonoxitron/gorilla/bot"
	"github.com/xonoxitron/gorilla/config"
	"github.com/xonoxitron/gorilla/storage"
	"github.com/xonoxitron/gorilla/utils"
)

func StartJob() {
	config := config.Get()
	utils.Log("starting bot")
	go bot.Start()

	utils.Log("starting scheduler")
	scheduler := gocron.NewScheduler(time.UTC)
	scheduler.Every(config.Interval).Seconds().Do(func() {
		// fetch new data
		newAssets := utils.ExtractAssets(api.GetKrakenPublicAssets())
		assetPairs := utils.ExtractAssetPairs(api.GetKrakenPublicAssetPairs(), config.Currency)
		newTickersData := utils.ExtractTickerData(api.GetKrakenPublicTicker(strings.Join(assetPairs, ",")))

		// output new data
		if config.Debug {
			utils.Log("public assets: " + strings.Join(newAssets, ","))
			utils.Log("asset pairs: " + strings.Join(assetPairs, ","))
			utils.Log("tickers: " + strings.Join(newTickersData, "\r"))
		}

		// fetch old data
		oldAssets := strings.Split(storage.Get("assets"), ",")
		oldTickersData := strings.Split(storage.Get("tickers"), "\r")

		// evaluate data
		assetsEvaluation := utils.EvaluateAssets(oldAssets, newAssets)
		tickersEvaluation := utils.EvaluateTickers(oldTickersData, newTickersData)

		// notify
		if assetsEvaluation != "" {
			go bot.Notify(assetsEvaluation)
		}
		if tickersEvaluation != "" {
			go bot.Notify(tickersEvaluation)
		}

		// store new data
		storage.Update("assets", strings.Join(newAssets, ","), true)
		storage.Update("tickers", strings.Join(newTickersData, "\r"), true)

	})
	scheduler.StartBlocking()
}
