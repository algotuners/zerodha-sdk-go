package pkg

import (
	"github.com/mayank-sheoran/zerodha-sdk-go/pkg/constants"
	"net/http"
)

func KiteConnect(encToken string) *KiteHttpClient {
	client := &KiteHttpClient{}
	client.SetHTTPClient(&http.Client{
		Timeout: constants.RequestTimeout,
	})
	client.SetBaseURI(constants.BaseURI)
	client.SetEncToken(encToken)
	return client
}
