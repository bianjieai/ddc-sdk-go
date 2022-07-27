package dto

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app/listener"
)

type ConfigInfo struct {
	conn              *ethclient.Client
	opbGatewayAddress string
	headerKey         string
	headerValue       string
	signEventListener listener.SignEventListener
	gasPrice          uint64
	authorityAddress  common.Address
	chargeAddress     common.Address
	ddc721Address     common.Address
	ddc1155Address    common.Address
}

func NewConfigInfo() *ConfigInfo {
	return &ConfigInfo{}
}

func (c *ConfigInfo) Conn() *ethclient.Client {
	return c.conn
}

func (c *ConfigInfo) SetConn(conn *ethclient.Client) {
	c.conn = conn
}

func (c *ConfigInfo) OpbGatewayAddress() string {
	return c.opbGatewayAddress
}

func (c *ConfigInfo) SetOpbGatewayAddress(opbGatewayAddress string) {
	c.opbGatewayAddress = opbGatewayAddress
}

func (c *ConfigInfo) HeaderKey() string {
	return c.headerKey
}

func (c *ConfigInfo) SetHeaderKey(headerKey string) {
	c.headerKey = headerKey
}

func (c *ConfigInfo) HeaderValue() string {
	return c.headerValue
}

func (c *ConfigInfo) SetHeaderValue(headerValue string) {
	c.headerValue = headerValue
}

func (c *ConfigInfo) SignEventListener() listener.SignEventListener {
	return c.signEventListener
}

func (c *ConfigInfo) SetSignEventListener(signEventListener listener.SignEventListener) {
	c.signEventListener = signEventListener
}

func (c *ConfigInfo) GasPrice() uint64 {
	return c.gasPrice
}

func (c *ConfigInfo) SetGasPrice(gasPrice uint64) {
	c.gasPrice = gasPrice
}

func (c *ConfigInfo) AuthorityAddress() common.Address {
	return c.authorityAddress
}

func (c *ConfigInfo) SetAuthorityAddress(authorityAddress string) {
	c.authorityAddress = common.HexToAddress(authorityAddress)
}

func (c *ConfigInfo) ChargeAddress() common.Address {
	return c.chargeAddress
}

func (c *ConfigInfo) SetChargeAddress(chargeAddress string) {
	c.chargeAddress = common.HexToAddress(chargeAddress)
}

func (c *ConfigInfo) Ddc721Address() common.Address {
	return c.ddc721Address
}

func (c *ConfigInfo) SetDdc721Address(ddc721Address string) {
	c.ddc721Address = common.HexToAddress(ddc721Address)
}

func (c *ConfigInfo) Ddc1155Address() common.Address {
	return c.ddc1155Address
}

func (c *ConfigInfo) SetDdc1155Address(ddc1155Address string) {
	c.ddc1155Address = common.HexToAddress(ddc1155Address)
}
