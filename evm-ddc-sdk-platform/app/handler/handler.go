package handler

import (
	"context"
	"net/http"
	"net/url"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/config"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/contracts"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
)

type myString string

// GetConn 获取连接实体
func GetConn() (conn *ethclient.Client) {
	var ctx context.Context
	var err error
	ctx = context.Background()

	u, err := url.Parse(config.Info.OpbGatewayAddress())
	if err != nil {
		log.Error.Printf("failed to parse url :%v", err.Error())
		return nil
	}

	switch u.Scheme {
	case "http", "https":
		var client *rpc.Client
		client, err = rpc.DialHTTPWithClient(config.Info.OpbGatewayAddress(), new(http.Client))
		//设置请求头参数
		if config.Info.HeaderKey() != "" && config.Info.HeaderValue() != "" {
			client.SetHeader(config.Info.HeaderKey(), config.Info.HeaderValue())
		}
		conn = ethclient.NewClient(client)
	default:
		//预留
		if config.Info.HeaderKey() != "" && config.Info.HeaderValue() != "" {
			ctx = context.WithValue(ctx, myString(config.Info.HeaderKey()), config.Info.HeaderValue())
		}
		conn, err = ethclient.DialContext(ctx, config.Info.OpbGatewayAddress())
	}

	if err != nil {
		log.Error.Printf("failed to connect to the rpc server :%v", err.Error())
	}

	return conn
}

// GetAuthority 获取Authority合约实体
func GetAuthority() *contracts.Authority {
	authority, err := contracts.NewAuthority(config.Info.AuthorityAddress(), config.Info.Conn())
	if err != nil {
		log.Error.Printf("failed to get contract Authority:%v", err)
		return nil
	}
	return authority
}

// GetCharge 获取Charge合约实体
func GetCharge() *contracts.Charge {
	charge, err := contracts.NewCharge(config.Info.ChargeAddress(), config.Info.Conn())
	if err != nil {
		log.Error.Printf("failed to get contract Charge:%v", err)
		return nil
	}
	return charge
}

// GetDDC721 获取DDC721合约实体
func GetDDC721() *contracts.DDC721 {
	ddc721, err := contracts.NewDDC721(config.Info.Ddc721Address(), config.Info.Conn())
	if err != nil {
		log.Error.Printf("failed to get contract DDC721:%v", err)
	}
	return ddc721
}

// GetDDC1155 获取DDC1155合约实体
func GetDDC1155() *contracts.DDC1155 {
	ddc1155, err := contracts.NewDDC1155(config.Info.Ddc1155Address(), config.Info.Conn())
	if err != nil {
		log.Error.Printf("failed to get contract DDC1155:%v", err)
	}
	return ddc1155
}
