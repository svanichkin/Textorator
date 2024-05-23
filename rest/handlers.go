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

	l := c.Query("lang")
	s := c.Query("style")
	f := c.Query("format")

	// Validate cache id

	cid := c.Query("cid")
	if len(cid) != 0 {
		cached := data.GetFromCache(cid)
		if len(cached) > 0 {
			response(c, http.StatusOK, gin.H{"text": stringToBase64(cached), "cid": cid}, rt)
			return
		}
		if len(l) == 0 && len(s) == 0 && len(f) == 0 {
			response(c, http.StatusOK, gin.H{"text": "", "cid": cid}, rt)
			return
		}
	}

	// Check input text

	t, err := url.QueryUnescape(c.Query("text"))
	if err != nil {
		responseError(c, fmt.Errorf("bad request: need valid text, %v", err), rt)
		return
	}

	// Check one parameter for do

	if len(l) == 0 && len(s) == 0 && len(f) == 0 {
		responseError(c, fmt.Errorf("bad request: need valid lang, style or format"), rt)
		return
	}

	// Main switch

	var do string

	if len(s) > 0 {
		do += "Convert to style (" + s + ")"
	}

	if len(l) > 0 {
		if len(do) != 0 {
			do += ", then translate to (" + l + ")"
		} else {
			do += "Translate to (" + l + ")"
		}
	}

	if len(f) > 0 {
		if len(do) != 0 {
			do += ", then convert to format (" + f + ")"
		} else {
			do += "Convert to format (" + f + ")"
		}
	}

	do += ", this text (" + t + ")"

	// Request to ai

	transform, err := open.DoTransform(do)
	if err != nil {
		responseError(c, fmt.Errorf("bad request: %v", err), rt)
		return
	}

	// Generate cid

	str := s + l + f + t
	hasher := sha256.New()
	hasher.Write([]byte(str))
	cid = hex.EncodeToString(hasher.Sum(nil))

	// Add to cache

	time := checkInt(c.Query("cache"))
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

	// Check lang

	l := c.Query("lang")
	if len(l) == 0 {
		responseError(c, fmt.Errorf("bad request: lang is not valid"), rt)
		return
	}

	// Check style

	s := c.Query("style")
	if len(s) == 0 {
		responseError(c, fmt.Errorf("bad request: style is not valid"), rt)
		return
	}

	// Check text

	t := c.Query("text")
	if len(t) == 0 {
		responseError(c, fmt.Errorf("bad request: text is not valid"), rt)
		return
	}

	// Check len

	ln := checkInt(c.Query("len"))
	if len(c.Query("len")) > 0 && ln <= 0 {
		responseError(c, fmt.Errorf("bad request: len is not valid"), rt)
		return
	}

	// Generate ai

	do := "Generate text about (" + t + "), in language (" + l + "), in style (" + s + ")"

	// Format if needed

	f := c.Query("format")
	if len(f) > 0 {
		do += ", then convert to format (" + f + ")"
	}

	// Length of text

	if ln > 0 {
		do += ", and please limit text to (" + fmt.Sprint(ln) + ") symbols."
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

	return base64.StdEncoding.EncodeToString([]byte(str))

}
