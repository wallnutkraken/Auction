# Install guide to be able to build this Server

## Setting up a Go environment

### Installing Go
 - Go to the [Golang website](https://golang.org/dl/)
 - Pick the package for your OS
 - Download and install Go

### PATH
 If you're using Windows, make sure to add Go's binaries to your PATH. [Here's an article that might help.](http://www.computerhope.com/issues/ch000549.htm)

 The path you want to add to your PATH is `[GoInstallDir]/bin`, where `[GoInstallDir]` is the directory where Go was installed.

### GOPATH

Along with setting your PATH variable, you need to set up a `GOPATH` variable for your environment where Go will store downloaded packages.

Use the [previous article](http://www.computerhope.com/issues/ch000549.htm) as reference and add a new environmental variable called `GOPATH`.
It should point to an empty directory of the user's choosing.

## Installing the Server

 - Open up a command line (git bash on Windows)
 - Type in `go get https://github.com/wallnutkraken/Auction` and press ENTER
 - Wait for the package to be downloaded
 - Navigate to `$GOPATH/src/github.com/wallnutkraken/Auction/Server`, where `$GOPATH` is your GOPATH variable
 - Type in `go build` and press ENTER