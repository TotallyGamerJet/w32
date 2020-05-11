About w32
==========

w32 is a wrapper of windows apis for the Go Programming Language.

It wraps win32 apis to "Go style" to make them easier to use.

## Notes 

This is a fork of [JamesHovious/w32](https://github.com/JamesHovious/w32) which hasn't been maintained and used CGO. This library doesn't use CGO and aims to mirror the win32 api and other Windows system dlls, without additional abstractions built on top of it. It attempts to be as organized/documented as possible. 

This mirror has some of my own additions plus updates from other forks of the original project. I've attempted 
to document where I've pulled code from someone else. 

I add new API functions in if my current project needs them. If your project needs a particular function 
please submit a PR or issue. I also add in additional functions as I see other forks, or Go libraries that have them.


Example
=====
```
package main

import (
	"github.com/TotallyGamerJet/w32"
)

func main() {
	w32.MessageBox(0, "Hello World!", "Hello, World!", 0)
}
```

For more examples, look at the example folder.

Setup
=====

1. From the command line, type `go get github.com/TotallyGamerJet/w32`
2. Create a new file, and try the example above.

Contribute
==========

Any type of contributions are welcome. Please make sure to follow Go formatting as well as include documentation for 
exported functions, constants and types.

Thank You!
