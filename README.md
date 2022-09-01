# search [![Test Status](https://github.com/imbue11235/search/workflows/Go/badge.svg)](https://github.com/imbue11235/search/actions?query=workflow:Go) [![Go Reference](https://pkg.go.dev/badge/github.com/imbue11235/search.svg)](https://pkg.go.dev/github.com/imbue11235/search)

> A generic simple binary search library to allow finding the index of any element in an ordered slice of comparable elements

## 🛠  Installation

Make sure to have Go installed (Version `1.18` or higher).

Install `search` with `go get`:

```sh
$ go get -u github.com/imbue11235/search
```

## 💻  Usage

### Using [constraints.Ordered](https://pkg.go.dev/golang.org/x/exp/constraints#Ordered)

Search slices of primitive types like int, float etc.

```go
search.Slice([]int{1, 2, 3, 4, 5, 6}, 2) // => index: 1
```

### Using [Comparable[T]](https://pkg.go.dev/github.com/imbue11235/search#Comparable)

Search slices of custom structs

```go
type ComparableStruct struct {
	value int
}

func (c ComparableStruct) CompareTo(other ComparableStruct) int {
	return c.value
}

element := ComparableStruct{2}
list := []ComparableStruct{{1}, {2}, {3}, {4}, {5}, {6}}

search.ComparableSlice(list, element) // => index: 1
```

## 📜 License

This project is licensed under the [MIT license](LICENSE).