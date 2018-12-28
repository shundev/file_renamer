package renamer

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

func check(err error) {
	if err != nil {
		_ = fmt.Errorf(err.Error())
		panic(err)
	}
}

// Rename all files whose name matches to the matchPtn in a directory
func Rename(dpath string, matchPtn string, replacePtn string) {
	files, _ := filepath.Glob(fmt.Sprintf("%s/*", dpath))
	re := regexp.MustCompile(matchPtn)

	numReplaced := 0
	for _, f := range files {
		if m := re.MatchString(f); m {
			err := os.Rename(f, re.ReplaceAllString(f, "${1}_copy.txt"))
			check(err)
			numReplaced++
		}
	}

	fmt.Println(numReplaced, "files were replaced.")

	return
}
