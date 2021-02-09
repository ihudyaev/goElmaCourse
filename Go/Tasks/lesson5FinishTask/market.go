package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//Duration Time - seconds
const DurationTime = 2

// Market Interface. Return read only channel
type MarketInterface interface {
	getMarketData() <-chan MarketDataUnion
}

// Union data for two Markets
// Using same fields
// Source - trade source name
type MarketDataUnion struct {
	Source string
	ID     int
	Date   string
	Price  float64
	Amount float64
}

// Get data from api
func getDataBodyFromApi(url string) (*http.Response, error) {
	response, err := http.Get(url)
	fmt.Println(response.Body)
	if err != nil {
		return nil, errors.New("Request Error")
	}

	return response, nil
}

//Market Share
type MarketDataStore []MarketDataUnion

//Function for REault Output (Console, SQL and etc)
func (ecd *MarketDataStore) showResults() {
	for _, item := range *ecd {
		fmt.Println(item)
	}
}

//Short form conversion
func parseFloatLocal(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}

//Type for Poloniex Trade Response
type PoloniexTradeResponseItem struct {
	GlobalTradeID int    `json:"globalTradeID"`
	TradeID       int    `json:"tradeID"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	Amount        string `json:"amount"`
	Total         string `json:"total"`
}

type PoloniexTradeResponse []PoloniexTradeResponseItem

//Type for Binance Trade Response
type BinanceTradeResponseItem struct {
	ID           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quaoteQty"`
	Time         int    `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

//getResponseJson - получение данных с биржи poloniex чепез api
func (poloniex *PoloniexTradeResponse) getResponseJson() <-chan MarketDataUnion {
	out := make(chan MarketDataUnion)
	response, err := getDataBodyFromApi("https://poloniex.com/public?command=returnTradeHistory&currencyPair=BTC_ETH")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(poloniex)
	go func() {
		defer close(out)
		for _, item := range *poloniex {
			fmt.Println(item)
			out <- MarketDataUnion{
				Source: "Poloniex",
				ID:     item.GlobalTradeID,
				Date:   item.Date,
				Price:  parseFloatLocal(item.Date),
				Amount: parseFloatLocal(item.Amount),
			}
		}
	}()
	return out
}

//getResponseJson - получение данных с биржи binance чепез api
func (binance *BinanceTradeResponse) getResponseJson() <-chan MarketDataUnion {
	out := make(chan MarketDataUnion)
	response, err := getDataBodyFromApi("https://api.binance.com/api/v1/trades?symbol=ETHBTC")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	json.NewDecoder(response.Body).Decode(binance)
	go func() {
		defer close(out)
		for _, item := range *binance {
			out <- MarketDataUnion{
				Source: "Binance",
				ID:     item.ID,
				Date:   strconv.Itoa(item.Time),
				Price:  parseFloatLocal(item.Price),
				Amount: parseFloatLocal(item.Qty),
			}
		}
	}()

	return out
}

type BinanceTradeResponse []BinanceTradeResponseItem

func main() {
	// get Channel for markets
	var binance = new(BinanceTradeResponse)
	binOut := binance.getResponseJson()
	var poloniex = new(PoloniexTradeResponse)
	polOut := poloniex.getResponseJson()

	//Shared Data Store
	var marketData MarketDataStore
	//non blocking read
	go func() {
		for {
			select {
			case b := <-binOut:
				marketData = append(marketData, b)
			case p := <-polOut:
				marketData = append(marketData, p)
			}
			//Market request timeout
			time.Sleep(DurationTime * time.Millisecond)
		}
	}()
	//Program timeout for gorutine
	time.Sleep(DurationTime * time.Second)
	//Show result for gorutine work
	marketData.showResults()
}
