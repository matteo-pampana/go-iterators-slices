# Understanding Iterators In Go 1.23

This project demonstrates how to create and use iterators for slices in Go. It includes two types of iterators: one that yields both the index and value of each element in the slice, and another that yields only the value.

You can read the medium article here: [Understanding Iterators In Go 1.23](https://medium.com/@matteopampana/understanding-iterators-in-go-1-23-eac5308dd49b)

Support me here: [Buy me a coffee](https://www.buymeacoffee.com/matteo.pampana)

## Requirements

- Go 1.23.3 or later

## Installation

Clone the repository and navigate to the project directory:

```sh
git clone https://github.com/matteo-pampana/go-iterators-slices.git
cd go-iterators-slices
```

## Usage

### sliceIteratorWithIndex

This function returns an iterator that yields both the index and value of each element in the slice.

```go
func sliceIteratorWithIndex(slice []int) iter.Seq2[int, int]
```

### sliceIterator

This function returns an iterator that yields only the value of each element in the slice.

```go
func sliceIterator(slice []int) iter.Seq[int]
```

### Example

The `main` function demonstrates how to use both iterators:

```go
func main() {
    mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

    for index, value := range sliceIteratorWithIndex(mySlice) {
        fmt.Println("Index:", index, "Value:", value)
    }

    fmt.Println("--------------------")

    for value := range sliceIterator(mySlice) {
        fmt.Println("Value:", value)
    }
}
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

Made with ❤️ by Matteo Pampana - 2024
