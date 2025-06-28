regenerate:
	@bash ./scripts/fetch-latest-protocol.sh && \
	echo 'Regenerating...\n' && \
	go run ./cmd/protocol-gen || \
	echo "Skipped code generation (no changes)."
