package log

import (
	"fmt"
	"runtime"
)

// GetStackTrace スタックトレースを取得する
func GetStackTrace() string {

	stackTrace := ""
	for depth := 0; ; depth++ {
		_, file, line, ok := runtime.Caller(depth)
		if !ok {
			break
		}
		stackTrace += fmt.Sprintf("%02d: %v:%d\n", depth+1, file, line)
	}

	return stackTrace
}
