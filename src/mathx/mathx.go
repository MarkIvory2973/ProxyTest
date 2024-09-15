package mathx

import (
	"cmp"
	"slices"

	"golang.org/x/exp/constraints"
	"gonum.org/v1/gonum/stat"
)

type SliceType[T any] interface {
	~[]T
}

func Std(x []float64) float64 {
	return stat.StdDev(x, nil)
}

func Median(x []float64) float64 {
	slices.Sort(x)
	return stat.Quantile(0.5, stat.Empirical, x, nil)
}

func Norm(x []float64) []float64 {
	stdX := Std(x)

	var y []float64
	for _, xi := range x {
		y = append(y, xi/stdX)
	}

	return y
}

func Argsort[T constraints.Ordered, S SliceType[T]](x S, reverse bool) []int {
	type Item struct {
		index int
		value T
	}

	var items []Item
	for index, value := range x {
		items = append(items, Item{index: index, value: value})
	}

	slices.SortStableFunc(items, func(a, b Item) int {
		if reverse {
			return cmp.Compare(b.value, a.value)
		} else {
			return cmp.Compare(a.value, b.value)
		}
	})

	var indexs []int
	for _, item := range items {
		indexs = append(indexs, item.index)
	}

	return indexs
}
