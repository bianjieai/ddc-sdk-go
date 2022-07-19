package service

import (
	"context"
	"math/big"
	"runtime"
	"strconv"
	"sync"
	"time"

	geth "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	gethethclient "github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	gethrpc "github.com/ethereum/go-ethereum/rpc"

	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/constant"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/handler"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/app/models/dto"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/config"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/contracts"
	"github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/ddc-sdk-platform-go/pkg/types"
)

const CtxTimeout = 10 * time.Second

type BlockService struct {
	abiAuthority *abi.ABI
	abiCharge    *abi.ABI
	abiDDC721    *abi.ABI
	abiDDC1155   *abi.ABI

	ethClient  *gethethclient.Client
	gethCli    *gethclient.Client
	gethRPCCli *gethrpc.Client
}

func NewBlockService() (*BlockService, error) {
	ctx, cancel := context.WithTimeout(context.Background(), CtxTimeout)
	defer cancel()
	rpcClient, err := gethrpc.DialContext(ctx, config.Info.OpbGatewayAddress())
	if err != nil {
		log.Error.Printf("failed to get rpcClient: %v", err.Error())
		return nil, err
	}

	abiAuthority, err := contracts.AuthorityMetaData.GetAbi()
	if err != nil {
		log.Error.Printf("failed to get Authority abi: %v", err.Error())
		return nil, err
	}
	abiCharge, err := contracts.ChargeMetaData.GetAbi()
	if err != nil {
		log.Error.Printf("failed to get Charge abi: %v", err.Error())
		return nil, err
	}
	abiDDC721, err := contracts.DDC721MetaData.GetAbi()
	if err != nil {
		log.Error.Printf("failed to get DDC721 abi: %v", err.Error())
		return nil, err
	}
	abiDDC1155, err := contracts.DDC1155MetaData.GetAbi()
	if err != nil {
		log.Error.Printf("failed to get DDC1155 abi: %v", err.Error())
		return nil, err
	}

	ethClient := gethethclient.NewClient(rpcClient)
	gethCli := gethclient.New(rpcClient)

	return &BlockService{
		abiAuthority: abiAuthority,
		abiCharge:    abiCharge,
		abiDDC721:    abiDDC721,
		abiDDC1155:   abiDDC1155,

		ethClient:  ethClient,
		gethCli:    gethCli,
		gethRPCCli: rpcClient,
	}, nil
}

func (b *BlockService) GetBlockByNumber(blockNumber int64) (*gethtypes.Block, error) {
	return b.ethClient.BlockByNumber(context.Background(), big.NewInt(blockNumber))
}

func (b *BlockService) GetBlockEvents(blockNumber int64) (*dto.BlockEventBean, error) {
	var result []interface{}

	logs, time, err := b.getLogs(blockNumber)
	if err != nil {
		log.Error.Printf("get logs failed :%v", err.Error())
		return &dto.BlockEventBean{}, err
	}

	logChs := make(chan gethtypes.Log, len(logs))
	resultChs := make(chan interface{}, len(logs))
	errorChs := make(chan error, 1)
	for _, v := range logs {
		logChs <- v
	}
	close(logChs) // 关闭管道，避免后面协程读阻塞！

	if len(logChs) < 1 {
		return &dto.BlockEventBean{}, nil
	}

	var wg sync.WaitGroup
	wg.Add(runtime.NumCPU() * 2)
	for i := 0; i < runtime.NumCPU()*2; i++ {
		go b.parseLogs(logChs, resultChs, errorChs, &wg)
	}
	wg.Wait()
	close(resultChs)
	for {
		ch, ok := <-resultChs
		if !ok {
			break
		}
		result = append(result, ch)
	}

	err = <-errorChs
	if err != nil {
		log.Error.Printf("failed to parse log: %v", err.Error())
		return &dto.BlockEventBean{}, err
	}

	return &dto.BlockEventBean{
		Events:    result,
		Timestamp: strconv.FormatUint(time, 10),
	}, nil
}

