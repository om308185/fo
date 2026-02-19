# 🎺 Function Calling Utilities for Go

![Go Version](https://img.shields.io/badge/Go-1.18%2B-brightgreen)
![License](https://img.shields.io/badge/License-MIT-blue)
![Build Status](https://img.shields.io/badge/build-passing-brightgreen)

Welcome to the **Function Calling Utilities** library, a Go library designed for ease of use with function calls. This library leverages Go 1.18+ generics to provide a simple yet powerful way to handle function invocations. It includes features such as Go2-like error handling and context invocation, making it an essential tool for modern Go developers.

## Table of Contents

1. [Features](#features)
2. [Installation](#installation)
3. [Usage](#usage)
4. [Examples](#examples)
5. [Contributing](#contributing)
6. [License](#license)
7. [Releases](#releases)

## Features

- **Generics Support**: Utilize the power of Go's generics to create reusable functions.
- **Error Handling**: Implement Go2-like error handling to manage errors more effectively.
- **Context Invocation**: Call functions with context support to handle cancellations and timeouts.
- **Lightweight**: The library is designed to be simple and efficient, ensuring minimal overhead.

## Installation

To install the library, use the following command:

```bash
go get github.com/om308185/fo
```

## Usage

Here’s a quick overview of how to use the Function Calling Utilities library in your Go projects.

### Basic Example

```go
package main

import (
    "context"
    "fmt"
    "github.com/om308185/fo"
)

func main() {
    ctx := context.Background()
    result, err := fo.Invoke(ctx, myFunction, "Hello, World!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}

func myFunction(input string) (string, error) {
    return input, nil
}
```

### Advanced Usage

You can also use the library to handle more complex scenarios, such as passing multiple arguments and managing timeouts.

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/om308185/fo"
)

func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    result, err := fo.Invoke(ctx, myAdvancedFunction, "Hello", "Go!")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}

func myAdvancedFunction(a, b string) (string, error) {
    return fmt.Sprintf("%s %s", a, b), nil
}
```

## Examples

For more examples, check the [examples directory](https://github.com/om308185/fo/tree/main/examples) in the repository. You can find various use cases demonstrating the library's capabilities.

## Contributing

We welcome contributions to enhance the library. To contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature/YourFeature`).
6. Create a pull request.

Please ensure your code adheres to the existing style and includes tests where applicable.

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/om308185/fo/blob/main/LICENSE) file for details.

## Releases

To download the latest release, visit [Releases](https://github.com/om308185/fo/releases). Here you can find the binaries and source code for the latest version. Make sure to download the appropriate files for your system and follow the instructions to execute them.

For detailed release notes and version history, refer to the [Releases section](https://github.com/om308185/fo/releases).

## Conclusion

The Function Calling Utilities library offers a robust solution for handling function calls in Go. Its use of generics, context support, and improved error handling makes it a valuable addition to any Go project. 

Feel free to explore the repository and utilize the features it provides. If you have any questions or need further assistance, don't hesitate to reach out.

Happy coding!