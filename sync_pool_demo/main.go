package main

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"
)

// 对象池可以减少内存分配，降低CG压力
// 对象池使用是较简单的，但原生的sync.Pool有个较大的问题：不能自由控制Pool中元素的数量，
// 放进Pool中的对象每次GC发生时都会被清理掉。这使得sync.Pool做简单的对象池还可以，但做连接池就有点心有余而力不足了，
// 比如：在高并发的情景下一旦Pool中的连接被GC清理掉，那每次连接DB都需要重新三次握手建立连接，这个代价就较大了。
var defaultSpliter = "/"

var defaultReplacer = "-"

var ErrType = errors.New("类型错误")

type Dao struct {
	bp sync.Pool
}

func New() (d *Dao) {
	d = &Dao{
		bp: sync.Pool{
			New: func() interface{} {
				return &bytes.Buffer{}
			},
		},
	}
	return
}

func (d *Dao) Infoc(args ...string) (value string, err error) {
	if len(args) == 0 {
		return
	}
	// fetch a buf from bufpool
	buf, ok := d.bp.Get().(*bytes.Buffer)
	if !ok {
		return "", ErrType
	}
	// append first arg
	if _, err := buf.WriteString(args[0]); err != nil {
		return "", err
	}
	for _, arg := range args[1:] {
		// append ,arg
		if _, err := buf.WriteString(defaultSpliter); err != nil {
			return "", err
		}
		if _, err := buf.WriteString(strings.Replace(arg, defaultSpliter, defaultReplacer, -1)); err != nil {
			return "", err
		}
	}
	value = buf.String()
	buf.Reset()
	d.bp.Put(buf)
	return
}

func main() {
	dao := New()
	str, err := dao.Infoc("make", "build", "test")
	if err != nil {
		fmt.Println("some error:", err)
	}
	fmt.Println("result:", str)
}
