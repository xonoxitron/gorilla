package models

type Ticker struct {
	Name                       string
	Ask                        []float32 // a = ask array(<price>, <whole lot volume>, <lot volume>)
	Bid                        []float32 // b = bid array(<price>, <whole lot volume>, <lot volume>)
	LastTradeClosed            []float32 // c = last trade closed array(<price>, <lot volume>)
	Volume                     []float32 // v = volume array(<today>, <last 24 hours>)
	VolumeWeightedAveragePrice []float32 // p = volume weighted average price array(<today>, <last 24 hours>)
	NumberOfTrades             []float32 // t = number of trades array(<today>, <last 24 hours>)
	Low                        []float32 // l = low array(<today>, <last 24 hours>)
	High                       []float32 // h = high array(<today>, <last 24 hours>)
	OpeningPrice               float32   // o = today's opening price
}
