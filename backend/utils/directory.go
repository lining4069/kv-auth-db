package utils

import "os"

// PathExists 判断路劲是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, err
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
