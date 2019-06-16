all: install

install: go.sum
		GO111MODULE=on go install ./cmd/nftcli
		GO111MODULE=on go install ./cmd/nftd

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		# GO111MODULE=on go mod verify
