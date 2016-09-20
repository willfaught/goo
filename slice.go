package goo

var _ Iterator = &sliceIterator{}

type Slice interface {
	Append(v ...interface{}) Slice

	AppendSlice(s Slice) Slice

	Cap() int

	Copy(s Slice) int

	Get(i int) interface{}

	Len() int

	Make(l, c int) Slice

	Set(i int, v interface{})

	Slice(i, j int) Slice

	SliceCap(i, j, m int) Slice
}

// TODO: All, And, Any, Combine, CombineRepeats, Concat, Contains, Cut,
// Difference, Drop, DropEnd, DropWith, DropWithEnd, Each, Expand, Extend,
// Filter, FindIndex, FindIndexEnd, FindIndexes, First, Foldl, Foldl1, Foldr,
// Foldr1, Greatest, Group, Head, Index, Indexes, Infix, Init, Inits, Insert,
// Insert, Intersect, Intersperse, Last, Least, Map, Or, Partition, Permute,
// Prefix, Product, Remove, RemoveAll, RemoveFast, RemoveFirst, RemoveLast,
// RemoveRange, Replicate, Replicate, Reverse, Sort, SortStable, Split, Subset,
// Subslice, Subslices, Suffix, Sum, Tail, Tails, Take, TakeEnd, TakeWith,
// TakeWithEnd, Transpose, Union, Zip, Zip3-7, ZipWith, ZipWith3-7

func Dequeue(s Slice) (Slice, interface{}) {
	return s.Slice(1, s.Len()-1), s.Get(0)
}

func Enqueue(s Slice, v interface{}) Slice {
	return s.Append(v)
}

func SliceIterator(s Slice) Iterator {
	return &sliceIterator{n: s.Len(), s: s}
}

func Pop(s Slice) (Slice, interface{}) {
	var l = s.Len()

	return s.Slice(0, l-2), s.Get(l - 1)
}

func Push(s Slice, v interface{}) Slice {
	return s.Append(v)
}

type sliceIterator struct {
	i int

	n int

	s Slice
}

func (i sliceIterator) More() bool {
	return i.i < i.n
}

func (i *sliceIterator) Next() interface{} {
	var v = i.s.Get(i.i)

	i.i++

	return v
}
