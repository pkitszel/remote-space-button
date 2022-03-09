package main

import (
	_ "embed"
	"fmt"
	"io"
	"net/http"
	"os/exec"
	"strings"
)

func click(q string) bool {
	err := exec.Command("./winclick.exe", q).Run()
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
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
	http.ListenAndServe("0.0.0.0:20133", nil)
}

//go:embed html/begin.html
var pageBegin string

//go:embed html/end.html
var pageEnd string
