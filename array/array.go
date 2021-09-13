package _array

import (
	"math/rand"
	"strings"
	"time"
)

// Distinct 去除重复的元素
func Distinct(source []int) (target []int) {
	if source == nil {
		return nil
	}
	distinctMap := make(map[int]int)
	for _, item := range source {
		_, exist := distinctMap[item]
		if !exist {
			target = append(target, item)
			distinctMap[item] = 1
		}
	}
	return target
}

// GetShuffleArray 生成包含count个元素的打乱后的数组 数组元素为[0,count)中的数字
func GetShuffleArray(count int) (shuffleArr []int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffleArr = r.Perm(count)
	return shuffleArr
}

// Contain 数组中是否包含指定的字符串
func Contain(source []string, target string) (exist bool) {
	for _, item := range source {
		if item == target {
			exist = true
			return exist
		}
	}
	return exist
}

// ContainStr 字符串通过字符切分后是否指定的字符串
func ContainStr(source string, target string, seq string) (exist bool) {
	spliteResult := strings.Split(source, seq)
	return Contain(spliteResult, target)
}

// ContainStrsOne 字符串通过字符切分后是否包含指定的多个字符串中的一个或多个
func ContainStrsOne(source string, targets []string, seq string) (exist bool) {
	spliteResult := strings.Split(source, seq)
	for _, target := range targets {
		if len(target) == 0 {
			continue
		}
		exist = Contain(spliteResult, target)
		if exist {
			return exist
		}
	}
	return exist
}

// ContainStrsAll 字符串通过字符切分后是否包含指定的N个字符串
// 如果targets为空数组或者空字符串则返回false
func ContainStrsAll(source string, targets []string, seq string) (exist bool) {
	if len(targets) == 0 {
		return false
	}
	spliteResult := strings.Split(source, seq)
	for _, target := range targets {
		if len(target) == 0 {
			return false
		}
		exist = Contain(spliteResult, target)
		if exist == false {
			return exist
		}
	}
	return true
}

// StrEqual 判断两个字符串数组是否相等,包括顺序
func StrEqual(left, right []string) bool {
	if left == nil && right == nil {
		return true
	}
	leftLen := len(left)
	rightLen := len(right)
	if leftLen != rightLen {
		return false
	}
	for index, value := range left {
		if value != right[index] {
			return false
		}
	}
	return true
}
