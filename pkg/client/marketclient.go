package client

import (
	"../../internal"
	"../../internal/requestbuilder"
	"../getrequest"
	"../response/market"
	"encoding/json"
	"errors"
	"strconv"
)

type MarketClient struct {
	publicUrlBuilder *requestbuilder.PublicUrlBuilder
}

func (p *MarketClient) Init(host string) *MarketClient {
	p.publicUrlBuilder = new(requestbuilder.PublicUrlBuilder).Init(host)
	return p
}

func (client *MarketClient) GetCandlestick(symbol string, optionalRequest getrequest.GetCandlestickOptionalRequest) ([]market.Candlestick, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)
	if optionalRequest.Period != "" {
		request.AddParam("period", optionalRequest.Period)
	}
	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	}

	url := client.publicUrlBuilder.Build("/market/history/kline", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetCandlestickResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}
	return nil, errors.New(getResp)

}

func (client *MarketClient) GetLast24hCandlestickAskBid(symbol string) (*market.CandlestickAskBid, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)

	url := client.publicUrlBuilder.Build("/market/detail/merged", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLast24hCandlestickAskBidResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)

}

func (client *MarketClient) GetLast24hCandlesticks() ([]market.SymbolCandlestick, error) {

	request := new(getrequest.GetRequest).Init()

	url := client.publicUrlBuilder.Build("/market/tickers", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLast24hCandlesticksResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)

}

func (client *MarketClient) GetDepth(symbol string, step string, optionalRequest getrequest.GetDepthOptionalRequest) (*market.Depth, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)
	request.AddParam("type", step)
	if optionalRequest.Size != 0 {
		request.AddParam("depth", strconv.Itoa(optionalRequest.Size))
	}

	url := client.publicUrlBuilder.Build("/market/depth", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetDepthResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)

}
func (client *MarketClient) GetLatestTrade(symbol string) (*market.TradeTick, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)

	url := client.publicUrlBuilder.Build("/market/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLatestTradeResponse{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}
func (client *MarketClient) GetHistoricalTrade(symbol string, optionalRequest getrequest.GetHistoricalTradeOptionalRequest) ([]market.TradeTick, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)
	if optionalRequest.Size != 0 {
		request.AddParam("size", strconv.Itoa(optionalRequest.Size))
	}

	url := client.publicUrlBuilder.Build("/market/history/trade", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetHistoricalTradeResponse{}

	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Data != nil {

		return result.Data, nil
	}

	return nil, errors.New(getResp)

}

func (client *MarketClient) GetLast24hCandlestick(symbol string) (*market.Candlestick, error) {

	request := new(getrequest.GetRequest).Init()
	request.AddParam("symbol", symbol)

	url := client.publicUrlBuilder.Build("/market/detail", request)
	getResp, getErr := internal.HttpGet(url)
	if getErr != nil {
		return nil, getErr
	}

	result := market.GetLast24hCandlestick{}
	jsonErr := json.Unmarshal([]byte(getResp), &result)
	if jsonErr != nil {
		return nil, jsonErr
	}
	if result.Status == "ok" && result.Tick != nil {

		return result.Tick, nil
	}

	return nil, errors.New(getResp)
}