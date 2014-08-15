/*
 * filename   : utils.go
 * created at : 2014-08-15 12:03:10
 * author     : Jianing Yang <jianingy.yang@gmail.com>
 */

package utils

import (
    "bytes"
    "fmt"
    "os"

)

func Log(severity string, args ...interface{}) {
    var buf bytes.Buffer
	fmt.Fprintf(&buf, "[%s] ", severity)
    if len(args) > 1 {
        fmt.Fprintf(&buf, args[0].(string), args[1:]...)
    } else {
        fmt.Fprint(&buf, args...)
    }
	fmt.Fprint(os.Stdout, buf.String())
}
