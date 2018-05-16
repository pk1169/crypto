package sm4



func SM4cbcEncrypt(key [16]byte, mode, length int, iv [16]byte, input, output []byte) {
	ctx := new(sm4_context)
	sm4_setkey_enc(ctx, key)
	Sm4_crypt_cbc(ctx,1, length, iv, input, output)
}

func SM4cbcDec(key [16]byte, mode, length int, iv [16]byte, input, output []byte) {
	ctx := new(sm4_context)
	sm4_setkey_dec(ctx, key)
	Sm4_crypt_cbc(ctx, 0, length, iv, input, output)
}
