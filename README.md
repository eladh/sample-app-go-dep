# sample-app-go-dep
A sample [Go](https://golang.org/) application  using dep [dep](https://github.com/golang/dep)

This basic example uses a custom router and serves 2 endpoints
`/health` and `/` it uses the environment var `VERSION` defined in the
[manifest.yml](manifest.yml) the one later is returned when querying the
`/health` endpoint.

The application has the following structure:

    |-- app handlers
    |     `--hello
    |       `--hello.go
    |-- health.go
    |-- main.go

# How to use it

Install and set your [GO environment](https://golang.org/doc/install)

For example using $HOME/go for your workspace

    $ export GOPATH=$HOME/go

    $ git clone https://github.com/eladh/sample-app-go-dep.git

> [git](https://git-scm.com/) is required to clone the repository.

To test the app locally before deploying it , build it by just typing make
After make finishes you can run the app:

    $ ./app
    2018/07/06 13:39:12 Adding path: /health [ALL]
    2018/07/06 13:39:12 Adding path: / [ALL]

> This steps requires [dep](https://github.com/golang/dep)

Visit [http://localhost:8080](http://localhost:8080)


# How to convert to GO Modules

Import dependencies from Gopkg.lock

    $ go mod init .


Remove unnecessary imports and add indirect imports

    $ go mod tidy


Delete the vendor folder

    $ rm –rf vendor/

Delete Gopkg files

    $ rm Gopkg.*
    
    
Checkout gocenter GOPROXY performance

    $ run ./performance.sh  
