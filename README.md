## Coolog

### Install

```
go get github.com/5046312/coolog
```

### Usage

```
import "github.com/5046312/coolog"
...
// Set Log Config
conf := coolog.FileConfig()
conf.Single = true
conf.Ext = ".bin"
conf.Path = "./runtime/lll/"
coolog.SetFileLog(fc)

// Write Log
coolog.Debug("Write Debug in file")
```

Or call directly, which will operate using the default configuration of the system


```
import "github.com/5046312/coolog"
...
// Not Set Config, Write Log With Default Config
coolog.Debug("Write Debug in file")
```