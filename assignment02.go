package assignment02

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {
	random_num := rand.Intn(10000)
	temp_hash := CalculateHash(blockData, random_num)
	hx := hex.EncodeToString([]byte(temp_hash))
	num, err := strconv.ParseInt(hx[0:4], 16, 16)
	if err != nil {
		fmt.Println(err)
		fmt.Println("\nError in Hexadecimal - Decimal Conversion!!!")
		return random_num
	}

	return int(num)
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	mynonce := GenerateNonce(blockData)
	current_hash := CalculateHash(blockData, mynonce)
	previous_hash := ""
	if chainHead == nil {
		previous_hash = "-"
	} else {
		previous_hash = CalculateHash(chainHead.BlockData, chainHead.Nonce)
	}

	myblock := Block{Nonce: mynonce, BlockData: blockData, PrevPointer: chainHead, PrevHash: previous_hash, CurrentHash: current_hash}
	return &myblock
}

func ListBlocks(chainHead *Block) {
	if chainHead == nil {
		fmt.Println("Blockchain Empty")
		return
	}

	for {
		fmt.Print("\nNonce : ")
		fmt.Println(chainHead.Nonce)
		fmt.Print("Current Hash : ")
		fmt.Println(chainHead.CurrentHash)
		fmt.Print("Previous Hash : ")
		fmt.Println(chainHead.PrevHash)
		fmt.Print("Previous Pointer : ")
		fmt.Println(chainHead.PrevPointer)
		fmt.Println("::::: Transactions :::::")
		DisplayTransactions(chainHead.BlockData)
		fmt.Print("::::: ---------------- :::::\n")

		if chainHead.PrevPointer == nil {
			break
		}
		chainHead = chainHead.PrevPointer
	}
}

func DisplayTransactions(blockData []Transaction) {
	for i := 0; i < len(blockData); i++ {
		fmt.Printf("Transaction #%d\n", i+1)
		fmt.Println("Transaction : ", blockData[i])
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	return Transaction{TransactionID: "1", Sender: sender, Receiver: receiver, Amount: amount}
}
