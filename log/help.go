package log

import "log"

//OpenLogFile 开启调用输出
func OpenLogFile() {
	rwLock.Lock()
	defer rwLock.Unlock()
	log.SetFlags(log.LstdFlags | log.Llongfile)
}
