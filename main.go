package main

import (
	"fmt"
	"time"

	"github.com/LarryBattle/nonce-golang"
	"github.com/google/uuid"
)

func uuidAsNonce() uuid.UUID {
	nonce := uuid.New()
	return nonce
}

func larrysNonce() string {
	nonce := nonce.NewToken()
	return nonce
}

func unixNanoAsNonce() int64 {
	nonce := time.Now().UnixNano()
	return nonce
}

var i int64 = 0

func incrementalNonce() int64 {
	nonce := i + 1
	return nonce
}

func main() {
	fmt.Printf("nonce (uuid): %v\n", uuidAsNonce())
	fmt.Printf("nonce (Larry): %v\n", larrysNonce())
	fmt.Printf("nonce (unixNano): %v\n", unixNanoAsNonce())
	fmt.Printf("nonce (incremental): %v\n", incrementalNonce())
}
