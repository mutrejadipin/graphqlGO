#!/bin/bash
set -e

echo "adding missing entries for gqlgen"

go get github.com/99designs/gqlgen/codegen/config@v0.17.90
go get github.com/99designs/gqlgen/internal/imports@v0.17.90
go get github.com/99designs/gqlgen/api@v0.17.90
go get github.com/99designs/gqlgen/codegen/config@v0.17.90
go get github.com/99designs/gqlgen/codegen@v0.17.90
go get github.com/99designs/gqlgen@v0.17.90



echo "starting generate command"

go run github.com/99designs/gqlgen generate


echo "Done!"