package main

import (
	"sync"
)

var blockOrderCache = struct {
	//m map[blockId]
	m map[uint64]*blockFragment
	sync.RWMutex
}{m: make(map[uint64]*blockFragment)}

var blockCommitCache = struct {
	m map[uint64]*blockFragment
	sync.RWMutex
}{m: make(map[uint64]*blockFragment)}

type blockFragment struct {
	sync.RWMutex
	hashOfEntriesInBlock []byte
	//Collect postEntry replies and send orderEntry
	//cryptSigOfLeaderEntry	*[]byte //Use cryptSigOfLeaderPostEntry as a base in reply

	entriesInBlock map[int]Entry

	//The coordinator of the order phase accumulates threshold
	//signatures from others, append it to orderedThreshSig,
	//which is going to be converted to one threshold signature
	//by function PenRecovery.
	concatThreshSig [][]byte

	//counter increments when receiving a postReply.
	//When 2f+1 identical postReplies received, 2f+1 servers have
	//agreed on the posted order.
	counter int
}

var vcCache = struct {
	cacheLock sync.RWMutex
	blockLock sync.RWMutex
	cache     map[ViewNumber]vcBlockFragment
	block     map[ViewNumber]vcBlockPermanent
}{
	cache: make(map[ViewNumber]vcBlockFragment),
	block: make(map[ViewNumber]vcBlockPermanent),
}

type vcBlockFragment struct {
	//used by voters during the election
	voted             bool
	votedFor          ServerId
	commonEndorsement []byte

	//used by candidate
	concatThreshSigOfVotes [][]byte
	counter                int
}

func constructCommonEndorsement(voteMeMsg VoteMe) []byte {
	// endorse can be changed differently
	// endorse cannot include partial signature!
	endorse := VoteMe{
		ServerId:               voteMeMsg.ServerId,
		View:                   voteMeMsg.View,
		LatestBlockId:          voteMeMsg.LatestBlockId,
		LatestBlockHash:        voteMeMsg.LatestBlockHash,
		LatestBlockCertificate: voteMeMsg.LatestBlockCertificate,
		RepPenaltyScore:        voteMeMsg.RepPenaltyScore,
		FoundNonce:             voteMeMsg.FoundNonce,
		PartialSignature:       nil,
	}

	b, err := getHashOfMsg(endorse)
	if err != nil {
		log.Errorf("getHashOfMsg failed | err: %v", err)
		return nil
	}

	return b
}

// vcBlockPermanent is the view-change block, which is updated upon new view is formed
// and keeps track of servers reputation penalty
type vcBlockPermanent struct {
	sync.RWMutex
	leaderId          ServerId
	reputationPenalty map[ServerId]int
	consumedTxBlocks  map[ServerId]int
	newViewMsg        NewView
}
