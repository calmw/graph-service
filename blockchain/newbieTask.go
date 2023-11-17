package blockchain

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/status-im/keycard-go/hexutils"
	intoBinding "graph-service/binding"
	"graph-service/blockchain/utils"
	"graph-service/log"
	"math/big"
	"os"
)

type NewbieTaskConfig struct {
	ChainId                 int64
	Rpc                     string
	PrivateKey              string
	ContractAddress         string
	ServerAddress           string
	PledgePoolAddress       string
	EarlyBirdAddress        string
	ToxAddress              string
	IntoSocialMiningAddress string
	IntoBindingAddress      string
	FaceAuthAddress         string
}

// NewbieTaskConfigTestnet testnet
var NewbieTaskConfigTestnet = NewbieTaskConfig{
	ChainId:                 9001,
	Rpc:                     "https://testnet-rpc.d2ao.com/",
	PrivateKey:              os.Getenv("HUZHI"),
	ContractAddress:         "0xb210c8Dad361386ecC6F55717E783808d8F62031",
	ServerAddress:           "",
	PledgePoolAddress:       "0x2bF7D41C7E76B9118dA229B98F56C8e5C490C94B",
	EarlyBirdAddress:        "0xf4697c567901342f3742Ea8F2E6Cc5F59F805f8C",
	ToxAddress:              "0x3eE243ff68074502b1D9D65443fa97b99f634570",
	IntoSocialMiningAddress: "0xD8c1d40a6FF4E53577389C8008343081949C373D",
	IntoBindingAddress:      "0x505611a42703243162353385770DdA9ABa788973",
	FaceAuthAddress:         "0xD712221FE8b655C981Ac47385aEd078E4E497d3A",
}

// NewbieTaskConfigMainnet mainnet
var NewbieTaskConfigMainnet = NewbieTaskConfig{
	ChainId:                 9001,
	Rpc:                     "https://rpc.matchscan.io/",
	PrivateKey:              os.Getenv("HUZHI"),
	ContractAddress:         "",
	ServerAddress:           "",
	PledgePoolAddress:       "",
	EarlyBirdAddress:        "0x73e5D54d076fF1f93391A50B7a9B50018C9FE531",
	ToxAddress:              "0x96397347Ea2beE29713Bc48749eB277D6A36A407",
	IntoSocialMiningAddress: "0x0007B44b6Ca810EBff3ED4560cD7d997b08BA104",
	IntoBindingAddress:      "0x15f414e10101afDeDA12Ba5608795eb4e14f2D3d",
	FaceAuthAddress:         "0x424CaA5beA9bDfeF9F49796D7C7005632fAbE2E8",
}

func Client(c NewbieTaskConfig) *ethclient.Client {
	client, err := ethclient.Dial(c.Rpc)
	if err != nil {
		panic("dail failed")
	}
	return client
}

func GetAuth(cli *ethclient.Client, chainId int64) (error, *bind.TransactOpts) {
	privateKey := os.Getenv("HUZHI")
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKeyEcdsa, big.NewInt(chainId))
	if err != nil {
		log.Logger.Sugar().Error(err)
		return err, nil
	}

	//gasLimit := uint64(21000)
	return nil, &bind.TransactOpts{
		From:      auth.From,
		Nonce:     nil,
		Signer:    auth.Signer, // Method to use for signing the transaction (mandatory)
		Value:     big.NewInt(0),
		GasPrice:  nil,
		GasFeeCap: nil,
		GasTipCap: nil,
		GasLimit:  0,
		Context:   context.Background(),
		NoSend:    false, // Do all transact steps but do not send the transaction
	}
}

func DeploySet(c NewbieTaskConfig) {
	cli := Client(c)
	_, auth := GetAuth(cli, c.ChainId)
	task, err := intoBinding.NewNewbieTask(common.HexToAddress(c.ContractAddress), cli)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetFaceAuth(auth, common.HexToAddress(c.FaceAuthAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetTox(auth, common.HexToAddress(c.ToxAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetIntoBinding(auth, common.HexToAddress(c.IntoBindingAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetIntoSocialMining(auth, common.HexToAddress(c.IntoSocialMiningAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetServerAddress(auth, common.HexToAddress(c.ServerAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	_, err = task.AdminSetPledgePool(auth, common.HexToAddress(c.ServerAddress))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	// 设置奖励的类型，以及奖励的数量
	_, err = task.AdminSetClaimType(auth, big.NewInt(1), big.NewInt(2))
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	fmt.Println("init success")
}

func Sign() {
	privateKey := os.Getenv("HUZHI")
	//privateKey := os.Getenv("META_ACCOUNT")
	fmt.Println(privateKey)
	fmt.Println("0x94b627F4F829Ac5E97fDc556B5BEeeFf9beF417e")
	privateKeyEcdsa, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	//MessageId := utils.Byte32ToByteSlice(hexutils.HexToBytes("0xe76e2daf62ae13665edfe35b36765cf5fe8e524dc4a1b9ad71d1502178a3c9e5"))
	MessageId := hexutils.HexToBytes("e76e2daf62ae13665edfe35b36765cf5fe8e524dc4a1b9ad71d1502178a3c9e5")
	signature, err := utils.Sign(MessageId, privateKeyEcdsa)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	fmt.Println(hexutils.BytesToHex(signature))

	//hash := crypto.Keccak256Hash(hexutils.HexToBytes("e76e2daf62ae13665edfe35b36765cf5fe8e524dc4a1b9ad71d1502178a3c9e5"))
	//fmt.Printf("Hash: %x\n", hash.Bytes())
	//fmt.Printf("\n=== Now using Ecrecover ===\n")
	////signature, _ := crypto.Sign(hash.Bytes(), privateKey)
	//signature, err = utils.Sign(hash.Bytes(), privateKeyEcdsa)
	//
	//sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	//fmt.Printf(string(sigPublicKey), err)

}
