package zerodha_sdk_go

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"github.com/mayank-sheoran/zerodha-sdk-go/constants"
	"github.com/mayank-sheoran/zerodha-sdk-go/httpUtils"
	"github.com/zerodha/gokiteconnect/v4/models"
	"net/http"
	"net/url"
)

type Order struct {
	AccountID string `json:"account_id"`
	PlacedBy  string `json:"placed_by"`

	OrderID                 string                 `json:"order_id"`
	ExchangeOrderID         string                 `json:"exchange_order_id"`
	ParentOrderID           string                 `json:"parent_order_id"`
	Status                  string                 `json:"status"`
	StatusMessage           string                 `json:"status_message"`
	StatusMessageRaw        string                 `json:"status_message_raw"`
	OrderTimestamp          models.Time            `json:"order_timestamp"`
	ExchangeUpdateTimestamp models.Time            `json:"exchange_update_timestamp"`
	ExchangeTimestamp       models.Time            `json:"exchange_timestamp"`
	Variety                 string                 `json:"variety"`
	Modified                bool                   `json:"modified"`
	Meta                    map[string]interface{} `json:"meta"`

	Exchange        string `json:"exchange"`
	TradingSymbol   string `json:"tradingsymbol"`
	InstrumentToken uint32 `json:"instrument_token"`

	OrderType         string  `json:"order_type"`
	TransactionType   string  `json:"transaction_type"`
	Validity          string  `json:"validity"`
	ValidityTTL       int     `json:"validity_ttl"`
	Product           string  `json:"product"`
	Quantity          float64 `json:"quantity"`
	DisclosedQuantity float64 `json:"disclosed_quantity"`
	Price             float64 `json:"price"`
	TriggerPrice      float64 `json:"trigger_price"`

	AveragePrice      float64 `json:"average_price"`
	FilledQuantity    float64 `json:"filled_quantity"`
	PendingQuantity   float64 `json:"pending_quantity"`
	CancelledQuantity float64 `json:"cancelled_quantity"`

	AuctionNumber string `json:"auction_number"`

	Tag  string   `json:"tag"`
	Tags []string `json:"tags"`
}

type Orders []Order

