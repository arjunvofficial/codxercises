package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var filePath = "bank_details.csv"

func main() {

	// created a simple handler just to give a json response back
	http.HandleFunc("/", handler)
	log.Print("server started")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	data := csvToJSON()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func csvToJSON() interface{} {
	desc, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	reader := csv.NewReader(bufio.NewReader(desc))
	// usually in csv files there will be coma (,) seperated value
	// but in some cases semicolon (;) seperated value are found
	// so change the (,) value to (;)
	// reader.Comma = ';'

	var bankInfo BankDetails
	var bankDetails []BankDetails
	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for count, line := range csvData {
		if count == 0 {
			continue
		}
		bankInfo.Age, _ = strconv.ParseInt(line[0], 0, 64)
		bankInfo.Job = line[1]
		bankInfo.Marital = line[2]
		bankInfo.Education = line[3]
		bankInfo.Default = line[4]
		bankInfo.Balance, _ = strconv.ParseInt(line[5], 0, 64)
		bankInfo.Housing = line[6]
		bankInfo.Loan = line[7]
		bankInfo.Contact = line[8]
		bankInfo.Day, _ = strconv.ParseInt(line[9], 0, 64)
		bankInfo.Month = line[10]
		bankInfo.Duration, _ = strconv.ParseInt(line[11], 0, 64)
		bankInfo.Campaign, _ = strconv.ParseInt(line[12], 0, 64)
		bankInfo.Pdays, _ = strconv.ParseInt(line[13], 0, 64)
		bankInfo.Previous, _ = strconv.ParseInt(line[14], 0, 64)
		bankInfo.Poutcome = line[15]
		bankDetails = append(bankDetails, bankInfo)
	}

	// used to create a json file from output
	writeToFile("bank_details", bankDetails)
	return bankDetails
}

// BankDetails contains the information about bank
type BankDetails struct {
	Age       int64  `json:"age"`
	Job       string `json:"job"`
	Marital   string `json:"marital"`
	Education string `json:"education"`
	Default   string `json:"default"`
	Balance   int64  `json:"balance"`
	Housing   string `json:"housing"`
	Loan      string `json:"loan"`
	Contact   string `json:"contact"`
	Day       int64  `json:"day"`
	Month     string `json:"month"`
	Duration  int64  `json:"duration"`
	Campaign  int64  `json:"campaign"`
	Pdays     int64  `json:"pdays"`
	Previous  int64  `json:"previous"`
	Poutcome  string `json:"poutcome"`
}

func writeToFile(name string, data interface{}) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(name+".json", file, 0644)
}
