// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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
	_ = abi.ConvertType
)

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctionCounter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"auctions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createAuction\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"durationHours\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"delistNFT\",\"inputs\":[{\"name\":\"listingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"endAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ethUsdPriceFeed\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractAggregatorV3Interface\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeRecipient\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getActiveListings\",\"inputs\":[],\"outputs\":[{\"name\":\"listingIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuction\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"startingBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidder\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAuctionPriceInUSD\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"startingBidUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"highestBidUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getETHPrice\",\"inputs\":[],\"outputs\":[{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getListing\",\"inputs\":[{\"name\":\"listingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getListingPriceInUSD\",\"inputs\":[{\"name\":\"listingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"priceInUSD\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPaymentTokenAddress\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_feeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_wethAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ethUsdPriceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaymentMethodSupported\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listNFT\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"listingCount\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"listings\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seller\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"},{\"name\":\"isActive\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingReturns\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"placeBid\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"platformFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"purchaseNFT\",\"inputs\":[{\"name\":\"nftContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"listingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"royaltyEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFeeRecipient\",\"inputs\":[{\"name\":\"_feeRecipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPaymentMethodSupported\",\"inputs\":[{\"name\":\"method\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"},{\"name\":\"supported\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPlatformFee\",\"inputs\":[{\"name\":\"_feeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setPriceFeed\",\"inputs\":[{\"name\":\"_priceFeed\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRoyaltyEnabled\",\"inputs\":[{\"name\":\"_enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWETHAddress\",\"inputs\":[{\"name\":\"_wethAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedPaymentMethods\",\"inputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateListingPrice\",\"inputs\":[{\"name\":\"listingId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"wethAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawBidRefund\",\"inputs\":[{\"name\":\"auctionId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTAuctionBidPlaced\",\"inputs\":[{\"name\":\"bidder\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"bidAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTAuctionCancelled\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTAuctionDelisted\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTAuctionEnded\",\"inputs\":[{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"finalBidAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTAuctionListed\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"startingBid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"endTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTDelisted\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTListed\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTListingPriceUpdated\",\"inputs\":[{\"name\":\"seller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"newPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NFTPurchased\",\"inputs\":[{\"name\":\"buyer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"tokenId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"paymentMethod\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentToken.PaymentMethod\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PriceInUSD\",\"inputs\":[{\"name\":\"listingId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"priceInWei\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"priceInUSD\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Marketplace.Contract.UPGRADEINTERFACEVERSION(&_Marketplace.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Marketplace *MarketplaceCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Marketplace.Contract.UPGRADEINTERFACEVERSION(&_Marketplace.CallOpts)
}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Marketplace *MarketplaceCaller) AuctionCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "auctionCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Marketplace *MarketplaceSession) AuctionCounter() (*big.Int, error) {
	return _Marketplace.Contract.AuctionCounter(&_Marketplace.CallOpts)
}

// AuctionCounter is a free data retrieval call binding the contract method 0xa7e76644.
//
// Solidity: function auctionCounter() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) AuctionCounter() (*big.Int, error) {
	return _Marketplace.Contract.AuctionCounter(&_Marketplace.CallOpts)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCaller) Auctions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "auctions", arg0)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		StartingBid   *big.Int
		HighestBid    *big.Int
		HighestBidder common.Address
		EndTime       *big.Int
		PaymentMethod uint8
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartingBid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighestBid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PaymentMethod = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceSession) Auctions(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.Auctions(&_Marketplace.CallOpts, arg0)
}

// Auctions is a free data retrieval call binding the contract method 0x571a26a0.
//
// Solidity: function auctions(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCallerSession) Auctions(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.Auctions(&_Marketplace.CallOpts, arg0)
}

// EthUsdPriceFeed is a free data retrieval call binding the contract method 0x42f6fb29.
//
// Solidity: function ethUsdPriceFeed() view returns(address)
func (_Marketplace *MarketplaceCaller) EthUsdPriceFeed(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "ethUsdPriceFeed")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EthUsdPriceFeed is a free data retrieval call binding the contract method 0x42f6fb29.
//
// Solidity: function ethUsdPriceFeed() view returns(address)
func (_Marketplace *MarketplaceSession) EthUsdPriceFeed() (common.Address, error) {
	return _Marketplace.Contract.EthUsdPriceFeed(&_Marketplace.CallOpts)
}

