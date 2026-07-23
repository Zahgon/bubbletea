package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"charm.land/bubbles/v2/progress"
	tea "charm.land/bubbletea/v2"
)

var p *tea.Program

type progressWriter struct {
	total      int
	downloaded int
	file       *os.File
	reader     io.Reader
	onProgress func(float64)
}

func (pw *progressWriter) Start() { _ = "STUB: not implemented"; return }

func (pw *progressWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func getResponse(url string) (*http.Response, error) { _ = "STUB: not implemented"; return nil, nil }

func main() {
	url := flag.String("url", "", "url for the file to download")
	flag.Parse()

	if *url == "" {
		flag.Usage()
		os.Exit(1)
	}

	resp, err := getResponse(*url)
	if err != nil {
		fmt.Println("could not get response", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.ContentLength <= 0 {
		fmt.Println("can't parse content length, aborting download")
		os.Exit(1)
	}

	filename := filepath.Base(*url)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("could not create file:", err)
		os.Exit(1)
	}
	defer file.Close()

	pw := &progressWriter{
		total:  int(resp.ContentLength),
		file:   file,
		reader: resp.Body,
		onProgress: func(ratio float64) {
			p.Send(progressMsg(ratio))
		},
	}

	m := model{
		pw:       pw,
		progress: progress.New(progress.WithDefaultBlend()),
	}

	p = tea.NewProgram(m)

	go pw.Start()

	if _, err := p.Run(); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
