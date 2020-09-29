// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/kkohtaka/go-bitflyer/pkg/api"
	"github.com/kkohtaka/go-bitflyer/pkg/api/auth"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type httpClient struct {
	authConfig *auth.AuthConfig
}

func New() *httpClient {
	return &httpClient{}
}

func (hc *httpClient) Auth(authConfig *auth.AuthConfig) *httpClient {
	hc.authConfig = authConfig
	return hc
}

func (hc *httpClient) Request(api api.API, req api.Request, result interface{}) error {
	u, err := api.BaseURL()
	if err != nil {
		return errors.Wrapf(err, "set base URI")
	}
	payload := req.Payload()
	u.RawQuery = req.Query()

	var body io.Reader
	if len(payload) > 0 {
		body = bytes.NewReader(payload)
	}
	rawReq, err := http.NewRequest(req.Method(), u.String(), body)
	if err != nil {
		return errors.Wrapf(err, "create POST request from url: %s", u.String())
	}
	if hc.authConfig != nil {
		header, err := auth.GenerateAuthHeaders(hc.authConfig, time.Now(), api, req)
		if err != nil {
			return errors.Wrap(err, "generate auth header")
		}
		rawReq.Header = *header
	}
	if len(payload) > 0 {
		rawReq.Header.Set("Content-Type", "application/json")
	}

	// fmt.Printf("[%s] > %s\n", rawReq.Method, rawReq.URL)

	c := &http.Client{}
	resp, err := c.Do(rawReq)
	if err != nil {
		return errors.Wrapf(err, "send HTTP request with url: %s", u.String())
	}

	fmt.Printf("[response code] %d\n", resp.StatusCode)
	fmt.Println("TODO: check if error")
	defer resp.Body.Close()

	var buf bytes.Buffer
	tee := io.TeeReader(resp.Body, &buf)

	// TODO: Don't use ioutil.ReadAll()
	dec := json.NewDecoder(tee)
	dec.DisallowUnknownFields()
	err = dec.Decode(result)

	if err != nil {
		fmt.Printf("[err] %s\n", err.Error())
		data, readErr := ioutil.ReadAll(&buf)
		if readErr != nil {
			return errors.Wrapf(err, "unmarshal response body: failed to read buffer")
		}
		return errors.Wrapf(err, "unmarshal response body: %s", string(data))
	}
	return nil
}
