package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/config"
)

var chargeService = client.GetChargeService()

func TestSelfRecharge(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(chargeService.SelfRecharge(opts, 1e5))
}
func TestRecharge(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(chargeService.Recharge(opts, v1, 1e10))
}
func TestBalanceOf(t *testing.T) {
	balance, _ := chargeService.BalanceOf(v1)
	fmt.Println("balance:", balance)
}
func TestQueryFee(t *testing.T) {
	fmt.Println(chargeService.QueryFee(config.Info.Ddc1155Address().Hex(), [4]byte{0x63, 0x57, 0x03, 0x55}))
}
func TestSetFee(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(chargeService.SetFee(opts, config.Info.AuthorityAddress().String(), [4]byte{0xe3, 0xf0, 0x0c, 0x3a}, 10))
}
func TestAcc(t *testing.T) {
	fmt.Println(common.HexToAddress("0x1") == common.HexToAddress("1"))
	fmt.Println(common.IsHexAddress(common.HexToAddress("0").Hex()))
}
