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
	"time"
	"utils"
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

func getStockList() []interface{} {
	list := PostBody{Token, "", nil, ""}
	list.API_NAME = "stock_basic"
	list.PARAMS = map[string]string{
		"list_status": "L",
	}
	list.FIELDS = "ts_code"

	jsonBytes, err := json.Marshal(list)
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
		return data["items"].([]interface{})
	}
	return nil
}

func getData(num string, post *PostBody) (numbers []string) {
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
		b, err := json.MarshalIndent(data["items"].([]interface{}), "", " ")
		if err != nil {
			panic(err)
		}
		utils.ExportFileS("/home/jake/Desktop/data/", num+"-"+post.API_NAME+".json", string(b))

		for _, items := range data["items"].([]interface{}) {
			for _, v := range items.([]interface{}) {
				switch v.(type) {
				case float64:
					if f, ok := v.(float64); ok {
						if f != 0. {
							numbers = append(numbers, strconv.FormatFloat(math.Abs(f)*10000, 'f', -1, 64))
						}
					}
				}
			}
		}
	}
	return
}

func main() {
	//格力000651 万向钱潮 000559 金刚玻璃300093 茅台600519
	//康美药业 600518 欧菲光002456 奥瑞德600666
	//康得新002450
	num := "000651.SZ"
	//period := "20181231"
	start_date := "19901219"
	end_date := "20191231"

	//利润表
	income := PostBody{Token, "", nil, ""}
	income.API_NAME = "income"
	income.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	income.FIELDS = "int_income"

	//资产负载表
	balancesheet := PostBody{Token, "", nil, ""}
	balancesheet.API_NAME = "balancesheet"
	balancesheet.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	balancesheet.FIELDS = ""

	//现金流量表
	cashflow := PostBody{Token, "", nil, ""}
	cashflow.API_NAME = "cashflow"
	cashflow.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	cashflow.FIELDS = ""

	//财务指标数据
	indicator := PostBody{Token, "", nil, ""}
	indicator.API_NAME = "fina_indicator"
	indicator.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//indicator.FIELDS = ""

	//主营业务
	mainbz := PostBody{Token, "", nil, ""}
	mainbz.API_NAME = "fina_mainbz"
	mainbz.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	mainbz.FIELDS = ""

	//业绩预告
	forecast := PostBody{Token, "", nil, ""}
	forecast.API_NAME = "forecast"
	forecast.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	forecast.FIELDS = ""

	//业绩快报
	express := PostBody{Token, "", nil, ""}
	express.API_NAME = "express"
	express.PARAMS = map[string]string{
		"ts_code":    num,
		"start_date": start_date,
		"end_date":   end_date,
	}
	//	express.FIELDS = ""

	//		report := PostBody{Token, "", nil, ""}
	//		report.API_NAME = "disclosure_period"
	//		report.PARAMS = map[string]string{
	//			"ts_code": num,
	//			"period":  period,
	//		}

	for _, s := range getStockList() {
		for _, v := range s.([]interface{}) {
			num = fmt.Sprint(v)
		}

		result := getData(num, &income)
		result = append(result, getData(num, &balancesheet)...)
		result = append(result, getData(num, &cashflow)...)
		result = append(result, getData(num, &indicator)...)
		result = append(result, getData(num, &mainbz)...)
		result = append(result, getData(num, &forecast)...)
		result = append(result, getData(num, &express)...)
		result = sortDeleteDuplicates(result)
		var raw [9]int
		var sum int
		for _, v := range result {
			i := v[0] - 49
			raw[i] = raw[i] + 1
			sum++
		}
		var rate [9]float64
		for i := 0; i < len(rate); i++ {
			rate[i] = 100 * float64(raw[i]) / float64(sum)
		}

		fmt.Printf("lenght: %v %v\n", sum, raw)
		fmt.Printf("rate: %.2f\n", rate)
		time.Sleep(time.Duration(1) * time.Second)

	}
}
