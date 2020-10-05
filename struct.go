package assignment01IBC

type Block struct {
	transactions []string
	prevPointer  *Block
	prevHash     string
	currentHash  string
}


