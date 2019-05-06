package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sort"
	"strconv"
)

const Token = "b5dee807f02c4eccc01525e7bcd82b02214c0576573549d970a77b"
const ApiUrl = "http://api.tushare.pro"

type PostBody struct {
	TOKEN    string            `json:"token"`
	API_NAME string            `json:"api_name"`
	PARAMS   map[string]string `json:"params"`
	FIELDS   string            `json:"fields"`
}

func sortDeleteDuplicates(slice []string) []string {
	sort.Strings(slice)
	i := 0
	var j int
	for {
		if i > len(slice)-1 {
			break
		}
		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {

		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice
	/*
		for _, s := range slice {
			if len(result) < 1 {
				result = append(result, s)
			} else {
				in := false
				for _, s1 := range result {
					if s1 == s {
						in = true
					}
				}
				if !in {
					result = append(result, s)
				}
			}
		}
		return
	*/
}

func getData(post *PostBody) (numbers []string) {
	jsonBytes, err := json.Marshal(&post)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", ApiUrl, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(body), &result); err == nil {
		data, ok := result["data"].(map[string]interface{})
		if !ok {
			panic(err)
		}

		for _, items := range data["items"].([]interface{}) {
			for _, v := range items.([]interface{}) {
				switch v.(type) {
				case float64:
					if f, ok := v.(float64); ok {
						if f != 0. {
							numbers = append(numbers, strconv.FormatFloat(math.Abs(f)*1000, 'f', -1, 64))
						}
					}
				}
			}
		}
	}
	return
}

func main() {

	num := "002011.SZ"
	date := "20181231"

	//利润表
	income := PostBody{Token, "", nil, ""}
	income.API_NAME = "income"
	income.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	income.FIELDS = "int_income"

	//资产负载表
	balancesheet := PostBody{Token, "", nil, ""}
	balancesheet.API_NAME = "balancesheet"
	balancesheet.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	balancesheet.FIELDS = ""

	//现金流量表
	cashflow := PostBody{Token, "", nil, ""}
	cashflow.API_NAME = "cashflow"
	cashflow.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	cashflow.FIELDS = ""

	//财务指标数据
	indicator := PostBody{Token, "", nil, ""}
	indicator.API_NAME = "fina_indicator"
	indicator.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//indicator.FIELDS = ""

	//主营业务
	mainbz := PostBody{Token, "", nil, ""}
	mainbz.API_NAME = "fina_mainbz"
	mainbz.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	mainbz.FIELDS = ""

	//业绩预告
	forecast := PostBody{Token, "", nil, ""}
	forecast.API_NAME = "forecast"
	forecast.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	forecast.FIELDS = ""

	//业绩快报
	express := PostBody{Token, "", nil, ""}
	express.API_NAME = "express"
	express.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	//	express.FIELDS = ""

	//		report := PostBody{Token, "", nil, ""}
	//		report.API_NAME = "disclosure_date"
	//		report.PARAMS = map[string]string{
	//			"ts_code": num,
	//			"period":  date,
	//		}

	result := getData(&income)
	result = append(result, getData(&balancesheet)...)
	result = append(result, getData(&cashflow)...)
	result = append(result, getData(&indicator)...)
	result = append(result, getData(&mainbz)...)
	result = append(result, getData(&forecast)...)
	result = append(result, getData(&express)...)
	result = sortDeleteDuplicates(result)
	var raw [9]int
	for _, v := range result {
		i := v[0] - 49
		raw[i] = raw[i] + 1
	}
	var sum float64
	for _, v := range raw {
		sum = sum + float64(v)
	}
	var rate [9]float64
	for i := 0; i < len(rate); i++ {
		rate[i] = float64(raw[i]) * 100.0 / sum
	}

	fmt.Printf("lenght: %v %v\n", sum, raw)
	fmt.Printf("rate: %.2f\n", rate)
}
