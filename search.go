package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/JasonKhew96/biliroaming-go-server/entity/bstar"
	"github.com/JasonKhew96/biliroaming-go-server/entity/web"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"golang.org/x/net/idna"
)

func (b *BiliroamingGo) addSearchAds(data []byte, clientType ClientType) ([]byte, error) {
	if b.config.CustomSearch.Data == "" {
		return data, nil
	}

	var customData interface{}
	if err := json.Unmarshal([]byte(b.config.CustomSearch.Data), &customData); err != nil {
		return nil, err
	}

	switch clientType {
	case ClientTypeBstarA:
		searchResult := &bstar.SearchResult{}
		if err := easyjson.Unmarshal(data, searchResult); err != nil {
			return nil, err
		}
		searchResult.Data.Items = append([]interface{}{customData}, searchResult.Data.Items...)
		searchResult.Data.Total++
		return easyjson.Marshal(searchResult)
	default:
		searchResult := &web.SearchResult{}
		if err := easyjson.Unmarshal(data, searchResult); err != nil {
			return nil, err
		}
		searchResult.Data.Items = append([]interface{}{customData}, searchResult.Data.Items...)
		searchResult.Data.Total++
		return easyjson.Marshal(searchResult)
	}
}

func (b *BiliroamingGo) handleAndroidSearch(ctx *fasthttp.RequestCtx) {
	queryArgs := ctx.URI().QueryArgs()
	args := b.processArgs(queryArgs)

	if args.area == "" || args.area == "th" {
		writeErrorJSON(ctx, -688, []byte("地理区域限制"))
		return
	}

	client := b.getClientByArea(args.area)

	if args.keyword == "" {
		writeErrorJSON(ctx, -688, []byte("地理区域限制"))
		return
	}

	v := url.Values{}
	v.Set("access_key", args.accessKey)
	v.Set("area", args.area)
	v.Set("build", "6400000")
	v.Set("highlight", "1")
	v.Set("keyword", args.keyword)
	v.Set("type", "7")
	v.Set("mobi_app", "android")
	v.Set("platform", "android")
	v.Set("pn", strconv.Itoa(args.pn))

	params, err := SignParams(v, ClientTypeAndroid)
	if err != nil {
		b.sugar.Error(err)
		ctx.Error(
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}

	reverseProxy := b.getReverseSearchProxyByArea(args.area)
	if reverseProxy == "" {
		reverseProxy = "app.bilibili.com"
	}
	domain, err := idna.New().ToASCII(reverseProxy)
	if err != nil {
		b.sugar.Error(err)
		ctx.Error(
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}

	url := fmt.Sprintf("https://%s/x/v2/search/type?%s", domain, params)
	b.sugar.Debug("New url: ", url)

	data, err := b.doRequestJson(ctx, client, url, []byte(http.MethodGet))
	if err != nil {
		b.processError(ctx, err)
		b.updateHealth(b.getSearchHealth(args.area), -500, "服务器错误")
		return
	}

	if isLimited, err := isResponseLimited(data); err != nil {
		b.sugar.Error(err)
	} else if isLimited {
		b.updateHealth(b.getSearchHealth(args.area), -412, "请求被拦截")
	} else {
		b.updateHealth(b.getSearchHealth(args.area), 0, "0")
	}

	dataByte, err := b.addSearchAds([]byte(data), ClientTypeAndroid)
	if err != nil {
		b.processError(ctx, err)
		return
	}

	setDefaultHeaders(ctx)
	ctx.Write(dataByte)
}

func (b *BiliroamingGo) handleBstarAndroidSearch(ctx *fasthttp.RequestCtx) {
	queryArgs := ctx.URI().QueryArgs()
	args := b.processArgs(queryArgs)

	if args.area == "" {
		args.area = "th"
		// writeErrorJSON(ctx, -688, []byte("地理区域限制"))
		// return
	}

	client := b.getClientByArea(args.area)

	if args.keyword == "" {
		writeErrorJSON(ctx, -400, []byte("请求错误"))
		return
	}

	if b.getAuthByArea(args.area) {
		if ok, _ := b.doAuth(ctx, args.accessKey, args.area); !ok {
			return
		}
	}

	v := url.Values{}
	v.Set("access_key", args.accessKey)
	v.Set("area", args.area)
	v.Set("build", "1080003")
	v.Set("keyword", args.keyword)
	v.Set("s_locale", "zh_SG")
	v.Set("type", "7")
	v.Set("mobi_app", "bstar_a")
	v.Set("pn", strconv.Itoa(args.pn))

	params, err := SignParams(v, ClientTypeBstarA)
	if err != nil {
		b.sugar.Error(err)
		ctx.Error(
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}

	reverseProxy := b.getReverseProxyByArea(args.area)
	if reverseProxy == "" {
		reverseProxy = "app.biliintl.com"
	}
	domain, err := idna.New().ToASCII(reverseProxy)
	if err != nil {
		b.sugar.Error(err)
		ctx.Error(
			fasthttp.StatusMessage(fasthttp.StatusInternalServerError),
			fasthttp.StatusInternalServerError,
		)
		return
	}

	url := fmt.Sprintf("https://%s/intl/gateway/v2/app/search/type?%s", domain, params)
	b.sugar.Debug("New url: ", url)

	data, err := b.doRequestJson(ctx, client, url, []byte(http.MethodGet))
	if err != nil {
		b.processError(ctx, err)
		b.updateHealth(b.HealthSearchTH, -500, "服务器错误")
		return
	}

	if isLimited, err := isResponseLimited(data); err != nil {
		b.sugar.Error(err)
	} else if isLimited {
		b.updateHealth(b.HealthSearchTH, -412, "请求被拦截")
	} else {
		b.updateHealth(b.HealthSearchTH, 0, "0")
	}

	dataByte, err := b.addSearchAds([]byte(data), ClientTypeBstarA)
	if err != nil {
		b.processError(ctx, err)
		return
	}

	setDefaultHeaders(ctx)
	ctx.Write(dataByte)
}
