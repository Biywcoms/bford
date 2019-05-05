package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const Token = "b5dee807f02c4eccc01525e7bcd82b02214c0576573549d970a77b"
const ApiUrl = "http://api.tushare.pro"

type PostBody struct {
	TOKEN    string            `json:"token"`
	API_NAME string            `json:"api_name"`
	PARAMS   map[string]string `json:"params"`
	FIELDS   string            `json:"fields"`
}

func getData(post *PostBody) {
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

		//		fmt.Println(data["data"])
		data, ok := result["data"].(map[string]interface{})
		if !ok {
			panic(err)
		}
		fmt.Println(data["items"])
		//		for _, v := range data["data"] {
		//			fmt.Println(v)
		//		}

		//		fmt.Println(dd["itmes"])
		//		fmt.Printf("Type:%T", dd["items"])
	}

}

func main() {

	num := "000002.SZ"
	date := "20190331"

	//利润表
	income := PostBody{Token, "", nil, ""}
	income.API_NAME = "income"
	income.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	income.FIELDS = ""

	//资产负载表
	balancesheet := PostBody{Token, "", nil, ""}
	balancesheet.API_NAME = "balancesheet"
	balancesheet.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	balancesheet.FIELDS = ""

	//现金流量表
	cashflow := PostBody{Token, "", nil, ""}
	cashflow.API_NAME = "cashflow"
	cashflow.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	cashflow.FIELDS = ""

	//业绩预告
	forecast := PostBody{Token, "", nil, ""}
	forecast.API_NAME = "forecast"
	forecast.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	forecast.FIELDS = ""

	//业绩快报
	express := PostBody{Token, "", nil, ""}
	express.API_NAME = "express"
	express.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	express.FIELDS = ""

	//财务指标数据
	indicator := PostBody{Token, "", nil, ""}
	indicator.API_NAME = "fina_indicator"
	indicator.PARAMS = map[string]string{
		"ts_code": num,
		"period":  date,
	}
	indicator.FIELDS = ""

	//主营业务
	mainbz := PostBody{Token, "", nil, ""}
	mainbz.API_NAME = "fina_mainbz"
	mainbz.PARAMS = map[string]string{
		"ts_code": num,
	}
	mainbz.FIELDS = ""

	report := PostBody{Token, "", nil, ""}
	report.API_NAME = "disclosure_date"
	report.PARAMS = map[string]string{
		"ts_code": num,
	}

	getData(&mainbz)
	//	fmt.Println("response Body:", string(body))

}
