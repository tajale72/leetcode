package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Bitcoin struct {
	Date        string `json:"date"`
	Currency    string `json:"currency"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	LastUpdated string `json:"lastUpdated"`
	Symbol      string `json::"sysmbol"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/btc", GetBitcoinData)
	r.Run()
}

func GetBitcoinData(c *gin.Context) {
	str, err := GetApiData()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	re := regexp.MustCompile(`"([^"]+)"`)

	// Find all matches of the regular expression in the input string
	matches := re.FindAllStringSubmatch(*str, -1)

	// Print the matched text for each match
	var bitcoindata []string
	for _, match := range matches {
		bitcoindata = append(bitcoindata, match[1])
	}

	usddbitcoindata := Bitcoin{
		LastUpdated: bitcoindata[2],
		Currency:    bitcoindata[12],
		Name:        bitcoindata[10],
		Price:       "$" + bitcoindata[18],
		Date:        "US " + time.Now().Format("Jan 2, 2006 at 3:04pm"),
	}

	loc, err := time.LoadLocation("Europe/London")
	if err != nil {
		panic(err)
	}
	uktime := time.Now().In(loc).Format("Jan 2, 2006 at 3:04pm")
	poundsbitcoindata := Bitcoin{
		LastUpdated: bitcoindata[6],
		Currency:    bitcoindata[30],
		Name:        bitcoindata[10],
		Price:       "£" + bitcoindata[28],
		Date:        "UK " + uktime,
	}

	var btcData []Bitcoin
	btcData = append(btcData, poundsbitcoindata, usddbitcoindata)

	c.JSON(http.StatusAccepted, btcData)

}

func GetApiData() (*string, error) {
	api := "https://api.coindesk.com/v1/bpi/currentprice.json"

	res, err := http.Get(api)
	if err != nil {
		log.Println("error from getting the data from the api")
		return nil, err

	}

	resp, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error from reading the data from the api")
		return nil, err
	}
	str := string(resp)
	return &str, nil
}
