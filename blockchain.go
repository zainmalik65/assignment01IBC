package assignment01IBC

import (
	"crypto/sha256"
	"fmt"
)

func CalculateHash(inputBlock *Block) string {

	hashes := sha256.New()
	hashes.Write([]byte(fmt.Sprintf("%v", inputBlock.transactions)))
	return fmt.Sprintf("%x", hashes.Sum(nil))
}

func InsertBlock(transactionsToInsert []string, chainHead *Block) *Block {


	var sum string
	for i:=0;i<len(transactionsToInsert);i++{
		sum+=transactionsToInsert[i]

	}
	hashsum :=sha256.Sum256([]byte(sum))
	var newBlock Block
	if chainHead==nil {
		newBlock= Block{transactionsToInsert, nil, "nil", string(hashsum[:])}
	}else{

		newBlock=Block{transactionsToInsert, chainHead, chainHead.currentHash, string(hashsum[:])}
		}

	fmt.Println("New Block is inserted into blockchain")
	return &newBlock



}

func ChangeBlock(oldTrans string, newTrans string, chainHead *Block) {


	newBlock:=chainHead
	check:=true
	for newBlock!=nil {

	 	for i:=0;i<len(newBlock.transactions);i++{
	 		if newBlock.transactions[i]==oldTrans{
	 			newBlock.transactions[i]=newTrans
	 			newBlock.currentHash=CalculateHash(newBlock)
				fmt.Println("Block value is changed")
	 			check=false
	 			break
			}

	 	}
		if check==false{
			break
		}
	 	newBlock=newBlock.prevPointer
	}
	if check==true {
		fmt.Println(oldTrans,"is not found in any block of current blockchain")
	}
	
}

func ListBlocks(chainHead *Block) {
	genesisBlock :=chainHead
	fmt.Println("Current blocks in the BlockChain are:")
	for genesisBlock!=nil{
		fmt.Println(genesisBlock.transactions)
		genesisBlock=genesisBlock.prevPointer
	}
}

func VerifyChain(chainHead *Block) {
	blockChainHead:=chainHead
	prevBlock:=blockChainHead.prevPointer
	checkInt:=true
	for prevBlock!=nil{

		if blockChainHead.prevHash!=prevBlock.currentHash {
			checkInt=false

			break
		}
		blockChainHead=prevBlock
		prevBlock=prevBlock.prevPointer


	}
	if checkInt==true {
		fmt.Println("Blockchain is intact")
	}else{
		fmt.Println("Blockchain is modified")
	}


}

