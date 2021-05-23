package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/opensourcez/godirwalk"
)

var skipList = make(map[string]bool)
var baseDIR = "/mnt/c/Users/Notandi/go/src/github.com/zveinn"

func main() {
	err := os.Chdir("/mnt/c/Users/Notandi/go/src/github.com/zveinn/cher")
	if err != nil {
		panic(err)
	}
}
func main1() {
	log.Println("LOOKING FOR:", os.Args)
	skipList["node_modules"] = true

	err := godirwalk.Walk(baseDIR, &godirwalk.Options{
		Callback: func(osPathname string, info *godirwalk.Dirent) error {
			// fmt.Print(".")
			if info.IsDir() {
				split := strings.Split(osPathname, "/")
				var ok bool
				lastDir := split[len(split)-1]
				ok, _ = skipList[lastDir]
				if ok {
					// log.Println("in list:", osPathname)
					return godirwalk.SkipThis
				}
				if lastDir[0] == byte('.') {
					// log.Println("has dota:", osPathname)
					return godirwalk.SkipThis
				}

				if strings.Contains(osPathname, os.Args[1]) {
					home, _ := os.UserHomeDir()
					x := filepath.Join(home, "goproject2")
					log.Println("FOUND:", osPathname, x)
					err := os.Chdir("/mnt/c/Users/Notandi/go/src/github.com/zveinn/cher")
					if err != nil {
						panic(err)
					}
					os.Exit(1)
				}
			}

			return nil
		},
		ErrorCallback: func(osPathname string, err error) godirwalk.ErrorAction {
			// log.Println("error:", osPathname, err)
			return 1
		},
		// Unsorted: true,
	})

	if err != nil {
		// Permissions error happens here, outside the walk
		os.Exit(1)
	}

}
