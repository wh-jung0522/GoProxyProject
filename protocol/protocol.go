package protocol

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var dump_dir = "C:\\Users\\kwdg1\\Desktop\\GoProxyProject\\jsonDB"

type Ticker struct {
	Timestamp     int64  `json:"timestamp"`
	Last          string `json:"last"`
	Open          string `json:"open,omitempty"`
	Bid           string `json:"bid,omitempty"`
	Ask           string `json:"ask,omitempty"`
	Low           string `json:"low,omitempty"`
	High          string `json:"high,omitempty"`
	Volume        string `json:"volume,omitempty"`
	Change        string `json:"change,omitempty"`
	ChangePercent string `json:"changePercent,omitempty"`
}

func (t *Ticker) dump() error {
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		return err
	}
	save_dir := dump_dir + "\\ticker"
	if _, err := os.Stat(save_dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(save_dir, 0777)
	}
	save_path := save_dir + "\\" + string(time.Now().Format("20060102150405")) + ".json"
	err = ioutil.WriteFile(save_path, file, 0777)
	if err != nil {
		return err
	}
	return nil
}

type Ticker_list map[string]Ticker

func (t *Ticker_list) dump() error {
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		return err
	}
	save_dir := dump_dir + "\\ticker_detailed_all"
	if _, err := os.Stat(save_dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(save_dir, 0777)
	}
	save_path := save_dir + "\\" + string(time.Now().Format("20060102150405")) + ".json"
	err = ioutil.WriteFile(save_path, file, 0777)
	if err != nil {
		return err
	}
	return nil
}

type Orderbook struct {
	Timestamp int64       `json:"timestamp"`
	Bids      [][3]string `json:"bids"`
	Asks      [][3]string `json:"asks"`
}

func (t *Orderbook) dump() error {
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		return err
	}
	save_dir := dump_dir + "\\orderbook"
	if _, err := os.Stat(save_dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(save_dir, 0777)
	}
	save_path := save_dir + "\\" + string(time.Now().Format("20060102150405")) + ".json"
	err = ioutil.WriteFile(save_path, file, 0777)
	if err != nil {
		return err
	}
	return nil
}

type Transaction struct {
	Timestamp int64  `json:"timestamp"`
	Tid       string `json:"tid"`
	Amount    string `json:"amount"`
	Type      string `json:"type"`
}

type Transactions []Transaction

func (t *Transactions) dump() error {
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		return err
	}
	save_dir := dump_dir + "\\transactions"
	if _, err := os.Stat(save_dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(save_dir, 0777)
	}
	save_path := save_dir + "\\" + string(time.Now().Format("20060102150405")) + ".json"
	err = ioutil.WriteFile(save_path, file, 0777)
	if err != nil {
		return err
	}
	return nil
}

type Mode int

type CandleStick struct {
	Market                  string  `json:"market"`
	Candle_date_time_utc    string  `json:"candle_date_time_utc"`
	Candle_date_time_kst    string  `json:"candle_date_time_kst"`
	Opening_price           int     `json:"opening_price"`
	High_price              int     `json:"high_price"`
	Low_price               int     `json:"low_price"`
	Trade_price             int     `json:"trade_price"`
	Timestamp               int64   `json:"timestamp"`
	Candle_acc_trade_price  float64 `json:"candle_acc_trade_price"`
	Candle_acc_trade_volume float64 `json:"candle_acc_trade_volume"`
	Unit                    int     `json:"unit"`
	/*
		[
			{
		    "market": "KRW-BTC",
		    "candle_date_time_utc": "2021-10-31T23:59:00",
		    "candle_date_time_kst": "2021-11-01T08:59:00",
		    "opening_price": 72431000,
		    "high_price": 72431000,
		    "low_price": 72289000,
		    "trade_price": 72391000,
		    "timestamp": 1635724799706,
		    "candle_acc_trade_price": 431731840.11912,
		    "candle_acc_trade_volume": 5.96484479,
		    "unit": 1
			},
		]
	*/
}
type CandleSticks []CandleStick

func (t *CandleSticks) dump() error {
	file, err := json.MarshalIndent(*t, "", " ")
	if err != nil {
		return err
	}
	save_dir := dump_dir + "\\candlesticks"
	if _, err := os.Stat(save_dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		os.Mkdir(save_dir, 0777)
	}
	save_path := save_dir + "\\" + string(time.Now().Format("20060102150405")) + ".json"
	err = ioutil.WriteFile(save_path, file, 0777)
	if err != nil {
		return err
	}
	return nil
}

const (
	/* Public Mode */
	ERR Mode = iota
	TICKER
	TICKER_DETAIL
	TICKER_DETAIL_ALL
	ORDERBOOK
	TRANSACTION
	/* Additional Mode Candle Stick */
	CANDLESTICK
	/* Private Mode */
)

