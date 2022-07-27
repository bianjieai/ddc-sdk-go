package test

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/config"
)

var authorityService = client.GetAuthorityService()

func TestAddConsumerByOperator(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(authorityService.AddAccountByOperator(opts, platform, "ddcplatform", "did:ddcplatform", ""))
}
func TestGetAccount(t *testing.T) {
	fmt.Println(authorityService.GetAccount(platform))
}
func TestAddFunction(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	//safeMint
	//fmt.Println(authorityService.AddFunction(opts, 2, config.Info.Ddc721Address().Hex(), [4]byte{0xf6, 0xdd, 0xa9, 0x36}))
	fmt.Println(authorityService.AddFunction(opts, 2, config.Info.Ddc721Address().Hex(), [4]byte{0x06, 0xfd, 0xde, 0x03}))

}
func TestUpdateAccState(t *testing.T) {
	opts.From = common.HexToAddress(operator)
	fmt.Println(authorityService.UpdateAccState(opts, platform, 1, true))
}
func TestHasFunctionPermission(t *testing.T) {

	fmt.Println(authorityService.HasFunctionPermission("0x228D38202B6311B6D011E839B10E0CD99A21D0C8", "0x1C917baf05863417391acCfe85d305Eae41401Ec", [4]byte{0x06, 0xfd, 0xde, 0x03}))
}
