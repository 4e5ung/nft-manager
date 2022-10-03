package service

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"manage-template/config"
	"manage-template/contracts/build"
	"math/big"
	"os"
	"time"
)

type TokenApiService struct{}

func (m TokenApiService) TransferTokenCancel(from string, to string, value big.Int) (common.Hash, error) {

	hash := common.HexToHash("")

	client, err := ethclient.Dial(config.EdgeTestnet.RpcURL)
	if err != nil {
		fmt.Printf("TransferTokenCancel: Dial(%s)\n", err.Error())
		return hash, err
	}

	var address = common.HexToAddress(config.EdgeTestnet.SmartContractAddress)

	instance, err := build.NewManage(address, client)
	if err != nil {
		fmt.Printf("TransferTokenCancel: NewManage(%s)\n", err.Error())
		return hash, err
	}

	pvKeyECDSA, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("TransferTokenCancel: HexToECDSA(%s)\n", err.Error())
		return hash, err
	}

	gasLimit := uint64(9999999)
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("TransferTokenCancel: SuggestGasPrice(%s)\n", err.Error())
		return hash, err
	}

	fmt.Printf("gasLimit(%v)  suggestedGasPrice(%v)\n", gasLimit, suggestedGasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(pvKeyECDSA, big.NewInt(config.EdgeTestnet.ChainId))

	tx, err := instance.TransferToken(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: suggestedGasPrice,
		GasLimit: gasLimit,
	}, common.HexToAddress(from), common.HexToAddress(to), &value)

	if err != nil {
		fmt.Printf("TransferTokenCancel: TransferToken(%s)\n", err.Error())
		return hash, err
	}

	newTx, err := cancelTransaction(client, tx, pvKeyECDSA)

	if err != nil {
		fmt.Printf("TransferTokenCancel: cancelTransaction(%s)\n", err.Error())
		return hash, err
	}

	return newTx.Hash(), nil
}

func (m TokenApiService) TransferTokenPolling(from string, to string, value big.Int) (common.Hash, uint64, error) {

	hash := common.HexToHash("")

	client, err := ethclient.Dial(config.EdgeTestnet.RpcURL)
	if err != nil {
		fmt.Printf("TransferTokenPolling: Dial(%s)\n", err.Error())
		return hash, 0, err
	}

	var address = common.HexToAddress(config.EdgeTestnet.SmartContractAddress)

	instance, err := build.NewManage(address, client)
	if err != nil {
		fmt.Printf("TransferTokenPolling: NewManage(%s)\n", err.Error())
		return hash, 0, err
	}

	pvKeyECDSA, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("TransferTokenPolling: HexToECDSA(%s)\n", err.Error())
		return hash, 0, err
	}

	gasLimit := uint64(9999999)
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("TransferTokenPolling: SuggestGasPrice(%s)\n", err.Error())
		return hash, 0, err
	}

	fmt.Printf("gasLimit(%v)  suggestedGasPrice(%v)\n", gasLimit, suggestedGasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(pvKeyECDSA, big.NewInt(config.EdgeTestnet.ChainId))

	fmt.Printf("value:%s, from:%s, to:%s \n", value.String(), from, to)

	transaction, err := instance.TransferToken(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: suggestedGasPrice,
		GasLimit: gasLimit,
	}, common.HexToAddress(from), common.HexToAddress(to), &value)

	status, err := waitTransaction(context.Background(), client, transaction)

	if err != nil {
		fmt.Printf("TransferTokenPolling: waitTransaction(%s)\n", err.Error())
		return hash, status, err
	}

	return transaction.Hash(), status, nil
}

