package utils

import (
	"encoding/base64"
	"fmt"
	"sync"
	"time"
)

const mx = 65536

var cur int32 = 0
var mutex sync.Mutex

// NextID 生成下一个本机唯一id
func NextID() string {
	var id int32

	mutex.Lock()
	id = cur
	cur++
	if cur >= mx {
		cur = 0
	}
	mutex.Unlock()

	now := time.Now().Unix()

	data := []byte(fmt.Sprintf("%016x-%04x", now, id))
	return base64.URLEncoding.EncodeToString(data)
}
