package control

import (
	"fmt"
	"sync"

	"github.com/jamesblockk/Start-Go/HelloGo/newqueue/bigline"
)

type Control struct {
	Q *bigline.BigLine
}

type BigLineProtocol struct {
	bigline.BigLineProtocol
}

var instance *Control
var once sync.Once

func Shared() *Control {
	once.Do(func() {
		instance = &Control{}
	})
	return instance
}

func (ctl *Control) Init() {
	ctl.Q = bigline.New()
	ctl.Q.Delegate = new(BigLineProtocol)
}

func (ctl *Control) Start() {
	ctl.Q.Start()
}

func (ctl *Control) Push(items ...interface{}) {
	ctl.Q.Push(items)
}

func (line BigLineProtocol) BigLinePop(result []interface{}, err error) {
	fmt.Println(result, err)
	fmt.Println("BigLinePoped ")
}
