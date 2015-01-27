package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	bts, _ := ioutil.ReadAll(r)
	msg := string(bts)

	re := regexp.MustCompile(`([a-zA-Z0-9\-_]+\.go:[0-9]+)`)
	panic_file_matches := re.FindAllString(msg, -1)
	msg = re.ReplaceAllString(msg, "\033[1;31m$1\033[0m")

	running_re := regexp.MustCompile(`\[running\]`)
	n_goroutines := len(running_re.FindAllString(msg, -1))

	fmt.Println(msg)

	if n_goroutines > 0 {
		fmt.Println("colar:\n\033[1;31m", n_goroutines, "goroutines were running during the panic\033[0m")
		for _, s := range panic_file_matches {
			fmt.Println("   \033[1;31m" + s + "\033[0m")
		}
	}
}
