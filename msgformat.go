package main

type ClientProposeEntry struct {
	ClientId    int
	Timestamp   uint64
	Transaction []byte
	Macs        []byte
}

type LeaderPostEntry struct {
	BlockId          uint64
	PartialSignature []byte

	//HashOfBatchedEntries is the common input of the penSign function
	HashOfBatchedEntries []byte
	//Entries map[int]Entry

	//Private:
	hashedForSig []byte
}

type Entry struct {
	TimeStamp uint64
	ClientId  int
	Tx        []byte
}

//Sent to Coordinator 1
type WorkerRelayClientRequest struct {
	ClientProposeEntry
	Relay bool
}

// Sent to Coordinator 2
type WorkerPostReply struct {
	BlockId          uint64
	PartialSignature []byte //this is a threshold signature
}

type LeaderOrderEntry struct {
	BlockId            uint64
	CombinedSignatures []byte //this is a combined threshold signature
	Entries            map[int]Entry
}

// Sent to Coordinator 3
type WorkerOrderReply struct {
	BlockId          uint64
	PartialSignature []byte
}

type LeaderCommitEntry struct {
	BlockId            uint64
	CombinedSignatures []byte //this is a combined threshold signature
}

// ClientConfirm is sent to client
type ClientConfirm struct {
	Timestamp   uint64
	BlockId     uint64
	InBlockTxId int
}

// ViewChangeRequest is the only message type for all view-change messages.
// MsgType distinguishes the differences of requests, types are as follows:
// MsgType: 0 --> VoteMe, sent from the nominee to all workers;
// MsgType: 1 --> GrantVote, sent from workers to nominee;
// MsgType: 2 --> NewView, sent from the new leader (wining nominee) to all
type ViewChangeRequest struct {
	MsgType      int
	VoteMeMsg    VoteMe
	GrantVoteMsg GrantVote
	NewViewMsg   NewView
}

type VoteMe struct {
	ServerId               int
	View                   uint64
	LatestBlockId          uint64
	LatestBlockHash        []byte
	LatestBlockCertificate []byte

	// RepPenaltyScore is the threshold value
	RepPenaltyScore int
	FoundNonce      []byte

	//
	PartialSignature []byte
}

type GrantVote struct {
	ServerId         int
	View             uint64
	IVoteU           bool
	PartialSignature []byte
}

type NewView struct {
	ServerId          int
	View              uint64
	CombinedSignature []byte //converted threshold signatures, only one []byte
}
