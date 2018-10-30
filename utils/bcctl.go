// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command sampleapp-boo prints out "hello, sample app boo".
// Don't put hyphens, that's ugly https://news.ycombinator.com/item?id=7762482
// Basically using hyphens in package names is No-No but for the app it's ok.
// But probably still not a good idea
//
// Usage
//
//    godoc2md $PACKAGE > $GOPATH/src/$PACKAGE/README.md
//
// Another line with description and maybe additional usage sample
//    hello $SOMEPARAM
package main

import "fmt"

func main() {

	fmt.Printf("bcctl utiliy started\n")
}
