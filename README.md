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
    a. ```go get -v -u github.com/gen2brain/raylib-go/raylib```
### Setting up Environment Variables
1. Create a file named ```secrets.txt``` in the root directory of the project.
2. Enter ```HACKMAN_API_KEY=key``` and replace key with your Clemson Hackman API key.
### Cross Compiling for Windows
1. Install mingw-w64
2. run ```make build_windows``` in the root directory of the project.