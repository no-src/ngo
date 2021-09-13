package _strings

import (
	"testing"
)

func TestTrim(t *testing.T) {
	testTrim(t, "hello 世界", "hello 世界")
	testTrim(t, "hello 世界  ", "hello 世界")
	testTrim(t, "  hello 世界", "hello 世界")
	testTrim(t, "  hello 世界  ", "hello 世界")

	testTrim(t, "hello 世界\r\n", "hello 世界")
	testTrim(t, "\r\nhello 世界", "hello 世界")
	testTrim(t, "\r\nhello 世界\r\n", "hello 世界")

	testTrim(t, "hello 世界\n", "hello 世界")
	testTrim(t, "\nhello 世界", "hello 世界")
	testTrim(t, "\nhello 世界\n", "hello 世界")
}

func testTrim(t *testing.T, src, expect string) {
	result := String(src).Trim()
	if expect != result.ToString() {
		t.Fail()
		t.Logf("TestTrim Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestTrimStart(t *testing.T) {
	testTrimStart(t, "hello 世界", "hello 世界")
	testTrimStart(t, "hello 世界  ", "hello 世界  ")
	testTrimStart(t, "  hello 世界", "hello 世界")
	testTrimStart(t, "  hello 世界  ", "hello 世界  ")

	testTrimStart(t, "hello 世界\r\n", "hello 世界\r\n")
	testTrimStart(t, "\r\nhello 世界", "hello 世界")
	testTrimStart(t, "\r\nhello 世界\r\n", "hello 世界\r\n")

	testTrimStart(t, "hello 世界\n", "hello 世界\n")
	testTrimStart(t, "\nhello 世界", "hello 世界")
	testTrimStart(t, "\nhello 世界\n", "hello 世界\n")
}

func testTrimStart(t *testing.T, src, expect string) {
	result := String(src).TrimStart()
	if expect != result.ToString() {
		t.Fail()
		t.Logf("TrimStart Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestTrimEnd(t *testing.T) {
	testTrimEnd(t, "hello 世界", "hello 世界")
	testTrimEnd(t, "hello 世界  ", "hello 世界")
	testTrimEnd(t, "  hello 世界", "  hello 世界")
	testTrimEnd(t, "  hello 世界  ", "  hello 世界")

	testTrimEnd(t, "hello 世界\r\n", "hello 世界")
	testTrimEnd(t, "\r\nhello 世界", "\r\nhello 世界")
	testTrimEnd(t, "\r\nhello 世界\r\n", "\r\nhello 世界")

	testTrimEnd(t, "hello 世界\n", "hello 世界")
	testTrimEnd(t, "\nhello 世界", "\nhello 世界")
	testTrimEnd(t, "\nhello 世界\n", "\nhello 世界")
}

func testTrimEnd(t *testing.T, src, expect string) {
	result := String(src).TrimEnd()
	if expect != result.ToString() {
		t.Fail()
		t.Logf("TrimEnd Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestToString(t *testing.T) {
	testToString(t, "hello", "hello")
	testToString(t, "世界", "世界")
	testToString(t, "hello 世界", "hello 世界")
	testToString(t, "hello 世界\r\n", "hello 世界\r\n")
}

func testToString(t *testing.T, src, expect string) {
	result := String(src).ToString()
	if result != expect {
		t.Fail()
		t.Logf("ToString Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestToLower(t *testing.T) {
	testToLower(t, "Hello", "hello")
	testToLower(t, "世界", "世界")
	testToLower(t, "heLlo 世界", "hello 世界")
	testToLower(t, "HELLO 世界\r\n", "hello 世界\r\n")
}

func testToLower(t *testing.T, src, expect string) {
	result := String(src).ToLower()
	if result.ToString() != expect {
		t.Fail()
		t.Logf("ToLower Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestToUpper(t *testing.T) {
	testToUpper(t, "hello", "HELLO")
	testToUpper(t, "世界", "世界")
	testToUpper(t, "Hello 世界", "HELLO 世界")
	testToUpper(t, "HELLO 世界\r\n", "HELLO 世界\r\n")
}

func testToUpper(t *testing.T, src, expect string) {
	result := String(src).ToUpper()
	if result.ToString() != expect {
		t.Fail()
		t.Logf("ToUpper Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestSubstring(t *testing.T) {
	testSubstring(t, "helloworld", 0, 5, "hello")
	testSubstring(t, "helloworld", 5, 9, "worl")
	testSubstring(t, "helloworld", 5, 10, "world")
	testSubstring(t, "helloworld", 5, 11, "world")
	testSubstring(t, "helloworld", 5, -1, "world")

	testSubstring(t, "你好Golang,测试", 0, 5, "你好Gol")
	testSubstring(t, "你好Golang,测试", 5, 9, "ang,")
	testSubstring(t, "你好Golang,测试", 5, 10, "ang,测")
	testSubstring(t, "你好Golang,测试", 5, 11, "ang,测试")
	testSubstring(t, "你好Golang,测试", 5, 12, "ang,测试")
	testSubstring(t, "你好Golang,测试", 5, -1, "ang,测试")
}

func testSubstring(t *testing.T, src string, start int, end int, expect string) {
	result := String(src).Substring(start, end)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Substring Failed src=[%s] start=%d end=%d expect=[%s] result=[%s] \r\n", src, start, end, expect, result)
	}
}

func TestSplit(t *testing.T) {
	testSplit(t, "你好,世界,hello,world", ",", []string{"你好", "世界", "hello", "world"})
	testSplit(t, "你好,,世界,hello,world", ",", []string{"你好", Empty, "世界", "hello", "world"})
	testSplit(t, "你好,世界,hello,world", "|", []string{"你好,世界,hello,world"})
	testSplit(t, Empty, ",", []string{Empty})
	testSplit(t, ",", ",", []string{Empty, Empty})
}

func testSplit(t *testing.T, src, sep string, expect []string) {
	result := String(src).Split(sep)

	if len(result) != len(expect) {
		t.Fail()
		t.Logf("Split Failed src=[%s] expect=[%v] result=[%v] \r\n", src, expect, result)
		return
	}

	for k, v := range result {
		if v != expect[k] {
			t.Fail()
			t.Logf("Split Failed src=[%s] expect=[%v] result=[%v] \r\n", src, expect, result)
			return
		}
	}
}

func TestReplace(t *testing.T) {
	testReplace(t, "hello", "l", "S", "heSSo")
	testReplace(t, "世界你好", "你好", "再见", "世界再见")
	testReplace(t, "世界你好", "再见", "bye", "世界你好")
}

func testReplace(t *testing.T, src, oldStr, newStr, expect string) {
	result := String(src).Replace(oldStr, newStr)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Replace Failed src=[%s] oldStr=[%s] newStr=[%s] expect=[%s] result=[%s] \r\n", src, oldStr, newStr, expect, result)
	}
}

func TestRemove(t *testing.T) {
	// 截取到最后
	testRemove(t, "helloworld", 5, 5, "hello")
	testRemove(t, "helloworld", 0, 10, Empty)
	testRemove(t, "helloworld", -1, 10, Empty)
	testRemove(t, "helloworld", 100, 10, "helloworld")

	testRemove(t, "你好golang", -1, 8, Empty)
	testRemove(t, "你好golang", 0, 8, Empty)
	testRemove(t, "你好golang", 1, 7, "你")
	testRemove(t, "你好golang", 2, 6, "你好")
	testRemove(t, "你好golang", 3, 5, "你好g")
	testRemove(t, "你好golang", 100, 8, "你好golang")
	// 部分
	testRemove(t, "helloworld", 5, 3, "hellold")
	testRemove(t, "helloworld", 0, 3, "loworld")
	testRemove(t, "helloworld", -1, 300, Empty)
	testRemove(t, "helloworld", 100, 3, "helloworld")

	testRemove(t, "你好golang", -1, 3, "olang")
	testRemove(t, "你好golang", 0, 3, "olang")
	testRemove(t, "你好golang", 1, 3, "你lang")
	testRemove(t, "你好golang", 2, 3, "你好ang")
	testRemove(t, "你好golang", 3, 3, "你好gng")
	testRemove(t, "你好golang", 100, 3, "你好golang")
}

func testRemove(t *testing.T, src string, startIndex int, count int, expect string) {
	result := String(src).Remove(startIndex, count)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Remove Failed src=[%s] startIndex=%d count=%d expect=[%s] result=[%s] \r\n", src, startIndex, count, expect, result)
	}
}

func TestIndexOf(t *testing.T) {
	testIndexOf(t, "helloworld", "", 0)
	testIndexOf(t, "helloworld", "hello", 0)
	testIndexOf(t, "helloworld", "world", 5)
	testIndexOf(t, "helloworld", "好", -1)
	testIndexOf(t, "helloworld", "o", 4)
	testIndexOf(t, "helloworld", "worldhello666", -1)
	testIndexOf(t, "helloworld", "中文", -1)
	testIndexOf(t, "hello", "中文字符搜索-中文", -1)

	testIndexOf(t, "你好golang", "", 0)
	testIndexOf(t, "你好golang", "你", 0)
	testIndexOf(t, "你好golang", "你好", 0)
	testIndexOf(t, "你好golang", "g", 2)
	testIndexOf(t, "你好golang", "golang", 2)
	testIndexOf(t, "你好golang", "x", -1)
	testIndexOf(t, "你好golang", "好g", 1)
	testIndexOf(t, "你好golang", "你好golang", 0)
	testIndexOf(t, "你好golang", "再见2golang", -1)
}

func testIndexOf(t *testing.T, src, findStr string, expect int) {
	result := String(src).IndexOf(findStr)
	if result != expect {
		t.Fail()
		t.Logf("IndexOf Failed src=[%s] findStr=[%s] expect=[%d] result=[%d] \r\n", src, findStr, expect, result)
	}
}

func TestLastIndexOf(t *testing.T) {
	testLastIndexOf(t, "helloworld", "", 10)
	testLastIndexOf(t, "helloworld", "hello", 0)
	testLastIndexOf(t, "helloworld", "world", 5)
	testLastIndexOf(t, "helloworld", "好", -1)
	testLastIndexOf(t, "helloworld", "o", 6)
	testLastIndexOf(t, "helloworld", "worldhello", -1)
	testLastIndexOf(t, "helloworld", "helloworld", 0)
	testLastIndexOf(t, "helloworld", "worldhello666", -1)

	testLastIndexOf(t, "你好golang中国", "", 10)
	testLastIndexOf(t, "你好golang", "你", 0)
	testLastIndexOf(t, "你好golang", "你好", 0)
	testLastIndexOf(t, "你好golang", "g", 7)
	testLastIndexOf(t, "你好golang", "golang", 2)
	testLastIndexOf(t, "你好golang", "x", -1)
	testLastIndexOf(t, "你好golang", "好g", 1)
	testLastIndexOf(t, "你好golang", "你好golang", 0)
	testLastIndexOf(t, "你好golang你好golang", "你好golang", 8)
	testLastIndexOf(t, "你好golang", "再见2golang", -1)
}

func testLastIndexOf(t *testing.T, src, findStr string, expect int) {
	result := String(src).LastIndexOf(findStr)
	if result != expect {
		t.Fail()
		t.Logf("LastIndexOf Failed src=[%s] findStr=[%s] expect=[%d] result=[%d] \r\n", src, findStr, expect, result)
	}
}

func TestIsNullOrEmpty(t *testing.T) {
	testIsNullOrEmpty(t, "hello", false)
	testIsNullOrEmpty(t, Empty, true)
	testIsNullOrEmpty(t, " ", false)
	testIsNullOrEmpty(t, "\n", false)
	testIsNullOrEmpty(t, "\r\n", false)
	testIsNullOrEmpty(t, "\t", false)
}

func testIsNullOrEmpty(t *testing.T, src string, expect bool) {
	result := String(src).IsNullOrEmpty()
	if result != expect {
		t.Fail()
		t.Logf("IsNullOrEmpty Failed src=[%s] expect=[%v] result=[%v] \r\n", src, expect, result)
	}
}

func TestIsNullOrWhiteSpace(t *testing.T) {
	testIsNullOrWhiteSpace(t, "hello", false)
	testIsNullOrWhiteSpace(t, Empty, true)
	testIsNullOrWhiteSpace(t, " ", true)
	testIsNullOrWhiteSpace(t, "\n", true)
	testIsNullOrWhiteSpace(t, "\r\n", true)
	testIsNullOrWhiteSpace(t, "\t", true)
}

func testIsNullOrWhiteSpace(t *testing.T, src string, expect bool) {
	result := String(src).IsNullOrWhiteSpace()
	if result != expect {
		t.Fail()
		t.Logf("IsNullOrWhiteSpace Failed src=[%s] expect=[%v] result=[%v] \r\n", src, expect, result)
	}
}

func TestInsert(t *testing.T) {
	testInsert(t, "helloworld", 5, "golang", "hellogolangworld")
	testInsert(t, "中文测试", 2, "golang", "中文golang测试")

	testInsert(t, "你好golang", 0, "测试", "测试你好golang")
	testInsert(t, "你好golang", -1, "测试", "测试你好golang")
	testInsert(t, "你好golang", 7, "测试", "你好golan测试g")
	testInsert(t, "你好golang", 8, "测试", "你好golang测试")
	testInsert(t, "你好golang", 9, "测试", "你好golang测试")
}

func testInsert(t *testing.T, src string, startIndex int, value string, expect string) {
	result := String(src).Insert(startIndex, value)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Insert Failed src=[%s] startIndex=%d value=[%s] expect=[%s] result=[%s] \r\n", src, startIndex, value, expect, result)
	}
}

func TestStartsWith(t *testing.T) {
	testStartsWith(t, "HelloWorld", "Hello", true)
	testStartsWith(t, "你好世界", "Hello", false)
	testStartsWith(t, "你好世界", "你好", true)
	testStartsWith(t, "你好世界", "世界", false)
	testStartsWith(t, " 你好世界", Empty, true)
	testStartsWith(t, "你好世界", Empty, true)

	testStartsWith(t, " 你好世界", " ", true)
	testStartsWith(t, "你好世界", " ", false)

	testStartsWith(t, "\r\n你好世界", "\r\n", true)
	testStartsWith(t, "你好世界", "\r\n", false)

	testStartsWith(t, "\t你好世界", "\t", true)
	testStartsWith(t, "你好世界", "\t", false)
}

func testStartsWith(t *testing.T, src, value string, expect bool) {
	result := String(src).StartsWith(value)
	if result != expect {
		t.Fail()
		t.Logf("StartsWith Failed src=[%s] value=[%s] expect=[%v] result=[%v] \r\n", src, value, expect, result)
	}
}

func TestEndsWith(t *testing.T) {
	testEndsWith(t, "HelloWorld", "World", true)
	testEndsWith(t, "你好世界", "Hello", false)
	testEndsWith(t, "你好世界", "你好", false)
	testEndsWith(t, "你好世界", "世界", true)
	testEndsWith(t, " 你好世界", Empty, true)
	testEndsWith(t, "你好世界", Empty, true)

	testEndsWith(t, "你好世界 ", " ", true)
	testEndsWith(t, "你好世界", " ", false)

	testEndsWith(t, "你好世界\r\n", "\r\n", true)
	testEndsWith(t, "你好世界", "\r\n", false)

	testEndsWith(t, "你好世界\t", "\t", true)
	testEndsWith(t, "你好世界", "\t", false)
}

func testEndsWith(t *testing.T, src, value string, expect bool) {
	result := String(src).EndsWith(value)
	if result != expect {
		t.Fail()
		t.Logf("EndsWith Failed src=[%s] value=[%s] expect=[%v] result=[%v] \r\n", src, value, expect, result)
	}
}

func TestContains(t *testing.T) {
	testContains(t, "HelloWorld", "World", true)
	testContains(t, "你好世界", "Hello", false)
	testContains(t, "你好世界", "你好", true)
	testContains(t, "你好世界", "世界", true)
	testContains(t, " 你好世界", Empty, true)
	testContains(t, "你好世界", Empty, true)

	testContains(t, "你好世界 ", " ", true)
	testContains(t, "你好世界", " ", false)

	testContains(t, "你好世界\r\n", "\r\n", true)
	testContains(t, "你好世界", "\r\n", false)

	testContains(t, "你好世界\t", "\t", true)
	testContains(t, "你好世界", "\t", false)
}

func testContains(t *testing.T, src, value string, expect bool) {
	result := String(src).Contains(value)
	if result != expect {
		t.Fail()
		t.Logf("Contains Failed src=[%s] value=[%s] expect=[%v] result=[%v] \r\n", src, value, expect, result)
	}
}

func TestLength(t *testing.T) {
	testLength(t, "Hello", 5)
	testLength(t, "你好Golang", 8)
	testLength(t, "你好 Golang", 9)
	testLength(t, "你好 Golang\r\n", 11)
	testLength(t, "你好 Golang\n", 10)
	testLength(t, "你好 Golang\t", 10)
}

func testLength(t *testing.T, src string, expect int) {
	result := String(src).Length()
	if result != expect {
		t.Fail()
		t.Logf("Length Failed src=[%s] expect=[%d] result=[%d] \r\n", src, expect, result)
	}
}

func TestFirst(t *testing.T) {
	testFirst(t, "", "")
	testFirst(t, "Hello", "H")
	testFirst(t, "\r\nHello", "\r")
	testFirst(t, "\tHello", "\t")
	testFirst(t, "测试Hello", "测")
	testFirst(t, "你好世界", "你")
}

func testFirst(t *testing.T, src string, expect string) {
	result := String(src).First()
	if result.ToString() != expect {
		t.Fail()
		t.Logf("First Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestLast(t *testing.T) {
	testLast(t, "Hello", "o")
	testLast(t, "Hello\r\n", "\n")
	testLast(t, "Hello\t", "\t")
	testLast(t, "Hello测试", "试")
	testLast(t, "你好世界", "界")
}

func testLast(t *testing.T, src string, expect string) {
	result := String(src).Last()
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Last Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestAt(t *testing.T) {
	testAt(t, "Hello\r\n", -1, "")
	testAt(t, "Hello\r\n", 0, "H")
	testAt(t, "Hello\t", 1, "e")
	testAt(t, "Hello", 2, "l")
	testAt(t, "Hello", 10, "")

	testAt(t, "世界Hello", -1, "")
	testAt(t, "世界Hello", 0, "世")
	testAt(t, "世界Hello", 1, "界")
	testAt(t, "世界Hello", 2, "H")
	testAt(t, "世界Hello", 10, "")
}

func testAt(t *testing.T, src string, index int, expect string) {
	result := String(src).At(index)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("At Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestToArray(t *testing.T) {
	str := String("世界Hi")
	strArr := str.ToArray()
	if strArr[0] == "世" && strArr[1] == "界" && strArr[2] == "H" && strArr[3] == "i" {
		//t.Logf("ToArray Success str=[%s] strArr=[%v] \r\n", str, strArr)
	} else {
		t.Logf("ToArray Failed str=[%s] strArr=[%v] \r\n", str, strArr)
		t.Fail()
	}
}

func TestReverse(t *testing.T) {
	testReverse(t, "Hello", "olleH")
	testReverse(t, "Hello\r\n", "\n\rolleH")
	testReverse(t, "Hello\t", "\tolleH")
	testReverse(t, "Hello测试", "试测olleH")
	testReverse(t, "你好世界", "界世好你")
}

func testReverse(t *testing.T, src string, expect string) {
	result := String(src).Reverse()
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Reverse Failed src=[%s] expect=[%s] result=[%s] \r\n", src, expect, result)
	}
}

func TestSkip(t *testing.T) {
	testSkip(t, "helloworld", 5, "world")
	testSkip(t, "helloworld", 0, "helloworld")
	testSkip(t, "helloworld", -1, "helloworld")
	testSkip(t, "helloworld", 100, "")

	testSkip(t, "你好golang", -1, "你好golang")
	testSkip(t, "你好golang", 0, "你好golang")
	testSkip(t, "你好golang", 1, "好golang")
	testSkip(t, "你好golang", 2, "golang")
	testSkip(t, "你好golang", 3, "olang")
	testSkip(t, "你好golang", 100, "")
}

func testSkip(t *testing.T, src string, count int, expect string) {
	result := String(src).Skip(count)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Skip Failed src=[%s] count=%d expect=[%s] result=[%s] \r\n", src, count, expect, result)
	}
}

func TestTake(t *testing.T) {
	testTake(t, "helloworld", 5, "hello")
	testTake(t, "helloworld", 0, Empty)
	testTake(t, "helloworld", -1, Empty)
	testTake(t, "helloworld", 100, "helloworld")

	testTake(t, "你好golang", -1, Empty)
	testTake(t, "你好golang", 0, Empty)
	testTake(t, "你好golang", 1, "你")
	testTake(t, "你好golang", 2, "你好")
	testTake(t, "你好golang", 3, "你好g")
	testTake(t, "你好golang", 100, "你好golang")
}

func testTake(t *testing.T, src string, count int, expect string) {
	result := String(src).Take(count)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Take Failed src=[%s] count=%d expect=[%s] result=[%s] \r\n", src, count, expect, result)
	}
}

func TestAppend(t *testing.T) {
	testAppend(t, "Hello", "World", "HelloWorld")
	testAppend(t, "Hello", "World\r\n", "HelloWorld\r\n")
	testAppend(t, "Hello", "世界", "Hello世界")
	testAppend(t, "Hello", "\t世界", "Hello\t世界")
}

func testAppend(t *testing.T, str string, appendStr string, expect string) {
	result := String(str).Append(appendStr)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Append Failed str=[%s] expect=[%s] result=[%s] \r\n", str, expect, result)
	}
}

func TestRepeat(t *testing.T) {
	testRepeat(t, "Hello世界", 2, "Hello世界Hello世界")
	testRepeat(t, "Hello世界", 1, "Hello世界")
	testRepeat(t, "Hello", 0, "")
	testRepeat(t, "Hello", -1, "")
}

func testRepeat(t *testing.T, str string, n int, expect string) {
	result := String(str).Repeat(n)
	if result.ToString() != expect {
		t.Fail()
		t.Logf("Repeat Failed str=[%s] expect=[%s] result=[%s] \r\n", str, expect, result)
	}
}
