package main

type BlockHeader struct {
	prevBlockHash [32]byte
	thisBlockHash [32]byte
}

type serverId int

type TxBlock struct {
	blockHeader BlockHeader
	n           uint64
	v           uint64
	orderingQC  []byte
	commitQC    []byte
	tx          map[uint64][]byte
}

type VcBlock struct {
	v        uint64
	leaderId serverId
	nonce    []byte
	vcQC     []byte
	penalty  map[serverId]int
	utxIndex map[serverId]uint64
}
