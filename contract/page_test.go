package _contract

import "testing"

func TestGetPage(t *testing.T) {
	testGetPage(t, 0, 0, 0, 0, 0)
	testGetPage(t, -1, -1, -1, 0, 0)
	testGetPage(t, 5, 1, 3, 3, 5)
	testGetPage(t, 5, 2, 3, 0, 0)
	testGetPage(t, 100, 2, 5, 10, 15)
	testGetPage(t, 100, -2, 5, 0, 0)
	testGetPage(t, 100, 0, 5, 0, 5)
}

func testGetPage(t *testing.T, total, pageIndex, pageSize, expectStartIndex, expectEndIndex int64) {
	startIndex, endIndex := GetPage(total, pageIndex, pageSize)
	if startIndex != expectStartIndex || endIndex != expectEndIndex {
		t.Logf("testGetPage Failed total=%d pageIndex=%d  pageSize=%d expectStartIndex=%d expectEndIndex=%d startIndex=%d endIndex=%d", total, pageIndex, pageSize, expectStartIndex, expectEndIndex, startIndex, endIndex)
		t.Fail()
	}
}
