package common

import (
	"errors"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

const (
	hexadecimalNumber = 16
)

func GetMethodName(abi *abi.ABI, inputData string) (string, error) {
	if len(inputData) < 10 {
		return "", errors.New("InputData invalid length")
	}

	methodSig := inputData[:10]
	method, err := abi.MethodById(common.FromHex(methodSig))
	if err != nil {
		return "", errors.Unwrap(err)
	}
	return method.Name, nil
}
