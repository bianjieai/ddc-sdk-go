package test

import (
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/app"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

var (
	clientBuilder = app.DDCSdkClientBuilder{}
	client        = clientBuilder.
			SetSignEventListener(new(SignListener)). //注册实现了签名接口的结构体
			SetGasPrice(1).                          //设置gasPrice
			SetGatewayURL("https://opbningxia.bsngate.com:18602/api/[ProjectId]/evmrpc").
			SetGatewayAPIKey("x-api-key").
			SetGatewayAPIValue("[ProjectKey]").
			SetAuthorityAddress("0xa7FC5B0F4A0085c5Ce689b919a866675Ce37B66b").
			SetChargeAddress("0x3BBb01B38958d4dbF1e004611EbB3c65979B0511").
			SetDDC721Address("0x3B09b7A00271C5d9AE84593850dE3A526b8BF96e").
			SetDDC1155Address("0xe5d3b9E7D16E03A4A1060c72b5D1cb7806DD9070").
			RegisterLog("./log.log"). //设置日志输出的文件路径
			Build()

	opts = new(bind.TransactOpts)

	platform    = "0x7FAF93F524FFDD1FB36BEC0ED6A167E8CE55BC4E" //ddcID：2、4、6
	platformPri = "0xED43B9686AB520C896BC33A1461BAF163EDBF0DBC4D3199E77793A49B9BB2568"
	//无gas
	v1    = "0x07B7BE76ED588CCEFB4C4A573CB28A7D2A1403CC"
	v1Pri = "0x1B8C36A57CB8D7FA20594283498EF310DCA9DFECDF6E9FDD04E992A5DA164E0B"
	//无gas
	v2    = "0x918F7F275A6C2D158E5B76F769D3F1678958A334" //ddcID:5
	v2Pri = "0x2F6976C530CFD2D0CC19EFC1868BD6A0AA1886D0BFCFA5A59D9B8899BE9B7241"

	//账户和私钥的映射
	senderPri = map[string]string{
		platform: platformPri,
		v1:       v1Pri,
		v2:       v2Pri,
	}
)
