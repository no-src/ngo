package _strings

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type String string

// Empty 空白字符串
const Empty = ""

// Trim 删除字符前后的空白字符
func (s String) Trim() String {
	return String(strings.TrimSpace(string(s)))
}

// TrimStart 去除字符前端的空白字符
func (s String) TrimStart() String {
	return String(strings.TrimLeftFunc(string(s), unicode.IsSpace))
}

// TrimEnd 去除字符后端的空白字符
func (s String) TrimEnd() String {
	return String(strings.TrimRightFunc(string(s), unicode.IsSpace))
}

// ToString 输出字符串
func (s String) ToString() string {
	return string(s)
}

// ToLower 全部转换为小写
func (s String) ToLower() String {
	return String(strings.ToLower(string(s)))
}

// ToUpper 全部转换为大写
func (s String) ToUpper() String {
	return String(strings.ToUpper(string(s)))
}

// Substring 截取指定位置的字符串
// start 开始截取的索引，结果包含该索引位置的字符，如果start<0，则start=0，如果start超过最大索引，则直接返回空字符串
// end 截取的结束索引，结果不包含该索引位置的字符，如果end<0或者end大于字符长度，则end=字符总长度
func (s String) Substring(start int, end int) String {
	str := string(s)
	length := s.Length()
	if length == 0 {
		return String(str)
	}
	if start < 0 {
		start = 0
	}
	if start > length-1 {
		return Empty
	}
	if end < 0 || end > length {
		end = length
	}
	if start == 0 && end == 0 {
		return Empty
	}
	str = string([]rune(str)[start:end])
	return String(str)
}

// Split 切分字符串
func (s String) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

// Replace 替换字符串中的字符
func (s String) Replace(oldStr, newStr string) String {
	str := string(s)
	str = strings.Replace(str, oldStr, newStr, -1)
	return String(str)
}

// Remove 移除指定索引开始的N个字符
func (s String) Remove(startIndex int, count int) String {
	length := s.Length()
	if length == 0 {
		return Empty
	}
	if startIndex > length {
		startIndex = length
	}
	if startIndex < 0 {
		startIndex = 0
	}
	left := s.Substring(0, startIndex)
	var right String = ""
	if startIndex+count <= length {
		right = s.Substring(startIndex+count, length)
	}
	return left.Append(right.ToString())
}

// IndexOf 指定字符串出现的第一个索引位置
func (s String) IndexOf(findStr string) int {
	findIndex := strings.Index(string(s), findStr)
	if findIndex == 0 || findIndex == -1 {
		return findIndex
	}
	length := s.Length()
	if length != len(string(s)) {
		str := s.ToString()
		prefix := str[:findIndex]
		findIndex = String(prefix).Length()
	}
	return findIndex
}

// LastIndexOf 最后一个匹配的字符索引
func (s String) LastIndexOf(findStr string) int {
	findLastIndex := strings.LastIndex(string(s), findStr)
	if findLastIndex == 0 || findLastIndex == -1 {
		return findLastIndex
	}
	length := s.Length()
	if length != len(string(s)) {
		str := s.ToString()
		prefix := str[:findLastIndex]
		findLastIndex = String(prefix).Length()
	}
	return findLastIndex
}

// IsNullOrEmpty 字符串是否为空
func (s String) IsNullOrEmpty() bool {
	if len(string(s)) == 0 {
		return true
	}
	return false
}

// IsNullOrWhiteSpace 字符串是否为空或空白字符
func (s String) IsNullOrWhiteSpace() bool {
	str := s.Trim()
	return String(str).IsNullOrEmpty()
}

// Insert 在指定位置插入字符串
func (s String) Insert(startIndex int, value string) String {
	length := s.Length()
	if startIndex > length {
		startIndex = length
	}
	if startIndex < 0 {
		startIndex = 0
	}
	left := s.Substring(0, startIndex)
	right := s.Substring(startIndex, -1)
	str := left.ToString() + value + right.ToString()
	return String(str)
}

// StartsWith 是否以某个字符串开头
func (s String) StartsWith(value string) bool {
	return strings.HasPrefix(string(s), value)
}

// EndsWith 是否以某个字符串结尾
func (s String) EndsWith(value string) bool {
	return strings.HasSuffix(string(s), value)
}

// Contains 是否包含指定的字符串
func (s String) Contains(value string) bool {
	return strings.Contains(string(s), value)
}

// Length 返回字符串的长度（根据utf8编码，一个中文算一个字符）
func (s String) Length() int {
	length := utf8.RuneCountInString(string(s))
	return length
}

// First 获取字符串中第一个字符
func (s String) First() String {
	length := s.Length()
	if length <= 1 {
		return s
	}
	return s.Substring(0, 1)
}

// Last 获取字符串中最后一个字符
func (s String) Last() String {
	length := s.Length()
	if length <= 1 {
		return s
	}
	return s.Substring(length-1, -1)
}

// At 获取指定索引的字符
func (s String) At(index int) String {
	if index < 0 {
		return Empty
	}
	length := s.Length()
	if length <= index {
		return Empty
	}
	return s.Substring(index, index+1)
}

// ToArray 将字符串转化为字符数组
func (s String) ToArray() []string {
	var strArr []string
	length := s.Length()
	if length == 0 {
		return strArr
	}
	for i := 0; i < length; i++ {
		strArr = append(strArr, s.At(i).ToString())
	}
	return strArr
}

// Reverse 反转字符串
func (s String) Reverse() String {
	strArr := s.ToArray()
	arrLen := len(strArr)
	if arrLen == 0 {
		return Empty
	}
	if arrLen == 1 {
		return String(strArr[0])
	}
	str := ""
	for i := arrLen - 1; i >= 0; i-- {
		str += strArr[i]
	}
	return String(str)
}

// Skip 返回忽略前几个字符后的字符串
func (s String) Skip(count int) String {
	if count <= 0 {
		count = 0
	}
	length := s.Length()
	if length == 0 || length <= count {
		return Empty
	}
	return s.Substring(count, -1)
}

// Take 返回指定的前N个字符
func (s String) Take(count int) String {
	length := s.Length()
	if count > length {
		count = length
	}
	return s.Remove(count, length-count)
}

// Append 将字符串附加到当前字符串的末尾
func (s String) Append(str string) String {
	return String(s.ToString() + str)
}

// Repeat 将当前字符串重复n次返回，如果n小于等于0，则返回空字符串
func (s String) Repeat(n int) String {
	if n <= 0 {
		return Empty
	}
	str := strings.Repeat(s.ToString(), n)
	return String(str)
}
