
tidy-all:
	cd app/cmd && go mod tidy
	cd libraries/libone && go mod tidy
	cd libraries/libtwo && go mod tidy
	go work sync

test-all:
	cd app/cmd && go test -race -cover ./...
	cd libraries/libone && go test -race -cover ./...
	cd libraries/libtwo && go test -race -cover ./...