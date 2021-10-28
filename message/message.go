package message

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"os"
)

const (
	CpuType = iota
	MemoryType
	MouseType
)

type Message struct {
	Key  int `json:"key"`
	Body interface{}
}

type Cpu struct {
	Core int `json:"core"`
}

type Memory struct {
	Size int `json:"size"`
}

type Mouse struct {
	Kind int `json:"kind"`
}

func Producer(no int) []byte {
	obj := Message{}
	switch no {
	case CpuType:
		obj.Body = Cpu{8}
	case MemoryType:
		obj.Body = Memory{3}
	case MouseType:
		obj.Body = Mouse{2}
	default:
		panic("no type match")
	}
	obj.Key = no
	buf, err := json.Marshal(&obj)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return buf
}

func Consumer(mes *Message) {
	arr, _ := json.Marshal(mes.Body)
	switch mes.Key {
	case CpuType:
		cpu := Cpu{}
		json.Unmarshal(arr, &cpu)
		fmt.Println("CPU", cpu)
	case MemoryType:
		memory := Memory{}
		json.Unmarshal(arr, &memory)
		fmt.Println("Memory", memory)
	case MouseType:
		mouse := Mouse{}
		json.Unmarshal(arr, &mouse)
		fmt.Println("Mouse", mouse)
	default:
		fmt.Println(" consuemr Nothing")
	}
}
