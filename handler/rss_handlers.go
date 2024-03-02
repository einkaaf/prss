package handler

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/einkaaf/prss/cache"
	"github.com/einkaaf/prss/constants"
	carir "github.com/einkaaf/prss/model/car_ir"
	"github.com/einkaaf/prss/model/digiato"
	"github.com/einkaaf/prss/model/khodrobank"
	"github.com/einkaaf/prss/model/tabnak"
	"github.com/einkaaf/prss/model/tasnim"
	"github.com/einkaaf/prss/model/yjc"
	"github.com/einkaaf/prss/model/zoomg"
	"github.com/einkaaf/prss/model/zoomit"
	"github.com/gin-gonic/gin"
)

func fetchData(c *gin.Context, url string, target interface{}) error {

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to fetch the RSS feed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %v", err)
	}

	if err := xml.Unmarshal(responseBody, target); err != nil {
		return fmt.Errorf("invalid XML format: %v", err)
	}

	return nil
}

func fetchJsonData(c *gin.Context, url string, target interface{}) error {
	resp, err := http.Post(url, "application/json; charset=utf-8", nil)
	if err != nil {
		return fmt.Errorf("failed to fetch the RSS feed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read the response body: %v", err)
	}

	if err := json.Unmarshal(responseBody, target); err != nil {
		return fmt.Errorf("invalid Json format: %v", err)
	}
	return nil
}

func CheckCache(key string) (interface{}, bool) {
	// check cache
	cache := cache.NewPrssCache()

	cacheResult, found := cache.GetCachedItem(key)

	if found {
		return cacheResult, true
	} else {
		return "", false
	}

}

func ZoomitHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.ZoomitKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}

	var data zoomit.Rss
	if err := fetchData(c, constants.ZoomitFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.ZoomitKey, data)

	c.JSON(http.StatusOK, data)
}

func DigiatoHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.DigiatoKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}
	var data digiato.Rss
	if err := fetchData(c, constants.DigiatoFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.DigiatoKey, data)

	c.JSON(http.StatusOK, data)
}

func TasnimHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.TasnimKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}

	var data tasnim.Rss
	if err := fetchData(c, constants.TasnimFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.TasnimKey, data)
	c.JSON(http.StatusOK, data)
}

func TabnakHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.TabnakKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}

	var data tabnak.Rss
	if err := fetchData(c, constants.TabnakFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.TabnakKey, data)
	c.JSON(http.StatusOK, data)
}

func YJCHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.YJCKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}

	var data yjc.Rss
	if err := fetchData(c, constants.YJCFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.YJCKey, data)
	c.JSON(http.StatusOK, data)
}

func ZoomgHandler(c *gin.Context) {

	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.ZoomgKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}
	var data zoomg.Rss
	if err := fetchData(c, constants.ZoomgFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.ZoomgKey, data)
	c.JSON(http.StatusOK, data)
}

func KhodroBankHandler(c *gin.Context) {
	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.KhodroBankKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}
	var data khodrobank.Rss
	if err := fetchData(c, constants.KhodroBankFeedURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.KhodroBankKey, data)
	c.JSON(http.StatusOK, data)
}

func CarPriceHandler(c *gin.Context) {
	cache := cache.NewPrssCache()
	if cacheRes, found := CheckCache(constants.CarIRPriceKey); found {
		c.JSON(http.StatusOK, cacheRes)
		return
	}
	var data carir.CarIR
	if err := fetchJsonData(c, constants.CarIRPriceURL, &data); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	cache.SetCachedItem(constants.CarIRPriceKey, data)
	c.JSON(http.StatusOK, data)
}

func handleError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, gin.H{"error": message})
}
