# hackman
A hangman game written in Go and Raylib for the Clemson Hackman Competition.
### Challenge
The challenge is simple, you must create a playable "hangman" game in any programming language/stack you'd like in 2 Weeks. 

## About the Game
All of the images in the game were made by me. Music and sound effects are from the public domain.

## Running
Download and extract the zip file for your operating system and architecture.
Double click to run.

## Development/ Compiling from Source
This tutorial assumes you are on macOS. If you are on Windows, see the Windows tutorial.
For local development and testing, a [Clemson Hackman API key](https://github.com/Jay-Madden/hackmanapi) is needed.
### Download necessary tools
1. Download and install [Go 1.17.6](https://go.dev/) or later.
2. Download and install [raylib-go](https://github.com/gen2brain/raylib-go)
    1. ```go get -v -u github.com/gen2brain/raylib-go/raylib```
### Setting up Environment Variables
1. Create a file named ```secrets.txt``` in ```PROJECTROOT/cmd/main```.
2. Enter ```HACKMAN_API_KEY=key``` at the top of the file and replace key with your Clemson Hackman API key.
### Cross-compiling for Windows
1. Install [Docker](https://www.docker.com/).
2. Run ```docker build -t hackman-builder:latest .``` in the project root.
3. Run ```docker run --rm -v ${PWD}:/src hackman-builder:latest make build_windows```
## Compiling on Windows (Not Tested)
1. Install [MSYS2](https://www.msys2.org/)
2. run in MSYS2 terminal: ```pacman -S mingw-w64-x86_64-go```
3. Setup Windows-specific enviroment variables.
    1. ```export GOROOT=/mingw64/lib/go```
    2. ```export GOPATH=/mingw64```
4. Open MinGW MSYS2 from start menu.
5. Change to project root.
6. Make sure to follow "Setting up Environment Variables" to add secrets.
7. Run ```make build_windows```
