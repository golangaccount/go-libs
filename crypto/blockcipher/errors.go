package blockcipher

import "errors"

var (
	//ErrMODENotExists 加密解密的模式不存在
	ErrMODENotExists = errors.New("MODE错误")
	//ErrAESPasswordTolang AES加密的秘钥长度过长，超过了32位
	ErrAESPasswordTolang = errors.New("AES加密的秘钥长度过长")
	//ErrAESIVLength AES的iv长度默认为16
	ErrAESIVLength = errors.New("AES iv 长度错误")
	//ErrAESStreamIv 从数据流获取iv数据失败
	ErrAESStreamIv = errors.New("获取iv数据失败")
)
