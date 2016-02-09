package main

import (
	"fmt"
	"os"
	"sort"
	"path/filepath"
)

/*
	Indicates if an item is in the list.
*/
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

/*
	Help message for prettypath command.
*/
const helpMsg string = `prettypath v 1.0 help:
		
prettypath [ -v | -h | -a | -j ]
		
-v : show version information.
-h : show help information.
-a : sort path information alphabetically.
-j : show java classpath.

`	

/*
	The main function.
*/
func main() {

	// command-line args.
	args := os.Args[1:]
	
	// show version info.
	if(stringInSlice("-v", args)) {
		fmt.Println("prettypath v 1.0")
		return
	}

	// display help information.
	if(stringInSlice("-h", args)) {
		fmt.Println(helpMsg)
		return
	}

	// configure PATH or CLASSPATH
	var pathVar string = "PATH"
	if(stringInSlice("-j", args)) {
		pathVar = "CLASSPATH"
	}

	// get the path variable info.
	path := os.Getenv(pathVar)
	path_entries := filepath.SplitList(path)
	
	// sort alphabetically if requested.
	if(stringInSlice("-a", args)) {
		sort.Strings(path_entries)
	}

	// pretty print the path.
	for _, element := range path_entries {
	  fmt.Println(element)
	}

}

