package regexp

import "regexp"

const (
	//REGEXP regexp验证
	REGEXP = "regexp"
	//EMAIL 邮箱验证
	EMAIL = "email"
	//URL url地址
	URL = "url"
	//MOBILPHONE 移动电话
	MOBILPHONE = "mobilphone"
)

var (
	//EMAILREG email 正则
	EMAILREG = regexp.MustCompile("[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?")
	//URLREG url正则
	URLREG = regexp.MustCompile(`(https?://)?([^\s\.]{1,}\.){1,}[^\s\.]{1}`)
	//MOBILPHONEREG 移动电话正则
	MOBILPHONEREG = regexp.MustCompile("1\\d{10}")
)
