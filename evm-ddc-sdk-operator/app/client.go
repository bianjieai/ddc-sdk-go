package app

import (
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/handler"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/listener"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app/service"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/config"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/pkg/log"
)

type DDCSdkClient struct {
	service.Base
	signEventListener listener.SignEventListener
	gasPrice          uint64
	opbGatewayAddress string
	headerKey         string
	headerValue       string
	authorityAddress  string
	chargeAddress     string
	ddc721Address     string
	ddc1155Address    string
	logFile           string
}

func (c *DDCSdkClient) GetAuthorityService() *service.AuthorityService {
	return service.NewAuthorityService()
}
func (c *DDCSdkClient) GetChargeService() *service.ChargeService {
	return service.NewChargeService()
}
func (c *DDCSdkClient) GetDDC721Service() *service.DDC721Service {
	return service.NewDDC721Service()
}
func (c *DDCSdkClient) GetDDC1155Service() *service.DDC1155Service {
	return service.NewDDC1155Service()
}

// DDCSdkClientBuilder 建造者模式
type DDCSdkClientBuilder struct {
	//被建造的变量
	client *DDCSdkClient
}

func (d *DDCSdkClientBuilder) SetSignEventListener(signEventListener listener.SignEventListener) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.signEventListener = signEventListener
	return d
}

func (d *DDCSdkClientBuilder) SetGasPrice(gasPrice uint64) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.gasPrice = gasPrice
	return d
}

func (d *DDCSdkClientBuilder) SetGatewayURL(url string) *DDCSdkClientBuilder {

	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.opbGatewayAddress = url
	return d
}
func (d *DDCSdkClientBuilder) SetGatewayAPIKey(apiKey string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.headerKey = apiKey
	return d
}
func (d *DDCSdkClientBuilder) SetGatewayAPIValue(apiValue string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.headerValue = apiValue
	return d
}
func (d *DDCSdkClientBuilder) SetAuthorityAddress(authorityAddress string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.authorityAddress = authorityAddress
	return d
}
func (d *DDCSdkClientBuilder) SetChargeAddress(chargeAddress string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.chargeAddress = chargeAddress
	return d
}
func (d *DDCSdkClientBuilder) SetDDC721Address(ddc721Address string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.ddc721Address = ddc721Address
	return d
}
func (d *DDCSdkClientBuilder) SetDDC1155Address(ddc1155Address string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.ddc1155Address = ddc1155Address
	return d
}
func (d *DDCSdkClientBuilder) RegisterLog(filePath string) *DDCSdkClientBuilder {
	if d.client == nil {
		d.client = &DDCSdkClient{}
	}
	d.client.logFile = filePath
	return d
}

func (d *DDCSdkClientBuilder) Build() *DDCSdkClient {
	if d.client.signEventListener == nil {
		log.Error.Println("signEventListener must be set")
		panic("signEventListener must be set")
	}
	// 将数据写入全局变量
	config.Init(d.client.signEventListener,
		d.client.gasPrice,
		d.client.opbGatewayAddress,
		d.client.headerKey,
		d.client.headerValue,
		d.client.authorityAddress,
		d.client.chargeAddress,
		d.client.ddc721Address,
		d.client.ddc1155Address)
	//设置logFile
	log.RegisterLog(d.client.logFile)

	//获取连接
	config.Info.SetConn(handler.GetConn())
	blockService, err := service.NewBlockService()
	if err != nil {
		panic(err)
	}
	d.client.BlockService = *blockService
	return d.client
}
