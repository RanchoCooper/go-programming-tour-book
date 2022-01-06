package util

import (
    "crypto/md5"
    "encoding/hex"
    "path/filepath"
    "runtime"
)

func GetCurrentPath() string {
    _, file, _, _ := runtime.Caller(1)
    return filepath.Dir(file)
}

func EncodeMD5(value string) string {
    m := md5.New()
    m.Write([]byte(value))

    return hex.EncodeToString(m.Sum(nil))
}
