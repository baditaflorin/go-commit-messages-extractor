package main

import (
	"encoding/json"
	"flag"
	"fmt"
	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"log"
)

type CommitMessage struct {
	Hash    string `json:"hash,omitempty"`
	Message string `json:"message"`
}

func main() {
	// Define flags for input and output options
	repoPath := flag.String("repo", "", "Path to the local repository (required)")
	outputFormat := flag.String("format", "line", "Output format: 'line' for concatenated lines or 'json' for JSON output")
	includeHash := flag.Bool("hash", true, "Include commit hash in the output (true/false)")
	flag.Parse()

	if *repoPath == "" {
		log.Fatal("Repository path is required. Use --repo to specify the path.")
	}

	// Open the repository
	repo, err := git.PlainOpen(*repoPath)
	if err != nil {
		log.Fatalf("Failed to open repository: %v", err)
	}

	// Get the commit messages
	ref, err := repo.Head()
	if err != nil {
		log.Fatalf("Failed to get HEAD: %v", err)
	}

	commitIter, err := repo.Log(&git.LogOptions{From: ref.Hash()})
	if err != nil {
		log.Fatalf("Failed to retrieve commits: %v", err)
	}

	var commitMessages []CommitMessage

	err = commitIter.ForEach(func(c *object.Commit) error {
		cm := CommitMessage{
			Message: c.Message,
		}
		if *includeHash {
			cm.Hash = c.Hash.String()
		}
		commitMessages = append(commitMessages, cm)
		return nil
	})
	if err != nil {
		log.Fatalf("Failed to iterate over commits: %v", err)
	}

	// Output the commit messages in the desired format
	switch *outputFormat {
	case "line":
		for _, cm := range commitMessages {
			if *includeHash {
				fmt.Printf("%s: %s\n", cm.Hash, cm.Message)
			} else {
				fmt.Println(cm.Message)
			}
		}
	case "json":
		output, err := json.MarshalIndent(commitMessages, "", "  ")
		if err != nil {
			log.Fatalf("Failed to marshal commit messages to JSON: %v", err)
		}
		fmt.Println(string(output))
	default:
		log.Fatalf("Invalid output format: %s. Use 'line' or 'json'.", *outputFormat)
	}
}
