package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
)

func click(q string) bool {
	err := exec.Command("./winclick.exe", q).Run()
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}

func printIP() {
	// resolver from "net" pkg uses cgo, so just exec `ipconfig`
	out, err := exec.Command("ipconfig").Output()
	if err != nil {
		fmt.Println(err)
	}
	re := regexp.MustCompile(`IPv[46] ?Address[ .]+([0-9.:]+)`)
	matches := re.FindAllSubmatch(out, -1)
	switch len(matches) {
	case 0:
		return
	case 1:
		fmt.Println("Open following page on your smartphone:")
	default:
		fmt.Println("Open one of following pages on your smartphone:")
	}
	for _, m := range matches {
		fmt.Printf("\thttp://%s:20133 \r\n", string(m[1]))
	}
}

func main() {
	var failCnt int

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.RawQuery
		q = strings.ReplaceAll(q, "_", " ")
		q = strings.ReplaceAll(q, "=", "")
		if q != "" {
			if !click(q) {
				failCnt += 1
			}
		}
		io.WriteString(w, pageBegin)
		if failCnt != 0 {
			fmt.Fprintf(w, "Error counter: %v\n", failCnt)
		}
		io.WriteString(w, pageEnd)
	})
	printIP()
	http.ListenAndServe("0.0.0.0:20133", nil)
}

//go:embed html/begin.html
var pageBegin string

//go:embed html/end.html
var pageEnd string
