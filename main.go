package main

import (
	_ "embed"
	"fmt"
	"log"

	"github.com/alessio-palumbo/lifxprotocol-go/decode"
	"github.com/alessio-palumbo/lifxprotocol-go/generate"
)

const (
	protocolGenerateDir = "gen/protocol"
)

//go:embed src/protocol.yml
var protocolYAML []byte

func main() {
	protocolSpec, err := decode.DecodeProtocol(protocolYAML)
	if err != nil {
		log.Fatalf("Failed to parse embedded protocol.yml: %v", err)
	}

	if err := generate.GenerateProtocol(protocolSpec, protocolGenerateDir); err != nil {
		log.Fatalf("Failed to generate Protocol: %v", err)
	}

	fmt.Println("Code generation completed successfully.")
}
