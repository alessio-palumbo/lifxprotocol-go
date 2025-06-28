package main

import (
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/alessio-palumbo/lifxprotocol-go/decode"
	"github.com/alessio-palumbo/lifxprotocol-go/generate"
)

const (
	protocolGenerateDir = "gen/protocol"
)

//go:embed src/protocol.yml
var protocolYAML []byte

//go:embed src/protocol_commit.txt
var protocolCommit []byte

func main() {
	sourceCommit := strings.TrimSpace(string(protocolCommit))
	if sourceCommit == "" {
		log.Fatal("Source commit not found")
	}

	protocolSpec, err := decode.DecodeProtocol(protocolYAML)
	if err != nil {
		log.Fatalf("Failed to parse embedded protocol.yml: %v", err)
	}

	if err := generate.GenerateProtocol(protocolSpec, protocolGenerateDir, sourceCommit); err != nil {
		log.Fatalf("Failed to generate Protocol: %v", err)
	}

	fmt.Println("Code generation completed successfully.")
}
