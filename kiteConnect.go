package zerodha_sdk_go

import (
	"github.com/mayank-sheoran/zerodha-sdk-go/constants"
	"net/http"
)

func KiteConnect(encToken string) *KiteHttpClient {
	client := &KiteHttpClient{}
	client.SetHTTPClient(&http.Client{
		Timeout: constants.RequestTimeout,
	})
	client.SetBaseURI(constants.KiteBaseURI)
	client.SetEncToken(encToken)
	return client
}
