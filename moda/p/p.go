package p

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func init() {
	fmt.Println("moda package init")
}

func ABI() *abi.ABI {
    fmt.Println("moda p call ABI")	
    return nil
}
