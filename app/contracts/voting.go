// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

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

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"akhiriPemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"detailKandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"detailPemilih\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_Admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminEmail\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AkhiriPemilihan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_MulaiPemilihan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_Organisasi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_PemilihanDeskripsi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_PemilihanTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_TotalKandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_TotalPemilih\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mulaiPemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adminNama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_adminEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_adminTitle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electionDeskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electionTitle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_organizationTitle\",\"type\":\"string\"}],\"name\":\"post_AdminCreatePemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address_public\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"}],\"name\":\"post_AdminPemilihBaru\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_deskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_slogan\",\"type\":\"string\"}],\"name\":\"post_AdminTambahKandidat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_verifedStatus\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"post_AdminVerifikasiPemilih\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"}],\"name\":\"post_RegisterPemilih\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"post_votePemilih\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317815560018190556002556003805461ffff19169055611378806100456000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c806376217974116100c3578063a0b84ba91161007c578063a0b84ba914610293578063a29f93da1461029b578063a96a2574146102ae578063da58c7d9146102b9578063f28d48f6146102cc578063f851a440146102df57600080fd5b8063762179741461024e5780637bfe4204146102615780637d312d861461026957806385bb3b8b146102715780638cdf64da14610279578063a05121881461028157600080fd5b80633b208753116101155780633b208753146101b957806358bc8c3e146101c15780635a27298f146101dd5780635c9fc7021461020257806361392bee146102155780637059a5861461022857600080fd5b80630625a9981461015257806308d36f6c1461016757806313ca4d8a1461016f5780631cbe1d121461018d578063268ef0d614610195575b600080fd5b610165610160366004610d64565b6102f2565b005b6101656103c7565b6101776103ee565b6040516101849190610ddd565b60405180910390f35b610177610483565b6101a86101a3366004610df7565b610495565b604051610184959493929190610e10565b61017761065c565b600354610100900460ff165b6040519015158152602001610184565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610184565b610165610210366004610f05565b61066e565b610165610223366004610ffa565b610735565b61023b610236366004611093565b610894565b60405161018497969594939291906110ae565b61016561025c36600461111a565b610a7b565b610165610abd565b610177610ae5565b610177610af7565b610177610b09565b6001545b604051908152602001610184565b600254610285565b6101656102a936600461114b565b610b1a565b60035460ff166101cd565b6101ea6102c7366004610df7565b610c6a565b6101656102da36600461114b565b610c94565b6000546101ea906001600160a01b031681565b6001600160a01b0381166000908152600c6020526040902060040154610100900460ff161561032057600080fd5b6001600160a01b0381166000908152600c602052604090206004015460ff16151560011461034d57600080fd5b60035460ff16151560011461036157600080fd5b600354610100900460ff161561037657600080fd5b60008281526004602081905260408220018054600192906103989084906111d3565b90915550506001600160a01b03166000908152600c60205260409020600401805461ff00191661010017905550565b6000546001600160a01b031633146103de57600080fd5b6003805461ffff19166001179055565b606060056004018054610400906111f9565b80601f016020809104026020016040519081016040528092919081815260200182805461042c906111f9565b80156104795780601f1061044e57610100808354040283529160200191610479565b820191906000526020600020905b81548152906001019060200180831161045c57829003601f168201915b5050505050905090565b606060056000018054610400906111f9565b600460205260009081526040902080546001820180549192916104b7906111f9565b80601f01602080910402602001604051908101604052809291908181526020018280546104e3906111f9565b80156105305780601f1061050557610100808354040283529160200191610530565b820191906000526020600020905b81548152906001019060200180831161051357829003601f168201915b505050505090806002018054610545906111f9565b80601f0160208091040260200160405190810160405280929190818152602001828054610571906111f9565b80156105be5780601f10610593576101008083540402835291602001916105be565b820191906000526020600020905b8154815290600101906020018083116105a157829003601f168201915b5050505050908060030180546105d3906111f9565b80601f01602080910402602001604051908101604052809291908181526020018280546105ff906111f9565b801561064c5780601f106106215761010080835404028352916020019161064c565b820191906000526020600020905b81548152906001019060200180831161062f57829003601f168201915b5050505050908060040154905085565b606060056001018054610400906111f9565b6000546001600160a01b0316331461068557600080fd5b6040805160c08101825287815260208101879052908101859052606081018490526080810183905260a081018290526005806106c18982611282565b50602082015160018201906106d69082611282565b50604082015160028201906106eb9082611282565b50606082015160038201906107009082611282565b50608082015160048201906107159082611282565b5060a0820151600582019061072a9082611282565b505050505050505050565b6040805160e0810182526001600160a01b0386811680835260208084018881528486018890526060850187905260006080860181905260a08601819052600160c08701819052938152600c909252949020835181546001600160a01b0319169316929092178255925191928392908201906107b09082611282565b50604082015160028201906107c59082611282565b50606082015160038201906107da9082611282565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ff00199415159490941661ffff19909216919091179290921792909216179055600b80546001818101835560009283527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db990910180546001600160a01b0389166001600160a01b031990911617905560028054919290916108889084906111d3565b90915550505050505050565b600c60205260009081526040902080546001820180546001600160a01b0390921692916108c0906111f9565b80601f01602080910402602001604051908101604052809291908181526020018280546108ec906111f9565b80156109395780601f1061090e57610100808354040283529160200191610939565b820191906000526020600020905b81548152906001019060200180831161091c57829003601f168201915b50505050509080600201805461094e906111f9565b80601f016020809104026020016040519081016040528092919081815260200182805461097a906111f9565b80156109c75780601f1061099c576101008083540402835291602001916109c7565b820191906000526020600020905b8154815290600101906020018083116109aa57829003601f168201915b5050505050908060030180546109dc906111f9565b80601f0160208091040260200160405190810160405280929190818152602001828054610a08906111f9565b8015610a555780601f10610a2a57610100808354040283529160200191610a55565b820191906000526020600020905b815481529060010190602001808311610a3857829003601f168201915b5050506004909301549192505060ff808216916101008104821691620100009091041687565b6000546001600160a01b03163314610a9257600080fd5b6001600160a01b03166000908152600c60205260409020600401805460ff1916911515919091179055565b6000546001600160a01b03163314610ad457600080fd5b6003805461ffff1916610100179055565b606060056003018054610400906111f9565b606060056002018054610400906111f9565b6060600580018054610400906111f9565b6040805160e0810182523380825260208083018781528385018790526060840186905260006080850181905260a08501819052600160c08601819052938152600c909252939020825181546001600160a01b0319166001600160a01b039091161781559251919283929091820190610b929082611282565b5060408201516002820190610ba79082611282565b5060608201516003820190610bbc9082611282565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ff00199415159490941661ffff19909216919091179290921792909216179055600b80546001818101835560009283527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db990910180546001600160a01b031916331790556002805491929091610c5f9084906111d3565b909155505050505050565b600b8181548110610c7a57600080fd5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b03163314610cab57600080fd5b6040805160a081018252600180548083526020808401888152848601889052606085018790526000608086018190529283526004909152939020825181559251919283929091820190610cfe9082611282565b5060408201516002820190610d139082611282565b5060608201516003820190610d289082611282565b50608082015181600401559050506001806000828254610c5f91906111d3565b80356001600160a01b0381168114610d5f57600080fd5b919050565b60008060408385031215610d7757600080fd5b82359150610d8760208401610d48565b90509250929050565b6000815180845260005b81811015610db657602081850181015186830182015201610d9a565b81811115610dc8576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610df06020830184610d90565b9392505050565b600060208284031215610e0957600080fd5b5035919050565b85815260a060208201526000610e2960a0830187610d90565b8281036040840152610e3b8187610d90565b90508281036060840152610e4f8186610d90565b9150508260808301529695505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e8957600080fd5b813567ffffffffffffffff80821115610ea457610ea4610e62565b604051601f8301601f19908116603f01168101908282118183101715610ecc57610ecc610e62565b81604052838152866020858801011115610ee557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060008060c08789031215610f1e57600080fd5b863567ffffffffffffffff80821115610f3657600080fd5b610f428a838b01610e78565b97506020890135915080821115610f5857600080fd5b610f648a838b01610e78565b96506040890135915080821115610f7a57600080fd5b610f868a838b01610e78565b95506060890135915080821115610f9c57600080fd5b610fa88a838b01610e78565b94506080890135915080821115610fbe57600080fd5b610fca8a838b01610e78565b935060a0890135915080821115610fe057600080fd5b50610fed89828a01610e78565b9150509295509295509295565b6000806000806080858703121561101057600080fd5b61101985610d48565b9350602085013567ffffffffffffffff8082111561103657600080fd5b61104288838901610e78565b9450604087013591508082111561105857600080fd5b61106488838901610e78565b9350606087013591508082111561107a57600080fd5b5061108787828801610e78565b91505092959194509250565b6000602082840312156110a557600080fd5b610df082610d48565b6001600160a01b038816815260e0602082018190526000906110d290830189610d90565b82810360408401526110e48189610d90565b905082810360608401526110f88188610d90565b9515156080840152505091151560a0830152151560c090910152949350505050565b6000806040838503121561112d57600080fd5b8235801515811461113d57600080fd5b9150610d8760208401610d48565b60008060006060848603121561116057600080fd5b833567ffffffffffffffff8082111561117857600080fd5b61118487838801610e78565b9450602086013591508082111561119a57600080fd5b6111a687838801610e78565b935060408601359150808211156111bc57600080fd5b506111c986828701610e78565b9150509250925092565b600082198211156111f457634e487b7160e01b600052601160045260246000fd5b500190565b600181811c9082168061120d57607f821691505b60208210810361122d57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561127d57600081815260208120601f850160051c8101602086101561125a5750805b601f850160051c820191505b8181101561127957828155600101611266565b5050505b505050565b815167ffffffffffffffff81111561129c5761129c610e62565b6112b0816112aa84546111f9565b84611233565b602080601f8311600181146112e557600084156112cd5750858301515b600019600386901b1c1916600185901b178555611279565b600085815260208120601f198616915b82811015611314578886015182559484019460019091019084016112f5565b50858210156113325787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212207f6d4cba2f72f635e686e29c28acd56e7c99598a6defbe6edfdda47f0f94c4b664736f6c634300080f0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Api *ApiCallerSession) Admin() (common.Address, error) {
	return _Api.Contract.Admin(&_Api.CallOpts)
}

