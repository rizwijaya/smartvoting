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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"akhiriPemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"detailKandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"deskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"slogan\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"detailPemilih\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isVerified\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isRegistered\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_Admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminEmail\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AdminTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_AkhiriPemilihan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_MulaiPemilihan\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_Organisasi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_PemilihanDeskripsi\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_PemilihanTitle\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_TotalKandidat\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_TotalPemilih\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mulaiPemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_adminNama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_adminEmail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_adminTitle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electionDeskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_electionTitle\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_organizationTitle\",\"type\":\"string\"}],\"name\":\"post_AdminCreatePemilihan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address_public\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"}],\"name\":\"post_AdminPemilihBaru\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nama\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_deskripsi\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_slogan\",\"type\":\"string\"}],\"name\":\"post_AdminTambahKandidat\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_verifedStatus\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"voterAddress\",\"type\":\"address\"}],\"name\":\"post_AdminVerifikasiPemilih\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"candidateId\",\"type\":\"uint256\"}],\"name\":\"post_GeneralVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_home_address\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"}],\"name\":\"post_RegisterPemilih\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317815560018190556002556003805461ffff19169055611330806100456000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c80637bfe4204116100c3578063a29f93da1161007c578063a29f93da14610288578063a51563071461029b578063a96a2574146102ae578063da58c7d9146102b9578063f28d48f6146102cc578063f851a440146102df57600080fd5b80637bfe42041461024e5780637d312d861461025657806385bb3b8b1461025e5780638cdf64da14610266578063a05121881461026e578063a0b84ba91461028057600080fd5b806358bc8c3e1161011557806358bc8c3e146101ae5780635a27298f146101ca5780635c9fc702146101ef57806361392bee146102025780637059a58614610215578063762179741461023b57600080fd5b806308d36f6c1461015257806313ca4d8a1461015c5780631cbe1d121461017a578063268ef0d6146101825780633b208753146101a6575b600080fd5b61015a6102f2565b005b610164610319565b6040516101719190610d70565b60405180910390f35b6101646103ae565b610195610190366004610d8a565b6103c0565b604051610171959493929190610da3565b610164610587565b600354610100900460ff165b6040519015158152602001610171565b6000546001600160a01b03165b6040516001600160a01b039091168152602001610171565b61015a6101fd366004610e98565b610599565b61015a610210366004610fa9565b610660565b610228610223366004611042565b6107b4565b604051610171979695949392919061105d565b61015a6102493660046110c9565b61099b565b61015a6109dd565b610164610a05565b610164610a17565b610164610a29565b6001545b604051908152602001610171565b600254610272565b61015a610296366004611103565b610a3a565b61015a6102a9366004610d8a565b610b8a565b60035460ff166101ba565b6101d76102c7366004610d8a565b610c45565b61015a6102da366004611103565b610c6f565b6000546101d7906001600160a01b031681565b6000546001600160a01b0316331461030957600080fd5b6003805461ffff19166001179055565b60606005600401805461032b9061118b565b80601f01602080910402602001604051908101604052809291908181526020018280546103579061118b565b80156103a45780601f10610379576101008083540402835291602001916103a4565b820191906000526020600020905b81548152906001019060200180831161038757829003601f168201915b5050505050905090565b60606005600001805461032b9061118b565b600460205260009081526040902080546001820180549192916103e29061118b565b80601f016020809104026020016040519081016040528092919081815260200182805461040e9061118b565b801561045b5780601f106104305761010080835404028352916020019161045b565b820191906000526020600020905b81548152906001019060200180831161043e57829003601f168201915b5050505050908060020180546104709061118b565b80601f016020809104026020016040519081016040528092919081815260200182805461049c9061118b565b80156104e95780601f106104be576101008083540402835291602001916104e9565b820191906000526020600020905b8154815290600101906020018083116104cc57829003601f168201915b5050505050908060030180546104fe9061118b565b80601f016020809104026020016040519081016040528092919081815260200182805461052a9061118b565b80156105775780601f1061054c57610100808354040283529160200191610577565b820191906000526020600020905b81548152906001019060200180831161055a57829003601f168201915b5050505050908060040154905085565b60606005600101805461032b9061118b565b6000546001600160a01b031633146105b057600080fd5b6040805160c08101825287815260208101879052908101859052606081018490526080810183905260a081018290526005806105ec8982611214565b50602082015160018201906106019082611214565b50604082015160028201906106169082611214565b506060820151600382019061062b9082611214565b50608082015160048201906106409082611214565b5060a082015160058201906106559082611214565b505050505050505050565b6040805160e0810182526001600160a01b03868116825260208083018781528385018790526060840186905260006080850181905260a08501819052600160c08601819052338252600c90935294909420835181546001600160a01b0319169316929092178255925191928392908201906106db9082611214565b50604082015160028201906106f09082611214565b50606082015160038201906107059082611214565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ff00199415159490941661ffff19909216919091179290921792909216179055600b80546001818101835560009283527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db990910180546001600160a01b0319163317905560028054919290916107a89084906112d4565b90915550505050505050565b600c60205260009081526040902080546001820180546001600160a01b0390921692916107e09061118b565b80601f016020809104026020016040519081016040528092919081815260200182805461080c9061118b565b80156108595780601f1061082e57610100808354040283529160200191610859565b820191906000526020600020905b81548152906001019060200180831161083c57829003601f168201915b50505050509080600201805461086e9061118b565b80601f016020809104026020016040519081016040528092919081815260200182805461089a9061118b565b80156108e75780601f106108bc576101008083540402835291602001916108e7565b820191906000526020600020905b8154815290600101906020018083116108ca57829003601f168201915b5050505050908060030180546108fc9061118b565b80601f01602080910402602001604051908101604052809291908181526020018280546109289061118b565b80156109755780601f1061094a57610100808354040283529160200191610975565b820191906000526020600020905b81548152906001019060200180831161095857829003601f168201915b5050506004909301549192505060ff808216916101008104821691620100009091041687565b6000546001600160a01b031633146109b257600080fd5b6001600160a01b03166000908152600c60205260409020600401805460ff1916911515919091179055565b6000546001600160a01b031633146109f457600080fd5b6003805461ffff1916610100179055565b60606005600301805461032b9061118b565b60606005600201805461032b9061118b565b606060058001805461032b9061118b565b6040805160e0810182523380825260208083018781528385018790526060840186905260006080850181905260a08501819052600160c08601819052938152600c909252939020825181546001600160a01b0319166001600160a01b039091161781559251919283929091820190610ab29082611214565b5060408201516002820190610ac79082611214565b5060608201516003820190610adc9082611214565b5060808201516004909101805460a084015160c0909401511515620100000262ff0000199415156101000261ff00199415159490941661ffff19909216919091179290921792909216179055600b80546001818101835560009283527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db990910180546001600160a01b031916331790556002805491929091610b7f9084906112d4565b909155505050505050565b336000908152600c6020526040902060040154610100900460ff1615610baf57600080fd5b336000908152600c602052604090206004015460ff161515600114610bd357600080fd5b60035460ff161515600114610be757600080fd5b600354610100900460ff1615610bfc57600080fd5b6000818152600460208190526040822001805460019290610c1e9084906112d4565b9091555050336000908152600c60205260409020600401805461ff00191661010017905550565b600b8181548110610c5557600080fd5b6000918252602090912001546001600160a01b0316905081565b6000546001600160a01b03163314610c8657600080fd5b6040805160a081018252600180548083526020808401888152848601889052606085018790526000608086018190529283526004909152939020825181559251919283929091820190610cd99082611214565b5060408201516002820190610cee9082611214565b5060608201516003820190610d039082611214565b50608082015181600401559050506001806000828254610b7f91906112d4565b6000815180845260005b81811015610d4957602081850181015186830182015201610d2d565b81811115610d5b576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610d836020830184610d23565b9392505050565b600060208284031215610d9c57600080fd5b5035919050565b85815260a060208201526000610dbc60a0830187610d23565b8281036040840152610dce8187610d23565b90508281036060840152610de28186610d23565b9150508260808301529695505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e1c57600080fd5b813567ffffffffffffffff80821115610e3757610e37610df5565b604051601f8301601f19908116603f01168101908282118183101715610e5f57610e5f610df5565b81604052838152866020858801011115610e7857600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060008060008060c08789031215610eb157600080fd5b863567ffffffffffffffff80821115610ec957600080fd5b610ed58a838b01610e0b565b97506020890135915080821115610eeb57600080fd5b610ef78a838b01610e0b565b96506040890135915080821115610f0d57600080fd5b610f198a838b01610e0b565b95506060890135915080821115610f2f57600080fd5b610f3b8a838b01610e0b565b94506080890135915080821115610f5157600080fd5b610f5d8a838b01610e0b565b935060a0890135915080821115610f7357600080fd5b50610f8089828a01610e0b565b9150509295509295509295565b80356001600160a01b0381168114610fa457600080fd5b919050565b60008060008060808587031215610fbf57600080fd5b610fc885610f8d565b9350602085013567ffffffffffffffff80821115610fe557600080fd5b610ff188838901610e0b565b9450604087013591508082111561100757600080fd5b61101388838901610e0b565b9350606087013591508082111561102957600080fd5b5061103687828801610e0b565b91505092959194509250565b60006020828403121561105457600080fd5b610d8382610f8d565b6001600160a01b038816815260e06020820181905260009061108190830189610d23565b82810360408401526110938189610d23565b905082810360608401526110a78188610d23565b9515156080840152505091151560a0830152151560c090910152949350505050565b600080604083850312156110dc57600080fd5b823580151581146110ec57600080fd5b91506110fa60208401610f8d565b90509250929050565b60008060006060848603121561111857600080fd5b833567ffffffffffffffff8082111561113057600080fd5b61113c87838801610e0b565b9450602086013591508082111561115257600080fd5b61115e87838801610e0b565b9350604086013591508082111561117457600080fd5b5061118186828701610e0b565b9150509250925092565b600181811c9082168061119f57607f821691505b6020821081036111bf57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561120f57600081815260208120601f850160051c810160208610156111ec5750805b601f850160051c820191505b8181101561120b578281556001016111f8565b5050505b505050565b815167ffffffffffffffff81111561122e5761122e610df5565b6112428161123c845461118b565b846111c5565b602080601f831160018114611277576000841561125f5750858301515b600019600386901b1c1916600185901b17855561120b565b600085815260208120601f198616915b828110156112a657888601518255948401946001909101908401611287565b50858210156112c45787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b600082198211156112f557634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220cdf56ae03cd56cc88a266a3f80962a187fc3bda39b62ddc43aed2970ab4e1cde64736f6c634300080f0033",
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

// PostGeneralVote is a paid mutator transaction binding the contract method 0xa5156307.
//
// Solidity: function post_GeneralVote(uint256 candidateId) returns()
func (_Api *ApiTransactor) PostGeneralVote(opts *bind.TransactOpts, candidateId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "post_GeneralVote", candidateId)
}

// PostGeneralVote is a paid mutator transaction binding the contract method 0xa5156307.
//
// Solidity: function post_GeneralVote(uint256 candidateId) returns()
func (_Api *ApiSession) PostGeneralVote(candidateId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PostGeneralVote(&_Api.TransactOpts, candidateId)
}

// PostGeneralVote is a paid mutator transaction binding the contract method 0xa5156307.
//
// Solidity: function post_GeneralVote(uint256 candidateId) returns()
func (_Api *ApiTransactorSession) PostGeneralVote(candidateId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PostGeneralVote(&_Api.TransactOpts, candidateId)
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
