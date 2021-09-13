package _array

import "testing"

func TestDistinct(t *testing.T) {
	srcArr := []int{1, 2, 3, 4, 4, 5, 6, 6, 7}
	targetArr := Distinct(srcArr)
	if len(targetArr) != 7 {
		t.Fail()
	}
}

func TestGetShuffleArray(t *testing.T) {
	shuffleArr := GetShuffleArray(5)
	if len(shuffleArr) != 5 {
		t.Fail()
	}
}

func TestContain(t *testing.T) {
	testContain(t, []string{"hello", "world", "\r\n", "\t", "测试", "\b"}, "world", true)
	testContain(t, []string{"hello", "world", "\r\n", "\t", "测试", "\b"}, "orld", false)
	testContain(t, []string{"hello", "world", "\r\n", "\t", "测试", "\b"}, "测试", true)
	testContain(t, []string{"hello", "world", "\r\n", "\t", "测试", "\b"}, "b", false)
	testContain(t, []string{"hello", "world", "\r\n", "\t", "测试", "\b"}, "\r\n", true)
}

func testContain(t *testing.T, source []string, target string, expect bool) {
	result := Contain(source, target)
	if expect != result {
		t.Fail()
		t.Logf("TestContain Failed source=[%v] target=[%s] expect=[%v] result=[%v] \r\n", source, target, expect, result)
	}
}

func TestContainStr(t *testing.T) {
	testContainStr(t, "hello,world,再见\r\n", "world", ",", true)
	testContainStr(t, "hello,world,再见\r\n", "再见", ",", false)
	testContainStr(t, "hello,world,再见\r\n", "hello", ",", true)
	testContainStr(t, "hello,world,再见\r\n", "hell", ",", false)
}

func testContainStr(t *testing.T, source string, target string, seq string, expect bool) {
	result := ContainStr(source, target, seq)
	if expect != result {
		t.Fail()
		t.Logf("TestContainStr Failed source=[%s] target=[%s] seq=[%s] expect=[%v] result=[%v] \r\n", source, target, seq, expect, result)
	}
}

func TestContainStrsOne(t *testing.T) {
	testContainStrsOne(t, "hello,world,再见\r\n", []string{"world", "再见"}, ",", true)
	testContainStrsOne(t, "hello,world,再见\r\n", []string{"he", "再见"}, ",", false)
	testContainStrsOne(t, "hello,world,再见\r\n", []string{""}, ",", false)
	testContainStrsOne(t, "hello,world,再见\r\n", []string{"\r\n"}, ",", false)
}

func testContainStrsOne(t *testing.T, source string, target []string, seq string, expect bool) {
	result := ContainStrsOne(source, target, seq)
	if expect != result {
		t.Fail()
		t.Logf("TestContainStrsOne Failed source=[%s] target=[%v] seq=[%s] expect=[%v] result=[%v] \r\n", source, target, seq, expect, result)
	}
}

func TestContainStrsAll(t *testing.T) {
	testContainStrsAll(t, "hello,world,再见\r\n", []string{"world", "再见"}, ",", false)
	testContainStrsAll(t, "hello,world,再见\r\n", []string{"world", "hello"}, ",", true)
	testContainStrsAll(t, "hello,world,再见\r\n", []string{"hello", "hello"}, ",", true)
	testContainStrsAll(t, "hello,world,再见\r\n", []string{"he", "再见"}, ",", false)
	testContainStrsAll(t, "hello,world,再见\r\n", []string{""}, ",", false)
	testContainStrsAll(t, "hello,world,再见\r\n", []string{"\r\n"}, ",", false)
}

func testContainStrsAll(t *testing.T, source string, target []string, seq string, expect bool) {
	result := ContainStrsAll(source, target, seq)
	if expect != result {
		t.Fail()
		t.Logf("TestContainStrsAll Failed source=[%s] target=[%v] seq=[%s] expect=[%v] result=[%v] \r\n", source, target, seq, expect, result)
	}
}

func TestStrEqual(t *testing.T) {
	testStrEqual(t, []string{"world", "再见"}, []string{"world", "再见"}, true)
	testStrEqual(t, []string{"world", "bye"}, []string{"world", "再见"}, false)
	testStrEqual(t, []string{"world", "再见"}, []string{"world", "再见\r\n"}, false)
	testStrEqual(t, []string{"再见", "world"}, []string{"world", "再见"}, false)
}

func testStrEqual(t *testing.T, left []string, right []string, expect bool) {
	result := StrEqual(left, right)
	if expect != result {
		t.Fail()
		t.Logf("TestStrEqual Failed left=[%v] right=[%v] expect=[%v] result=[%v] \r\n", left, right, expect, result)
	}
}