func (b *BlockService) getLogs(blockNumber int64) ([]gethtypes.Log, uint64, error) {
	//查找区块中所有交易
	block, err := b.ethClient.BlockByNumber(context.Background(), big.NewInt(blockNumber))
	if err != nil {
		return nil, 0, err
	}

	addresses := []common.Address{config.Info.AuthorityAddress(), config.Info.ChargeAddress(), config.Info.Ddc721Address(), config.Info.Ddc1155Address()}
	topics := [][]common.Hash{{b.abiAuthority.Events[constant.EventAddAccount].ID,
		b.abiAuthority.Events[constant.EventUpdateAccountState].ID,
		b.abiAuthority.Events[constant.EventCrossPlatformApproval].ID,
		b.abiCharge.Events[constant.EventRecharge].ID,
		b.abiCharge.Events[constant.EventPay].ID,
		b.abiCharge.Events[constant.EventSetFee].ID,
		b.abiCharge.Events[constant.EventDelFee].ID,
		b.abiCharge.Events[constant.EventDelDDC].ID,
		b.abiDDC721.Events[constant.EventTransfer].ID,
		b.abiDDC721.Events[constant.EventPay].ID,
		b.abiDDC721.Events[constant.EventEnterBlacklist].ID,
		b.abiDDC721.Events[constant.EventExitBlacklist].ID,
		b.abiDDC721.Events[constant.EventSetURI].ID,
		b.abiDDC1155.Events[constant.EventTransferSingle].ID,
		b.abiDDC1155.Events[constant.EventTransferBatch].ID,
		b.abiDDC1155.Events[constant.EventEnterBlacklist].ID,
		b.abiDDC1155.Events[constant.EventExitBlacklist].ID,
		b.abiDDC1155.Events[constant.EventSetURI].ID,
	}}

	filter := geth.FilterQuery{
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
		Addresses: addresses,
		Topics:    topics,
	}
	logs, err := b.ethClient.FilterLogs(context.Background(), filter)
	if err != nil {
		return nil, 0, err
	}

	return logs, block.Time(), nil
}

func (b BlockService) parseLogs(logs chan gethtypes.Log, resChs chan interface{}, errChs chan error, wg *sync.WaitGroup) {
	var err error

	defer func() {
		wg.Done()
		errChs <- err
	}()

	for {
		if len(logs) < 1 { // 先预过滤一批协程
			return
		}
		// 取出tx数据
		l, ok := <-logs
		if !ok { // 管道数据已被读完，当前协程直接退出
			return
		}
		authority := handler.GetAuthority()
		charge := handler.GetCharge()
		ddc721 := handler.GetDDC721()
		ddc1155 := handler.GetDDC1155()

		var event interface{}
		//2.匹配对应的事件
		switch l.Topics[0] {
		case b.abiAuthority.Events[constant.EventAddAccount].ID:
			event, err = authority.ParseAddAccount(l)
		case b.abiAuthority.Events[constant.EventUpdateAccountState].ID:
			event, err = authority.ParseUpdateAccountState(l)
		case b.abiAuthority.Events[constant.EventCrossPlatformApproval].ID:
			event, err = authority.ParseCrossPlatformApproval(l)
		case b.abiCharge.Events[constant.EventRecharge].ID:
			event, err = charge.ParseRecharge(l)
		case b.abiCharge.Events[constant.EventPay].ID:
			event, err = charge.ParsePay(l)
		case b.abiCharge.Events[constant.EventSetFee].ID:
			event, err = charge.ParseSetFee(l)
		case b.abiCharge.Events[constant.EventDelFee].ID:
			event, err = charge.ParseDelFee(l)
		case b.abiCharge.Events[constant.EventDelDDC].ID:
			event, err = charge.ParseDelDDC(l)
		case b.abiDDC721.Events[constant.EventTransfer].ID:
			event, err = ddc721.ParseTransfer(l)
		case b.abiDDC721.Events[constant.EventPay].ID:
			event, err = ddc721.ParseOwnershipTransferred(l)
		case b.abiDDC721.Events[constant.EventEnterBlacklist].ID:
			event, err = ddc721.ParseEnterBlacklist(l)
		case b.abiDDC721.Events[constant.EventExitBlacklist].ID:
			event, err = ddc721.ParseExitBlacklist(l)
		case b.abiDDC721.Events[constant.EventSetURI].ID:
			event, err = ddc721.ParseSetURI(l)
		case b.abiDDC1155.Events[constant.EventTransferSingle].ID:
			event, err = ddc1155.ParseTransferSingle(l)
		case b.abiDDC1155.Events[constant.EventTransferBatch].ID:
			event, err = ddc1155.ParseTransferBatch(l)
		case b.abiDDC1155.Events[constant.EventEnterBlacklist].ID:
			event, err = ddc1155.ParseEnterBlacklist(l)
		case b.abiDDC1155.Events[constant.EventExitBlacklist].ID:
			event, err = ddc1155.ParseExitBlacklist(l)
		case b.abiDDC1155.Events[constant.EventSetURI].ID:
			event, err = ddc1155.ParseSetURI(l)
		}
		if err != nil {
			return
		}
		resChs <- event
	}
}

