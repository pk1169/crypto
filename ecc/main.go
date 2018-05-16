package main

/*
#include "./keccak256/sha3.h"
#include "./keccak256/sha3.cpp"
*/
import "C"

import (
	"crypto/elliptic"
	"crypto/ecdsa"
	"crypto/rand"

	"go-ethereum-master/crypto"
	"fmt"
	"go-ethereum-master/crypto/ecies"
	"unsafe"
)

type cbyte C.uchar

type cint C.int

func main() {

	// go自带的椭圆曲线
	curve := elliptic.P256()
	privKey, err := ecdsa.GenerateKey(curve, rand.Reader)
	// 打印出公钥和私钥
	if err == nil {
		fmt.Println(privKey.D)
		fmt.Println(privKey.PublicKey)
	}

	plain := []byte{3, 4, 76, 91, 03, 87}

	// 以太坊的hash算法，并且对hash值进行签名，签名算法用的是go自带的椭圆曲线
	hash := crypto.Keccak256(plain)
	fmt.Println(hash)
	fmt.Println()
	r, s , err := ecdsa.Sign(rand.Reader, privKey, hash)

	if err == nil {
		fmt.Println(r, " ", s)
	}

	fmt.Println(privKey.PublicKey.Y)
	f := ecdsa.Verify(&privKey.PublicKey, hash, r, s)
	fmt.Println(f)



	// 以太坊用的是secp256k1, 首先生成私钥，
	privKey1, err1 := crypto.GenerateKey()

	if err1 == nil {
		fmt.Println(privKey1)
	}

	sig, err := crypto.Sign(hash, privKey1) // r || s || v.  sig[64] = v
	fmt.Println(len(sig))
	fmt.Println(sig[64])

	pubKey, err := crypto.Ecrecover(hash, sig)
	isSigTrue:= crypto.VerifySignature(pubKey, hash, sig[:64])
	fmt.Println(isSigTrue)

	//eciePriv := ecies.ImportECDSA(privKey1)
	// ecie加密体系
	eciePub := ecies.ImportECDSAPublic(&privKey1.PublicKey)
	d, err2 := ecies.Encrypt(rand.Reader, eciePub, hash, []byte{0, 1, 2, 7}, []byte{9,1, 2, 3})
	if err2 == nil {
		fmt.Println(d)
	}

	//encode := rlp.Encode(os.Stdout, hash)
	//fmt.Println(encode)

	// c语言sha3的库调用，对字符串进行hash
	var ss  = "zxcvbnmasdfghjklqwertyuiop123456"
	var input1 [32]cbyte
	var output1 = make([]cbyte, 32)
	// 将字符串转化成数组，go里面的字符串不支持修改
	for k, value := range ss {
		input1[k] = cbyte(value)
	}

	// 调用c语言sha3库
	C.sha3(unsafe.Pointer(&input1[0]), C.ulong(len(input1)), unsafe.Pointer(&output1[0]), C.int(len(output1)))
	fmt.Println(output1)

}
