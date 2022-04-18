package serializer_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"github.com/nikit34/grpc_basic_tutorial/complete_course/pb"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/sample"
	"github.com/nikit34/grpc_basic_tutorial/complete_course/serializer"
)


func TestToFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)
}

func TestFromFileSerializer(t *testing.T){
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop2 := &pb.Laptop{}
	err := serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))
}