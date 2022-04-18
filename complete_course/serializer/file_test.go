package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/sample"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/serializer"
)


func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

}