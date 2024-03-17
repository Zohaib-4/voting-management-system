package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Vote struct {
	VoterID   int
	Candidate string
}

type Block struct {
	PrevHash    string
	CurrentHash string
	Votes       []Vote
}

var Blockchain []Block
var Candidates map[string]int
var Voters map[int]bool

func RegisterVoter(voterID int) {
	Voters[voterID] = false
	fmt.Printf("Voter %d registered.\n", voterID)
}

func CastVote(voterID int, candidate string) {
	if _, ok := Voters[voterID]; !ok {
		fmt.Println("Invalid voter ID: ", voterID)
		return
	}

	if Voters[voterID] {
		fmt.Println("Voter", voterID, "has already cast a vote.")
		return
	}

	if _, ok := Candidates[candidate]; !ok {
		fmt.Println("Candidate does not exist.")
		return
	}

	vote := Vote{VoterID: voterID, Candidate: candidate}
	lastBlock := Blockchain[len(Blockchain)-1]
	newBlock := Block{
		PrevHash: lastBlock.CurrentHash,
		Votes:    append(lastBlock.Votes, vote),
	}
	newBlock.CurrentHash = calculateHash(newBlock)
	Blockchain = append(Blockchain, newBlock)

	Candidates[candidate]++
	Voters[voterID] = true

	fmt.Printf("Vote cast by Voter %d for %s is recorded.\n", voterID, candidate)
}

func calculateHash(block Block) string {
	record := block.PrevHash
	for _, vote := range block.Votes {
		record += string(vote.VoterID) + vote.Candidate
	}
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func CalculateElectionResults() {
	fmt.Println("\nElection Results:")
	var winner string
	maxVotes := -1

	for candidate, votes := range Candidates {
		if votes > maxVotes {
			winner = candidate
			maxVotes = votes
		} else if votes == maxVotes {
			winner = "Tie"
		}
	}

	if winner != "Tie" {
		fmt.Printf("Winner: %s\n", winner)
	} else {
		fmt.Println("Election resulted in a tie.")
	}
}

func main() {
	genesisBlock := Block{PrevHash: "", CurrentHash: "", Votes: nil}
	Blockchain = append(Blockchain, genesisBlock)

	Candidates = make(map[string]int)
	Candidates["Candidate A"] = 0
	Candidates["Candidate B"] = 0

	Voters = make(map[int]bool)
	for i := 1; i <= 10; i++ {
		RegisterVoter(i)
	}

	CastVote(1, "Candidate A")
	CastVote(2, "Candidate B")
	CastVote(3, "Candidate A")
	CastVote(3, "Candidate B")
	CastVote(4, "Candidate B")
	CastVote(5, "Candidate A")
	CastVote(5, "Candidate A")
	CastVote(6, "Candidate B")
	CastVote(7, "Candidate C")
	CastVote(11, "Candidate B")

	CalculateElectionResults()

	fmt.Println("\nBlockchain:")
	for i, block := range Blockchain {
		fmt.Printf("Block %d\n", i)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("CurrentHash: %s\n", block.CurrentHash)
		fmt.Printf("Votes: %v\n\n", block.Votes)
	}
}
