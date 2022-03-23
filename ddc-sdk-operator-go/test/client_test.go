package test

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-operator-go/app"
)

var (
	clientBuilder = app.DDCSdkClientBuilder{}
	client        = clientBuilder.
			SetSignEventListener(new(SignListener)).
			SetGasPrice(1e10).
			SetGatewayUrl("http://192.168.150.42:8545").
			SetAuthorityAddress("0x6a3B24042dA7Bb5F2CBF1BCB2ABE0C632590C580").
			SetChargeAddress("0x95aDFbA9050C5D886419334Ae478b9844f413eF2").
			SetDDC721Address("0x1C917baf05863417391acCfe85d305Eae41401Ec").
			SetDDC1155Address("0x02A25C69843E197e3063Ed848f6FEA512633CB8E").
			RegisterLog("./log.log").
			Build()
	//client        = clientBuilder.
	//		SetSignEventListener(new(SignListener)).
	//		SetGasPrice(1e10).
	//		SetGatewayUrl("http://192.168.150.42:8545").
	//			SetAuthorityAddress("0x0754D6B28cB82cce0E19FE495e9eca445F1431aF").
	//		SetChargeAddress("0x95aDFbA9050C5D886419334Ae478b9844f413eF2").
	//		SetDDC721Address("0x1C917baf05863417391acCfe85d305Eae41401Ec").
	//		SetDDC1155Address("0x02A25C69843E197e3063Ed848f6FEA512633CB8E").
	//		RegisterLog("./log.log").
	//		Build()

	opts     = new(bind.TransactOpts)
	sender   = "0x7FAF93F524FFDD1FB36BEC0ED6A167E8CE55BC4E"
	senderPr = "0xED43B9686AB520C896BC33A1461BAF163EDBF0DBC4D3199E77793A49B9BB2568"

	genV1  = "0x4F3B9C50A29562E3B71144A62E8447AC0DE7EC33"
	genV1P = "0x199a46b9bb8a68a2c1c9bc160cf1ee9614585d088826c48635c707ef18775366"

	pl  = "0xEDEDA0F1C1FAF1F34C9126975267FC6F95BE75F9"
	plP = "0xDCF313CD386BCF3787ECF9968F603729B278C2A0DCA18E940E3EF59EA600FD8E"

	operator    = "0x02CEB40D892061D457E7FA346988D0FF329935DF"
	operatorPri = "0xE253AB375A5806FA331E7DB32EDE524BD7D998475A60C957806066F14F479C25"
	platform    = "0x7FAF93F524FFDD1FB36BEC0ED6A167E8CE55BC4E" //ddcId：2、4、6
	platformPri = "0xED43B9686AB520C896BC33A1461BAF163EDBF0DBC4D3199E77793A49B9BB2568"
	//无gas
	v1    = "0x07B7BE76ED588CCEFB4C4A573CB28A7D2A1403CC"
	v1Pri = "0x1B8C36A57CB8D7FA20594283498EF310DCA9DFECDF6E9FDD04E992A5DA164E0B"
	//无gas
	v2    = "0x918F7F275A6C2D158E5B76F769D3F1678958A334" //ddcId:5
	v2Pri = "0x2F6976C530CFD2D0CC19EFC1868BD6A0AA1886D0BFCFA5A59D9B8899BE9B7241"

	//账户和私钥的映射
	senderPri = map[string]string{
		genV1:    genV1P,
		pl:       plP,
		sender:   senderPr,
		operator: operatorPri,
		platform: platformPri,
		v1:       v1Pri,
		v2:       v2Pri,
	}
)
