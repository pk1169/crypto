package main

import (
	"fmt"
	"crypto_learning/mycrypto"
	"io/ioutil"
	"time"
	"os"
)

// 对des，3des，aes的性能测试

func main() {
	//h, time, e := mycrypto.FileKeccak256("./Go并发编程实战 第2版.pdf")
	//
	//if e != nil {
	//	fmt.Println(e)
	//} else {
	//	fmt.Println(time)
	//	fmt.Println(h)
	//}


	content, err := ioutil.ReadFile("./Go并发编程实战 第2版.pdf")
	checkErr(err)



	desKey := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	aesKey := make([]byte, 0)

	// AesKey
	for i := 0; i < 3; i++ {
		aesKey = append(aesKey, desKey...)
	}
	tdesKey := aesKey
	if err != nil {
		fmt.Println(err)
	} else {
		var n int
		for {
			fmt.Scanf("%d", &n)
			switch n {
			case 1:
				t1 := time.Now()
				_, err := mycrypto.DesEncrypt(content, desKey)
				checkErr(err)
				t2 := time.Now()
				fmt.Println("time cost: ", t2.Sub(t1))
			case 2:
				t3 := time.Now()
				_, err := mycrypto.TripleDesEncrypt(content, tdesKey)
				checkErr(err)
				t4 := time.Now()
				fmt.Println("time cost: ", t4.Sub(t3))
			case 3:
				t5 := time.Now()
				_, err := mycrypto.AesEncrypt(content, aesKey)
				checkErr(err)
				t6 := time.Now()
				fmt.Println("time cost: ", t6.Sub(t5))
			case 0:
				os.Exit(-1)
			}
		}
	}


	//origdata := []byte{34, 56, 12, 3, 5, 9, 89, 90, 00, 70}
	//
	//d, e1 := mycrypto.DesEncrypt(origdata, key)
	//
	//if e1 != nil {
	//	fmt.Println(e1)
	//}
	//
	//fmt.Println(d)
	//
	//p, e2 := mycrypto.DesDecrypt(d, key)
	//
	//if e2 != nil {
	//	fmt.Println(e2)
	//}
	//
	//fmt.Println(p)
	//
	//key1 := make([]byte, 0)
	//for i := 0; i < 3; i++ {
	//	key1 = append(key1, key...)
	//}
	//
	//d1, e3 := mycrypto.AesEncrypt(origdata, key1)
	//if e3 != nil {
	//	fmt.Println(e3)
	//}
	//fmt.Println(d1)
	//
	//p1, e4 := mycrypto.AesDecrypt(d1, key1)
	//if e4 != nil {
	//	fmt.Println(e4)
	//}
	//fmt.Println(p1)

}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}