package serializer

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/neel229/gRPC-introduction/pb/pb"
	"github.com/neel229/gRPC-introduction/sample"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"

	laptop1 := sample.NewLaptop()
	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)
	require.True(t, proto.Equal(laptop1, laptop2))

	jsonFile := "../tmp/laptop.json"
	err = WriteProtobuToJSONFile(laptop2, jsonFile)
	require.NoError(t, err)
}
