package jarr

import (
	"crypto/md5"
	"fmt"
)

func EleInArr(ele string, arr []string) bool {
	for _, v := range arr {
		if v == ele {
			return true
		}
	}
	return false
}

func UniqueString(arr []string) (unique_arr []string) {
	seenLines := make(map[[16]byte]bool)
	for _, v := range arr {
		lineBytes := []byte(v)
		hash := calculateHash(lineBytes)

		if !seenLines[hash] {
			seenLines[hash] = true
			unique_arr = append(unique_arr, v)
		}
	}
	return
}

func UniqueInt(arr []int) (unique_arr []int) {
	seenLines := make(map[[16]byte]bool)
	for _, v := range arr {
		lineBytes := []byte(fmt.Sprintf("%d", v))
		hash := calculateHash(lineBytes)

		if !seenLines[hash] {
			seenLines[hash] = true
			unique_arr = append(unique_arr, v)
		}
	}
	return
}

func calculateHash(data []byte) [16]byte {
	return md5.Sum(data)
}

func GetLongestEls(arr []string) (longest string) {
	// 初始化最长字符串为空字符串
	// 遍历字符串数组
	for _, str := range arr {
		// 检查当前字符串的长度是否大于最长字符串的长度
		if len(str) > len(longest) {
			longest = str
		}
	}
	return
}
