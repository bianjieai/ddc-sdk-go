package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/config"
)

var chargeService = client.GetChargeService()

func TestRecharge(t *testing.T) {
	opts.From = common.HexToAddress(platform)
	fmt.Println(chargeService.Recharge(opts, v1, 1e10))
}
func TestBalanceOf(t *testing.T) {
	balance, _ := chargeService.BalanceOf(v1)
	fmt.Println("balance:", balance)
}
func TestQueryFee(t *testing.T) {
	fmt.Println(chargeService.QueryFee(config.Info.Ddc1155Address().Hex(), [4]byte{0x63, 0x57, 0x03, 0x55}))
}