type OrderParams struct {
	Exchange        string `url:"exchange,omitempty"`
	Tradingsymbol   string `url:"tradingsymbol,omitempty"`
	Validity        string `url:"validity,omitempty"`
	ValidityTTL     int    `url:"validity_ttl,omitempty"`
	Product         string `url:"product,omitempty"`
	OrderType       string `url:"order_type,omitempty"`
	TransactionType string `url:"transaction_type,omitempty"`

	Quantity          int     `url:"quantity,omitempty"`
	DisclosedQuantity int     `url:"disclosed_quantity,omitempty"`
	Price             float64 `url:"price,omitempty"`
	TriggerPrice      float64 `url:"trigger_price,omitempty"`

	Squareoff        float64 `url:"squareoff,omitempty"`
	Stoploss         float64 `url:"stoploss,omitempty"`
	TrailingStoploss float64 `url:"trailing_stoploss,omitempty"`

	IcebergLegs int `url:"iceberg_legs,omitempty"`
	IcebergQty  int `url:"iceberg_quantity,omitempty"`

	AuctionNumber string `url:"auction_number,omitempty"`

	Tag string `json:"tag" url:"tag,omitempty"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
}

type Trade struct {
	AveragePrice      float64     `json:"average_price"`
	Quantity          float64     `json:"quantity"`
	TradeID           string      `json:"trade_id"`
	Product           string      `json:"product"`
	FillTimestamp     models.Time `json:"fill_timestamp"`
	ExchangeTimestamp models.Time `json:"exchange_timestamp"`
	ExchangeOrderID   string      `json:"exchange_order_id"`
	OrderID           string      `json:"order_id"`
	TransactionType   string      `json:"transaction_type"`
	TradingSymbol     string      `json:"tradingsymbol"`
	Exchange          string      `json:"exchange"`
	InstrumentToken   uint32      `json:"instrument_token"`
}

type Trades []Trade

func (kiteHttpClient *KiteHttpClient) GetOrders() (Orders, error) {
	var orders Orders
	var successEnvelope httpUtils.HttpSuccessEnvelope
	var errorEnvelope httpUtils.HttpErrorEnvelope
	err := kiteHttpClient.doEnvelope(http.MethodGet, constants.URIGetOrders, nil, nil, &orders, errorEnvelope, successEnvelope)
	return orders, err
}

func (kiteHttpClient *KiteHttpClient) GetTrades() (Trades, error) {
	var trades Trades
	var successEnvelope httpUtils.HttpSuccessEnvelope
	var errorEnvelope httpUtils.HttpErrorEnvelope
	err := kiteHttpClient.doEnvelope(http.MethodGet, constants.URIGetTrades, nil, nil, &trades, errorEnvelope, successEnvelope)
	return trades, err
}

func (kiteHttpClient *KiteHttpClient) GetOrderHistory(OrderID string) ([]Order, error) {
	var orderHistory []Order
	var successEnvelope httpUtils.HttpSuccessEnvelope
	var errorEnvelope httpUtils.HttpErrorEnvelope
	err := kiteHttpClient.doEnvelope(http.MethodGet, fmt.Sprintf(constants.URIGetOrderHistory, OrderID), nil, nil, &orderHistory, errorEnvelope, successEnvelope)
	return orderHistory, err
}

func (kiteHttpClient *KiteHttpClient) GetOrderTrades(OrderID string) ([]Trade, error) {
	var orderTrades []Trade
	var successEnvelope httpUtils.HttpSuccessEnvelope
	var errorEnvelope httpUtils.HttpErrorEnvelope
	err := kiteHttpClient.doEnvelope(http.MethodGet, fmt.Sprintf(constants.URIGetOrderTrades, OrderID), nil, nil, &orderTrades, errorEnvelope, successEnvelope)
	return orderTrades, err
}

func (kiteHttpClient *KiteHttpClient) PlaceOrder(variety string, orderParams OrderParams) (OrderResponse, error) {
	var (
		orderResponse   OrderResponse
		params          url.Values
		err             error
		successEnvelope httpUtils.HttpSuccessEnvelope
		errorEnvelope   httpUtils.HttpErrorEnvelope
	)

	if params, err = query.Values(orderParams); err != nil {
		return orderResponse, httpUtils.NewErrorHelper(httpUtils.InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	}

	err = kiteHttpClient.doEnvelope(http.MethodPost, fmt.Sprintf(constants.URIPlaceOrder, variety), params, nil, &orderResponse, errorEnvelope, successEnvelope)
	return orderResponse, err
}

func (kiteHttpClient *KiteHttpClient) ModifyOrder(variety string, orderID string, orderParams OrderParams) (OrderResponse, error) {
	var (
		orderResponse   OrderResponse
		params          url.Values
		err             error
		successEnvelope httpUtils.HttpSuccessEnvelope
		errorEnvelope   httpUtils.HttpErrorEnvelope
	)

	if params, err = query.Values(orderParams); err != nil {
		return orderResponse, httpUtils.NewErrorHelper(httpUtils.InputError, fmt.Sprintf("Error decoding order params: %v", err), nil)
	}

	err = kiteHttpClient.doEnvelope(http.MethodPut, fmt.Sprintf(constants.URIModifyOrder, variety, orderID), params, nil, &orderResponse, errorEnvelope, successEnvelope)
	return orderResponse, err
}

func (kiteHttpClient *KiteHttpClient) CancelOrder(variety string, orderID string, parentOrderID *string) (OrderResponse, error) {
	var (
		orderResponse   OrderResponse
		params          url.Values
		successEnvelope httpUtils.HttpSuccessEnvelope
		errorEnvelope   httpUtils.HttpErrorEnvelope
	)

	if parentOrderID != nil {
		// initialize the params map first
		params := url.Values{}
		params.Add("parent_order_id", *parentOrderID)
	}

	err := kiteHttpClient.doEnvelope(http.MethodDelete, fmt.Sprintf(constants.URICancelOrder, variety, orderID), params, nil, &orderResponse, errorEnvelope, successEnvelope)
	return orderResponse, err
}

func (kiteHttpClient *KiteHttpClient) ExitOrder(variety string, orderID string, parentOrderID *string) (OrderResponse, error) {
	return kiteHttpClient.CancelOrder(variety, orderID, parentOrderID)
}