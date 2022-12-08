package dataloaders

import "testing"

func TestSortByOrder(t *testing.T) {
	order := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	src := testGetExample()
	dst := SortByOrder(
		order, src,
		func(data *testStruct) int {
			return data.ID
		},
	)
	t.Log(src)
	t.Log(dst)
}

func TestSort2DByOrder(t *testing.T) {
	order := []string{"a", "b", "c"}
	src := testGetExample()
	dst := Sort2DByOrder(
		order, src,
		func(data *testStruct) string {
			return data.Name
		},
		func(a, b *testStruct) bool {
			return a.ID < b.ID
		},
	)
	t.Log(src)
	t.Log(dst)
}

type testStruct struct {
	ID   int
	Name string
}

func testGetExample() []*testStruct {
	return []*testStruct{
		&testStruct{
			ID:   1,
			Name: "a",
		},
		&testStruct{
			ID:   3,
			Name: "a",
		},
		&testStruct{
			ID:   5,
			Name: "b",
		},
		&testStruct{
			ID:   7,
			Name: "b",
		},
		&testStruct{
			ID:   9,
			Name: "b",
		},
	}
}
