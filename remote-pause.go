package main

import (
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

var pageBegin = `
<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>Remote SPACE button</title>
	<meta name="description" content="Remote SPACE button">
	<meta name="author" content="pkitszel">
	<style>
	.button {
		background-color: #e7e7e7;
		border: none;
		color: black;
		padding: 15px 32px;
		text-align: center;
		text-decoration: none;
		display: flex;
		justify-content: center;
		align-items: center;
		font-size: 96px;
		margin: 4px 2px;
		cursor: pointer;
	}
	.center {
		margin: 0;
		position: absolute;
		top: 50%;
		left: 50%;
		-ms-transform: translate(-50%, -50%);
		transform: translate(-50%, -50%);
	}
	</style>
</head>
<body>
`

var pageEnd = `
<center class="center">
		<a href="?=_"  class="button">▶️/⏸️</a></br>
		<a href="?=S"  class="button">Skip Intro</a></br>
		<a href="?=F"  class="button">Fullscreen</a></br>
		<a href="?=SF" class="button">JUST MOVIE!!!</a></br>
		<a href="?=M"  class="button">Mjót</a></br>
	</center>
</body>
</html>
`
