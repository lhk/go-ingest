## Building

`go build -o go-ingest`

You can also download a build from the tagged releases.

## Usage
The executable takes an arbitrary number of file patterns as command line parameters.
Here I'm piping the output to a file, to not flood the terminal:

`./go-ingest lklein/go-ingest/*.md lklein/go-ingest/*.go > output.txt`

On MacOS you can pipe to pbcopy and this fills your clipboard:

`./go-ingest <file pattern> | pbcopy`

output.txt will look like this
```
> cat output.txt
# Folder structure
└── lklein
    └── go-ingest
        ├── README.md
        └── main.go

# lklein/go-ingest/README.md
## Building

`go build -o go-ingest`

<... README inception :) ...>


# lklein/go-ingest/main.go
package main

import (
	"flag"
	"fmt"
	"os"

<... The go code of this repo ...>
```



