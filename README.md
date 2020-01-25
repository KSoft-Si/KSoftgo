# KSoftgo

## Getting Started
### Installing
```sh
go get github.com/Noctember/KSoftgo
```

## Usage
Import the package into your project
```go
import "github.com/Noctember/KSoftgo"
```

Construct a new KSoft session, this will be needed to access the API
```go
session, err := ksoftgo.New("fancy-token-here")
```

