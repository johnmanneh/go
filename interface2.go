package main

import(
	"fmt"
	"bytes"

)

type Writer interface{
	Write([]byte) (int,error)
}
type Closer interface{
	Close() error
}
type WriterCloser interface{
	Writer
	Closer
}

type BufferedWriterCloser struct{
	buffer *bytes.Buffer
}
func(bwc BufferedWriterCloser) Writer(data[]byte)(int error){
	_,err := bwc.buffer.Write(data)
	if err != nil{
		return err
	}
	val := make([]byte,8)
	for bwc.buffer.Len() > 8{
		_, err := bwc.buffer.Read(val)
		if err != nil{
			return err
		}
		_, err = fmt.Println(string(val))
		if err != nil{
			return err
		}
	}
	return nil
}

func (bwc BufferedWriterCloser) Closer() error{
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil{
			return err
		}
	}
	return nil
}

func NewBufferWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer:bytes.NewBuffer([]byte{}),
	}
}
func main(){
	
	var obj interface{} = NewBufferWriterCloser()
	if wc,ok := obj.(WriterCloser); ok{
		wc.Write([]byte("testfile interface"))
		wc.Close()
	}
}