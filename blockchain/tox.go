package blockchain

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	intoCityNode "graph-service/binding"
	"graph-service/log"
	"math/big"
)

func ApproveTox(contractAddress string, cli *ethclient.Client, chanId int64) {
	_, auth := GetAuth(cli, chanId)
	tox, err := intoCityNode.NewBgtToken(common.HexToAddress(contractAddress), cli)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	E18 := big.NewInt(1e18)
	amount := E18.Mul(E18, big.NewInt(100000))
	tx, err := tox.Approve(auth, common.HexToAddress(contractAddress), amount)

	if err != nil {
		log.Logger.Sugar().Error(err)
		return
	}
	fmt.Println(tx, err)
}
