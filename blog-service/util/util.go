package util

import (
    "path/filepath"
    "runtime"
)

/**
 * @author Rancho
 * @date 2021/12/8
 */

func GetCurrentPath() string {
    _, file, _, _ := runtime.Caller(1)
    return filepath.Dir(file)
}