func (m TokenApiService) TransferToken(from string, to string, value big.Int) (common.Hash, error) {

	hash := common.HexToHash("")

	client, err := ethclient.Dial(config.EdgeTestnet.RpcURL)
	if err != nil {
		fmt.Printf("TransferToken: Dial(%s)\n", err.Error())
		return hash, err
	}

	var address = common.HexToAddress(config.EdgeTestnet.SmartContractAddress)

	instance, err := build.NewManage(address, client)
	if err != nil {
		fmt.Printf("TransferToken: NewManage(%s)\n", err.Error())
		return hash, err
	}

	pvKeyECDSA, err := crypto.HexToECDSA(os.Getenv("PRIVATE_KEY"))
	if err != nil {
		fmt.Printf("TransferToken: HexToECDSA(%s)\n", err.Error())
		return hash, err
	}
	gasLimit := uint64(30000000)
	suggestedGasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Printf("TransferToken: SuggestGasPrice(%s)\n", err.Error())
		return hash, err
	}

	fmt.Printf("gasLimit(%v)  suggestedGasPrice(%v)\n", gasLimit, suggestedGasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(pvKeyECDSA, big.NewInt(config.EdgeTestnet.ChainId))

	fmt.Printf("value:%s, from:%s, to:%s \n", value.String(), from, to)

	transact, err := instance.TransferToken(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: suggestedGasPrice,
		GasLimit: gasLimit,
	}, common.HexToAddress(from), common.HexToAddress(to), &value)

	if err != nil {
		fmt.Printf("TransferToken: SuggestGasPrice(%s)\n", err.Error())
		return hash, err
	}

	return transact.Hash(), nil
}

func (m TokenApiService) BalanceToken(account string) (*big.Int, error) {

	client, err := ethclient.Dial(config.EdgeTestnet.RpcURL)
	if err != nil {
		fmt.Printf("BalanceToken: Dial(%s)\n", err.Error())
		return nil, err

	}

	var address = common.HexToAddress(config.EdgeTestnet.SmartContractAddress)

	instance, err := build.NewManage(address, client)
	if err != nil {
		fmt.Printf("BalanceToken: NewManage(%s)\n", err.Error())
		return nil, err
	}

	balance, err := instance.BalanceToken(&bind.CallOpts{}, common.HexToAddress(account))

	if err != nil {
		fmt.Printf("BalanceToken: BalanceToken(%s)\n", err.Error())
		return nil, err
	}

	return balance, nil
}

func (m TokenApiService) TransactionReceipt(hash string) (receipt *types.Receipt, err error) {

	client, err := ethclient.Dial(config.EdgeTestnet.RpcURL)
	if err != nil {
		fmt.Printf("TransactionReceipt: Dial(%s)\n", err.Error())
		return nil, err
	}

	receipt, err = client.TransactionReceipt(context.Background(), common.HexToHash(hash))

	if err != nil {
		fmt.Printf("TransactionReceipt: TransactionReceipt(%s)\n", err.Error())
		return nil, err
	}

	return receipt, err
}

func waitTransaction(ctx context.Context, client *ethclient.Client, tx *types.Transaction) (uint64, error) {
	queryTicker := time.NewTicker(time.Second * 1)

	defer queryTicker.Stop()

	for {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err == nil {
			return receipt.Status, nil
		}

		//if errors.Is(err, ethereum.NotFound) {
		//	fmt.Printf("TransactionReceipt(%s)\n", err.Error())
		//} else {
		//	fmt.Printf("TransactionReceipt(%s)\n", err.Error())
		//}

		fmt.Printf("TransactionReceipt(%s)\n", err.Error())

		// Wait for the next round.
		select {
		case <-ctx.Done():
			return 100, ctx.Err()
		case <-queryTicker.C:
		case <-time.After(time.Second * 10):
			return 101, nil
		}
	}
}

func cancelTransaction(client *ethclient.Client, transact *types.Transaction, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {

	calcPrice := big.NewInt(0).Div(transact.GasPrice(), big.NewInt(10))
	newGasPrice := big.NewInt(0).Add(transact.GasPrice(), calcPrice)

	fmt.Printf("newGasPrice: %d\n", newGasPrice)
	fmt.Printf("transact.Gas(): %d\n", transact.Gas())

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	var data []byte

	var newTx = types.NewTx(&types.LegacyTx{
		Nonce:    transact.Nonce(),
		GasPrice: newGasPrice,
		Gas:      transact.Gas(),
		To:       &fromAddress,
		Value:    big.NewInt(0),
		Data:     data,
	})

	signedTx, err := types.SignTx(newTx, types.NewEIP155Signer(big.NewInt(config.EdgeTestnet.ChainId)), privateKey)

	if err != nil {
		fmt.Printf("cancelTransaction: SignTx(%s)\n", err.Error())
		return nil, err
	}

	err = client.SendTransaction(context.Background(), signedTx)

	fmt.Printf("signedTx:%s\n", signedTx.Hash())

	if err != nil {
		fmt.Printf("cancelTransaction: SendTransaction(%s)\n", err.Error())
		return nil, err
	}

	return signedTx, nil
}
