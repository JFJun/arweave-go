package api

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/JFJun/arweave-go/utils"
	"testing"
)

var client, _ = Dial("http://54.238.245.19:1984")

func TestClient_GetBlockByHeight(t *testing.T) {
	block, err := client.GetBlockByHeight(context.Background(), 633720)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
	fmt.Println(len(block.Txs))
}

func TestClient_GetTransaction(t *testing.T) {
	tx, err := client.GetTransaction(context.Background(), "l9H7CGO3RH27fjG6K1gm3aK9n4DIpDaGz8KdyhCYTxM")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(tx)
	fmt.Println(string(d))
	if "" == tx.Target() {
		panic("假充值")
	}
	data, err := base64.RawURLEncoding.DecodeString(tx.Owner()) //
	h := sha256.New()
	h.Write(data)
	from := utils.EncodeToBase64(h.Sum(nil))
	fmt.Println(from)
	//转到哪而去
	to := tx.Target()
	fmt.Println(to)
	txid := tx.Hash()
	fmt.Println(txid)
}
