/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	regexp2 "regexp"
	"time"

	"github.com/valyala/fasthttp"
)

type origin struct {
	Origin string `json:"origin"`
}

type hookResp struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	LastRequest int    `json:"last_request"`
	LastStatus  int    `json:"last_status"`
	LastFail    int    `json:"last_fail"`
	Fails       int    `json:"fails"`
	Error       string `json:"error"`
}

func doRequest(uri string, method string, values map[string]string) ([]byte, error) {
	query := url.Values{}

	for k, v := range values {
		query.Set(k, v)
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI(uri)
	req.Header.SetMethod(method)
	req.Header.SetUserAgent("fasthook/" + version + " (+https://github.com/geosonic/fasthook)")
	req.SetBodyString(query.Encode())

	resp := fasthttp.AcquireResponse()
	client := &fasthttp.Client{}
	err := client.Do(req, resp)

	return resp.Body(), err
}

func Connect(config *Config) {
	time.Sleep(time.Second * 3)
	// При автопривязке порт по умолчанию 80
	config.Settings.Port = 80

	var origin origin
	res, err := doRequest("http://httpbin.org/ip", "GET", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(res, &origin)
	if err != nil {
		log.Fatalln(err)
	}

	var ip string
	regexp := regexp2.MustCompile(`([0-9]{1,3}[.]){3}[0-9]{1,3}`)
	if re := regexp.FindStringSubmatch(origin.Origin); len(re) != 0 {
		ip = re[0]
	} else {
		log.Fatalln("Error read ip")
	}

	var resp hookResp

	uri := "http://" + ip + config.Settings.Path

	res, err = doRequest("https://api.chatmanager.pro/", "POST", map[string]string{
		"method": "account.setWebHook",
		"token":  config.Settings.Token,
		"url":    uri,
	})
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(res, &resp)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.Error == "" {
		fmt.Println("Привязка сервера прошла успешно!")
	} else {
		log.Fatalln(resp.Error)
	}
}
