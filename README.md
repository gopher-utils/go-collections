# golang-linq

<img src="https://github.com/gopinathr143/golang-linq/actions/workflows/pr-actions.yml/badge.svg" alt="Build">
<a href="https://github.com/gopinathr143/golang-linq/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-purple" alt="License"></a>

Brings **L**anguage **In**tegrated **Q**uery features through basic collections into your go project.

**Features**:

<p>
&nbsp;&nbsp;&nbsp;&nbsp;✅&nbsp; List datastructure of both primitive types and structs<br />
&nbsp;&nbsp;&nbsp;&nbsp;❌&nbsp; Set datastructure of both primitive types and structs<br />
&nbsp;&nbsp;&nbsp;&nbsp;❌&nbsp; Map datastructure of both primitive types and structs<br />
&nbsp;&nbsp;&nbsp;&nbsp;❌&nbsp; Sorting and searching of collections<br />
</p>

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

```shell
go get github.com/gopinathr143/golang-linq
```

## Usage

Below is an example of how you can use golang-linq in your project.

```go
package main

import (
	"fmt"

	"github.com/gopinathr143/golang-linq/collections/list"
)

func main() {
	// Create a new empty list of integers with capacity 10
	emptyList := list.New[int](10)
	// Insert a number
	emptyList.Add(1)
	fmt.Println(emptyList) // Prints [1]

	// Create a list of integers with repeating ones of size 5
	repeatingList := list.Repeating(1, 5)
	fmt.Println(repeatingList) // Prints [1,1,1,1,1]

	// Create a list from array
	fromArrayList := list.From([]int{2, 3, 4})
	fmt.Println(fromArrayList) // Prints [2,3,4]

	// Concatenate two lists together
	emptyList.Extend(fromArrayList)
	fmt.Println(emptyList) // Prints [1,2,3,4]
}

```

Please check the [documentation](https://pkg.go.dev/github.com/gopinathr143/golang-linq) for more comprehensive usage.

## Contributing

Any contributions to this repo are always welcome. Please check the [Contribution Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
