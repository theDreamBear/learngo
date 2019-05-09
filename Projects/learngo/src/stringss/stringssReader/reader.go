package main

import (
	"fmt"
	"unicode/utf8"
)

type SsReader struct {
	Data  []byte
	token []byte
	Split SplitFunc
	start int
	end   int
	done  bool
}

type SplitFunc func(buf []byte, isEOF bool) (int, []byte)

func (sr *SsReader) Splitfunc(split SplitFunc) {
	sr.Split = split
}

func SplitWords(buf []byte, isEOF bool) (int, []byte) {
	start := 0
	for width := 0; start < len(buf); start += width {
		var r rune
		r, width = utf8.DecodeRune(buf[start:])
		if !isSpace(r) {
			break
		}
	}
	for width, i := 0, start; i < len(buf); i += width {
		var r rune
		r, width = utf8.DecodeRune(buf[i:])
		if isSpace(r) {
			return i + width, buf[start:i]
		}
	}

	//if isEOF && len(buf) > start {

	return len(buf), buf[start:]
	//}

	//return start, nil
}

func (sr *SsReader) Scan() bool {
	if sr.done {
		return false
	}
	for {
		if sr.end > sr.start {
			advance, buf := sr.Split(sr.Data[sr.start:sr.end], sr.done)
			sr.start += advance
			sr.token = buf
			if sr.start == sr.end {
				sr.done = true
			}
			break
		}

	}
	return true
}

func (sr *SsReader) left() int {
	return sr.end - sr.start
}

func NewReader(data string) SsReader {
	return SsReader{Data: []byte(data), start: 0, end: len(data)}
}

func (sr *SsReader) Text() string {
	return string(sr.token)
}

func (sr *SsReader) Read(p []byte) error {
	var waitForCopy = 0
	if len(p) >= sr.left() {
		waitForCopy = sr.left()
	} else if len(p) < sr.left() {
		waitForCopy = len(p)
	}
	copied := copy(p, sr.Data[sr.start:sr.start+waitForCopy])
	sr.start = sr.start + copied
	if sr.left() == 0 {
		sr.done = true
	}
	return nil
}

func isSpace(r rune) bool {
	if r <= '\u00FF' {
		// Obvious ASCII ones: \t through \r plus space. Plus two Latin-1 oddballs.
		switch r {
		case ' ', '\t', '\n', '\v', '\f', '\r':
			return true
		case '\u0085', '\u00A0':
			return true
		}
		return false
	}
	// High-valued ones.
	if '\u2000' <= r && r <= '\u200a' {
		return true
	}
	switch r {
	case '\u1680', '\u2028', '\u2029', '\u202f', '\u205f', '\u3000':
		return true
	}
	return false
}

func main() {
	var names = "nic silent tomy carpu seee  s"
	var sr = NewReader(names)
	//var buffer = make([]byte, 5)
	//for sr.done == false {
	//	sr.Read(buffer)
	//	fmt.Println(string(buffer))
	//}
	sr.Split = SplitWords
	for sr.Scan() {
		fmt.Println(sr.Text())
	}
}
