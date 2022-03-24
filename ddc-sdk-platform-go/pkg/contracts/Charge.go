// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ChargeMetaData contains all meta data concerning the Charge contract.
var ChargeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"DelDDC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"DelFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"Pay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recharge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address[]\",\"name\":\"toList\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"RechargeBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"SetFee\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accAddr\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accAddrs\",\"type\":\"address[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"}],\"name\":\"delDDC\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"delFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"ddcId\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"queryFee\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"toList\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"rechargeBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"selfRecharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"authorityProxyAddress\",\"type\":\"address\"}],\"name\":\"setAuthorityProxyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"internalType\":\"uint32\",\"name\":\"amount\",\"type\":\"uint32\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"ddcAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"settlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b50600054610100900460ff168061002e575060005460ff16155b6100955760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff161580156100b7576000805461ffff19166101011790555b80156100c9576000805461ff00191690555b50608051612585620000fb60003960008181610509015281816105490152818161067f01526106bf01526125856000f3fe6080604052600436106101135760003560e01c8063715018a6116100a0578063c9c45a0f11610064578063c9c45a0f146102d0578063d213fe4514610305578063ef18e3c914610325578063f1e8b16814610345578063f2fde38b1461036557600080fd5b8063715018a61461023e5780638129fc1c146102535780638a27a80d146102685780638da5cb5b14610288578063c5837d82146102b057600080fd5b80633659cfe6116100e75780633659cfe61461019e578063458c738e146101be5780634f1ef286146101eb57806363569189146101fe57806370a082311461021e57600080fd5b80620b8f7014610118578063093f28e01461013a57806318160ddd1461015a57806336351c7c1461017e575b600080fd5b34801561012457600080fd5b50610138610133366004611cf7565b610385565b005b34801561014657600080fd5b50610138610155366004611d2a565b610408565b34801561016657600080fd5b5060c9545b6040519081526020015b60405180910390f35b34801561018a57600080fd5b50610138610199366004611d5d565b6104a4565b3480156101aa57600080fd5b506101386101b9366004611cf7565b6104ff565b3480156101ca57600080fd5b506101de6101d9366004611e64565b6105c7565b6040516101759190611ea1565b6101386101f9366004611f0d565b610675565b34801561020a57600080fd5b50610138610219366004611f99565b61072e565b34801561022a57600080fd5b5061016b610239366004611cf7565b6107ee565b34801561024a57600080fd5b50610138610834565b34801561025f57600080fd5b5061013861086a565b34801561027457600080fd5b50610138610283366004611d5d565b6108e5565b34801561029457600080fd5b506033546040516001600160a01b039091168152602001610175565b3480156102bc57600080fd5b506101386102cb366004611cf7565b610a06565b3480156102dc57600080fd5b506102f06102eb366004611d2a565b610aab565b60405163ffffffff9091168152602001610175565b34801561031157600080fd5b50610138610320366004611fe9565b610b8a565b34801561033157600080fd5b50610138610340366004612002565b610c28565b34801561035157600080fd5b5061013861036036600461203e565b610d28565b34801561037157600080fd5b50610138610380366004611cf7565b610e60565b6001600160a01b0381166103b45760405162461bcd60e51b81526004016103ab906120f9565b60405180910390fd5b6103bc610ef8565b6001600160a01b038116600081815260ca6020526040808220600101805460ff19169055517f0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e89190a250565b6001600160a01b03821661042e5760405162461bcd60e51b81526004016103ab906120f9565b610436610ef8565b6001600160a01b038216600081815260ca602090815260408083206001600160e01b0319861680855290835292819020805463ffffffff19169055519182527f2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d910160405180910390a25050565b6104af338383610fc5565b6104ba338383611131565b6040518181526001600160a01b0383169033907f4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1906020015b60405180910390a35050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036105475760405162461bcd60e51b81526004016103ab90612126565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316610579611223565b6001600160a01b03161461059f5760405162461bcd60e51b81526004016103ab90612172565b6105a881611251565b604080516000808252602082019092526105c49183919061127b565b50565b60606000825167ffffffffffffffff8111156105e5576105e5611d87565b60405190808252806020026020018201604052801561060e578160200160208202803683370190505b50905060005b835181101561066e5761063f848281518110610632576106326121be565b60200260200101516107ee565b828281518110610651576106516121be565b602090810291909101015280610666816121ea565b915050610614565b5092915050565b6001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001630036106bd5760405162461bcd60e51b81526004016103ab90612126565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166106ef611223565b6001600160a01b0316146107155760405162461bcd60e51b81526004016103ab90612172565b61071e82611251565b61072a8282600161127b565b5050565b6001600160a01b0383166107545760405162461bcd60e51b81526004016103ab906120f9565b61075c610ef8565b6001600160a01b038316600081815260ca602090815260408083206001808201805460ff191690911790556001600160e01b0319871680855290835292819020805463ffffffff191663ffffffff87169081179091558151938452918301919091527f929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc910160405180910390a2505050565b60006001600160a01b03821681036108185760405162461bcd60e51b81526004016103ab906120f9565b506001600160a01b0316600090815260cb602052604090205490565b6033546001600160a01b0316331461085e5760405162461bcd60e51b81526004016103ab90612203565b61086860006113c6565b565b600054610100900460ff1680610883575060005460ff16155b61089f5760405162461bcd60e51b81526004016103ab90612238565b600054610100900460ff161580156108c1576000805461ffff19166101011790555b6108c9611418565b6108d161147f565b80156105c4576000805461ff001916905550565b806000036109355760405162461bcd60e51b815260206004820181905260248201527f6368617267653a206e6f207472616e73666572206973206e656365737361727960448201526064016103ab565b6001600160a01b0382163b1515801561096957506001600160a01b038216600090815260ca602052604090206001015460ff165b6109b55760405162461bcd60e51b815260206004820152601860248201527f6368617267653a206e6f742044444320636f6e7472616374000000000000000060448201526064016103ab565b6109bd610ef8565b6109c982335b83611131565b6040518181526001600160a01b0383169033907fca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be906020016104f3565b6033546001600160a01b03163314610a305760405162461bcd60e51b81526004016103ab90612203565b6001600160a01b038116600003610a895760405162461bcd60e51b815260206004820152601d60248201527f6368617267653a206175746820746865207a65726f206164647265737300000060448201526064016103ab565b60cc80546001600160a01b0319166001600160a01b0392909216919091179055565b60006001600160a01b0383168103610ad55760405162461bcd60e51b81526004016103ab906120f9565b6001600160a01b038316600090815260ca602052604090206001015460ff16610b4e5760405162461bcd60e51b815260206004820152602560248201527f6368617267653a6464632070726f787920636f6e747261637420756e617661696044820152646c61626c6560d81b60648201526084016103ab565b506001600160a01b038216600090815260ca602090815260408083206001600160e01b03198516845290915290205463ffffffff165b92915050565b80600003610bda5760405162461bcd60e51b815260206004820181905260248201527f6368617267653a206e6f207472616e73666572206973206e656365737361727960448201526064016103ab565b610be2610ef8565b610bed6000336109c3565b60405181815233906000907f4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb19060200160405180910390a350565b6001600160a01b038316610c4e5760405162461bcd60e51b81526004016103ab906120f9565b80600003610c955760405162461bcd60e51b815260206004820152601460248201527318da185c99d94e9a5b9d985b1a5908191918d25960621b60448201526064016103ab565b336000610ca28285610aab565b905063ffffffff811615610cc157610cc185838363ffffffff16611131565b604080516001600160e01b03198616815263ffffffff831660208201529081018490526001600160a01b0380841691908716907f750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab9060600160405180910390a35050505050565b8051825114610d725760405162461bcd60e51b81526020600482015260166024820152750c6d0c2e4ceca74d8cadccee8d040dad2e6dac2e8c6d60531b60448201526064016103ab565b60005b8251811015610e0a57610dbb33848381518110610d9457610d946121be565b6020026020010151848481518110610dae57610dae6121be565b6020026020010151610fc5565b610df833848381518110610dd157610dd16121be565b6020026020010151848481518110610deb57610deb6121be565b6020026020010151611131565b80610e02816121ea565b915050610d75565b5081604051610e199190612286565b6040518091039020610e283390565b6001600160a01b03167f744c96b20cd1add6104a750930b0e49658e6c5542ed4e9c9a90ae33359ceb8f9836040516104f39190611ea1565b6033546001600160a01b03163314610e8a5760405162461bcd60e51b81526004016103ab90612203565b6001600160a01b038116610eef5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103ab565b6105c4816113c6565b60cc546001600160a01b031663ed5cad643360006040518363ffffffff1660e01b8152600401610f299291906122db565b602060405180830381865afa158015610f46573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6a9190612316565b6108685760405162461bcd60e51b815260206004820152602660248201527f4444433732313a6e6f742061206f70657261746f7220726f6c65206f722064696044820152651cd8589b195960d21b60648201526084016103ab565b6001600160a01b0382166110125760405162461bcd60e51b81526020600482015260146024820152736368617267653a207a65726f206164647265737360601b60448201526064016103ab565b816001600160a01b0316836001600160a01b03160361107e5760405162461bcd60e51b815260206004820152602260248201527f6368617267653a20696e76616c6964207265636861726765206f70657261746960448201526137b760f11b60648201526084016103ab565b806000036110c75760405162461bcd60e51b815260206004820152601660248201527518da185c99d94e881a5b9d985b1a5908185b5bdd5b9d60521b60448201526064016103ab565b6110d183836114e6565b61112c5760405162461bcd60e51b815260206004820152602660248201527f6368617267653a20756e737570706f72746564207265636861726765206f70656044820152653930ba34b7b760d11b60648201526084016103ab565b505050565b6001600160a01b038316156111d9578061114a846107ee565b10156111a65760405162461bcd60e51b815260206004820152602560248201527f6368617267653a206163636f756e742062616c616e6365206973206e6f7420656044820152640dcdeeaced60db1b60648201526084016103ab565b6001600160a01b038316600090815260cb6020526040812080548392906111ce908490612338565b909155506111f19050565b8060c960008282546111eb919061234f565b90915550505b6001600160a01b038216600090815260cb60205260408120805483929061121990849061234f565b9091555050505050565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b6033546001600160a01b031633146105c45760405162461bcd60e51b81526004016103ab90612203565b6000611285611223565b905061129084611969565b60008351118061129d5750815b156112ae576112ac8484611a0e565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff166113bf57805460ff191660011781556040516001600160a01b038316602482015261132d90869060440160408051601f198184030181529190526020810180516001600160e01b0316631b2ce7f360e11b179052611a0e565b50805460ff1916815561133e611223565b6001600160a01b0316826001600160a01b0316146113b65760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201526e75727468657220757067726164657360881b60648201526084016103ab565b6113bf85611af0565b5050505050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff1680611431575060005460ff16155b61144d5760405162461bcd60e51b81526004016103ab90612238565b600054610100900460ff1615801561146f576000805461ffff19166101011790555b611477611b30565b6108d1611b9a565b600054610100900460ff1680611498575060005460ff16155b6114b45760405162461bcd60e51b81526004016103ab90612238565b600054610100900460ff161580156114d6576000805461ffff19166101011790555b6114de611b30565b6108d1611b30565b60006115266040805160e0810182526060808252602082015290810160008152606060208201526040016000815260200160008152602001606081525090565b60cc5460405163fbcbc0f160e01b81526001600160a01b0386811660048301529091169063fbcbc0f190602401600060405180830381865afa158015611570573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261159891908101906123fa565b5092935090918560408101606082016080830160a084018560018111156115c1576115c16122c5565b60018111156115d2576115d26122c5565b90528560018111156115e6576115e66122c5565b60018111156115f7576115f76122c5565b905285905285600281111561160e5761160e6122c5565b600281111561161f5761161f6122c5565b9052949094525060019250611632915050565b81608001516001811115611648576116486122c5565b14801561166a575060018160a001516001811115611668576116686122c5565b145b6116b65760405162461bcd60e51b815260206004820152601860248201527f6368617267653a206066726f6d602069732066726f7a656e000000000000000060448201526064016103ab565b6002816040015160028111156116ce576116ce6122c5565b0361171b5760405162461bcd60e51b815260206004820152601e60248201527f6368617267653a206e6f207265636861726765207065726d697373696f6e000060448201526064016103ab565b6117596040805160e0810182526060808252602082015290810160008152606060208201526040016000815260200160008152602001606081525090565b60cc5460405163fbcbc0f160e01b81526001600160a01b0386811660048301529091169063fbcbc0f190602401600060405180830381865afa1580156117a3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526117cb91908101906123fa565b5092935090918560408101606082016080830160a084018560018111156117f4576117f46122c5565b6001811115611805576118056122c5565b9052856001811115611819576118196122c5565b600181111561182a5761182a6122c5565b9052859052856002811115611841576118416122c5565b6002811115611852576118526122c5565b9052949094525060019250611865915050565b8160800151600181111561187b5761187b6122c5565b14801561189d575060018160a00151600181111561189b5761189b6122c5565b145b6118e25760405162461bcd60e51b815260206004820152601660248201527531b430b933b29d10303a37b01034b990333937bd32b760511b60448201526064016103ab565b6000826040015160028111156118fa576118fa6122c5565b148061191157506060810151825161191191611bfa565b8061196057506060808201519083015161192a91611bfa565b801561193e57508051825161193e91611bfa565b8015611960575060028160400151600281111561195d5761195d6122c5565b14155b95945050505050565b803b6119cd5760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084016103ab565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b611a6d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b60648201526084016103ab565b600080846001600160a01b031684604051611a8891906124d9565b600060405180830381855af49150503d8060008114611ac3576040519150601f19603f3d011682016040523d82523d6000602084013e611ac8565b606091505b5091509150611960828260405180606001604052806027815260200161252960279139611c9b565b611af981611969565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b600054610100900460ff1680611b49575060005460ff16155b611b655760405162461bcd60e51b81526004016103ab90612238565b600054610100900460ff161580156108d1576000805461ffff191661010117905580156105c4576000805461ff001916905550565b600054610100900460ff1680611bb3575060005460ff16155b611bcf5760405162461bcd60e51b81526004016103ab90612238565b600054610100900460ff16158015611bf1576000805461ffff19166101011790555b6108d1336113c6565b805182516000918491849114611c1557600092505050610b84565b815160005b81811015611c8e57828181518110611c3457611c346121be565b602001015160f81c60f81b6001600160f81b031916848281518110611c5b57611c5b6121be565b01602001516001600160f81b03191614611c7c576000945050505050610b84565b80611c86816121ea565b915050611c1a565b5060019695505050505050565b60608315611caa575081611cd4565b825115611cba5782518084602001fd5b8160405162461bcd60e51b81526004016103ab91906124f5565b9392505050565b80356001600160a01b0381168114611cf257600080fd5b919050565b600060208284031215611d0957600080fd5b611cd482611cdb565b80356001600160e01b031981168114611cf257600080fd5b60008060408385031215611d3d57600080fd5b611d4683611cdb565b9150611d5460208401611d12565b90509250929050565b60008060408385031215611d7057600080fd5b611d7983611cdb565b946020939093013593505050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715611dc657611dc6611d87565b604052919050565b600067ffffffffffffffff821115611de857611de8611d87565b5060051b60200190565b600082601f830112611e0357600080fd5b81356020611e18611e1383611dce565b611d9d565b82815260059290921b84018101918181019086841115611e3757600080fd5b8286015b84811015611e5957611e4c81611cdb565b8352918301918301611e3b565b509695505050505050565b600060208284031215611e7657600080fd5b813567ffffffffffffffff811115611e8d57600080fd5b611e9984828501611df2565b949350505050565b6020808252825182820181905260009190848201906040850190845b81811015611ed957835183529284019291840191600101611ebd565b50909695505050505050565b600067ffffffffffffffff821115611eff57611eff611d87565b50601f01601f191660200190565b60008060408385031215611f2057600080fd5b611f2983611cdb565b9150602083013567ffffffffffffffff811115611f4557600080fd5b8301601f81018513611f5657600080fd5b8035611f64611e1382611ee5565b818152866020838501011115611f7957600080fd5b816020840160208301376000602083830101528093505050509250929050565b600080600060608486031215611fae57600080fd5b611fb784611cdb565b9250611fc560208501611d12565b9150604084013563ffffffff81168114611fde57600080fd5b809150509250925092565b600060208284031215611ffb57600080fd5b5035919050565b60008060006060848603121561201757600080fd5b61202084611cdb565b925061202e60208501611d12565b9150604084013590509250925092565b6000806040838503121561205157600080fd5b823567ffffffffffffffff8082111561206957600080fd5b61207586838701611df2565b935060209150818501358181111561208c57600080fd5b85019050601f8101861361209f57600080fd5b80356120ad611e1382611dce565b81815260059190911b820183019083810190888311156120cc57600080fd5b928401925b828410156120ea578335825292840192908401906120d1565b80955050505050509250929050565b6020808252601390820152726368617267653a7a65726f206164647265737360681b604082015260600190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b19195b1959d85d1958d85b1b60a21b606082015260800190565b6020808252602c908201527f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060408201526b6163746976652070726f787960a01b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016121fc576121fc6121d4565b5060010190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b815160009082906020808601845b838110156122b95781516001600160a01b031685529382019390820190600101612294565b50929695505050505050565b634e487b7160e01b600052602160045260246000fd5b6001600160a01b0383168152604081016003831061230957634e487b7160e01b600052602160045260246000fd5b8260208301529392505050565b60006020828403121561232857600080fd5b81518015158114611cd457600080fd5b60008282101561234a5761234a6121d4565b500390565b60008219821115612362576123626121d4565b500190565b60005b8381101561238257818101518382015260200161236a565b83811115612391576000848401525b50505050565b600082601f8301126123a857600080fd5b81516123b6611e1382611ee5565b8181528460208386010111156123cb57600080fd5b611e99826020830160208701612367565b805160038110611cf257600080fd5b805160028110611cf257600080fd5b600080600080600080600060e0888a03121561241557600080fd5b875167ffffffffffffffff8082111561242d57600080fd5b6124398b838c01612397565b985060208a015191508082111561244f57600080fd5b61245b8b838c01612397565b975061246960408b016123dc565b965060608a015191508082111561247f57600080fd5b61248b8b838c01612397565b955061249960808b016123eb565b94506124a760a08b016123eb565b935060c08a01519150808211156124bd57600080fd5b506124ca8a828b01612397565b91505092959891949750929550565b600082516124eb818460208701612367565b9190910192915050565b6020815260008251806020840152612514816040850160208701612367565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220da29516e543dbce36780cc998c456dedf0ea0015c76338961e891cf9d726474664736f6c634300080d0033",
}

