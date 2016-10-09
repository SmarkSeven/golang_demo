package main

import (
	"fmt"
	"net"
)

func main() {
	// addrs, _ := net.InterfaceAddrs()
	// fmt.Println(addrs)
	// interfaces, _ := net.Interfaces()
	// fmt.Println(interfaces)
	// hp := net.JoinHostPort("127.0.0.1", "8080")
	// fmt.Println(hp) //127.0.0.1:8080,根据ip和端口组成一个addr字符串表示
	// lt, _ := net.LookupAddr("180.97.33.107")
	// fmt.Println(lt) //[www.a.shifen.com.],根据地址查找到改地址的一个映射列表
	// cname, _ := net.LookupCNAME("www.baidu.com")
	// fmt.Println(cname) //www.a.shifen.com,查找规范的dns主机名字
	// host, _ := net.LookupHost("www.baidu.com")
	// fmt.Println(host) //[180.97.33.107 180.97.33.107],查找给定域名的host名称
	// ip, _ := net.LookupIP("www.baidu.com")
	// fmt.Println(ip) //[180.97.33.107 180.97.33.107],查找给定域名的ip地址,可通过nslookup www.baidu.com进行查找操作.
	// mxs, _ := net.LookupMX("www.baidu.com")
	// fmt.Println(mxs)
	nss, _ := net.LookupNS("www.baidu.com")
	fmt.Println(nss)
	dnserror := net.DNSError{
		Err:       "dns error",
		Name:      "dnserr",
		Server:    "dns search",
		IsTimeout: true,
	}
	fmt.Println(dnserror.Error())     //lookup dnserr on dns search: dns error
	fmt.Println(dnserror.Temporary()) //true
	fmt.Println(dnserror.Timeout())   //true
	fmt.Println(dnserror.Server)      //dns search
	fmt.Println(dnserror.Err)         //dns error
	fmt.Println(dnserror.Name)        //dnserr
	fmt.Println(dnserror.IsTimeout)   //true

}
