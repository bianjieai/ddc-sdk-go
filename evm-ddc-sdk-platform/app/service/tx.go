package service

import (
	"context"
	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/config"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/log"
	types2 "github.com/bianjieai/ddc-sdk-go/evm-ddc-sdk-platform/pkg/types"
)

type TxService struct {
}

// GetTxByHash
// @Description: 运营方或平台方根据交易哈希对交易信息进行查询
// @receiver t
// @param txHash: 交易哈希
// @return *types.Transaction: 交易信息
// @return bool： pending状态
// @return error
func (t TxService) GetTxByHash(txHash string) (*types.Transaction, bool, error) {
	transaction, isPending, err := config.Info.Conn().TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTxByHash: %v", err.Error())
		return transaction, isPending, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	transaction.Data()

	return transaction, isPending, nil
}

// GetTxReceipt
// @Description: 运营方或平台方根据交易哈希对交易回执信息进行查询。
// @receiver t
// @param txHash: 交易哈希
// @return string： 交易回执
// @return error
func (t TxService) GetTxReceipt(txHash string) (*types.Receipt, error) {
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTxReceipt: %v", err.Error())
		return nil, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return receipt, nil
}

// GetTxStatus
// @Description: 运营方或平台方根据交易哈希查询交易状态是否成功
// @receiver t
// @param txHash: 交易哈希
// @return bool：交易是否成功
// @return error
func (t TxService) GetTxStatus(txHash string) (bool, error) {
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to execute GetTxStatus: %v", err.Error())
		return false, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	if receipt.Status == 0 {
		return false, nil
	}

	return true, nil
}

// GetTimeByTxHash
// @Description: 通过txHash查询所在块的时间
// @receiver t
// @param txHash 交易哈希
// @return uint64 所在块的时间
// @return error
func (t TxService) GetTimeByTxHash(txHash string) (uint64, error) {
	receipt, err := config.Info.Conn().TransactionReceipt(context.Background(), common.HexToHash(txHash))
	if err != nil {
		log.Error.Printf("failed to get TransactionReceipt: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}
	block, err := config.Info.Conn().BlockByNumber(context.Background(), receipt.BlockNumber)
	if err != nil {
		log.Error.Printf("failed to execute BlockByNumber: %v", err.Error())
		return 0, types2.NewSDKError(types2.QueryError.Error(), err.Error())
	}

	return block.Time(), nil
}