// ChargeABI is the input ABI used to generate the binding from.
// Deprecated: Use ChargeMetaData.ABI instead.
var ChargeABI = ChargeMetaData.ABI

// ChargeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ChargeMetaData.Bin instead.
var ChargeBin = ChargeMetaData.Bin

// DeployCharge deploys a new Ethereum contract, binding an instance of Charge to it.
func DeployCharge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Charge, error) {
	parsed, err := ChargeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ChargeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Charge{ChargeCaller: ChargeCaller{contract: contract}, ChargeTransactor: ChargeTransactor{contract: contract}, ChargeFilterer: ChargeFilterer{contract: contract}}, nil
}

// Charge is an auto generated Go binding around an Ethereum contract.
type Charge struct {
	ChargeCaller     // Read-only binding to the contract
	ChargeTransactor // Write-only binding to the contract
	ChargeFilterer   // Log filterer for contract events
}

// ChargeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChargeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChargeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChargeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChargeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChargeSession struct {
	Contract     *Charge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChargeCallerSession struct {
	Contract *ChargeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ChargeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChargeTransactorSession struct {
	Contract     *ChargeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChargeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChargeRaw struct {
	Contract *Charge // Generic contract binding to access the raw methods on
}

// ChargeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChargeCallerRaw struct {
	Contract *ChargeCaller // Generic read-only contract binding to access the raw methods on
}

// ChargeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChargeTransactorRaw struct {
	Contract *ChargeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharge creates a new instance of Charge, bound to a specific deployed contract.
func NewCharge(address common.Address, backend bind.ContractBackend) (*Charge, error) {
	contract, err := bindCharge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Charge{ChargeCaller: ChargeCaller{contract: contract}, ChargeTransactor: ChargeTransactor{contract: contract}, ChargeFilterer: ChargeFilterer{contract: contract}}, nil
}

// NewChargeCaller creates a new read-only instance of Charge, bound to a specific deployed contract.
func NewChargeCaller(address common.Address, caller bind.ContractCaller) (*ChargeCaller, error) {
	contract, err := bindCharge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeCaller{contract: contract}, nil
}

// NewChargeTransactor creates a new write-only instance of Charge, bound to a specific deployed contract.
func NewChargeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChargeTransactor, error) {
	contract, err := bindCharge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChargeTransactor{contract: contract}, nil
}

