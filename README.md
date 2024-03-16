# Palindrome

This Go package provides a convienent method to test whether given string is palindrome or not. 

## What is palindrome?

A palindrome is a word or characters that are read same from its front and back. It leads a way to very interesting intellectual challenges in computer programming. Read more at [wikipedia](https://en.wikipedia.org/wiki/Palindrome) 

## Usage

### Installation guide

```bash
go get github.com/ubbn/palindrome@v0.1.1
```

### Import it

```go
import github.com/ubbn/palindrome

func main() {
    truth, err := palindrome.IsPalindrome("ABBA")
    fmt.Println(truth) // true
}
```

### More

Package documentation can be found on [pkg.go.dev](https://pkg.go.dev/github.com/ubbn/palindrome)