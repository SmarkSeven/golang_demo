package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// 使用Context 在http.HandleFunc 之间传递数据
func handle1(w http.ResponseWriter, r *http.Request) {
	ctx := context.WithValue(r.Context(), "abc", "123") // 将数据写入到Context
	handle2(w, r.WithContext(ctx))
}

func handle2(w http.ResponseWriter, r *http.Request) {
	value, ok := r.Context().Value("abc").(string)
	if !ok {
		value = "no string"
	}
	w.Write([]byte("Context.abc = " + value))
}

// 利用Context处理超时请求
func handle3(w http.ResponseWriter, r *http.Request) {
	ctx, cancelFunc := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancelFunc()

	resCh := make(chan string, 1)
	go func() {
		// 业务超时
		time.Sleep(2 * time.Second)
		resCh <- r.FormValue("abc")
	}()

	select {
	case <-ctx.Done():
		w.WriteHeader(http.StatusGatewayTimeout)
		w.Write([]byte("http handle is timeout:" + ctx.Err().Error()))
	case r := <-resCh:
		w.Write([]byte("get:abc = " + r))
	}

}

func main() {
	http.HandleFunc("/", handle1)
	http.HandleFunc("/timeout", handle3)
	if err := http.ListenAndServe(":4001", nil); err != nil {
		fmt.Println("Start http server fail:", err)
	}
}
