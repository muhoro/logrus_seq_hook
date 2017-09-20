<img src="https://avatars2.githubusercontent.com/u/5023402?v=4&s=200" width="40" height="40"/>

# SEQ Hooks for [Logrus](https://github.com/sirupsen/logrus) <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:"/>

## Install

```shell
$ go get github.com/mundipagg/logrus_seq_hook
```

## Usage

```go
package main

import (
	. "github.com/mundipagg/logrus_seq_hook"
	"github.com/sirupsen/logrus"
)

func main() {
	seqHook, _ := NewSeqHook("http://localhost:5341", "API-TOKEN")
	logrus.AddHook(seqHook)
	logrus.WithFields(logrus.Fields{"key": "value"}).Info("Logging on SEQ")
	logrus.WithFields(logrus.Fields{"key": "value"}).Warn("Logging on SEQ")
	logrus.WithFields(logrus.Fields{"key": "value"}).Fatal("Logging on SEQ")

}
```

## License
*MIT*
