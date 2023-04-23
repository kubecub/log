# Project myproject

## Introduction

This project aims to provide a lightweight logging framework that supports multiple logging formats and output methods. It includes the following features:

- Built-in formatters: supports JSON and text formatting output methods.
- Extensible output methods: supports output to console, file, ELK, and other methods.
- Configurable log levels: supports setting different levels of log information for easy troubleshooting.
- Optional stack information: supports outputting call stack information for easy troubleshooting.

## Features

This framework supports the following features:

- Supports multiple log levels: `Trace`, `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`.
- Supports multiple output methods: `Console`, `File`, `ELK`.
- Supports multiple formatting methods: `JSON`, `Text`.
- Supports asynchronous log writing.
- Supports printing all environment variables at startup.
- Supports dynamically setting log levels.
- Supports outputting stack information.

## Getting Started

### Prerequisites

This project requires Go 2.0 or higher.

### Building

You can build the example program by running the following command in the project root directory:

```
$ make build

```

### Running

You can run the example program by running the following command in the project root directory:

```
$ ./bin/main

```

## Usage

Import the framework as follows:

```
import "github.com/kubecub/log"

```

### Code Examples

```
package main

import (
    "github.com/kubecub/log"
)

func main() {
    // Create a Logger instance
    logger := cuslog.NewLogger(cuslog.WithLevel(cuslog.InfoLevel))

    // Output log information
    logger.Info("Info message")

    // Output log information with fields
    logger.WithFields(cuslog.Fields{
        "animal": "walrus",
        "size":   10,
    }).Info("A group of walrus emerges from the ocean")
}

```

## Contributing

Contributions to this project are welcome! Please see CONTRIBUTING.md for details.

## Community (Optional)

Join our community! You can find us on the following channels:

- GitHub Issues: [https://github.com/kubecub/log/issues](https://github.com/kubecub/log/issues)

## Author

- Author: Xinwei Xiong(cubxxw)
- Email: [3293172751nss@gmail.com](mailto:3293172751nss@gmail.com)

## License

This project is licensed under the MIT License.