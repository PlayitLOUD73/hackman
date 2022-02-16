CROSS_COMPILER=x86_64-w64-mingw32-gcc
WINDOWS_EXECUTABLE=hackman.exe
EXECUTABLE=hackman

build:
	go build -o ${EXECUTABLE}

setup:
	- mkdir releases
	
build_windows: setup
	GOOS=windows CGO_ENABLED=1 CC=${CROSS_COMPILER} CGO_LDFLAGS="-static-libgcc" go build -o releases/${WINDOWS_EXECUTABLE} ./cmd/main
	zip releases/hackman_windows.zip releases/${WINDOWS_EXECUTABLE}
	rm releases/${WINDOWS_EXECUTABLE}

build_darwin_intel:	setup
	GOOS=darwin CGO_ENABLED=1 go build -o releases/${EXECUTABLE} ./cmd/main
	zip releases/hackman_macos_intel.zip releases/${EXECUTABLE}
	rm releases/${EXECUTABLE}

build_darwin_arm: setup
	GOOS=darwin CGO_ENABLED=1 GOARCH=arm64 go build -o releases/${EXECUTABLE} ./cmd/main
	zip releases/hackman_macos_arm.zip releases/${EXECUTABLE}
	rm releases/${EXECUTABLE}

build_linux: setup
	GOOS=linux CGO_ENABLED=1 CC=gcc GOARCH=amd64 go build -o releases/${EXECUTABLE} ./cmd/main
	zip releases/hackman_linux.zip releases/${EXECUTABLE}
	rm releases/${EXECUTABLE}

build_release: build_windows build_darwin_intel build_darwin_arm

run: build
	./hangman

clean:
	go clean
	- rm -rf releases
	- rm .DS_Store