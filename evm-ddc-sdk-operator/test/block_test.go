package test

import (
	"fmt"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/app/service"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-operator/pkg/contracts"
)

func TestGetBlockByNumber(t *testing.T) {

	fmt.Println(client.GetBlockByNumber(301751))
}
func TestGetEvents(t *testing.T) {

	fmt.Println(client.GetBlockEvents(15051000))
}
func TestGetEvent(t *testing.T) {

	events, _ := client.GetTxEvents(common.HexToHash("0xfff913ba0f63fc3b5ead83453d8cd8ad95c142735ef8f71c285e267112fe1676"))

	for _, e := range events {
		res, ok := e.(*contracts.DDC721Transfer)
		if !ok {
			continue
		}
		fmt.Println(res.DdcId)
	}
}

func TestEvents(t *testing.T) {

	sss, err := service.NewBlockService()
	if err != nil {
		fmt.Println("---------------err", err)
	}
	events, err := sss.GetBlockEvents(15057363)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(events)
}
