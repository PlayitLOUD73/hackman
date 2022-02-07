CROSS_COMPILER=/usr/local/bin/x86_64-w64-mingw32-gcc
WINDOWS_EXECUTABLE=hackman.exe
EXECUTABLE=hackman

build:
	go build -o ${EXECUTABLE}

build_release:
	mkdir releases
	GOOS=windows CGO_ENABLED=1 CC=${CROSS_COMPILER} go build -o releases/${WINDOWS_EXECUTABLE}
	go build -o releases/${EXECUTABLE}

run:
	./hangman

clean:
	go clean
	rm -rf releases