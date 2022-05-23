package main

import (
	"fmt"
	"time"
)
type Block struct{
	timestamp int64
	previousHash string
	transactions []string
	nounce int
}
func NewBlock(nounce int, hash string) *Block{
  b:=new(Block)
  b.nounce=nounce
  b.previousHash=hash
  b.timestamp=time.Now().UnixNano()
  return b
}
func (b *Block) print(){
	fmt.Println(b.nounce)
	fmt.Println(b.previousHash)
	fmt.Println(b.timestamp)
	fmt.Println(b.transactions)
}
func main(){
 b:=NewBlock(0,"Init State")
 b.print()
}