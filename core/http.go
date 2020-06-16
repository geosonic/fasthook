/*
 * Copyright (c) 2020. All rights reserved.
 */

package core

import (
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func Start(config *Config) {
	m := func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Server", "fasthook/"+version+" (+https://github.com/geosonic/fasthook)")
		switch string(ctx.Path()) {
		case config.Settings.Path:
			handle(ctx, config)
		default:
			ctx.Error(fasthttp.StatusMessage(fasthttp.StatusNotFound), fasthttp.StatusNotFound)
		}
	}

	if config.Settings.AutoConnect {
		go Connect(config)
	}

	err := fasthttp.ListenAndServe(fmt.Sprintf(":%d", config.Settings.Port), m)
	if err != nil {
		log.Fatalln(err)
	}
}
