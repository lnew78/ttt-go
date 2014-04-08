#Unbeatable Tic Tac Toe in Go

## Setup

If you don't have it already, install Go 1.2 by running the command below.

*Note* If you have a previous version of Go, you will need to [uninstall](http://golang.org/doc/install#uninstall) it first

```Go
brew install go
```

Set the GOPATH -- make sure your .profile or .bashprofile is updated as well

    ```
    $ mkdir $HOME/go
    $ export GOPATH=$HOME/go
    $ export PATH=$PATH:$GOPATH/bin
    ```

For additional info on setting the GOPATH, go [here](http://golang.org/doc/code.html#GOPATH)

## Project Installation

You should now be able to run the following commands from command line:

###Get the project

    ```
    $ go get github.com/lnew78/ttt-go/ttt

    $ cd $GOPATH/src/github.com/lnew78/ttt-go

    $ go install
    ```

###Get the testing framework

    ```
    go get github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
    go get github.com/onsi/gomega         # fetches the matcher library
    ```

##Running the Game

Simply enter the following from anywhere on the command line:

    ```
    ttt-go
    ```

##Running the tests

From the project root, run:

    ```
    ginkgo -v ttt      # for verbose
    ```

    or

    ```
    ginkgo ttt         # just for test results
    ```
