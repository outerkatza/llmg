//go:build tools

package tools

import (
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
    _ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)