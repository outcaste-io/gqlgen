package imports

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/outcaste-io/gqlgen/internal/code"
	"github.com/stretchr/testify/require"
)

func TestPrune(t *testing.T) {
	// prime the packages cache so that it's not considered uninitialized

	b, err := Prune("testdata/unused.go", mustReadFile("testdata/unused.go"), &code.Packages{})
	require.NoError(t, err)
	require.Equal(t, strings.Replace(string(mustReadFile("testdata/unused.expected.go")), "\r\n", "\n", -1), string(b))
}

func mustReadFile(filename string) []byte {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return b
}
