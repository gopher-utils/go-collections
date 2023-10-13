# go-collections
<p>
<img src="https://github.com/gopher-utils/go-collections/actions/workflows/build.yml/badge.svg" alt="Build">
<a href="https://github.com/gopher-utils/go-collections/blob/main/LICENSE"><img src="https://img.shields.io/badge/license-MIT-purple" alt="License"></a>
[![codecov](https://codecov.io/gh/gopher-utils/go-collections/graph/badge.svg?token=Z6SU2GOHDI)](https://codecov.io/gh/gopher-utils/go-collections)
</p>


Basic collections with Go using generics.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Installation

```shell	
go get github.com/gopher-utils/go-collections
```

## Usage

Below is an example of how you can use go-collections in your project.

```go
package main

import (
	"fmt"

	"github.com/gopher-utils/go-collections/collections"
)

func main() {
	// Create a new empty list of integers with capacity 10
	emptyList := collections.NewList[int](10)
	// Insert a number
	emptyList.Add(1)
	fmt.Println(emptyList) // Prints [1]

	// Create a list of integers with repeating ones of size 5
	repeatingList := collections.RepeatingList(1, 5)
	fmt.Println(repeatingList) // Prints [1,1,1,1,1]

	// Create a list from array
	fromArrayList := collections.ToList([]int{2, 3, 4})
	fmt.Println(fromArrayList) // Prints [2,3,4]

	// Concatenate two lists together
	emptyList.Extend(fromArrayList)
	fmt.Println(emptyList) // Prints [1,2,3,4]
}

```

Please check the [documentation](https://pkg.go.dev/github.com/gopher-utils/go-collections) for more comprehensive usage.

## Contributing

Any contributions to this repo are always welcome. Please check the [Contribution Guidelines](CONTRIBUTING.md) for details.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
