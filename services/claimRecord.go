package services

import (
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"graph-service/db"
	"graph-service/log"
	"graph-service/models/newbieTaskModel"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ClaimRecordResult struct {
	Data struct {
		ClaimRecords []struct {
			ID        string `json:"id"`
			Tid       string `json:"tid"`
			User      string `json:"user"`
			ClaimType string `json:"claimType"`
			Amount    string `json:"amount"`
			TxHash    string `json:"txHash"`
			Ctime     string `json:"ctime"`
		} `json:"claimRecords"`
	} `json:"data"`
}

func GetClaimRecord() error {
	url := "http://8.130.102.48:8000/subgraphs/name/into-newbie-task"
	method := "POST"
	query := fmt.Sprintf(`{"query": "query MyQuery {\n  claimRecords(first:200,orderBy:id,orderDirection:desc) { id\n    tid\n    user\n    claimType\n    amount\n   txHash\n    ctime\n    }\n}\n    ", "variables": null, "operationName": "MyQuery", "extensions": {"headers": null}}`)
	payload := strings.NewReader(query)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	var claimRecordResult ClaimRecordResult
	err = json.Unmarshal(body, &claimRecordResult)
	if err != nil {
		log.Logger.Error(err.Error())
		return err
	}
	if len(claimRecordResult.Data.ClaimRecords) <= 0 {
		return nil
	}

	recordTotal := len(claimRecordResult.Data.ClaimRecords) - 1
	for i := recordTotal; i >= 0; i-- {
		var claimRecord newbieTaskModel.ClaimRecord
		record := claimRecordResult.Data.ClaimRecords[i]
		logId := strings.Split(record.ID, "-")
		err = db.Mysql.Model(&newbieTaskModel.ClaimRecord{}).Where("user=? and tx_hash=? and log_index=? and block_height=?", record.User, record.TxHash, logId[1], logId[0]).First(&claimRecord).Error
		if err == gorm.ErrRecordNotFound {
			amount, _ := decimal.NewFromString(record.Amount)
			claimType, _ := strconv.ParseInt(record.ClaimType, 10, 64)
			ctime, _ := strconv.ParseInt(record.Ctime, 10, 64)
			db.Mysql.Model(&newbieTaskModel.ClaimRecord{}).Create(&newbieTaskModel.ClaimRecord{
				User:      record.User,
				Amount:    amount,
				TxHash:    record.TxHash,
				ClaimType: claimType,
				Ctime:     time.Unix(ctime, 0),
			})
		}
	}

	return nil
}
