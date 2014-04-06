#Unbeatable Tic Tac Toe in Go

## Setup
If you don't have it already, install Go 1.2 by running:

    ```
    brew install go
    ```

set the GOPATH -- make sure your .profile or .bashprofile is updated as well

```
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
```

for additional info on setting the GOPATH, go [here](http://golang.org/doc/code.html#GOPATH)

## Project Installation

You should now be able to run the following commands from command line:

###Get the project

    ```
    go get github.com/lnew78/ttt-go
    ```

###Get the testing framework

    ```
    go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
    go get github.com/onsi/gomega         # fetches the matcher library
    ```

##Running the Game
simply enter the following from anywhere on the command line:
   ```
   ttt-go
   ```

