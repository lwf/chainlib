# chainlib

chainlib is a library of helper functions for writing utilities which pass
state to other binaries using [chain loading](https://en.wikipedia.org/wiki/Chain_loading), 
a technique popularized by [DJB](http://cr.yp.to/).

## Minimal example
```go
package main

import (
	"time"

	"github.com/lwf/chainlib"
)

func realMain() (map[string]string, error) {
	return map[string]string{
		"time": time.Now().String(),
	}, nil
}

func main() {
	chainlib.Main("MYNAMESPACE", realMain)
}
```

```shell
# ./example /bin/sh -c 'echo $MYNAMESPACE_TIME'
2015-09-03 23:53:11.382559551 +0200 CEST
```

## License

Copyright 2015 Torbjörn Norinder

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
