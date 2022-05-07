package tool

import (
	"fmt"
	"strconv"
	"sync"
	"time"
	"tmp-code/data"
)

const (
	numberBits uint8 = 12
	numberMax  int64 = -1 ^ (-1 << numberBits)
	timeShift  uint8 = numberBits
	startTime  int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	number    int64
}

func NewWorker() *Worker {

	// 生成一个新节点
	return &Worker{
		timestamp: 0,
		number:    0,
	}
}

func (w *Worker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 获取毫秒
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {

			// 如果当前所有的序列号用光了，这里自旋等待
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.number))
	return ID
}

func (w *Worker) ShortId() string {

	//获取一个id
	num := w.GetId()

	// 最终返回的值
	var new_num_str string = ""

	var remainder int32
	var remainder_string string
	for num != 0 {

		// 求余数，使用int32存储，避免生成申请int64
		remainder = int32(num % 62)
		if 62 > remainder && remainder > 9 {
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(int(remainder))
		}
		new_num_str = remainder_string + new_num_str
		num = num / 62
	}
	return new_num_str
}

func HandleTime(expire string) time.Duration {
	fmt.Println(expire)
	switch expire {
	case data.Hour:
		return time.Hour
	case data.Day:
		return time.Hour * 24
	case data.Week:
		return time.Hour * 24 * 7
	case data.HalfMonth:
		return time.Hour * 24 * 15
	case data.Month:
		return time.Hour * 24 * 30
	}
	return time.Hour
}