// EthUsdPriceFeed is a free data retrieval call binding the contract method 0x42f6fb29.
//
// Solidity: function ethUsdPriceFeed() view returns(address)
func (_Marketplace *MarketplaceCallerSession) EthUsdPriceFeed() (common.Address, error) {
	return _Marketplace.Contract.EthUsdPriceFeed(&_Marketplace.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Marketplace *MarketplaceCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Marketplace *MarketplaceSession) FeeRecipient() (common.Address, error) {
	return _Marketplace.Contract.FeeRecipient(&_Marketplace.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Marketplace *MarketplaceCallerSession) FeeRecipient() (common.Address, error) {
	return _Marketplace.Contract.FeeRecipient(&_Marketplace.CallOpts)
}

// GetActiveListings is a free data retrieval call binding the contract method 0x87c35bc0.
//
// Solidity: function getActiveListings() view returns(uint256[] listingIds)
func (_Marketplace *MarketplaceCaller) GetActiveListings(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getActiveListings")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveListings is a free data retrieval call binding the contract method 0x87c35bc0.
//
// Solidity: function getActiveListings() view returns(uint256[] listingIds)
func (_Marketplace *MarketplaceSession) GetActiveListings() ([]*big.Int, error) {
	return _Marketplace.Contract.GetActiveListings(&_Marketplace.CallOpts)
}

// GetActiveListings is a free data retrieval call binding the contract method 0x87c35bc0.
//
// Solidity: function getActiveListings() view returns(uint256[] listingIds)
func (_Marketplace *MarketplaceCallerSession) GetActiveListings() ([]*big.Int, error) {
	return _Marketplace.Contract.GetActiveListings(&_Marketplace.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCaller) GetAuction(opts *bind.CallOpts, auctionId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getAuction", auctionId)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		StartingBid   *big.Int
		HighestBid    *big.Int
		HighestBidder common.Address
		EndTime       *big.Int
		PaymentMethod uint8
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartingBid = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.HighestBid = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HighestBidder = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.EndTime = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PaymentMethod = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceSession) GetAuction(auctionId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.GetAuction(&_Marketplace.CallOpts, auctionId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 auctionId) view returns(address seller, address nftContract, uint256 tokenId, uint256 startingBid, uint256 highestBid, address highestBidder, uint256 endTime, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCallerSession) GetAuction(auctionId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	HighestBid    *big.Int
	HighestBidder common.Address
	EndTime       *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.GetAuction(&_Marketplace.CallOpts, auctionId)
}

// GetAuctionPriceInUSD is a free data retrieval call binding the contract method 0x74a705b8.
//
// Solidity: function getAuctionPriceInUSD(uint256 auctionId) view returns(uint256 startingBidUSD, uint256 highestBidUSD)
func (_Marketplace *MarketplaceCaller) GetAuctionPriceInUSD(opts *bind.CallOpts, auctionId *big.Int) (struct {
	StartingBidUSD *big.Int
	HighestBidUSD  *big.Int
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getAuctionPriceInUSD", auctionId)

	outstruct := new(struct {
		StartingBidUSD *big.Int
		HighestBidUSD  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartingBidUSD = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.HighestBidUSD = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAuctionPriceInUSD is a free data retrieval call binding the contract method 0x74a705b8.
//
// Solidity: function getAuctionPriceInUSD(uint256 auctionId) view returns(uint256 startingBidUSD, uint256 highestBidUSD)
func (_Marketplace *MarketplaceSession) GetAuctionPriceInUSD(auctionId *big.Int) (struct {
	StartingBidUSD *big.Int
	HighestBidUSD  *big.Int
}, error) {
	return _Marketplace.Contract.GetAuctionPriceInUSD(&_Marketplace.CallOpts, auctionId)
}

// GetAuctionPriceInUSD is a free data retrieval call binding the contract method 0x74a705b8.
//
// Solidity: function getAuctionPriceInUSD(uint256 auctionId) view returns(uint256 startingBidUSD, uint256 highestBidUSD)
func (_Marketplace *MarketplaceCallerSession) GetAuctionPriceInUSD(auctionId *big.Int) (struct {
	StartingBidUSD *big.Int
	HighestBidUSD  *big.Int
}, error) {
	return _Marketplace.Contract.GetAuctionPriceInUSD(&_Marketplace.CallOpts, auctionId)
}

// GetETHPrice is a free data retrieval call binding the contract method 0xa607a8d9.
//
// Solidity: function getETHPrice() view returns(uint256 price)
func (_Marketplace *MarketplaceCaller) GetETHPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getETHPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetETHPrice is a free data retrieval call binding the contract method 0xa607a8d9.
//
// Solidity: function getETHPrice() view returns(uint256 price)
func (_Marketplace *MarketplaceSession) GetETHPrice() (*big.Int, error) {
	return _Marketplace.Contract.GetETHPrice(&_Marketplace.CallOpts)
}

// GetETHPrice is a free data retrieval call binding the contract method 0xa607a8d9.
//
// Solidity: function getETHPrice() view returns(uint256 price)
func (_Marketplace *MarketplaceCallerSession) GetETHPrice() (*big.Int, error) {
	return _Marketplace.Contract.GetETHPrice(&_Marketplace.CallOpts)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCaller) GetListing(opts *bind.CallOpts, listingId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListing", listingId)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		Price         *big.Int
		PaymentMethod uint8
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PaymentMethod = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceSession) GetListing(listingId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetListing is a free data retrieval call binding the contract method 0x107a274a.
//
// Solidity: function getListing(uint256 listingId) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCallerSession) GetListing(listingId *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.GetListing(&_Marketplace.CallOpts, listingId)
}

// GetListingPriceInUSD is a free data retrieval call binding the contract method 0x4adbc9df.
//
// Solidity: function getListingPriceInUSD(uint256 listingId) view returns(uint256 priceInUSD)
func (_Marketplace *MarketplaceCaller) GetListingPriceInUSD(opts *bind.CallOpts, listingId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListingPriceInUSD", listingId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPriceInUSD is a free data retrieval call binding the contract method 0x4adbc9df.
//
// Solidity: function getListingPriceInUSD(uint256 listingId) view returns(uint256 priceInUSD)
func (_Marketplace *MarketplaceSession) GetListingPriceInUSD(listingId *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.GetListingPriceInUSD(&_Marketplace.CallOpts, listingId)
}

// GetListingPriceInUSD is a free data retrieval call binding the contract method 0x4adbc9df.
//
// Solidity: function getListingPriceInUSD(uint256 listingId) view returns(uint256 priceInUSD)
func (_Marketplace *MarketplaceCallerSession) GetListingPriceInUSD(listingId *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.GetListingPriceInUSD(&_Marketplace.CallOpts, listingId)
}

// GetPaymentTokenAddress is a free data retrieval call binding the contract method 0x0bb182dd.
//
// Solidity: function getPaymentTokenAddress(uint8 method) view returns(address)
func (_Marketplace *MarketplaceCaller) GetPaymentTokenAddress(opts *bind.CallOpts, method uint8) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getPaymentTokenAddress", method)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPaymentTokenAddress is a free data retrieval call binding the contract method 0x0bb182dd.
//
// Solidity: function getPaymentTokenAddress(uint8 method) view returns(address)
func (_Marketplace *MarketplaceSession) GetPaymentTokenAddress(method uint8) (common.Address, error) {
	return _Marketplace.Contract.GetPaymentTokenAddress(&_Marketplace.CallOpts, method)
}

// GetPaymentTokenAddress is a free data retrieval call binding the contract method 0x0bb182dd.
//
// Solidity: function getPaymentTokenAddress(uint8 method) view returns(address)
func (_Marketplace *MarketplaceCallerSession) GetPaymentTokenAddress(method uint8) (common.Address, error) {
	return _Marketplace.Contract.GetPaymentTokenAddress(&_Marketplace.CallOpts, method)
}

// IsPaymentMethodSupported is a free data retrieval call binding the contract method 0x2f40762f.
//
// Solidity: function isPaymentMethodSupported(uint8 method) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsPaymentMethodSupported(opts *bind.CallOpts, method uint8) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isPaymentMethodSupported", method)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaymentMethodSupported is a free data retrieval call binding the contract method 0x2f40762f.
//
// Solidity: function isPaymentMethodSupported(uint8 method) view returns(bool)
func (_Marketplace *MarketplaceSession) IsPaymentMethodSupported(method uint8) (bool, error) {
	return _Marketplace.Contract.IsPaymentMethodSupported(&_Marketplace.CallOpts, method)
}

// IsPaymentMethodSupported is a free data retrieval call binding the contract method 0x2f40762f.
//
// Solidity: function isPaymentMethodSupported(uint8 method) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsPaymentMethodSupported(method uint8) (bool, error) {
	return _Marketplace.Contract.IsPaymentMethodSupported(&_Marketplace.CallOpts, method)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) ListingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) ListingCount() (*big.Int, error) {
	return _Marketplace.Contract.ListingCount(&_Marketplace.CallOpts)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) ListingCount() (*big.Int, error) {
	return _Marketplace.Contract.ListingCount(&_Marketplace.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Seller        common.Address
		NftContract   common.Address
		TokenId       *big.Int
		Price         *big.Int
		PaymentMethod uint8
		IsActive      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.NftContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PaymentMethod = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.IsActive = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod, bool isActive)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Seller        common.Address
	NftContract   common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	IsActive      bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// PendingReturns is a free data retrieval call binding the contract method 0x8ddac1ef.
//
// Solidity: function pendingReturns(uint256 , address ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) PendingReturns(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pendingReturns", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingReturns is a free data retrieval call binding the contract method 0x8ddac1ef.
//
// Solidity: function pendingReturns(uint256 , address ) view returns(uint256)
func (_Marketplace *MarketplaceSession) PendingReturns(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.PendingReturns(&_Marketplace.CallOpts, arg0, arg1)
}

// PendingReturns is a free data retrieval call binding the contract method 0x8ddac1ef.
//
// Solidity: function pendingReturns(uint256 , address ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) PendingReturns(arg0 *big.Int, arg1 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.PendingReturns(&_Marketplace.CallOpts, arg0, arg1)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Marketplace *MarketplaceCaller) PlatformFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "platformFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Marketplace *MarketplaceSession) PlatformFee() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFee(&_Marketplace.CallOpts)
}

// PlatformFee is a free data retrieval call binding the contract method 0x26232a2e.
//
// Solidity: function platformFee() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) PlatformFee() (*big.Int, error) {
	return _Marketplace.Contract.PlatformFee(&_Marketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceSession) ProxiableUUID() ([32]byte, error) {
	return _Marketplace.Contract.ProxiableUUID(&_Marketplace.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Marketplace *MarketplaceCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Marketplace.Contract.ProxiableUUID(&_Marketplace.CallOpts)
}

// RoyaltyEnabled is a free data retrieval call binding the contract method 0x55984b92.
//
// Solidity: function royaltyEnabled() view returns(bool)
func (_Marketplace *MarketplaceCaller) RoyaltyEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "royaltyEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RoyaltyEnabled is a free data retrieval call binding the contract method 0x55984b92.
//
// Solidity: function royaltyEnabled() view returns(bool)
func (_Marketplace *MarketplaceSession) RoyaltyEnabled() (bool, error) {
	return _Marketplace.Contract.RoyaltyEnabled(&_Marketplace.CallOpts)
}

// RoyaltyEnabled is a free data retrieval call binding the contract method 0x55984b92.
//
// Solidity: function royaltyEnabled() view returns(bool)
func (_Marketplace *MarketplaceCallerSession) RoyaltyEnabled() (bool, error) {
	return _Marketplace.Contract.RoyaltyEnabled(&_Marketplace.CallOpts)
}

// SupportedPaymentMethods is a free data retrieval call binding the contract method 0x7c012c73.
//
// Solidity: function supportedPaymentMethods(uint8 ) view returns(bool)
func (_Marketplace *MarketplaceCaller) SupportedPaymentMethods(opts *bind.CallOpts, arg0 uint8) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "supportedPaymentMethods", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportedPaymentMethods is a free data retrieval call binding the contract method 0x7c012c73.
//
// Solidity: function supportedPaymentMethods(uint8 ) view returns(bool)
func (_Marketplace *MarketplaceSession) SupportedPaymentMethods(arg0 uint8) (bool, error) {
	return _Marketplace.Contract.SupportedPaymentMethods(&_Marketplace.CallOpts, arg0)
}

// SupportedPaymentMethods is a free data retrieval call binding the contract method 0x7c012c73.
//
// Solidity: function supportedPaymentMethods(uint8 ) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) SupportedPaymentMethods(arg0 uint8) (bool, error) {
	return _Marketplace.Contract.SupportedPaymentMethods(&_Marketplace.CallOpts, arg0)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Marketplace *MarketplaceCaller) WethAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "wethAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Marketplace *MarketplaceSession) WethAddress() (common.Address, error) {
	return _Marketplace.Contract.WethAddress(&_Marketplace.CallOpts)
}

// WethAddress is a free data retrieval call binding the contract method 0x4f0e0ef3.
//
// Solidity: function wethAddress() view returns(address)
func (_Marketplace *MarketplaceCallerSession) WethAddress() (common.Address, error) {
	return _Marketplace.Contract.WethAddress(&_Marketplace.CallOpts)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4479b344.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startingBid, uint256 durationHours, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceTransactor) CreateAuction(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, startingBid *big.Int, durationHours *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createAuction", nftContract, tokenId, startingBid, durationHours, paymentMethod)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4479b344.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startingBid, uint256 durationHours, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startingBid *big.Int, durationHours *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateAuction(&_Marketplace.TransactOpts, nftContract, tokenId, startingBid, durationHours, paymentMethod)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x4479b344.
//
// Solidity: function createAuction(address nftContract, uint256 tokenId, uint256 startingBid, uint256 durationHours, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceTransactorSession) CreateAuction(nftContract common.Address, tokenId *big.Int, startingBid *big.Int, durationHours *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateAuction(&_Marketplace.TransactOpts, nftContract, tokenId, startingBid, durationHours, paymentMethod)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xcb917b9c.
//
// Solidity: function delistNFT(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactor) DelistNFT(opts *bind.TransactOpts, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "delistNFT", listingId)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xcb917b9c.
//
// Solidity: function delistNFT(uint256 listingId) returns()
func (_Marketplace *MarketplaceSession) DelistNFT(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.DelistNFT(&_Marketplace.TransactOpts, listingId)
}

// DelistNFT is a paid mutator transaction binding the contract method 0xcb917b9c.
//
// Solidity: function delistNFT(uint256 listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) DelistNFT(listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.DelistNFT(&_Marketplace.TransactOpts, listingId)
}

// EndAuction is a paid mutator transaction binding the contract method 0xcbf8455d.
//
// Solidity: function endAuction(uint256 auctionId, address nftContract) returns()
func (_Marketplace *MarketplaceTransactor) EndAuction(opts *bind.TransactOpts, auctionId *big.Int, nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "endAuction", auctionId, nftContract)
}

// EndAuction is a paid mutator transaction binding the contract method 0xcbf8455d.
//
// Solidity: function endAuction(uint256 auctionId, address nftContract) returns()
func (_Marketplace *MarketplaceSession) EndAuction(auctionId *big.Int, nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.EndAuction(&_Marketplace.TransactOpts, auctionId, nftContract)
}

// EndAuction is a paid mutator transaction binding the contract method 0xcbf8455d.
//
// Solidity: function endAuction(uint256 auctionId, address nftContract) returns()
func (_Marketplace *MarketplaceTransactorSession) EndAuction(auctionId *big.Int, nftContract common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.EndAuction(&_Marketplace.TransactOpts, auctionId, nftContract)
}

// Initialize is a paid mutator transaction binding the contract method 0x754d1d54.
//
// Solidity: function initialize(uint256 _feeBps, address _feeRecipient, address _wethAddress, address _ethUsdPriceFeed) returns()
func (_Marketplace *MarketplaceTransactor) Initialize(opts *bind.TransactOpts, _feeBps *big.Int, _feeRecipient common.Address, _wethAddress common.Address, _ethUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "initialize", _feeBps, _feeRecipient, _wethAddress, _ethUsdPriceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0x754d1d54.
//
// Solidity: function initialize(uint256 _feeBps, address _feeRecipient, address _wethAddress, address _ethUsdPriceFeed) returns()
func (_Marketplace *MarketplaceSession) Initialize(_feeBps *big.Int, _feeRecipient common.Address, _wethAddress common.Address, _ethUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _feeBps, _feeRecipient, _wethAddress, _ethUsdPriceFeed)
}

// Initialize is a paid mutator transaction binding the contract method 0x754d1d54.
//
// Solidity: function initialize(uint256 _feeBps, address _feeRecipient, address _wethAddress, address _ethUsdPriceFeed) returns()
func (_Marketplace *MarketplaceTransactorSession) Initialize(_feeBps *big.Int, _feeRecipient common.Address, _wethAddress common.Address, _ethUsdPriceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.Initialize(&_Marketplace.TransactOpts, _feeBps, _feeRecipient, _wethAddress, _ethUsdPriceFeed)
}

// ListNFT is a paid mutator transaction binding the contract method 0x2fa5546b.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceTransactor) ListNFT(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "listNFT", nftContract, tokenId, price, paymentMethod)
}

// ListNFT is a paid mutator transaction binding the contract method 0x2fa5546b.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceSession) ListNFT(nftContract common.Address, tokenId *big.Int, price *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ListNFT(&_Marketplace.TransactOpts, nftContract, tokenId, price, paymentMethod)
}

// ListNFT is a paid mutator transaction binding the contract method 0x2fa5546b.
//
// Solidity: function listNFT(address nftContract, uint256 tokenId, uint256 price, uint8 paymentMethod) returns(uint256)
func (_Marketplace *MarketplaceTransactorSession) ListNFT(nftContract common.Address, tokenId *big.Int, price *big.Int, paymentMethod uint8) (*types.Transaction, error) {
	return _Marketplace.Contract.ListNFT(&_Marketplace.TransactOpts, nftContract, tokenId, price, paymentMethod)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Marketplace *MarketplaceTransactor) PlaceBid(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "placeBid", auctionId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Marketplace *MarketplaceSession) PlaceBid(auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PlaceBid(&_Marketplace.TransactOpts, auctionId)
}

// PlaceBid is a paid mutator transaction binding the contract method 0x9979ef45.
//
// Solidity: function placeBid(uint256 auctionId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) PlaceBid(auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PlaceBid(&_Marketplace.TransactOpts, auctionId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address nftContract, uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactor) PurchaseNFT(opts *bind.TransactOpts, nftContract common.Address, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "purchaseNFT", nftContract, listingId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address nftContract, uint256 listingId) payable returns()
func (_Marketplace *MarketplaceSession) PurchaseNFT(nftContract common.Address, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseNFT(&_Marketplace.TransactOpts, nftContract, listingId)
}

// PurchaseNFT is a paid mutator transaction binding the contract method 0xa87e25ac.
//
// Solidity: function purchaseNFT(address nftContract, uint256 listingId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) PurchaseNFT(nftContract common.Address, listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseNFT(&_Marketplace.TransactOpts, nftContract, listingId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Marketplace *MarketplaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Marketplace.Contract.RenounceOwnership(&_Marketplace.TransactOpts)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Marketplace *MarketplaceTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Marketplace *MarketplaceSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeRecipient(&_Marketplace.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Marketplace *MarketplaceTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetFeeRecipient(&_Marketplace.TransactOpts, _feeRecipient)
}

// SetPaymentMethodSupported is a paid mutator transaction binding the contract method 0x5daa5581.
//
// Solidity: function setPaymentMethodSupported(uint8 method, bool supported) returns()
func (_Marketplace *MarketplaceTransactor) SetPaymentMethodSupported(opts *bind.TransactOpts, method uint8, supported bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setPaymentMethodSupported", method, supported)
}

// SetPaymentMethodSupported is a paid mutator transaction binding the contract method 0x5daa5581.
//
// Solidity: function setPaymentMethodSupported(uint8 method, bool supported) returns()
func (_Marketplace *MarketplaceSession) SetPaymentMethodSupported(method uint8, supported bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPaymentMethodSupported(&_Marketplace.TransactOpts, method, supported)
}

// SetPaymentMethodSupported is a paid mutator transaction binding the contract method 0x5daa5581.
//
// Solidity: function setPaymentMethodSupported(uint8 method, bool supported) returns()
func (_Marketplace *MarketplaceTransactorSession) SetPaymentMethodSupported(method uint8, supported bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPaymentMethodSupported(&_Marketplace.TransactOpts, method, supported)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceTransactor) SetPlatformFee(opts *bind.TransactOpts, _feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setPlatformFee", _feeBps)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceSession) SetPlatformFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPlatformFee(&_Marketplace.TransactOpts, _feeBps)
}

// SetPlatformFee is a paid mutator transaction binding the contract method 0x12e8e2c3.
//
// Solidity: function setPlatformFee(uint256 _feeBps) returns()
func (_Marketplace *MarketplaceTransactorSession) SetPlatformFee(_feeBps *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPlatformFee(&_Marketplace.TransactOpts, _feeBps)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Marketplace *MarketplaceTransactor) SetPriceFeed(opts *bind.TransactOpts, _priceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setPriceFeed", _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Marketplace *MarketplaceSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPriceFeed(&_Marketplace.TransactOpts, _priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address _priceFeed) returns()
func (_Marketplace *MarketplaceTransactorSession) SetPriceFeed(_priceFeed common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetPriceFeed(&_Marketplace.TransactOpts, _priceFeed)
}

// SetRoyaltyEnabled is a paid mutator transaction binding the contract method 0xf8e3e9e0.
//
// Solidity: function setRoyaltyEnabled(bool _enabled) returns()
func (_Marketplace *MarketplaceTransactor) SetRoyaltyEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setRoyaltyEnabled", _enabled)
}

// SetRoyaltyEnabled is a paid mutator transaction binding the contract method 0xf8e3e9e0.
//
// Solidity: function setRoyaltyEnabled(bool _enabled) returns()
func (_Marketplace *MarketplaceSession) SetRoyaltyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyEnabled(&_Marketplace.TransactOpts, _enabled)
}

// SetRoyaltyEnabled is a paid mutator transaction binding the contract method 0xf8e3e9e0.
//
// Solidity: function setRoyaltyEnabled(bool _enabled) returns()
func (_Marketplace *MarketplaceTransactorSession) SetRoyaltyEnabled(_enabled bool) (*types.Transaction, error) {
	return _Marketplace.Contract.SetRoyaltyEnabled(&_Marketplace.TransactOpts, _enabled)
}

// SetWETHAddress is a paid mutator transaction binding the contract method 0x6a4234eb.
//
// Solidity: function setWETHAddress(address _wethAddress) returns()
func (_Marketplace *MarketplaceTransactor) SetWETHAddress(opts *bind.TransactOpts, _wethAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setWETHAddress", _wethAddress)
}

// SetWETHAddress is a paid mutator transaction binding the contract method 0x6a4234eb.
//
// Solidity: function setWETHAddress(address _wethAddress) returns()
func (_Marketplace *MarketplaceSession) SetWETHAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetWETHAddress(&_Marketplace.TransactOpts, _wethAddress)
}

// SetWETHAddress is a paid mutator transaction binding the contract method 0x6a4234eb.
//
// Solidity: function setWETHAddress(address _wethAddress) returns()
func (_Marketplace *MarketplaceTransactorSession) SetWETHAddress(_wethAddress common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.SetWETHAddress(&_Marketplace.TransactOpts, _wethAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Marketplace *MarketplaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Marketplace.Contract.TransferOwnership(&_Marketplace.TransactOpts, newOwner)
}

// UpdateListingPrice is a paid mutator transaction binding the contract method 0xc4604943.
//
// Solidity: function updateListingPrice(uint256 listingId, uint256 newPrice) returns()
func (_Marketplace *MarketplaceTransactor) UpdateListingPrice(opts *bind.TransactOpts, listingId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "updateListingPrice", listingId, newPrice)
}

// UpdateListingPrice is a paid mutator transaction binding the contract method 0xc4604943.
//
// Solidity: function updateListingPrice(uint256 listingId, uint256 newPrice) returns()
func (_Marketplace *MarketplaceSession) UpdateListingPrice(listingId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateListingPrice(&_Marketplace.TransactOpts, listingId, newPrice)
}

// UpdateListingPrice is a paid mutator transaction binding the contract method 0xc4604943.
//
// Solidity: function updateListingPrice(uint256 listingId, uint256 newPrice) returns()
func (_Marketplace *MarketplaceTransactorSession) UpdateListingPrice(listingId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.UpdateListingPrice(&_Marketplace.TransactOpts, listingId, newPrice)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.UpgradeToAndCall(&_Marketplace.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Marketplace *MarketplaceTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Marketplace.Contract.UpgradeToAndCall(&_Marketplace.TransactOpts, newImplementation, data)
}

// WithdrawBidRefund is a paid mutator transaction binding the contract method 0x34fa4e04.
//
// Solidity: function withdrawBidRefund(uint256 auctionId) returns()
func (_Marketplace *MarketplaceTransactor) WithdrawBidRefund(opts *bind.TransactOpts, auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawBidRefund", auctionId)
}

// WithdrawBidRefund is a paid mutator transaction binding the contract method 0x34fa4e04.
//
// Solidity: function withdrawBidRefund(uint256 auctionId) returns()
func (_Marketplace *MarketplaceSession) WithdrawBidRefund(auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawBidRefund(&_Marketplace.TransactOpts, auctionId)
}

// WithdrawBidRefund is a paid mutator transaction binding the contract method 0x34fa4e04.
//
// Solidity: function withdrawBidRefund(uint256 auctionId) returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawBidRefund(auctionId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawBidRefund(&_Marketplace.TransactOpts, auctionId)
}

// MarketplaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Marketplace contract.
type MarketplaceInitializedIterator struct {
	Event *MarketplaceInitialized // Event containing the contract specifics and raw log

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
func (it *MarketplaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceInitialized)
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
		it.Event = new(MarketplaceInitialized)
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
func (it *MarketplaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceInitialized represents a Initialized event raised by the Marketplace contract.
type MarketplaceInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Marketplace *MarketplaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*MarketplaceInitializedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MarketplaceInitializedIterator{contract: _Marketplace.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Marketplace *MarketplaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MarketplaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceInitialized)
				if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Marketplace *MarketplaceFilterer) ParseInitialized(log types.Log) (*MarketplaceInitialized, error) {
	event := new(MarketplaceInitialized)
	if err := _Marketplace.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTAuctionBidPlacedIterator is returned from FilterNFTAuctionBidPlaced and is used to iterate over the raw logs and unpacked data for NFTAuctionBidPlaced events raised by the Marketplace contract.
type MarketplaceNFTAuctionBidPlacedIterator struct {
	Event *MarketplaceNFTAuctionBidPlaced // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTAuctionBidPlacedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTAuctionBidPlaced)
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
		it.Event = new(MarketplaceNFTAuctionBidPlaced)
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
func (it *MarketplaceNFTAuctionBidPlacedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTAuctionBidPlacedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTAuctionBidPlaced represents a NFTAuctionBidPlaced event raised by the Marketplace contract.
type MarketplaceNFTAuctionBidPlaced struct {
	Bidder    common.Address
	TokenId   *big.Int
	BidAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNFTAuctionBidPlaced is a free log retrieval operation binding the contract event 0x99a7095cf382ac930a9a37710a33aaad46f2b02c0fd130166b7217a4d2b5e8b3.
//
// Solidity: event NFTAuctionBidPlaced(address indexed bidder, uint256 indexed tokenId, uint256 bidAmount)
func (_Marketplace *MarketplaceFilterer) FilterNFTAuctionBidPlaced(opts *bind.FilterOpts, bidder []common.Address, tokenId []*big.Int) (*MarketplaceNFTAuctionBidPlacedIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTAuctionBidPlaced", bidderRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTAuctionBidPlacedIterator{contract: _Marketplace.contract, event: "NFTAuctionBidPlaced", logs: logs, sub: sub}, nil
}

// WatchNFTAuctionBidPlaced is a free log subscription operation binding the contract event 0x99a7095cf382ac930a9a37710a33aaad46f2b02c0fd130166b7217a4d2b5e8b3.
//
// Solidity: event NFTAuctionBidPlaced(address indexed bidder, uint256 indexed tokenId, uint256 bidAmount)
func (_Marketplace *MarketplaceFilterer) WatchNFTAuctionBidPlaced(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTAuctionBidPlaced, bidder []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTAuctionBidPlaced", bidderRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTAuctionBidPlaced)
				if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionBidPlaced", log); err != nil {
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

// ParseNFTAuctionBidPlaced is a log parse operation binding the contract event 0x99a7095cf382ac930a9a37710a33aaad46f2b02c0fd130166b7217a4d2b5e8b3.
//
// Solidity: event NFTAuctionBidPlaced(address indexed bidder, uint256 indexed tokenId, uint256 bidAmount)
func (_Marketplace *MarketplaceFilterer) ParseNFTAuctionBidPlaced(log types.Log) (*MarketplaceNFTAuctionBidPlaced, error) {
	event := new(MarketplaceNFTAuctionBidPlaced)
	if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionBidPlaced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTAuctionCancelledIterator is returned from FilterNFTAuctionCancelled and is used to iterate over the raw logs and unpacked data for NFTAuctionCancelled events raised by the Marketplace contract.
type MarketplaceNFTAuctionCancelledIterator struct {
	Event *MarketplaceNFTAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTAuctionCancelled)
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
		it.Event = new(MarketplaceNFTAuctionCancelled)
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
func (it *MarketplaceNFTAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTAuctionCancelled represents a NFTAuctionCancelled event raised by the Marketplace contract.
type MarketplaceNFTAuctionCancelled struct {
	Seller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTAuctionCancelled is a free log retrieval operation binding the contract event 0xb46058bf7976b4f620a776a89e173eaa83db8f3eae64e8f75f5e81f817138f4f.
//
// Solidity: event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterNFTAuctionCancelled(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTAuctionCancelledIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTAuctionCancelled", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTAuctionCancelledIterator{contract: _Marketplace.contract, event: "NFTAuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchNFTAuctionCancelled is a free log subscription operation binding the contract event 0xb46058bf7976b4f620a776a89e173eaa83db8f3eae64e8f75f5e81f817138f4f.
//
// Solidity: event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchNFTAuctionCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTAuctionCancelled, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTAuctionCancelled", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTAuctionCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionCancelled", log); err != nil {
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

// ParseNFTAuctionCancelled is a log parse operation binding the contract event 0xb46058bf7976b4f620a776a89e173eaa83db8f3eae64e8f75f5e81f817138f4f.
//
// Solidity: event NFTAuctionCancelled(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseNFTAuctionCancelled(log types.Log) (*MarketplaceNFTAuctionCancelled, error) {
	event := new(MarketplaceNFTAuctionCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTAuctionDelistedIterator is returned from FilterNFTAuctionDelisted and is used to iterate over the raw logs and unpacked data for NFTAuctionDelisted events raised by the Marketplace contract.
type MarketplaceNFTAuctionDelistedIterator struct {
	Event *MarketplaceNFTAuctionDelisted // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTAuctionDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTAuctionDelisted)
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
		it.Event = new(MarketplaceNFTAuctionDelisted)
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
func (it *MarketplaceNFTAuctionDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTAuctionDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTAuctionDelisted represents a NFTAuctionDelisted event raised by the Marketplace contract.
type MarketplaceNFTAuctionDelisted struct {
	Seller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTAuctionDelisted is a free log retrieval operation binding the contract event 0x8871e5d9db01a894a3c548c86f959087144c006783e688f63966eb5f7585207c.
//
// Solidity: event NFTAuctionDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterNFTAuctionDelisted(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTAuctionDelistedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTAuctionDelisted", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTAuctionDelistedIterator{contract: _Marketplace.contract, event: "NFTAuctionDelisted", logs: logs, sub: sub}, nil
}

// WatchNFTAuctionDelisted is a free log subscription operation binding the contract event 0x8871e5d9db01a894a3c548c86f959087144c006783e688f63966eb5f7585207c.
//
// Solidity: event NFTAuctionDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchNFTAuctionDelisted(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTAuctionDelisted, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTAuctionDelisted", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTAuctionDelisted)
				if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionDelisted", log); err != nil {
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

// ParseNFTAuctionDelisted is a log parse operation binding the contract event 0x8871e5d9db01a894a3c548c86f959087144c006783e688f63966eb5f7585207c.
//
// Solidity: event NFTAuctionDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseNFTAuctionDelisted(log types.Log) (*MarketplaceNFTAuctionDelisted, error) {
	event := new(MarketplaceNFTAuctionDelisted)
	if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTAuctionEndedIterator is returned from FilterNFTAuctionEnded and is used to iterate over the raw logs and unpacked data for NFTAuctionEnded events raised by the Marketplace contract.
type MarketplaceNFTAuctionEndedIterator struct {
	Event *MarketplaceNFTAuctionEnded // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTAuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTAuctionEnded)
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
		it.Event = new(MarketplaceNFTAuctionEnded)
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
func (it *MarketplaceNFTAuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTAuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTAuctionEnded represents a NFTAuctionEnded event raised by the Marketplace contract.
type MarketplaceNFTAuctionEnded struct {
	Winner         common.Address
	TokenId        *big.Int
	FinalBidAmount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNFTAuctionEnded is a free log retrieval operation binding the contract event 0x5f45bd830be663a7ec749a07692f3d7aa7e998ad8ea733bf07748f01597bbb76.
//
// Solidity: event NFTAuctionEnded(address indexed winner, uint256 indexed tokenId, uint256 finalBidAmount)
func (_Marketplace *MarketplaceFilterer) FilterNFTAuctionEnded(opts *bind.FilterOpts, winner []common.Address, tokenId []*big.Int) (*MarketplaceNFTAuctionEndedIterator, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTAuctionEnded", winnerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTAuctionEndedIterator{contract: _Marketplace.contract, event: "NFTAuctionEnded", logs: logs, sub: sub}, nil
}

// WatchNFTAuctionEnded is a free log subscription operation binding the contract event 0x5f45bd830be663a7ec749a07692f3d7aa7e998ad8ea733bf07748f01597bbb76.
//
// Solidity: event NFTAuctionEnded(address indexed winner, uint256 indexed tokenId, uint256 finalBidAmount)
func (_Marketplace *MarketplaceFilterer) WatchNFTAuctionEnded(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTAuctionEnded, winner []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTAuctionEnded", winnerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTAuctionEnded)
				if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionEnded", log); err != nil {
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

// ParseNFTAuctionEnded is a log parse operation binding the contract event 0x5f45bd830be663a7ec749a07692f3d7aa7e998ad8ea733bf07748f01597bbb76.
//
// Solidity: event NFTAuctionEnded(address indexed winner, uint256 indexed tokenId, uint256 finalBidAmount)
func (_Marketplace *MarketplaceFilterer) ParseNFTAuctionEnded(log types.Log) (*MarketplaceNFTAuctionEnded, error) {
	event := new(MarketplaceNFTAuctionEnded)
	if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionEnded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTAuctionListedIterator is returned from FilterNFTAuctionListed and is used to iterate over the raw logs and unpacked data for NFTAuctionListed events raised by the Marketplace contract.
type MarketplaceNFTAuctionListedIterator struct {
	Event *MarketplaceNFTAuctionListed // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTAuctionListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTAuctionListed)
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
		it.Event = new(MarketplaceNFTAuctionListed)
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
func (it *MarketplaceNFTAuctionListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTAuctionListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTAuctionListed represents a NFTAuctionListed event raised by the Marketplace contract.
type MarketplaceNFTAuctionListed struct {
	Seller        common.Address
	TokenId       *big.Int
	StartingBid   *big.Int
	EndTime       *big.Int
	PaymentMethod uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNFTAuctionListed is a free log retrieval operation binding the contract event 0xf0420338a19ce6611667836279d49c5e443a6f03eeb00c98f9f7f45007653729.
//
// Solidity: event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) FilterNFTAuctionListed(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTAuctionListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTAuctionListed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTAuctionListedIterator{contract: _Marketplace.contract, event: "NFTAuctionListed", logs: logs, sub: sub}, nil
}

// WatchNFTAuctionListed is a free log subscription operation binding the contract event 0xf0420338a19ce6611667836279d49c5e443a6f03eeb00c98f9f7f45007653729.
//
// Solidity: event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) WatchNFTAuctionListed(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTAuctionListed, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTAuctionListed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTAuctionListed)
				if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionListed", log); err != nil {
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

// ParseNFTAuctionListed is a log parse operation binding the contract event 0xf0420338a19ce6611667836279d49c5e443a6f03eeb00c98f9f7f45007653729.
//
// Solidity: event NFTAuctionListed(address indexed seller, uint256 indexed tokenId, uint256 startingBid, uint256 endTime, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) ParseNFTAuctionListed(log types.Log) (*MarketplaceNFTAuctionListed, error) {
	event := new(MarketplaceNFTAuctionListed)
	if err := _Marketplace.contract.UnpackLog(event, "NFTAuctionListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTDelistedIterator is returned from FilterNFTDelisted and is used to iterate over the raw logs and unpacked data for NFTDelisted events raised by the Marketplace contract.
type MarketplaceNFTDelistedIterator struct {
	Event *MarketplaceNFTDelisted // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTDelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTDelisted)
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
		it.Event = new(MarketplaceNFTDelisted)
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
func (it *MarketplaceNFTDelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTDelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTDelisted represents a NFTDelisted event raised by the Marketplace contract.
type MarketplaceNFTDelisted struct {
	Seller  common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNFTDelisted is a free log retrieval operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) FilterNFTDelisted(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTDelistedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTDelisted", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTDelistedIterator{contract: _Marketplace.contract, event: "NFTDelisted", logs: logs, sub: sub}, nil
}

// WatchNFTDelisted is a free log subscription operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) WatchNFTDelisted(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTDelisted, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTDelisted", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTDelisted)
				if err := _Marketplace.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
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

// ParseNFTDelisted is a log parse operation binding the contract event 0x62d66f5d387e08db35989a64daa510d22a7232d53039c6539e9bf747845f83e7.
//
// Solidity: event NFTDelisted(address indexed seller, uint256 indexed tokenId)
func (_Marketplace *MarketplaceFilterer) ParseNFTDelisted(log types.Log) (*MarketplaceNFTDelisted, error) {
	event := new(MarketplaceNFTDelisted)
	if err := _Marketplace.contract.UnpackLog(event, "NFTDelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTListedIterator is returned from FilterNFTListed and is used to iterate over the raw logs and unpacked data for NFTListed events raised by the Marketplace contract.
type MarketplaceNFTListedIterator struct {
	Event *MarketplaceNFTListed // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTListed)
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
		it.Event = new(MarketplaceNFTListed)
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
func (it *MarketplaceNFTListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTListed represents a NFTListed event raised by the Marketplace contract.
type MarketplaceNFTListed struct {
	Seller        common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNFTListed is a free log retrieval operation binding the contract event 0x36706a4b50544f20fc0b29aab34b81b81fe61145c673a6e73c4278ba90f6855d.
//
// Solidity: event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) FilterNFTListed(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTListedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTListed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTListedIterator{contract: _Marketplace.contract, event: "NFTListed", logs: logs, sub: sub}, nil
}

// WatchNFTListed is a free log subscription operation binding the contract event 0x36706a4b50544f20fc0b29aab34b81b81fe61145c673a6e73c4278ba90f6855d.
//
// Solidity: event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) WatchNFTListed(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTListed, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTListed", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTListed)
				if err := _Marketplace.contract.UnpackLog(event, "NFTListed", log); err != nil {
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

// ParseNFTListed is a log parse operation binding the contract event 0x36706a4b50544f20fc0b29aab34b81b81fe61145c673a6e73c4278ba90f6855d.
//
// Solidity: event NFTListed(address indexed seller, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) ParseNFTListed(log types.Log) (*MarketplaceNFTListed, error) {
	event := new(MarketplaceNFTListed)
	if err := _Marketplace.contract.UnpackLog(event, "NFTListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTListingPriceUpdatedIterator is returned from FilterNFTListingPriceUpdated and is used to iterate over the raw logs and unpacked data for NFTListingPriceUpdated events raised by the Marketplace contract.
type MarketplaceNFTListingPriceUpdatedIterator struct {
	Event *MarketplaceNFTListingPriceUpdated // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTListingPriceUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTListingPriceUpdated)
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
		it.Event = new(MarketplaceNFTListingPriceUpdated)
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
func (it *MarketplaceNFTListingPriceUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTListingPriceUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTListingPriceUpdated represents a NFTListingPriceUpdated event raised by the Marketplace contract.
type MarketplaceNFTListingPriceUpdated struct {
	Seller   common.Address
	TokenId  *big.Int
	NewPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNFTListingPriceUpdated is a free log retrieval operation binding the contract event 0xd3ce01f4cdf64c0873f65eb9c463c495d22b08cb09e73ee5116515360a524c53.
//
// Solidity: event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) FilterNFTListingPriceUpdated(opts *bind.FilterOpts, seller []common.Address, tokenId []*big.Int) (*MarketplaceNFTListingPriceUpdatedIterator, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTListingPriceUpdated", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTListingPriceUpdatedIterator{contract: _Marketplace.contract, event: "NFTListingPriceUpdated", logs: logs, sub: sub}, nil
}

// WatchNFTListingPriceUpdated is a free log subscription operation binding the contract event 0xd3ce01f4cdf64c0873f65eb9c463c495d22b08cb09e73ee5116515360a524c53.
//
// Solidity: event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) WatchNFTListingPriceUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTListingPriceUpdated, seller []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTListingPriceUpdated", sellerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTListingPriceUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "NFTListingPriceUpdated", log); err != nil {
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

// ParseNFTListingPriceUpdated is a log parse operation binding the contract event 0xd3ce01f4cdf64c0873f65eb9c463c495d22b08cb09e73ee5116515360a524c53.
//
// Solidity: event NFTListingPriceUpdated(address indexed seller, uint256 indexed tokenId, uint256 newPrice)
func (_Marketplace *MarketplaceFilterer) ParseNFTListingPriceUpdated(log types.Log) (*MarketplaceNFTListingPriceUpdated, error) {
	event := new(MarketplaceNFTListingPriceUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "NFTListingPriceUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceNFTPurchasedIterator is returned from FilterNFTPurchased and is used to iterate over the raw logs and unpacked data for NFTPurchased events raised by the Marketplace contract.
type MarketplaceNFTPurchasedIterator struct {
	Event *MarketplaceNFTPurchased // Event containing the contract specifics and raw log

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
func (it *MarketplaceNFTPurchasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceNFTPurchased)
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
		it.Event = new(MarketplaceNFTPurchased)
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
func (it *MarketplaceNFTPurchasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceNFTPurchasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceNFTPurchased represents a NFTPurchased event raised by the Marketplace contract.
type MarketplaceNFTPurchased struct {
	Buyer         common.Address
	TokenId       *big.Int
	Price         *big.Int
	PaymentMethod uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNFTPurchased is a free log retrieval operation binding the contract event 0x6af81c051535566e6f7c39c176ce0c7472ce9de1d4fa0c6e7101e567ce5d8568.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) FilterNFTPurchased(opts *bind.FilterOpts, buyer []common.Address, tokenId []*big.Int) (*MarketplaceNFTPurchasedIterator, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "NFTPurchased", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceNFTPurchasedIterator{contract: _Marketplace.contract, event: "NFTPurchased", logs: logs, sub: sub}, nil
}

// WatchNFTPurchased is a free log subscription operation binding the contract event 0x6af81c051535566e6f7c39c176ce0c7472ce9de1d4fa0c6e7101e567ce5d8568.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) WatchNFTPurchased(opts *bind.WatchOpts, sink chan<- *MarketplaceNFTPurchased, buyer []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "NFTPurchased", buyerRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceNFTPurchased)
				if err := _Marketplace.contract.UnpackLog(event, "NFTPurchased", log); err != nil {
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

// ParseNFTPurchased is a log parse operation binding the contract event 0x6af81c051535566e6f7c39c176ce0c7472ce9de1d4fa0c6e7101e567ce5d8568.
//
// Solidity: event NFTPurchased(address indexed buyer, uint256 indexed tokenId, uint256 price, uint8 paymentMethod)
func (_Marketplace *MarketplaceFilterer) ParseNFTPurchased(log types.Log) (*MarketplaceNFTPurchased, error) {
	event := new(MarketplaceNFTPurchased)
	if err := _Marketplace.contract.UnpackLog(event, "NFTPurchased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Marketplace contract.
type MarketplaceOwnershipTransferredIterator struct {
	Event *MarketplaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketplaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceOwnershipTransferred)
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
		it.Event = new(MarketplaceOwnershipTransferred)
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
func (it *MarketplaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceOwnershipTransferred represents a OwnershipTransferred event raised by the Marketplace contract.
type MarketplaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketplaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceOwnershipTransferredIterator{contract: _Marketplace.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Marketplace *MarketplaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketplaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceOwnershipTransferred)
				if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseOwnershipTransferred(log types.Log) (*MarketplaceOwnershipTransferred, error) {
	event := new(MarketplaceOwnershipTransferred)
	if err := _Marketplace.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePriceInUSDIterator is returned from FilterPriceInUSD and is used to iterate over the raw logs and unpacked data for PriceInUSD events raised by the Marketplace contract.
type MarketplacePriceInUSDIterator struct {
	Event *MarketplacePriceInUSD // Event containing the contract specifics and raw log

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
func (it *MarketplacePriceInUSDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePriceInUSD)
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
		it.Event = new(MarketplacePriceInUSD)
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
func (it *MarketplacePriceInUSDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePriceInUSDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePriceInUSD represents a PriceInUSD event raised by the Marketplace contract.
type MarketplacePriceInUSD struct {
	ListingId  *big.Int
	PriceInWei *big.Int
	PriceInUSD *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceInUSD is a free log retrieval operation binding the contract event 0xdb0f2ab6fb984b2a281e04bf4dd6396434cf21a3c357e7584d76255e00c48d61.
//
// Solidity: event PriceInUSD(uint256 indexed listingId, uint256 priceInWei, uint256 priceInUSD)
func (_Marketplace *MarketplaceFilterer) FilterPriceInUSD(opts *bind.FilterOpts, listingId []*big.Int) (*MarketplacePriceInUSDIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PriceInUSD", listingIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePriceInUSDIterator{contract: _Marketplace.contract, event: "PriceInUSD", logs: logs, sub: sub}, nil
}

// WatchPriceInUSD is a free log subscription operation binding the contract event 0xdb0f2ab6fb984b2a281e04bf4dd6396434cf21a3c357e7584d76255e00c48d61.
//
// Solidity: event PriceInUSD(uint256 indexed listingId, uint256 priceInWei, uint256 priceInUSD)
func (_Marketplace *MarketplaceFilterer) WatchPriceInUSD(opts *bind.WatchOpts, sink chan<- *MarketplacePriceInUSD, listingId []*big.Int) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PriceInUSD", listingIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePriceInUSD)
				if err := _Marketplace.contract.UnpackLog(event, "PriceInUSD", log); err != nil {
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

// ParsePriceInUSD is a log parse operation binding the contract event 0xdb0f2ab6fb984b2a281e04bf4dd6396434cf21a3c357e7584d76255e00c48d61.
//
// Solidity: event PriceInUSD(uint256 indexed listingId, uint256 priceInWei, uint256 priceInUSD)
func (_Marketplace *MarketplaceFilterer) ParsePriceInUSD(log types.Log) (*MarketplacePriceInUSD, error) {
	event := new(MarketplacePriceInUSD)
	if err := _Marketplace.contract.UnpackLog(event, "PriceInUSD", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Marketplace contract.
type MarketplaceUpgradedIterator struct {
	Event *MarketplaceUpgraded // Event containing the contract specifics and raw log

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
func (it *MarketplaceUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceUpgraded)
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
		it.Event = new(MarketplaceUpgraded)
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
func (it *MarketplaceUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceUpgraded represents a Upgraded event raised by the Marketplace contract.
type MarketplaceUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Marketplace *MarketplaceFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*MarketplaceUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceUpgradedIterator{contract: _Marketplace.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Marketplace *MarketplaceFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *MarketplaceUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceUpgraded)
				if err := _Marketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Marketplace *MarketplaceFilterer) ParseUpgraded(log types.Log) (*MarketplaceUpgraded, error) {
	event := new(MarketplaceUpgraded)
	if err := _Marketplace.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
