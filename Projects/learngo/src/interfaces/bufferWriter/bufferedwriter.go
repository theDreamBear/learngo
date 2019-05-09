package main

import (
	"fmt"
	"io"
)

type StdWriter struct {
	writer io.Writer
}

func (Fw *StdWriter) Write(p []byte) (int, error) {
	fmt.Printf("write content is: %v\n", string(p))
	return len(p), nil
}

type BufferedWriter struct {
	Buffer  []byte
	Cur     int
	Size    int
	Writer  io.Writer
	iotimes int
	err     error
}

func NewBufferedWriter(writer io.Writer, BufferSize int) *BufferedWriter {
	bw := new(BufferedWriter)
	bw.Buffer = make([]byte, BufferSize)
	bw.Cur = 0
	bw.Size = BufferSize
	bw.Writer = writer
	return bw
}

func (Bw *BufferedWriter) Write(p []byte) (n int, err error) {
	var pstart = 0
	if Bw.LeftSize() == 0 {
		Bw.writeTo()
	}
	for pstart < len(p) {

		if Bw.LeftSize() >= len(p)-pstart {
			Bw.writeToBuffer(p[pstart:])
			pstart = len(p)
			//Bw.writeTo()
		} else {
			Bw.writeToBuffer(p[pstart : pstart+Bw.LeftSize()])
			Bw.writeTo()
			pstart += Bw.LeftSize()
		}
	}

	return len(p), nil
}

func (Bw *BufferedWriter) Flush() {
	fmt.Printf("flush the last\n")
	Bw.writeTo()
}

func (Bw *BufferedWriter) LeftSize() int {
	return Bw.Size - Bw.Cur
}

func (Bw *BufferedWriter) writeToBuffer(p []byte) (int, error) {
	copy(Bw.Buffer[Bw.Cur:], p)
	Bw.Cur += int(len(p))
	return len(p), nil
}

func (Bw *BufferedWriter) writeTo() (int, error) {
	Bw.iotimes++
	fmt.Printf("io 次数:%v\n", Bw.iotimes)
	if _, err := Bw.Writer.Write(Bw.Buffer[0:Bw.Cur]); err != nil {
		Bw.err = err
	} else {
		Bw.Cur = 0

	}
	return Bw.Size, nil
}

func main() {
	var stdW = &StdWriter{}
	var bw = NewBufferedWriter(stdW, 5)
	bw.Write([]byte("a"))
	bw.Write([]byte("b"))
	bw.Write([]byte("c"))
	bw.Write([]byte("d"))
	bw.Write([]byte("e"))
	bw.Write([]byte("f"))
	bw.Write([]byte("g"))
	bw.Write([]byte("h"))
	bw.Write([]byte("i"))
	bw.Write([]byte("j"))
	bw.Write([]byte("hello world"))
	bw.Flush()

	//var bw = bufio.NewWriter()
}
