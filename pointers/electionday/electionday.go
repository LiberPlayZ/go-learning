package electionday

import (
	"fmt"
)

type ElectionResult struct {
	// Name of the candidate
	Name string
	// Votes of votes the candidate had
	Votes int
}

// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {
	var t int = initialVotes
	var p *int = &t
	return p
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}
	return *counter
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	if counter != nil {
		*counter = *counter + increment
	}
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {
	var election = &ElectionResult{Name: candidateName, Votes: votes}
	return election
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	s := fmt.Sprintf("%s (%d)", result.Name, result.Votes)
	return s
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for k := range results {
		if k == candidate {
			results[k] = results[k] - 1
		}
	}
}
