// json2type - convert json response into a Go type-struct
//
// Copyright 2011 The json2type Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is only a helping tool... do not trust it too much
// useful to run from inside Emacs with C-u M-|
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func camel(s string) string {
	var news string
	for _, w := range strings.Split(s, "_") {
		news += strings.ToUpper(w[0:1]) + w[1:]
	}
	return news
}

// reads from std input and dumps to std out
func main() {
	buf, _ := ioutil.ReadAll(os.Stdin)
	var obj interface{}
	err := json.Unmarshal(buf, &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println("type MyType ")
	switch v := obj.(type) {
	case []interface{}:
		gen(v[0], "")
	default:
		gen(obj, "")
	}
}

func jname(n string) string {
	if n == "" {
		return "\n"
	}
	return fmt.Sprintf("`json:\"%s\"`\n", n)
}
func gen(obj interface{}, name string) {
	switch v := obj.(type) {
	case float64:
		fmt.Printf("int64 %s\n", jname(name))
	case []interface{}:
		fmt.Print("[]")
		if len(v) < 1 {
			fmt.Printf("string %s", jname(name))
		} else {
			gen(v[0], name)
		}
	case map[string]interface{}:
		fmt.Printf("struct {")
		for k, i := range v {
			fmt.Printf("%s ", camel(k))
			gen(i, k)
		}
		fmt.Printf("} %s", jname(name))
	case string:
		fmt.Printf("string %s", jname(name))
	case bool:
		fmt.Printf("bool %s", jname(name))
	default:
		fmt.Printf("string %s", jname(name))
	}
}
