package bigline

import (
	"time"

	"github.com/golang-collections/go-datastructures/queue"
)

type BigLine struct {
	Master *queue.Queue
	//Time   time.Time

	Sub         *Queue
	MaxSubQueue int64

	SubExpireTime time.Duration
	Delegate      BigLineProtocol
	PopFrequency  time.Duration
}

type Queue struct {
	Queue *queue.Queue
	//Time  time.Time
}

func New() *BigLine {
	instance := &BigLine{}
	instance.Master = queue.New(60)
	//instance.Time = time.Now()

	instance.MaxSubQueue = 100000
	instance.SubExpireTime = 1000 * time.Millisecond
	instance.PopFrequency = 1 * time.Second
	instance.Sub = &Queue{}
	instance.Sub.Queue = queue.New(60)
	//instance.Sub.Time = time.Now()

	return instance
}

type BigLineProtocol interface {
	BigLinePop(result []interface{}, err error)
}
