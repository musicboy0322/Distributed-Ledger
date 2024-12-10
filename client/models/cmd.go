type CMD3Message struct {
	Command string
	FromWallet string
	ToWallet string
	Amount int 
}

type CMD5Message struct {
	Command string
	Sha256Content string
}