package set

type Set[T comparable] struct {
	data map[T]struct{}
}

func New[T comparable](ts ...T) *Set[T] {
	set := &Set[T]{data: make(map[T]struct{})}
	set.Add(ts...)
	return set
}

func (s *Set[T]) Add(ts ...T) *Set[T] {
	for _, t := range ts {
		s.data[t] = struct{}{}
	}
	return s
}

func (s *Set[T]) Remove(ts ...T) *Set[T] {
	for _, t := range ts {
		delete(s.data, t)
	}
	return s
}

func (s *Set[T]) Clear() *Set[T] {
	s.data = make(map[T]struct{})
	return s
}

func (s *Set[T]) Copy() *Set[T] {
	return New(s.List()...)
}

func (s *Set[T]) Contains(t T) bool {
	_, ok := s.data[t]
	return ok
}

func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	return s.Remove(other.List()...)
}

func (s *Set[T]) Intersection(other *Set[T]) *Set[T] {
	if s.Len() > other.Len() {
		return other.Intersection(s)
	}
	result := make([]T, 0, s.Len())
	for k := range s.data {
		if other.Contains(k) {
			result = append(result, k)
		}
	}
	return New(result...)
}

func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	return s.Add(other.List()...)
}

func (s *Set[T]) List() []T {
	res := make([]T, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

func (s *Set[T]) Map() map[T]struct{} {
	res := make(map[T]struct{})
	for k, v := range s.data {
		res[k] = v
	}
	return res
}

func (s *Set[T]) Len() int {
	return len(s.data)
}
