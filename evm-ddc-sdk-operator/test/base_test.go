package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/app/constant"
)

func TestEstimateGasLimit(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(client.EstimateGasLimit(opts, constant.ContrDDC721, constant.FuncMint, common.HexToAddress(platform), "www.123.com"))
}
func TestGetDDCId(t *testing.T) {
	fmt.Println(client.DDCIdByHash("0xc806461847abf17c899a027edddeaa8c0ea57b64f10580f63c1c385f9aa3b282"))
}
func TestCreateAccount(t *testing.T) {
	fmt.Println(client.CreateAccount())
}
func TestHexToBech32(t *testing.T) {
	fmt.Println(client.HexToBech32("0x5BD8DE9E786FED6A2755CF1BB1E0D322B0827612"))
}
