package main

import (
	"flag"
	"fmt"
)

// 第一种创建flag变量的方式
var name = flag.String("name", "World", "A name to say hell to.")

// 第二种创建flag变量的方式
var spanish bool
var num int

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language.")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language.")
	flag.IntVar(&num, "num", 1, "number of language.")
}

func main() {
	// 打印帮助文档
	// flag.PrintDefaults()

	// 自定义帮助文档
	flag.VisitAll(func(flag *flag.Flag) {
		format := "\t-%s: %s (Default: '%s')\n"
		fmt.Printf(format, flag.Name, flag.Usage, flag.DefValue)
	})
	// 解析实际值到flag变量
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