type Dumper interface {
	dump() error
}

func NewDumper(m Mode) Dumper {
	switch m {
	case TICKER:
		return &Ticker{}
	case TICKER_DETAIL:
		return &Ticker{}
	case TICKER_DETAIL_ALL:
		return &Ticker_list{}
	case ORDERBOOK:
		return &Orderbook{}
	case TRANSACTION:
		return &Transactions{}
	case CANDLESTICK:
		return &CandleSticks{}
	default:
		return nil
	}
}

func ApiHandler(w http.ResponseWriter, r *http.Request) {
	origin_URL := "https://api.korbit.co.kr/v1"
	upbit_url := "https://api.upbit.com/v1"
	notSupportError := errors.New("apiHandler: URL Not Support")
	notSupportCurrencyError := errors.New("apiHandler: Not Support Currency")
	getCurrency := func(raw_query string) (string, error) {
		prefix_query := "currency_pair="
		if !strings.HasPrefix(raw_query, prefix_query) {
			return "", notSupportCurrencyError
		}
		if len(raw_query) < len(prefix_query)+7 {
			return "", notSupportCurrencyError
		} else {
			return raw_query[len(prefix_query):], nil
		}
	}

	/* convCandleQuery
	Korbit Type : "http://localhost:8887/candles/minutes/1?currency_pair=eth_krw&to=2021-11-01T00%3A00%3A00&count=200"
	Upbit Type : ""https://api.upbit.com/v1/candles/minutes/200?market=KRW-BTC&to=2021-11-01%2000%3A00%3A00&count=200"
	*/
	convCandleQuery := func(raw_query string) string {
		var converted string
		converted_slice := []string{}
		slice := strings.Split(raw_query, "&")
		for _, chunk := range slice {
			temp, err := getCurrency(chunk)
			if len(temp) != 0 && err == nil {
				temp_slice := strings.Split(temp, "_")
				temp_slice[0], temp_slice[1] = strings.ToUpper(temp_slice[1]), strings.ToUpper(temp_slice[0])
				temp = "market=" + strings.Join(temp_slice, "-")
				converted_slice = append(converted_slice, temp)
			} else {
				converted_slice = append(converted_slice, chunk)
			}
		}
		converted = strings.Join(converted_slice, "&")
		return converted
	}
	getMode := func() (Mode, error) {
		urlpath := r.URL.Path
		raw_query := r.URL.RawQuery
		isticker := strings.HasPrefix(urlpath, "/ticker")
		isorderbook := strings.HasPrefix(urlpath, "/orderbook")
		istransaction := strings.HasPrefix(urlpath, "/transactions")
		iscandlestick := strings.HasPrefix(urlpath, "/candles")
		if isticker {
			next_path := urlpath[len("/ticker"):]
			isdetail := strings.HasPrefix(next_path, "/detailed")
			if isdetail {
				next_path = next_path[len("/detailed"):]
				isall := next_path == "/all"
				if isall {
					return TICKER_DETAIL_ALL, nil
					//TODO : korbit_get_detailed_ticker_all
				} else {
					_, err := getCurrency(raw_query)
					if err != nil {
						return ERR, err
					}
					return TICKER_DETAIL, nil
					//TODO : korbit_get_detailed_ticker Specific currency
				}
			} else if len(next_path) == 0 {
				_, err := getCurrency(raw_query)
				if err != nil {
					return ERR, err
				}
				return TICKER, nil
			} else {
				return ERR, notSupportError
			}
		} else if isorderbook {
			_, err := getCurrency(raw_query)
			if err != nil {
				return ERR, err
			}
			//TODO : korbit_get_orderbook
			return ORDERBOOK, nil
		} else if istransaction {
			_, err := getCurrency(raw_query)
			if err != nil {
				return ERR, err
			}
			//TODO : korbit_get_transactions
			return TRANSACTION, nil
		} else if iscandlestick {
			return CANDLESTICK, nil
		} else {
			return ERR, notSupportError
		}
	}
	switch r.Method {
	case http.MethodGet:
		mode, err := getMode()
		if err != nil {
			log.Println(err)
			return
		}
		req_url := origin_URL + r.RequestURI
		if mode == CANDLESTICK {
			req_url = upbit_url + r.URL.Path + "?" + convCandleQuery(r.URL.RawQuery)
		}
		resp, err := http.Get(req_url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		dumper := NewDumper(mode)
		json.Unmarshal(data, &dumper)
		err = dumper.dump()
		if err != nil {
			panic(err)
		}
		json.NewEncoder(w).Encode(dumper)

	case http.MethodPost:
		panic("Coming Soon")
	default:
		panic("Not Supply this method.")
	}
}
