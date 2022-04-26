package set

type ImmutableSet[T comparable] struct {
	data map[T]struct{}
}

func NewImmutable[T comparable](ts ...T) ImmutableSet[T] {
	set := ImmutableSet[T]{data: make(map[T]struct{})}
	for _, t := range ts {
		set.data[t] = struct{}{}
	}
	return set
}

func (s ImmutableSet[T]) Add(ts ...T) ImmutableSet[T] {
	return NewImmutable(append(s.List(), ts...)...)
}

func (s ImmutableSet[T]) Remove(ts ...T) ImmutableSet[T] {
	sc := s.Copy()
	for _, t := range ts {
		delete(sc.data, t)
	}
	return sc
}

func (s ImmutableSet[T]) Clear() ImmutableSet[T] {
	return NewImmutable[T]()
}

func (s ImmutableSet[T]) Copy() ImmutableSet[T] {
	return NewImmutable(s.List()...)
}

func (s ImmutableSet[T]) Contains(t T) bool {
	_, ok := s.data[t]
	return ok
}

func (s ImmutableSet[T]) Difference(other ImmutableSet[T]) ImmutableSet[T] {
	return s.Remove(other.List()...)
}

func (s ImmutableSet[T]) Intersection(other ImmutableSet[T]) ImmutableSet[T] {
	if s.Len() > other.Len() {
		return other.Intersection(s)
	}
	result := make([]T, 0, s.Len())
	for k := range s.data {
		if other.Contains(k) {
			result = append(result, k)
		}
	}
	return NewImmutable(result...)
}

func (s ImmutableSet[T]) Union(other ImmutableSet[T]) ImmutableSet[T] {
	return s.Add(other.List()...)
}

func (s ImmutableSet[T]) List() []T {
	res := make([]T, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}

func (s ImmutableSet[T]) Map() map[T]struct{} {
	res := make(map[T]struct{})
	for k, v := range s.data {
		res[k] = v
	}
	return res
}

func (s ImmutableSet[T]) Len() int {
	return len(s.data)
}
