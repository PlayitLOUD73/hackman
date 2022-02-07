# hackman
A hangman game written in Go and Raylib.

### Setting up Environment Variables
The application needs an API key to run.
To setup the application run<br />```export HACKMAN_API_KEY=YOUR_KEY```<br />
This will work as long as the shell is active.

### Cross Compiling for Windows
1. Install mingw-w64
2. Build by running <br />```GOOS=windows CGO_ENABLED=1 CC=/usr/local/bin/x86_64-w64-mingw32-gcc go build -o hackman.exe```<br />