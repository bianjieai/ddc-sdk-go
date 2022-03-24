package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/constant"
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
func TestAddress(t *testing.T) {
	fmt.Println(common.HexToAddress("0") == common.HexToAddress("0"))
}