// GetTxEvents
// @Description: 获取指定交易中的所有events
// @receiver b
// @param txHash 交易哈希
// @return events 查询出的事件
// @return err
func (b *BlockService) GetTxEvents(txHash common.Hash) (events []interface{}, err error) {
	//获取对应的交易回执
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Error.Printf("failed to execute TransactionReceipt: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	var event interface{}
	var abi *abi.ABI
	//获取交易的logs中的所有log对应的event
	for _, l := range receipt.Logs {
		//1.匹配对应的合约
		switch l.Address {
		case config.Info.AuthorityAddress():
			{
				authority := handler.GetAuthority()

				abi, err = contracts.AuthorityMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Authority abi: %v", err.Error())
					return nil, err
				}
				//2.匹配对应的事件
				switch l.Topics[0] {
				case abi.Events[constant.EventAddAccount].ID:
					event, err = authority.ParseAddAccount(*l)
				case abi.Events[constant.EventUpdateAccountState].ID:
					event, err = authority.ParseUpdateAccountState(*l)
				case abi.Events[constant.EventCrossPlatformApproval].ID:
					event, err = authority.ParseCrossPlatformApproval(*l)
				}
			}
		case config.Info.ChargeAddress():
			{
				charge := handler.GetCharge()
				abi, err = contracts.ChargeMetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get Charge abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventRecharge].ID:
					event, err = charge.ParseRecharge(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = charge.ParsePay(*l)
				case abi.Events[constant.EventSetFee].ID:
					event, err = charge.ParseSetFee(*l)
				case abi.Events[constant.EventDelFee].ID:
					event, err = charge.ParseDelFee(*l)
				case abi.Events[constant.EventDelDDC].ID:
					event, err = charge.ParseDelDDC(*l)
				}
			}
		case config.Info.Ddc721Address():
			{
				ddc721 := handler.GetDDC721()
				abi, err = contracts.DDC721MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC721 abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransfer].ID:
					event, err = ddc721.ParseTransfer(*l)
				case abi.Events[constant.EventPay].ID:
					event, err = ddc721.ParseOwnershipTransferred(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc721.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc721.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc721.ParseSetURI(*l)
				}
			}
		case config.Info.Ddc1155Address():
			{
				ddc1155 := handler.GetDDC1155()
				abi, err = contracts.DDC1155MetaData.GetAbi()
				if err != nil {
					log.Error.Printf("failed to get DDC1155 abi: %v", err.Error())
					return nil, err
				}
				switch l.Topics[0] {
				case abi.Events[constant.EventTransferSingle].ID:
					event, err = ddc1155.ParseTransferSingle(*l)
				case abi.Events[constant.EventTransferBatch].ID:
					event, err = ddc1155.ParseTransferBatch(*l)
				case abi.Events[constant.EventEnterBlacklist].ID:
					event, err = ddc1155.ParseEnterBlacklist(*l)
				case abi.Events[constant.EventExitBlacklist].ID:
					event, err = ddc1155.ParseExitBlacklist(*l)
				case abi.Events[constant.EventSetURI].ID:
					event, err = ddc1155.ParseSetURI(*l)
				}
			}
		}
		if err != nil {
			log.Error.Printf("failed to parse event: %v", err.Error())
			return nil, err
		}
		events = append(events, event)
	}

	return
}
