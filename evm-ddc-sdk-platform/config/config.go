package config

import (
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/listener"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/models/dto"
)

var (
	Info = dto.NewConfigInfo()
)

func Init(signEventListener listener.SignEventListener, gasPrice uint64, opbGatewayAddress, headerKey, headerValue, authorityAddress, chargeAddress, ddc721Address, ddc1155Address string) {
	//参数校验：start
	//...
	//参数校验：end
	Info.SetSignEventListener(signEventListener)
	Info.SetGasPrice(gasPrice)
	Info.SetOpbGatewayAddress(opbGatewayAddress)
	Info.SetHeaderKey(headerKey)
	Info.SetHeaderValue(headerValue)
	Info.SetAuthorityAddress(authorityAddress)
	Info.SetChargeAddress(chargeAddress)
	Info.SetDdc721Address(ddc721Address)
	Info.SetDdc1155Address(ddc1155Address)

}
