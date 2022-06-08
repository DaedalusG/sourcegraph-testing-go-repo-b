package gorepob

import (
	"strings"

	gorepoc "github.com/sourcegraph-testing/go-repo-c"
)

func FunctionInRepoB(names []string) string {
	var results []string

	for i, name := range names {
		result := gorepoc.FunctionInRepoC(name, i, i+1)
		results = append(results, result)
	}

	return strings.Join(results, ", ")
}
