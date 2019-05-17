## Coolog

### Install

```
go get github.com/5046312/coolog
```

### Usage

```
import "github.com/5046312/coolog"
...
fc := coolog.FileConfig()
logger := coolog.NewFileLog(fc)
logger.Debug("Write Debug in file")
```