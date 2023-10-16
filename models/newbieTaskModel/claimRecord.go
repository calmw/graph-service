package newbieTaskModel

import (
	"github.com/shopspring/decimal"
	"time"
)

type ClaimRecord struct {
	Id          int             `gorm:"column:id;primaryKey"`
	User        string          `json:"user" gorm:"column:user"`
	Amount      decimal.Decimal `json:"amount" gorm:"column:amount"`
	TxHash      string          `json:"tx_hash" gorm:"column:tx_hash"`
	ClaimType   int64           `json:"claim_type" gorm:"column:claim_type"`
	BlockHeight int64           `json:"block_height" gorm:"column:block_height"`
	LogIndex    int64           `json:"log_index" gorm:"column:log_index"`
	Ctime       time.Time       `json:"ctime" gorm:"column:ctime"`
}
