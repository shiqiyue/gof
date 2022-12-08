package sets

import (
	"github.com/pkg/errors"
	"sync"
)

type Int64SortSet struct {
	item []int64
	lock *sync.RWMutex
}

// 新建SortSet
func NewInt64SortSet() *Int64SortSet {

	return &Int64SortSet{
		item: make([]int64, 0),
		lock: new(sync.RWMutex)}
}

// 返回SortSet大小
func (s *Int64SortSet) Size() int {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return len(s.item)
}

// 返回数据
func (s *Int64SortSet) Data() []int64 {
	return s.item
}

// 切片
// j如果为-1，j将设置为切片长度
func (s *Int64SortSet) Sub(i, j int) ([]int64, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	l := len(s.item)
	if i < 0 {
		return nil, errors.New("i can not less than 0")
	}
	if i > l {
		return nil, errors.New("i is out of index")
	}
	if j < 0 {
		j = l
	}
	if j > l {
		return nil, errors.New("j is out of index")
	}
	if i > j {
		return nil, errors.New("i is large than j")
	}
	return s.item[i:j], nil
}

// 判断原始是否存在
func (s *Int64SortSet) exist(d int64) bool {

	return s.getPosition(d) != -1
}

// 添加
// 返回添加的元素的序号
// 返回值为-1时候，表示元素已存在
func (s *Int64SortSet) Add(i int64) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.exist(i) {
		return -1
	}
	pos := s.getAddPosition(i)
	s.addItem(i, pos)
	return pos
}

// 批量添加
func (s *Int64SortSet) AddAll(strs ...int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, v := range strs {
		if s.exist(v) {
			continue
		}
		pos := s.getAddPosition(v)
		s.addItem(v, pos)
	}
}

// 删除
// 返回删除的元素的序号
// 返回值为-1时候，表示元素不存在
func (s *Int64SortSet) Remove(i int64) int {
	s.lock.Lock()
	defer s.lock.Unlock()
	pos := s.getPosition(i)
	if pos != -1 {
		s.removeItem(pos)
	}
	return pos
}

// 批量删除
func (s *Int64SortSet) RemoveAll(strs ...int64) {
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, v := range strs {
		pos := s.getPosition(v)
		if pos != -1 {
			s.removeItem(pos)
		}
	}
}

// 添加元素
// 返回添加的元素的序号
func (s *Int64SortSet) addItem(i int64, pos int) {
	after := append([]int64{}, s.item[pos:]...)
	s.item = append(s.item[0:pos], i)
	s.item = append(s.item, after...)
}

// 删除元素
func (s *Int64SortSet) removeItem(pos int) {
	s.item = append(s.item[:pos], s.item[pos+1:]...)
}

// 获取元素的位置
// 如果元素不存在，则返回-1
func (s *Int64SortSet) getPosition(d int64) int {
	for i, v := range s.item {
		if d == v {
			return i
		}
	}
	return -1
}

// 获取新元素插入的位置
func (s *Int64SortSet) getAddPosition(d int64) int {
	itemLen := len(s.item)
	if itemLen == 0 {
		return 0
	}
	for i := 0; i < itemLen; i++ {
		if d < s.item[i] {
			return i
		}
	}
	return itemLen
}

// 交集
// 返回一个新的SortSet
func (s *Int64SortSet) Intersection(s2 *Int64SortSet) (*Int64SortSet, error) {
	if s2 == nil || s2.item == nil {
		return nil, errors.New("param can not be empty")
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := NewInt64SortSet()
	for _, v := range s.item {
		if s2.exist(v) {
			newSet.item = append(newSet.item, v)
		}
	}
	return newSet, nil
}

// 交集
// 返回一个新的SortSet
func (s *Int64SortSet) Intersections(s2 ...*Int64SortSet) (*Int64SortSet, error) {
	eset := s.Copy()
	var err error
	for _, set := range s2 {
		eset, err = eset.Intersection(set)
		if err != nil {
			return nil, err
		}
	}
	return eset, nil
}

// 复制
func (s *Int64SortSet) Copy() *Int64SortSet {
	eset := NewInt64SortSet()
	eset.AddAll(s.item...)
	return eset
}

// 并集
// 返回一个新的SortSet
func (s *Int64SortSet) Union(s2 *Int64SortSet) (*Int64SortSet, error) {
	if s2 == nil || s2.item == nil {
		return nil, errors.New("param can not be empty")
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := NewInt64SortSet()
	for _, v := range s.item {
		newSet.Add(v)
	}
	for _, v := range s2.item {
		newSet.Add(v)
	}
	return newSet, nil
}

// 差集
// 返回一个新的SortSet
func (s *Int64SortSet) Difference(s2 *Int64SortSet) (*Int64SortSet, error) {
	if s2 == nil || s2.item == nil {
		return nil, errors.New("param can not be empty")
	}
	s.lock.RLock()
	defer s.lock.RUnlock()
	newSet := NewInt64SortSet()
	for _, v := range s.item {
		if !s2.exist(v) {
			newSet.item = append(newSet.item, v)
		}
	}
	return newSet, nil
}
