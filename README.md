# go-examples
[![Build Status](https://img.shields.io/travis/romain-cotte/go-examples/master.svg?style=flat-square)](https://travis-ci.org/romain-cotte/go-examples)

Variables:
```
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go/gopath
```

Installation:
```
go get gopkg.in/pg.v5
go get github.com/spf13/viper
```

To format correctly, hit
```
go fmt ./...
```

## Tests

Functions starting with "Example" are checked against Output comments whereas
unit tests using testing module must start with "Test" and are not checked
with the Output comments.
