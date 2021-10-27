package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

const (
	CpuType = iota
	MemoryType
	MouseType
)

type Message struct {
	Key  int8
	Body interface{}
}

type Cpu struct {
	Core int8
}

type Memory struct {
	Size int8
}

type Mouse struct {
	Kind int8
}

func Producer(no int) []byte {
	buf := new(bytes.Buffer)
	var obj Cpu
	switch no {
	case CpuType:
		obj = Cpu{8}
	case MemoryType:
		obj = Cpu{3}
	case MouseType:
		obj = Cpu{2}
	default:
		panic("no type match")
	}
	err := binary.Write(buf, binary.LittleEndian, obj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buf.Bytes()
}

func Consumer(mes *Message) {
	switch mes.Key {
	case CpuType:
		fmt.Println(mes.Body.(*Cpu))
	case MemoryType:
		fmt.Println(mes.Body.(*Memory))
	case MouseType:
		fmt.Println(mes.Body.(*Mouse))
	default:
		fmt.Println(" consuemr Nothing")
	}
}