// DetailKandidat is a free data retrieval call binding the contract method 0x268ef0d6.
//
// Solidity: function detailKandidat(uint256 ) view returns(uint256 candidateId, string nama, string deskripsi, string slogan, uint256 voteCount)
func (_Api *ApiCaller) DetailKandidat(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CandidateId *big.Int
	Nama        string
	Deskripsi   string
	Slogan      string
	VoteCount   *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "detailKandidat", arg0)

	outstruct := new(struct {
		CandidateId *big.Int
		Nama        string
		Deskripsi   string
		Slogan      string
		VoteCount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CandidateId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Nama = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Deskripsi = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Slogan = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DetailKandidat is a free data retrieval call binding the contract method 0x268ef0d6.
//
// Solidity: function detailKandidat(uint256 ) view returns(uint256 candidateId, string nama, string deskripsi, string slogan, uint256 voteCount)
func (_Api *ApiSession) DetailKandidat(arg0 *big.Int) (struct {
	CandidateId *big.Int
	Nama        string
	Deskripsi   string
	Slogan      string
	VoteCount   *big.Int
}, error) {
	return _Api.Contract.DetailKandidat(&_Api.CallOpts, arg0)
}

// DetailKandidat is a free data retrieval call binding the contract method 0x268ef0d6.
//
// Solidity: function detailKandidat(uint256 ) view returns(uint256 candidateId, string nama, string deskripsi, string slogan, uint256 voteCount)
func (_Api *ApiCallerSession) DetailKandidat(arg0 *big.Int) (struct {
	CandidateId *big.Int
	Nama        string
	Deskripsi   string
	Slogan      string
	VoteCount   *big.Int
}, error) {
	return _Api.Contract.DetailKandidat(&_Api.CallOpts, arg0)
}

// DetailPemilih is a free data retrieval call binding the contract method 0x7059a586.
//
// Solidity: function detailPemilih(address ) view returns(address voterAddress, string name, string home_address, string phone, bool isVerified, bool hasVoted, bool isRegistered)
func (_Api *ApiCaller) DetailPemilih(opts *bind.CallOpts, arg0 common.Address) (struct {
	VoterAddress common.Address
	Name         string
	HomeAddress  string
	Phone        string
	IsVerified   bool
	HasVoted     bool
	IsRegistered bool
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "detailPemilih", arg0)

	outstruct := new(struct {
		VoterAddress common.Address
		Name         string
		HomeAddress  string
		Phone        string
		IsVerified   bool
		HasVoted     bool
		IsRegistered bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VoterAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.HomeAddress = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Phone = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IsVerified = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.HasVoted = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.IsRegistered = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// DetailPemilih is a free data retrieval call binding the contract method 0x7059a586.
//
// Solidity: function detailPemilih(address ) view returns(address voterAddress, string name, string home_address, string phone, bool isVerified, bool hasVoted, bool isRegistered)
func (_Api *ApiSession) DetailPemilih(arg0 common.Address) (struct {
	VoterAddress common.Address
	Name         string
	HomeAddress  string
	Phone        string
	IsVerified   bool
	HasVoted     bool
	IsRegistered bool
}, error) {
	return _Api.Contract.DetailPemilih(&_Api.CallOpts, arg0)
}

// DetailPemilih is a free data retrieval call binding the contract method 0x7059a586.
//
// Solidity: function detailPemilih(address ) view returns(address voterAddress, string name, string home_address, string phone, bool isVerified, bool hasVoted, bool isRegistered)
func (_Api *ApiCallerSession) DetailPemilih(arg0 common.Address) (struct {
	VoterAddress common.Address
	Name         string
	HomeAddress  string
	Phone        string
	IsVerified   bool
	HasVoted     bool
	IsRegistered bool
}, error) {
	return _Api.Contract.DetailPemilih(&_Api.CallOpts, arg0)
}

// GetAdmin is a free data retrieval call binding the contract method 0x5a27298f.
//
// Solidity: function get_Admin() view returns(address)
func (_Api *ApiCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_Admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x5a27298f.
//
// Solidity: function get_Admin() view returns(address)
func (_Api *ApiSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x5a27298f.
//
// Solidity: function get_Admin() view returns(address)
func (_Api *ApiCallerSession) GetAdmin() (common.Address, error) {
	return _Api.Contract.GetAdmin(&_Api.CallOpts)
}

// GetAdminEmail is a free data retrieval call binding the contract method 0x3b208753.
//
// Solidity: function get_AdminEmail() view returns(string)
func (_Api *ApiCaller) GetAdminEmail(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_AdminEmail")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAdminEmail is a free data retrieval call binding the contract method 0x3b208753.
//
// Solidity: function get_AdminEmail() view returns(string)
func (_Api *ApiSession) GetAdminEmail() (string, error) {
	return _Api.Contract.GetAdminEmail(&_Api.CallOpts)
}

// GetAdminEmail is a free data retrieval call binding the contract method 0x3b208753.
//
// Solidity: function get_AdminEmail() view returns(string)
func (_Api *ApiCallerSession) GetAdminEmail() (string, error) {
	return _Api.Contract.GetAdminEmail(&_Api.CallOpts)
}

// GetAdminName is a free data retrieval call binding the contract method 0x1cbe1d12.
//
// Solidity: function get_AdminName() view returns(string)
func (_Api *ApiCaller) GetAdminName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_AdminName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAdminName is a free data retrieval call binding the contract method 0x1cbe1d12.
//
// Solidity: function get_AdminName() view returns(string)
func (_Api *ApiSession) GetAdminName() (string, error) {
	return _Api.Contract.GetAdminName(&_Api.CallOpts)
}

// GetAdminName is a free data retrieval call binding the contract method 0x1cbe1d12.
//
// Solidity: function get_AdminName() view returns(string)
func (_Api *ApiCallerSession) GetAdminName() (string, error) {
	return _Api.Contract.GetAdminName(&_Api.CallOpts)
}

// GetAdminTitle is a free data retrieval call binding the contract method 0x85bb3b8b.
//
// Solidity: function get_AdminTitle() view returns(string)
func (_Api *ApiCaller) GetAdminTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_AdminTitle")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAdminTitle is a free data retrieval call binding the contract method 0x85bb3b8b.
//
// Solidity: function get_AdminTitle() view returns(string)
func (_Api *ApiSession) GetAdminTitle() (string, error) {
	return _Api.Contract.GetAdminTitle(&_Api.CallOpts)
}

// GetAdminTitle is a free data retrieval call binding the contract method 0x85bb3b8b.
//
// Solidity: function get_AdminTitle() view returns(string)
func (_Api *ApiCallerSession) GetAdminTitle() (string, error) {
	return _Api.Contract.GetAdminTitle(&_Api.CallOpts)
}

// GetAkhiriPemilihan is a free data retrieval call binding the contract method 0x58bc8c3e.
//
// Solidity: function get_AkhiriPemilihan() view returns(bool)
func (_Api *ApiCaller) GetAkhiriPemilihan(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_AkhiriPemilihan")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAkhiriPemilihan is a free data retrieval call binding the contract method 0x58bc8c3e.
//
// Solidity: function get_AkhiriPemilihan() view returns(bool)
func (_Api *ApiSession) GetAkhiriPemilihan() (bool, error) {
	return _Api.Contract.GetAkhiriPemilihan(&_Api.CallOpts)
}

// GetAkhiriPemilihan is a free data retrieval call binding the contract method 0x58bc8c3e.
//
// Solidity: function get_AkhiriPemilihan() view returns(bool)
func (_Api *ApiCallerSession) GetAkhiriPemilihan() (bool, error) {
	return _Api.Contract.GetAkhiriPemilihan(&_Api.CallOpts)
}

// GetMulaiPemilihan is a free data retrieval call binding the contract method 0xa96a2574.
//
// Solidity: function get_MulaiPemilihan() view returns(bool)
func (_Api *ApiCaller) GetMulaiPemilihan(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_MulaiPemilihan")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetMulaiPemilihan is a free data retrieval call binding the contract method 0xa96a2574.
//
// Solidity: function get_MulaiPemilihan() view returns(bool)
func (_Api *ApiSession) GetMulaiPemilihan() (bool, error) {
	return _Api.Contract.GetMulaiPemilihan(&_Api.CallOpts)
}

// GetMulaiPemilihan is a free data retrieval call binding the contract method 0xa96a2574.
//
// Solidity: function get_MulaiPemilihan() view returns(bool)
func (_Api *ApiCallerSession) GetMulaiPemilihan() (bool, error) {
	return _Api.Contract.GetMulaiPemilihan(&_Api.CallOpts)
}

// GetOrganisasi is a free data retrieval call binding the contract method 0x8cdf64da.
//
// Solidity: function get_Organisasi() view returns(string)
func (_Api *ApiCaller) GetOrganisasi(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_Organisasi")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetOrganisasi is a free data retrieval call binding the contract method 0x8cdf64da.
//
// Solidity: function get_Organisasi() view returns(string)
func (_Api *ApiSession) GetOrganisasi() (string, error) {
	return _Api.Contract.GetOrganisasi(&_Api.CallOpts)
}

// GetOrganisasi is a free data retrieval call binding the contract method 0x8cdf64da.
//
// Solidity: function get_Organisasi() view returns(string)
func (_Api *ApiCallerSession) GetOrganisasi() (string, error) {
	return _Api.Contract.GetOrganisasi(&_Api.CallOpts)
}

// GetPemilihanDeskripsi is a free data retrieval call binding the contract method 0x7d312d86.
//
// Solidity: function get_PemilihanDeskripsi() view returns(string)
func (_Api *ApiCaller) GetPemilihanDeskripsi(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_PemilihanDeskripsi")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPemilihanDeskripsi is a free data retrieval call binding the contract method 0x7d312d86.
//
// Solidity: function get_PemilihanDeskripsi() view returns(string)
func (_Api *ApiSession) GetPemilihanDeskripsi() (string, error) {
	return _Api.Contract.GetPemilihanDeskripsi(&_Api.CallOpts)
}

// GetPemilihanDeskripsi is a free data retrieval call binding the contract method 0x7d312d86.
//
// Solidity: function get_PemilihanDeskripsi() view returns(string)
func (_Api *ApiCallerSession) GetPemilihanDeskripsi() (string, error) {
	return _Api.Contract.GetPemilihanDeskripsi(&_Api.CallOpts)
}

// GetPemilihanTitle is a free data retrieval call binding the contract method 0x13ca4d8a.
//
// Solidity: function get_PemilihanTitle() view returns(string)
func (_Api *ApiCaller) GetPemilihanTitle(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_PemilihanTitle")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPemilihanTitle is a free data retrieval call binding the contract method 0x13ca4d8a.
//
// Solidity: function get_PemilihanTitle() view returns(string)
func (_Api *ApiSession) GetPemilihanTitle() (string, error) {
	return _Api.Contract.GetPemilihanTitle(&_Api.CallOpts)
}

// GetPemilihanTitle is a free data retrieval call binding the contract method 0x13ca4d8a.
//
// Solidity: function get_PemilihanTitle() view returns(string)
func (_Api *ApiCallerSession) GetPemilihanTitle() (string, error) {
	return _Api.Contract.GetPemilihanTitle(&_Api.CallOpts)
}

// GetTotalKandidat is a free data retrieval call binding the contract method 0xa0512188.
//
// Solidity: function get_TotalKandidat() view returns(uint256)
func (_Api *ApiCaller) GetTotalKandidat(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_TotalKandidat")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalKandidat is a free data retrieval call binding the contract method 0xa0512188.
//
// Solidity: function get_TotalKandidat() view returns(uint256)
func (_Api *ApiSession) GetTotalKandidat() (*big.Int, error) {
	return _Api.Contract.GetTotalKandidat(&_Api.CallOpts)
}

// GetTotalKandidat is a free data retrieval call binding the contract method 0xa0512188.
//
// Solidity: function get_TotalKandidat() view returns(uint256)
func (_Api *ApiCallerSession) GetTotalKandidat() (*big.Int, error) {
	return _Api.Contract.GetTotalKandidat(&_Api.CallOpts)
}

// GetTotalPemilih is a free data retrieval call binding the contract method 0xa0b84ba9.
//
// Solidity: function get_TotalPemilih() view returns(uint256)
func (_Api *ApiCaller) GetTotalPemilih(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_TotalPemilih")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalPemilih is a free data retrieval call binding the contract method 0xa0b84ba9.
//
// Solidity: function get_TotalPemilih() view returns(uint256)
func (_Api *ApiSession) GetTotalPemilih() (*big.Int, error) {
	return _Api.Contract.GetTotalPemilih(&_Api.CallOpts)
}

// GetTotalPemilih is a free data retrieval call binding the contract method 0xa0b84ba9.
//
// Solidity: function get_TotalPemilih() view returns(uint256)
func (_Api *ApiCallerSession) GetTotalPemilih() (*big.Int, error) {
	return _Api.Contract.GetTotalPemilih(&_Api.CallOpts)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Api *ApiCaller) Voters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "voters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Api *ApiSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Voters(&_Api.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xda58c7d9.
//
// Solidity: function voters(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Voters(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Voters(&_Api.CallOpts, arg0)
}

// AkhiriPemilihan is a paid mutator transaction binding the contract method 0x7bfe4204.
//
// Solidity: function akhiriPemilihan() returns()
func (_Api *ApiTransactor) AkhiriPemilihan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "akhiriPemilihan")
}

// AkhiriPemilihan is a paid mutator transaction binding the contract method 0x7bfe4204.
//
// Solidity: function akhiriPemilihan() returns()
func (_Api *ApiSession) AkhiriPemilihan() (*types.Transaction, error) {
	return _Api.Contract.AkhiriPemilihan(&_Api.TransactOpts)
}

// AkhiriPemilihan is a paid mutator transaction binding the contract method 0x7bfe4204.
//
// Solidity: function akhiriPemilihan() returns()
func (_Api *ApiTransactorSession) AkhiriPemilihan() (*types.Transaction, error) {
	return _Api.Contract.AkhiriPemilihan(&_Api.TransactOpts)
}

// MulaiPemilihan is a paid mutator transaction binding the contract method 0x08d36f6c.
//
// Solidity: function mulaiPemilihan() returns()
func (_Api *ApiTransactor) MulaiPemilihan(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "mulaiPemilihan")
}

// MulaiPemilihan is a paid mutator transaction binding the contract method 0x08d36f6c.
//
// Solidity: function mulaiPemilihan() returns()
func (_Api *ApiSession) MulaiPemilihan() (*types.Transaction, error) {
	return _Api.Contract.MulaiPemilihan(&_Api.TransactOpts)
}

// MulaiPemilihan is a paid mutator transaction binding the contract method 0x08d36f6c.
//
// Solidity: function mulaiPemilihan() returns()
func (_Api *ApiTransactorSession) MulaiPemilihan() (*types.Transaction, error) {
	return _Api.Contract.MulaiPemilihan(&_Api.TransactOpts)
}

// PostAdminCreatePemilihan is a paid mutator transaction binding the contract method 0x5c9fc702.
//
// Solidity: function post_AdminCreatePemilihan(string _adminNama, string _adminEmail, string _adminTitle, string _electionDeskripsi, string _electionTitle, string _organizationTitle) returns()
func (_Api *ApiTransactor) PostAdminCreatePemilihan(opts *bind.TransactOpts, _adminNama string, _adminEmail string, _adminTitle string, _electionDeskripsi string, _electionTitle string, _organizationTitle string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_AdminCreatePemilihan", _adminNama, _adminEmail, _adminTitle, _electionDeskripsi, _electionTitle, _organizationTitle)
}

// PostAdminCreatePemilihan is a paid mutator transaction binding the contract method 0x5c9fc702.
//
// Solidity: function post_AdminCreatePemilihan(string _adminNama, string _adminEmail, string _adminTitle, string _electionDeskripsi, string _electionTitle, string _organizationTitle) returns()
func (_Api *ApiSession) PostAdminCreatePemilihan(_adminNama string, _adminEmail string, _adminTitle string, _electionDeskripsi string, _electionTitle string, _organizationTitle string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminCreatePemilihan(&_Api.TransactOpts, _adminNama, _adminEmail, _adminTitle, _electionDeskripsi, _electionTitle, _organizationTitle)
}

// PostAdminCreatePemilihan is a paid mutator transaction binding the contract method 0x5c9fc702.
//
// Solidity: function post_AdminCreatePemilihan(string _adminNama, string _adminEmail, string _adminTitle, string _electionDeskripsi, string _electionTitle, string _organizationTitle) returns()
func (_Api *ApiTransactorSession) PostAdminCreatePemilihan(_adminNama string, _adminEmail string, _adminTitle string, _electionDeskripsi string, _electionTitle string, _organizationTitle string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminCreatePemilihan(&_Api.TransactOpts, _adminNama, _adminEmail, _adminTitle, _electionDeskripsi, _electionTitle, _organizationTitle)
}

// PostAdminPemilihBaru is a paid mutator transaction binding the contract method 0x61392bee.
//
// Solidity: function post_AdminPemilihBaru(address _address_public, string _name, string _home_address, string _phone) returns()
func (_Api *ApiTransactor) PostAdminPemilihBaru(opts *bind.TransactOpts, _address_public common.Address, _name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_AdminPemilihBaru", _address_public, _name, _home_address, _phone)
}

// PostAdminPemilihBaru is a paid mutator transaction binding the contract method 0x61392bee.
//
// Solidity: function post_AdminPemilihBaru(address _address_public, string _name, string _home_address, string _phone) returns()
func (_Api *ApiSession) PostAdminPemilihBaru(_address_public common.Address, _name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminPemilihBaru(&_Api.TransactOpts, _address_public, _name, _home_address, _phone)
}

// PostAdminPemilihBaru is a paid mutator transaction binding the contract method 0x61392bee.
//
// Solidity: function post_AdminPemilihBaru(address _address_public, string _name, string _home_address, string _phone) returns()
func (_Api *ApiTransactorSession) PostAdminPemilihBaru(_address_public common.Address, _name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminPemilihBaru(&_Api.TransactOpts, _address_public, _name, _home_address, _phone)
}

// PostAdminTambahKandidat is a paid mutator transaction binding the contract method 0xf28d48f6.
//
// Solidity: function post_AdminTambahKandidat(string _nama, string _deskripsi, string _slogan) returns()
func (_Api *ApiTransactor) PostAdminTambahKandidat(opts *bind.TransactOpts, _nama string, _deskripsi string, _slogan string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_AdminTambahKandidat", _nama, _deskripsi, _slogan)
}

// PostAdminTambahKandidat is a paid mutator transaction binding the contract method 0xf28d48f6.
//
// Solidity: function post_AdminTambahKandidat(string _nama, string _deskripsi, string _slogan) returns()
func (_Api *ApiSession) PostAdminTambahKandidat(_nama string, _deskripsi string, _slogan string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminTambahKandidat(&_Api.TransactOpts, _nama, _deskripsi, _slogan)
}

// PostAdminTambahKandidat is a paid mutator transaction binding the contract method 0xf28d48f6.
//
// Solidity: function post_AdminTambahKandidat(string _nama, string _deskripsi, string _slogan) returns()
func (_Api *ApiTransactorSession) PostAdminTambahKandidat(_nama string, _deskripsi string, _slogan string) (*types.Transaction, error) {
	return _Api.Contract.PostAdminTambahKandidat(&_Api.TransactOpts, _nama, _deskripsi, _slogan)
}

// PostAdminVerifikasiPemilih is a paid mutator transaction binding the contract method 0x76217974.
//
// Solidity: function post_AdminVerifikasiPemilih(bool _verifedStatus, address voterAddress) returns()
func (_Api *ApiTransactor) PostAdminVerifikasiPemilih(opts *bind.TransactOpts, _verifedStatus bool, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_AdminVerifikasiPemilih", _verifedStatus, voterAddress)
}

// PostAdminVerifikasiPemilih is a paid mutator transaction binding the contract method 0x76217974.
//
// Solidity: function post_AdminVerifikasiPemilih(bool _verifedStatus, address voterAddress) returns()
func (_Api *ApiSession) PostAdminVerifikasiPemilih(_verifedStatus bool, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.PostAdminVerifikasiPemilih(&_Api.TransactOpts, _verifedStatus, voterAddress)
}

// PostAdminVerifikasiPemilih is a paid mutator transaction binding the contract method 0x76217974.
//
// Solidity: function post_AdminVerifikasiPemilih(bool _verifedStatus, address voterAddress) returns()
func (_Api *ApiTransactorSession) PostAdminVerifikasiPemilih(_verifedStatus bool, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.PostAdminVerifikasiPemilih(&_Api.TransactOpts, _verifedStatus, voterAddress)
}

// PostRegisterPemilih is a paid mutator transaction binding the contract method 0xa29f93da.
//
// Solidity: function post_RegisterPemilih(string _name, string _home_address, string _phone) returns()
func (_Api *ApiTransactor) PostRegisterPemilih(opts *bind.TransactOpts, _name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_RegisterPemilih", _name, _home_address, _phone)
}

// PostRegisterPemilih is a paid mutator transaction binding the contract method 0xa29f93da.
//
// Solidity: function post_RegisterPemilih(string _name, string _home_address, string _phone) returns()
func (_Api *ApiSession) PostRegisterPemilih(_name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.Contract.PostRegisterPemilih(&_Api.TransactOpts, _name, _home_address, _phone)
}

// PostRegisterPemilih is a paid mutator transaction binding the contract method 0xa29f93da.
//
// Solidity: function post_RegisterPemilih(string _name, string _home_address, string _phone) returns()
func (_Api *ApiTransactorSession) PostRegisterPemilih(_name string, _home_address string, _phone string) (*types.Transaction, error) {
	return _Api.Contract.PostRegisterPemilih(&_Api.TransactOpts, _name, _home_address, _phone)
}

// PostVotePemilih is a paid mutator transaction binding the contract method 0x0625a998.
//
// Solidity: function post_votePemilih(uint256 candidateId, address voterAddress) returns()
func (_Api *ApiTransactor) PostVotePemilih(opts *bind.TransactOpts, candidateId *big.Int, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_votePemilih", candidateId, voterAddress)
}

// PostVotePemilih is a paid mutator transaction binding the contract method 0x0625a998.
//
// Solidity: function post_votePemilih(uint256 candidateId, address voterAddress) returns()
func (_Api *ApiSession) PostVotePemilih(candidateId *big.Int, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.PostVotePemilih(&_Api.TransactOpts, candidateId, voterAddress)
}

// PostVotePemilih is a paid mutator transaction binding the contract method 0x0625a998.
//
// Solidity: function post_votePemilih(uint256 candidateId, address voterAddress) returns()
func (_Api *ApiTransactorSession) PostVotePemilih(candidateId *big.Int, voterAddress common.Address) (*types.Transaction, error) {
	return _Api.Contract.PostVotePemilih(&_Api.TransactOpts, candidateId, voterAddress)
}
