# hackman
A hangman game written in Go and Raylib.

## Running
Download and extract the zip file for your operating system and architecture.
Double click to run.

## Development/ Compiling from Source
For local development and testing, a [Clemson Hackman API key](https://github.com/Jay-Madden/hackmanapi) is needed.
### Download necessary tools
1. Download and install [Go 1.17.6](https://go.dev/)
2. Download and install [raylib-go](https://github.com/gen2brain/raylib-go)
    1. ```go get -v -u github.com/gen2brain/raylib-go/raylib```
### Setting up Environment Variables
1. Create a file named ```secrets.txt``` in ```PROJECTROOT/cmd/main```.
2. Enter ```HACKMAN_API_KEY=key``` at the top of the file and replace key with your Clemson Hackman API key.
### Cross-compiling for Windows
1. Install [Docker](https://www.docker.com/).
2. Run ```docker build -t hackman-builder:latest .``` in the project root.
3. Run ```docker run --rm -v ${PWD}:/src hackman-builder:latest make build_windows```
