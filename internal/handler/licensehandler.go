package handler

import (
	"encoding/json"
	"maps"
	"slices"

	"github.com/gythialy/jb-helper/internal/license"
	"github.com/gythialy/jb-helper/internal/model"
	"github.com/gythialy/jb-helper/internal/response"
	"github.com/valyala/fasthttp"
)

func LicenseHandler(ctx *fasthttp.RequestCtx) {
	var licenseReq model.License
	err := json.Unmarshal(ctx.PostBody(), &licenseReq)
	if err != nil {
		response.Error(ctx, err)
		return
	}

	date := licenseReq.Products[0].FallbackDate

	licenseReq.Products = append(licenseReq.Products, model.Product{
		Code:         "PCWMP",
		FallbackDate: date,
		PaidUpTo:     date,
		Extended:     false,
	}, model.Product{
		Code:         "PRR",
		FallbackDate: date,
		PaidUpTo:     date,
		Extended:     false,
	}, model.Product{
		Code:         "PRAINBOWBRACKET",
		FallbackDate: date,
		PaidUpTo:     date,
		Extended:     false,
	})

	// remove duplicate items
	products := make(map[string]model.Product)

	for _, v := range licenseReq.Products {
		if _, ok := products[v.Code]; !ok {
			products[v.Code] = v
		}
	}
	licenseReq.Products = slices.Collect(maps.Values(products))
	// licenseStr, _ := json.Marshal(licenseReq)
	// log.Printf("licenseReq: %s\n", licenseStr)
	lic, err := license.GenerateLicense(&licenseReq)
	if err != nil {
		response.Error(ctx, err)
		return
	}
	response.Success(ctx, model.ResponseData{
		License: lic,
	})
}