// NewChargeFilterer creates a new log filterer instance of Charge, bound to a specific deployed contract.
func NewChargeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChargeFilterer, error) {
	contract, err := bindCharge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChargeFilterer{contract: contract}, nil
}

// bindCharge binds a generic wrapper to an already deployed contract.
func bindCharge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChargeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charge *ChargeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charge.Contract.ChargeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charge *ChargeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.Contract.ChargeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charge *ChargeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charge.Contract.ChargeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charge *ChargeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charge *ChargeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charge *ChargeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charge.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeCaller) BalanceOf(opts *bind.CallOpts, accAddr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "balanceOf", accAddr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _Charge.Contract.BalanceOf(&_Charge.CallOpts, accAddr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accAddr) view returns(uint256)
func (_Charge *ChargeCallerSession) BalanceOf(accAddr common.Address) (*big.Int, error) {
	return _Charge.Contract.BalanceOf(&_Charge.CallOpts, accAddr)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x458c738e.
//
// Solidity: function balanceOfBatch(address[] accAddrs) view returns(uint256[])
func (_Charge *ChargeCaller) BalanceOfBatch(opts *bind.CallOpts, accAddrs []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "balanceOfBatch", accAddrs)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x458c738e.
//
// Solidity: function balanceOfBatch(address[] accAddrs) view returns(uint256[])
func (_Charge *ChargeSession) BalanceOfBatch(accAddrs []common.Address) ([]*big.Int, error) {
	return _Charge.Contract.BalanceOfBatch(&_Charge.CallOpts, accAddrs)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x458c738e.
//
// Solidity: function balanceOfBatch(address[] accAddrs) view returns(uint256[])
func (_Charge *ChargeCallerSession) BalanceOfBatch(accAddrs []common.Address) ([]*big.Int, error) {
	return _Charge.Contract.BalanceOfBatch(&_Charge.CallOpts, accAddrs)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeSession) Owner() (common.Address, error) {
	return _Charge.Contract.Owner(&_Charge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Charge *ChargeCallerSession) Owner() (common.Address, error) {
	return _Charge.Contract.Owner(&_Charge.CallOpts)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeCaller) QueryFee(opts *bind.CallOpts, ddcAddr common.Address, sig [4]byte) (uint32, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "queryFee", ddcAddr, sig)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _Charge.Contract.QueryFee(&_Charge.CallOpts, ddcAddr, sig)
}

// QueryFee is a free data retrieval call binding the contract method 0xc9c45a0f.
//
// Solidity: function queryFee(address ddcAddr, bytes4 sig) view returns(uint32)
func (_Charge *ChargeCallerSession) QueryFee(ddcAddr common.Address, sig [4]byte) (uint32, error) {
	return _Charge.Contract.QueryFee(&_Charge.CallOpts, ddcAddr, sig)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Charge.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeSession) TotalSupply() (*big.Int, error) {
	return _Charge.Contract.TotalSupply(&_Charge.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Charge *ChargeCallerSession) TotalSupply() (*big.Int, error) {
	return _Charge.Contract.TotalSupply(&_Charge.CallOpts)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeTransactor) DelDDC(opts *bind.TransactOpts, ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "delDDC", ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.Contract.DelDDC(&_Charge.TransactOpts, ddcAddr)
}

// DelDDC is a paid mutator transaction binding the contract method 0x000b8f70.
//
// Solidity: function delDDC(address ddcAddr) returns()
func (_Charge *ChargeTransactorSession) DelDDC(ddcAddr common.Address) (*types.Transaction, error) {
	return _Charge.Contract.DelDDC(&_Charge.TransactOpts, ddcAddr)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeTransactor) DelFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "delFee", ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.Contract.DelFee(&_Charge.TransactOpts, ddcAddr, sig)
}

// DelFee is a paid mutator transaction binding the contract method 0x093f28e0.
//
// Solidity: function delFee(address ddcAddr, bytes4 sig) returns()
func (_Charge *ChargeTransactorSession) DelFee(ddcAddr common.Address, sig [4]byte) (*types.Transaction, error) {
	return _Charge.Contract.DelFee(&_Charge.TransactOpts, ddcAddr, sig)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeSession) Initialize() (*types.Transaction, error) {
	return _Charge.Contract.Initialize(&_Charge.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Charge *ChargeTransactorSession) Initialize() (*types.Transaction, error) {
	return _Charge.Contract.Initialize(&_Charge.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeTransactor) Pay(opts *bind.TransactOpts, payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "pay", payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Pay(&_Charge.TransactOpts, payer, sig, ddcId)
}

// Pay is a paid mutator transaction binding the contract method 0xef18e3c9.
//
// Solidity: function pay(address payer, bytes4 sig, uint256 ddcId) returns()
func (_Charge *ChargeTransactorSession) Pay(payer common.Address, sig [4]byte, ddcId *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Pay(&_Charge.TransactOpts, payer, sig, ddcId)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeTransactor) Recharge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "recharge", to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Recharge(&_Charge.TransactOpts, to, amount)
}

// Recharge is a paid mutator transaction binding the contract method 0x36351c7c.
//
// Solidity: function recharge(address to, uint256 amount) returns()
func (_Charge *ChargeTransactorSession) Recharge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Recharge(&_Charge.TransactOpts, to, amount)
}

// RechargeBatch is a paid mutator transaction binding the contract method 0xf1e8b168.
//
// Solidity: function rechargeBatch(address[] toList, uint256[] amounts) returns()
func (_Charge *ChargeTransactor) RechargeBatch(opts *bind.TransactOpts, toList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "rechargeBatch", toList, amounts)
}

// RechargeBatch is a paid mutator transaction binding the contract method 0xf1e8b168.
//
// Solidity: function rechargeBatch(address[] toList, uint256[] amounts) returns()
func (_Charge *ChargeSession) RechargeBatch(toList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Charge.Contract.RechargeBatch(&_Charge.TransactOpts, toList, amounts)
}

// RechargeBatch is a paid mutator transaction binding the contract method 0xf1e8b168.
//
// Solidity: function rechargeBatch(address[] toList, uint256[] amounts) returns()
func (_Charge *ChargeTransactorSession) RechargeBatch(toList []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Charge.Contract.RechargeBatch(&_Charge.TransactOpts, toList, amounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Charge.Contract.RenounceOwnership(&_Charge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Charge *ChargeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Charge.Contract.RenounceOwnership(&_Charge.TransactOpts)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeTransactor) SelfRecharge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "selfRecharge", amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.SelfRecharge(&_Charge.TransactOpts, amount)
}

// SelfRecharge is a paid mutator transaction binding the contract method 0xd213fe45.
//
// Solidity: function selfRecharge(uint256 amount) returns()
func (_Charge *ChargeTransactorSession) SelfRecharge(amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.SelfRecharge(&_Charge.TransactOpts, amount)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeTransactor) SetAuthorityProxyAddress(opts *bind.TransactOpts, authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "setAuthorityProxyAddress", authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.Contract.SetAuthorityProxyAddress(&_Charge.TransactOpts, authorityProxyAddress)
}

// SetAuthorityProxyAddress is a paid mutator transaction binding the contract method 0xc5837d82.
//
// Solidity: function setAuthorityProxyAddress(address authorityProxyAddress) returns()
func (_Charge *ChargeTransactorSession) SetAuthorityProxyAddress(authorityProxyAddress common.Address) (*types.Transaction, error) {
	return _Charge.Contract.SetAuthorityProxyAddress(&_Charge.TransactOpts, authorityProxyAddress)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeTransactor) SetFee(opts *bind.TransactOpts, ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "setFee", ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.Contract.SetFee(&_Charge.TransactOpts, ddcAddr, sig, amount)
}

// SetFee is a paid mutator transaction binding the contract method 0x63569189.
//
// Solidity: function setFee(address ddcAddr, bytes4 sig, uint32 amount) returns()
func (_Charge *ChargeTransactorSession) SetFee(ddcAddr common.Address, sig [4]byte, amount uint32) (*types.Transaction, error) {
	return _Charge.Contract.SetFee(&_Charge.TransactOpts, ddcAddr, sig, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeTransactor) Settlement(opts *bind.TransactOpts, ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "settlement", ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Settlement(&_Charge.TransactOpts, ddcAddr, amount)
}

// Settlement is a paid mutator transaction binding the contract method 0x8a27a80d.
//
// Solidity: function settlement(address ddcAddr, uint256 amount) returns()
func (_Charge *ChargeTransactorSession) Settlement(ddcAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Charge.Contract.Settlement(&_Charge.TransactOpts, ddcAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Charge.Contract.TransferOwnership(&_Charge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Charge *ChargeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Charge.Contract.TransferOwnership(&_Charge.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeTo(&_Charge.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Charge *ChargeTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeTo(&_Charge.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeToAndCall(&_Charge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Charge *ChargeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Charge.Contract.UpgradeToAndCall(&_Charge.TransactOpts, newImplementation, data)
}

// ChargeAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Charge contract.
type ChargeAdminChangedIterator struct {
	Event *ChargeAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeAdminChanged represents a AdminChanged event raised by the Charge contract.
type ChargeAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ChargeAdminChangedIterator, error) {

	logs, sub, err := _Charge.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ChargeAdminChangedIterator{contract: _Charge.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ChargeAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Charge.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeAdminChanged)
				if err := _Charge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Charge *ChargeFilterer) ParseAdminChanged(log types.Log) (*ChargeAdminChanged, error) {
	event := new(ChargeAdminChanged)
	if err := _Charge.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Charge contract.
type ChargeBeaconUpgradedIterator struct {
	Event *ChargeBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeBeaconUpgraded represents a BeaconUpgraded event raised by the Charge contract.
type ChargeBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ChargeBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ChargeBeaconUpgradedIterator{contract: _Charge.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeBeaconUpgraded)
				if err := _Charge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Charge *ChargeFilterer) ParseBeaconUpgraded(log types.Log) (*ChargeBeaconUpgraded, error) {
	event := new(ChargeBeaconUpgraded)
	if err := _Charge.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeDelDDCIterator is returned from FilterDelDDC and is used to iterate over the raw logs and unpacked data for DelDDC events raised by the Charge contract.
type ChargeDelDDCIterator struct {
	Event *ChargeDelDDC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeDelDDCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeDelDDC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeDelDDC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeDelDDCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeDelDDCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeDelDDC represents a DelDDC event raised by the Charge contract.
type ChargeDelDDC struct {
	DdcAddr common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelDDC is a free log retrieval operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) FilterDelDDC(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeDelDDCIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeDelDDCIterator{contract: _Charge.contract, event: "DelDDC", logs: logs, sub: sub}, nil
}

// WatchDelDDC is a free log subscription operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) WatchDelDDC(opts *bind.WatchOpts, sink chan<- *ChargeDelDDC, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "DelDDC", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeDelDDC)
				if err := _Charge.contract.UnpackLog(event, "DelDDC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelDDC is a log parse operation binding the contract event 0x0ba05d508af447342f624920545278b6e2d2320ee40cb9eff56b89d21e1b25e8.
//
// Solidity: event DelDDC(address indexed ddcAddr)
func (_Charge *ChargeFilterer) ParseDelDDC(log types.Log) (*ChargeDelDDC, error) {
	event := new(ChargeDelDDC)
	if err := _Charge.contract.UnpackLog(event, "DelDDC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeDelFeeIterator is returned from FilterDelFee and is used to iterate over the raw logs and unpacked data for DelFee events raised by the Charge contract.
type ChargeDelFeeIterator struct {
	Event *ChargeDelFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeDelFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeDelFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeDelFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeDelFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeDelFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeDelFee represents a DelFee event raised by the Charge contract.
type ChargeDelFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDelFee is a free log retrieval operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) FilterDelFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeDelFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeDelFeeIterator{contract: _Charge.contract, event: "DelFee", logs: logs, sub: sub}, nil
}

// WatchDelFee is a free log subscription operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) WatchDelFee(opts *bind.WatchOpts, sink chan<- *ChargeDelFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "DelFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeDelFee)
				if err := _Charge.contract.UnpackLog(event, "DelFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelFee is a log parse operation binding the contract event 0x2f93e067617701594eddb2443d90441c5bb959df555ae8e4d150f0a8bf8b006d.
//
// Solidity: event DelFee(address indexed ddcAddr, bytes4 sig)
func (_Charge *ChargeFilterer) ParseDelFee(log types.Log) (*ChargeDelFee, error) {
	event := new(ChargeDelFee)
	if err := _Charge.contract.UnpackLog(event, "DelFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Charge contract.
type ChargeOwnershipTransferredIterator struct {
	Event *ChargeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeOwnershipTransferred represents a OwnershipTransferred event raised by the Charge contract.
type ChargeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChargeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChargeOwnershipTransferredIterator{contract: _Charge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChargeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeOwnershipTransferred)
				if err := _Charge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Charge *ChargeFilterer) ParseOwnershipTransferred(log types.Log) (*ChargeOwnershipTransferred, error) {
	event := new(ChargeOwnershipTransferred)
	if err := _Charge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargePayIterator is returned from FilterPay and is used to iterate over the raw logs and unpacked data for Pay events raised by the Charge contract.
type ChargePayIterator struct {
	Event *ChargePay // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargePayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargePay)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargePay)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargePayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargePayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargePay represents a Pay event raised by the Charge contract.
type ChargePay struct {
	Payer  common.Address
	Payee  common.Address
	Sig    [4]byte
	Amount uint32
	DdcId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPay is a free log retrieval operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) FilterPay(opts *bind.FilterOpts, payer []common.Address, payee []common.Address) (*ChargePayIterator, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return &ChargePayIterator{contract: _Charge.contract, event: "Pay", logs: logs, sub: sub}, nil
}

// WatchPay is a free log subscription operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) WatchPay(opts *bind.WatchOpts, sink chan<- *ChargePay, payer []common.Address, payee []common.Address) (event.Subscription, error) {

	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var payeeRule []interface{}
	for _, payeeItem := range payee {
		payeeRule = append(payeeRule, payeeItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Pay", payerRule, payeeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargePay)
				if err := _Charge.contract.UnpackLog(event, "Pay", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePay is a log parse operation binding the contract event 0x750e56f33a72767cd99db8943f4d04ca416c55fb783003107a869f5d5062dbab.
//
// Solidity: event Pay(address indexed payer, address indexed payee, bytes4 sig, uint32 amount, uint256 ddcId)
func (_Charge *ChargeFilterer) ParsePay(log types.Log) (*ChargePay, error) {
	event := new(ChargePay)
	if err := _Charge.contract.UnpackLog(event, "Pay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeRechargeIterator is returned from FilterRecharge and is used to iterate over the raw logs and unpacked data for Recharge events raised by the Charge contract.
type ChargeRechargeIterator struct {
	Event *ChargeRecharge // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeRechargeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeRecharge)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeRecharge)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeRechargeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeRechargeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeRecharge represents a Recharge event raised by the Charge contract.
type ChargeRecharge struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecharge is a free log retrieval operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) FilterRecharge(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ChargeRechargeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ChargeRechargeIterator{contract: _Charge.contract, event: "Recharge", logs: logs, sub: sub}, nil
}

// WatchRecharge is a free log subscription operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) WatchRecharge(opts *bind.WatchOpts, sink chan<- *ChargeRecharge, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Recharge", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeRecharge)
				if err := _Charge.contract.UnpackLog(event, "Recharge", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRecharge is a log parse operation binding the contract event 0x4ade20d82044693c0d3331ba1d2a482d103734f274166989491c8d30d3f2ecb1.
//
// Solidity: event Recharge(address indexed from, address indexed to, uint256 amount)
func (_Charge *ChargeFilterer) ParseRecharge(log types.Log) (*ChargeRecharge, error) {
	event := new(ChargeRecharge)
	if err := _Charge.contract.UnpackLog(event, "Recharge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeRechargeBatchIterator is returned from FilterRechargeBatch and is used to iterate over the raw logs and unpacked data for RechargeBatch events raised by the Charge contract.
type ChargeRechargeBatchIterator struct {
	Event *ChargeRechargeBatch // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeRechargeBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeRechargeBatch)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeRechargeBatch)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeRechargeBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeRechargeBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeRechargeBatch represents a RechargeBatch event raised by the Charge contract.
type ChargeRechargeBatch struct {
	From    common.Address
	ToList  []common.Address
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRechargeBatch is a free log retrieval operation binding the contract event 0x744c96b20cd1add6104a750930b0e49658e6c5542ed4e9c9a90ae33359ceb8f9.
//
// Solidity: event RechargeBatch(address indexed from, address[] indexed toList, uint256[] amounts)
func (_Charge *ChargeFilterer) FilterRechargeBatch(opts *bind.FilterOpts, from []common.Address, toList [][]common.Address) (*ChargeRechargeBatchIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toListRule []interface{}
	for _, toListItem := range toList {
		toListRule = append(toListRule, toListItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "RechargeBatch", fromRule, toListRule)
	if err != nil {
		return nil, err
	}
	return &ChargeRechargeBatchIterator{contract: _Charge.contract, event: "RechargeBatch", logs: logs, sub: sub}, nil
}

// WatchRechargeBatch is a free log subscription operation binding the contract event 0x744c96b20cd1add6104a750930b0e49658e6c5542ed4e9c9a90ae33359ceb8f9.
//
// Solidity: event RechargeBatch(address indexed from, address[] indexed toList, uint256[] amounts)
func (_Charge *ChargeFilterer) WatchRechargeBatch(opts *bind.WatchOpts, sink chan<- *ChargeRechargeBatch, from []common.Address, toList [][]common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toListRule []interface{}
	for _, toListItem := range toList {
		toListRule = append(toListRule, toListItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "RechargeBatch", fromRule, toListRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeRechargeBatch)
				if err := _Charge.contract.UnpackLog(event, "RechargeBatch", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRechargeBatch is a log parse operation binding the contract event 0x744c96b20cd1add6104a750930b0e49658e6c5542ed4e9c9a90ae33359ceb8f9.
//
// Solidity: event RechargeBatch(address indexed from, address[] indexed toList, uint256[] amounts)
func (_Charge *ChargeFilterer) ParseRechargeBatch(log types.Log) (*ChargeRechargeBatch, error) {
	event := new(ChargeRechargeBatch)
	if err := _Charge.contract.UnpackLog(event, "RechargeBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeSetFeeIterator is returned from FilterSetFee and is used to iterate over the raw logs and unpacked data for SetFee events raised by the Charge contract.
type ChargeSetFeeIterator struct {
	Event *ChargeSetFee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeSetFeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeSetFee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeSetFee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeSetFeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeSetFeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeSetFee represents a SetFee event raised by the Charge contract.
type ChargeSetFee struct {
	DdcAddr common.Address
	Sig     [4]byte
	Amount  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetFee is a free log retrieval operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) FilterSetFee(opts *bind.FilterOpts, ddcAddr []common.Address) (*ChargeSetFeeIterator, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeSetFeeIterator{contract: _Charge.contract, event: "SetFee", logs: logs, sub: sub}, nil
}

// WatchSetFee is a free log subscription operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) WatchSetFee(opts *bind.WatchOpts, sink chan<- *ChargeSetFee, ddcAddr []common.Address) (event.Subscription, error) {

	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "SetFee", ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeSetFee)
				if err := _Charge.contract.UnpackLog(event, "SetFee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSetFee is a log parse operation binding the contract event 0x929dc21675623ba3d42ef9e085962b7d88bf57ba159fe8f0a86d6785195e2acc.
//
// Solidity: event SetFee(address indexed ddcAddr, bytes4 sig, uint32 amount)
func (_Charge *ChargeFilterer) ParseSetFee(log types.Log) (*ChargeSetFee, error) {
	event := new(ChargeSetFee)
	if err := _Charge.contract.UnpackLog(event, "SetFee", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the Charge contract.
type ChargeSettlementIterator struct {
	Event *ChargeSettlement // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeSettlement)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeSettlement)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeSettlement represents a Settlement event raised by the Charge contract.
type ChargeSettlement struct {
	AccAddr common.Address
	DdcAddr common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) FilterSettlement(opts *bind.FilterOpts, accAddr []common.Address, ddcAddr []common.Address) (*ChargeSettlementIterator, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return &ChargeSettlementIterator{contract: _Charge.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *ChargeSettlement, accAddr []common.Address, ddcAddr []common.Address) (event.Subscription, error) {

	var accAddrRule []interface{}
	for _, accAddrItem := range accAddr {
		accAddrRule = append(accAddrRule, accAddrItem)
	}
	var ddcAddrRule []interface{}
	for _, ddcAddrItem := range ddcAddr {
		ddcAddrRule = append(ddcAddrRule, ddcAddrItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Settlement", accAddrRule, ddcAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeSettlement)
				if err := _Charge.contract.UnpackLog(event, "Settlement", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSettlement is a log parse operation binding the contract event 0xca2ce982d63c71a419517d389df253be4b0d6f4da85fe1491e49608b139ee0be.
//
// Solidity: event Settlement(address indexed accAddr, address indexed ddcAddr, uint256 amount)
func (_Charge *ChargeFilterer) ParseSettlement(log types.Log) (*ChargeSettlement, error) {
	event := new(ChargeSettlement)
	if err := _Charge.contract.UnpackLog(event, "Settlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChargeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Charge contract.
type ChargeUpgradedIterator struct {
	Event *ChargeUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ChargeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChargeUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ChargeUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ChargeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChargeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChargeUpgraded represents a Upgraded event raised by the Charge contract.
type ChargeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ChargeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Charge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ChargeUpgradedIterator{contract: _Charge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ChargeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Charge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChargeUpgraded)
				if err := _Charge.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Charge *ChargeFilterer) ParseUpgraded(log types.Log) (*ChargeUpgraded, error) {
	event := new(ChargeUpgraded)
	if err := _Charge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
