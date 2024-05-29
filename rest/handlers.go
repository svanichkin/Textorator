package rest

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"main/data"
	"main/open"

	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

func transformHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Get all parameters

	var lang, style, format, cid, text, cache string
	if c.Request.Method == "GET" {
		lang = c.Query("lang")
		style = c.Query("style")
		format = c.Query("format")
		cid = c.Query("cid")
		text = c.Query("text")
		cache = c.Query("cache")
	} else if c.Request.Method == "POST" {
		lang = c.PostForm("lang")
		style = c.PostForm("style")
		format = c.PostForm("format")
		cid = c.PostForm("cid")
		text = c.PostForm("text")
		cache = c.PostForm("cache")
	}

	// Check cache, if only cid parameter

	if len(cid) >= 32 && len(lang) == 0 && len(style) == 0 && len(format) == 0 {
		response(c, http.StatusOK, gin.H{"text": stringToBase64(data.GetFromCache(cid)), "cid": cid}, rt)
		return
	}

	// Check input text

	text, err := url.QueryUnescape(text)
	if err != nil {
		responseError(c, fmt.Errorf("bad request: need valid text, %v", err), rt)
		return
	}

	// Check one parameter for do

	if len(lang) == 0 && len(style) == 0 && len(format) == 0 {
		responseError(c, fmt.Errorf("bad request: need valid lang, style or format"), rt)
		return
	}

	// Generate cid if needed, check cache

	if len(cid) < 32 {
		cid = stringToCid(style + lang + format + text)
		str := stringToBase64(data.GetFromCache(cid))
		if len(str) > 0 {
			response(c, http.StatusOK, gin.H{"text": stringToBase64(data.GetFromCache(cid)), "cid": cid}, rt)
			return
		}
	}

	// Main switch

	var do string
	do += ""

	if len(style) > 0 {
		do += "Стиль → " + style + ". "
	}

	if len(lang) > 0 {
		do += "Язык → " + lang + ". "
	}

	if len(format) > 0 {
		do += "Формат → " + format + ". "
	}

	do += "Текст для обработки → " + text + ""

	// Request to ai

	transform, err := open.DoTransform(do)
	if err != nil {
		responseError(c, fmt.Errorf("bad request: %v", err), rt)
		return
	}

	// Add to cache

	time := checkInt(cache)
	if time > 0 {
		data.AddToCache(cid, transform, time)
	} else {
		data.RemoveFromCache(cid)
	}

	// Return result

	response(c, http.StatusOK, gin.H{"text": stringToBase64(transform), "cid": cid}, rt)

}

func generateHandler(c *gin.Context) {

	// Check return type

	rt := checkType(c)

	// Get all parameters

	var lang, style, format, text, length string
	if c.Request.Method == "GET" {
		lang = c.Query("lang")
		style = c.Query("style")
		format = c.Query("format")
		text = c.Query("text")
		length = c.Query("len")
	} else if c.Request.Method == "POST" {
		lang = c.PostForm("lang")
		style = c.PostForm("style")
		format = c.PostForm("format")
		text = c.PostForm("text")
		length = c.PostForm("len")
	}

	// Check lang

	if len(lang) == 0 {
		responseError(c, fmt.Errorf("bad request: lang is not valid"), rt)
		return
	}

	// Check style

	if len(style) == 0 {
		responseError(c, fmt.Errorf("bad request: style is not valid"), rt)
		return
	}

	// Check text

	if len(text) == 0 {
		responseError(c, fmt.Errorf("bad request: text is not valid"), rt)
		return
	}

	// Check input text

	text, err := url.QueryUnescape(text)
	if err != nil {
		responseError(c, fmt.Errorf("bad request: need valid text, %v", err), rt)
		return
	}

	// Check len

	ln := checkInt(length)
	if len(length) > 0 && ln <= 0 {
		responseError(c, fmt.Errorf("bad request: len is not valid"), rt)
		return
	}

	// Generate ai

	do := "Сгенерируй текст по запросу (" + text + "), на языке (" + lang + "), в стилистике (" + style + ")"

	// Format if needed

	if len(format) > 0 {
		do += ", затем приведи к формату (" + format + ")"
	}

	// Length of text

	if ln > 0 {
		do += ", и пожалуйста ограничь длину текста до (" + fmt.Sprint(ln) + ") символов."
	}

	// Request to ai

	generate, err := open.DoGenerate(do)
	if err != nil {
		responseError(c, fmt.Errorf("bad request: %v", err), rt)
		return
	}

	// Return result

	response(c, http.StatusOK, gin.H{"text": stringToBase64(generate)}, rt)

}

func assistHandler(c *gin.Context) {

	// Check return type

	// rt := checkType(c)

	// // Check cid

	// cid := c.Query("cid")
	// if len(cid) == 0 {
	// 	responseError(c, fmt.Errorf("bad request: lang is not valid"), rt)
	// 	return
	// }

	// // Check style

	// s := c.Query("style")
	// if len(s) == 0 {
	// 	responseError(c, fmt.Errorf("bad request: style is not valid"), rt)
	// 	return
	// }

	// // Check text

	// t := c.Query("text")
	// if len(t) == 0 {
	// 	responseError(c, fmt.Errorf("bad request: text is not valid"), rt)
	// 	return
	// }

	// // Check len

	// ln := checkInt(c.Query("len"))
	// if len(c.Query("len")) > 0 && ln <= 0 {
	// 	responseError(c, fmt.Errorf("bad request: len is not valid"), rt)
	// 	return
	// }

	// // Generate ai

	// do := "Generate text about (" + t + "), in language (" + l + "), in style (" + s + ")"

	// // Format if needed

	// f := c.Query("format")
	// if len(f) > 0 {
	// 	do += ", then convert to format (" + f + ")"
	// }

	// // Length of text

	// if ln > 0 {
	// 	do += ", and please limit text to (" + fmt.Sprint(ln) + ") symbols."
	// }

	// // Request to ai

	// generate, err := open.DoGenerate(do)
	// if err != nil {
	// 	responseError(c, fmt.Errorf("bad request: %v", err), rt)
	// 	return
	// }

	// // Return result

	// response(c, http.StatusOK, gin.H{"text": stringToBase64(generate)}, rt)

}

func stringToBase64(str string) string {

	if len(str) == 0 {
		return ""
	}
	return base64.StdEncoding.EncodeToString([]byte(str))

}

func stringToCid(str string) string {

	hasher := sha256.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))

}
