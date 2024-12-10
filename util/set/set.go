package set

type Set[T comparable] map[T]struct{}

func New[T comparable](items ...T) Set[T] {
	s := make(Set[T], len(items))
	for _, i := range items {
		s.Add(i)
	}
	return s
}

func (s Set[T]) Exists(val T) bool {
	_, ok := s[val]
	return ok
}

func (s Set[T]) Add(val T) {
	s[val] = struct{}{}
}

func (s Set[T]) Remove(val T) {
	delete(s, val)
}
