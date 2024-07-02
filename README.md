# github.com/zostay/replacer

This is a drop in replacement for `strings.Replacer` which misbehaves in at least on edge case:

```go
package main

import (
    "fmt"
	"strings"
)

func main() {
	// DO NOT DO THIS: This will give you inconsistent results.
	r := strings.NewReplacer(
	  "modifiedbyid", "modified_by_id",
	  "modifiedby", "modified_by")
	fmt.Println(r.Replace("modifiedbyid"))
	// Output may be: modified_by_id
	//            OR: modified_byid
}
```

If you replace the above with this library, your behavior will always be consistent:

```go
package main

import (
	"fmt"
	"github.com/zostay/replacer"
)

func main() {
	// DO THIS! This will always give you the same result.
	r := replacer.New(
		"modifiedbyid", "modified_by_id",
		"modifiedby", "modified_by")
	fmt.Println(r.Replace("modifiedbyid"))
	// Output will always be: modified_by_id
}
```

This module always selects the longest match.

## TODO

This is not a complete replacement for `strings.Replacement`. The following is
needed to be a complete replacement:

* Implement `WriteString(w io.Writer, s string) (n int, err error)`

## Caveat Emptor

This library is a first release. As of this writing, I cooked up and
released it in a couple hours to meet the needs of a client. It may not perform
well and while it has 100% test coverage, that doesn't mean I know what it will
do in very edge case.

# LICENSE

Copyright 2024 Andrew Sterling Hanenkamp

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
