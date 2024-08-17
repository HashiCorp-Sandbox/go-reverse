[![PkgGoDev](https://pkg.go.dev/badge/github.com/hashicorp-sandbox/go-reverse)](https://pkg.go.dev/github.com/hashicorp-sandbox/go-reverse)

# go-reverse

> [!IMPORTANT]
> This library is specifically created for a workshop on building Terraform providers using the [Terraform Plugin Framework](https://github.com/hashicorp/terraform-plugin-framework). It is designed to be as simple as possible to facilitate understanding of the core concepts and is not intended for general use.

## Usage

```go
package main

import (
    "fmt"
    "log"

    "github.com/hashicorp-sandbox/go-reverse"
)

func main() {
    result := reverse.ReverseString("HashiCorp")
    fmt.Println(result)
}
```

## Terraform's type system

Because Terraform's type system is different from Go's, the `tfreverse` package provides a function that requires the type system implemented in the Terraform Plugin Framework.

```go
package main

import (
	"fmt"

	"github.com/hashicorp-sandbox/go-reverse/tfreverse"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func main() {
	result := tfreverse.ReverseString(types.StringValue("HashiCorp"))
	fmt.Println(result.ValueString())
}
```

## License

This library is licensed under the Mozilla Public License 2.0. For more information, see the LICENSE file.
