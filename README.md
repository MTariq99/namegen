NameGen Package
The namegen package generates unique names from a predefined list using a customizable seed for randomness. This README provides an overview of how to use the package, its uniqueness, and details about the generated names.

Installation
To use the namegen package, ensure you have Go installed. You can install the package with go get:

bash
```
go get -u github.com/mtariq99/namegen
```

Usage
Importing the Package
Import the namegen package in your Go files:

```go
import (
    "fmt"
    "math/rand"
    "github.com/MTariq99/namegen"
)
```
Initializing the Generator
Initialize a new name generator with a seed:

```go
seed := int64(42) // Replace with your desired seed
generator := namegen.NewGenerator(seed)
Generating Unique Names
//Use the GetUniqueName method to generate unique names:
```

```go
name1 := generator.GetUniqueName()
name2 := generator.GetUniqueName()

fmt.Println(name1)
fmt.Println(name2)
//Each call to GetUniqueName returns a unique name based on the seed provided.
```

Uniqueness
The uniqueness of the names generated by namegen is ensured through:

Customizable Seed: You can specify a seed to generate different sequences of names. Changing the seed will produce different unique names while maintaining the sequence order.
Large Name Pool: The package includes a substantial pool of names (425 in total) ensuring a wide variety of unique combinations.
Example
Here’s an example demonstrating the usage of the namegen package:

```go
package main

import (
    "fmt"
    "github.com/MTariq99/namegen"
)

func main() {
    seed := int64(12345) // Example seed
    generator := namegen.NewGenerator(seed)

    fmt.Println(generator.GetUniqueName())
}
```
Contributing
Contributions to the namegen package are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request on GitHub.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Notes:
Ensure to update the seed variable in your examples to demonstrate different outputs.
Include any additional details specific to your package or usage scenarios.
This structure with headings makes the README.md file clear and easy to navigate. Adjust the content further based on your specific package features and user requirements.
