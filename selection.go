package peco

import "github.com/google/btree"

// NewSelection creates a new empty Selection
func NewSelection() *Selection {
	s := &Selection{}
	s.Reset()
	return s
}

// Add adds a new line to the selection. If the line already
// exists in the selection, it is silently ignored
func (s *Selection) Add(l Line) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tree.ReplaceOrInsert(l)
}

// Remove removes the specified line from the selection
func (s *Selection) Remove(l Line) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tree.Delete(l)
}

func (s *Selection) Reset() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tree = btree.New(32)
}

func (s *Selection) Has(x Line) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.tree.Has(x)
}

func (s *Selection) Len() int {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	return s.tree.Len()
}

func (s *Selection) Ascend(i btree.ItemIterator) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.tree.Ascend(i)
}

