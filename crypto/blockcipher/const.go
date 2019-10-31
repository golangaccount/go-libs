package blockcipher

//MODE 分组加密模式类型
type MODE string

const (
	//CFB CFB
	CFB MODE = "CFB"
	//CTR CTR
	CTR MODE = "CTR"
	//OFB OFB
	OFB MODE = "OFB"
)
