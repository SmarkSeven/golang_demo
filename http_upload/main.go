package main

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
)

type MyHandler struct{}

func (MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if "POST" != r.Method {
		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, nil)
	} else {
		HttpHandle(w, r)
	}
}

func main() {
	handler := MyHandler{}
	if err := http.ListenAndServe(":8001", handler); err != nil {
		fmt.Println("Start http server fail:", err)
	}
}

func HttpHandle(w http.ResponseWriter, r *http.Request) {
	// 这里一定要记得 r.ParseMultipartForm(), 否则 r.MultipartForm 是空的
	// 调用 r.FormFile() 的时候会自动执行 r.ParseMultipartForm()
	r.ParseMultipartForm(32 << 20)
	// 写明缓冲的大小。如果超过缓冲，文件内容会被放在临时目录中，而不是内存。过大可能较多占用内存，过小可能增加硬盘 I/O
	// FormFile() 时调用 ParseMultipartForm() 使用的大小是 32 << 20，32MB
	file, fileHeader, err := r.FormFile("file") // file 是上传表单域的名字
	if err != nil {
		fmt.Println("get upload file fail:", err)
		w.WriteHeader(500)
		return
	}
	defer file.Close() // 此时上传内容的 IO 已经打开，需要手动关闭！！

	// fileHeader 有一些文件的基本信息
	fmt.Println(fileHeader.Header.Get("Content-Type"), fileHeader.Filename)

	// 打开目标地址，把上传的内容存进去
	f, err := os.OpenFile("saveto", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("save upload file fail:", err)
		w.WriteHeader(500)
		return
	}

	defer f.Close()
	if _, err = io.Copy(f, file); err != nil {
		fmt.Println("save upload file fail:", err)
		w.WriteHeader(500)
		return
	}

	size, _ := getUploadFileSize(file)
	w.Write([]byte("upload file:" + fileHeader.Filename + " - saveto : saveto" + " - size:" + strconv.FormatInt(size, 10) + "B"))
}

type fileSizer interface {
	Size() int64
}

// 从 multipart.File 获取文件大小
func getUploadFileSize(f multipart.File) (int64, error) {
	// 从内存读取出来
	// if return *http.sectionReader, it is alias to *io.SectionReader
	if s, ok := f.(fileSizer); ok {
		return s.Size(), nil
	}
	// 从临时文件读取出来
	// or *os.File
	if fp, ok := f.(*os.File); ok {
		fi, err := fp.Stat()
		if err != nil {
			return 0, err
		}
		return fi.Size(), nil
	}
	return 0, nil
}
