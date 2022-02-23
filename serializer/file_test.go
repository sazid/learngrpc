package serializer

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang/protobuf/proto"
	v1 "github.com/sazid/learngrpc/api/v1"
	"github.com/sazid/learngrpc/sample"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	f, err := os.CreateTemp("", "serializer_*")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	laptop1 := sample.NewLaptop()
	err = WriteProtobufToBinaryFile(laptop1, f.Name())
	require.NoError(t, err)

	laptop2 := &v1.Laptop{}
	err = ReadProtobufFromFile(f, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}

func TestJsonSerializer(t *testing.T) {
	t.Parallel()

	laptop1 := sample.NewLaptop()
	jsonStr, err := WriteProtobufToJson(laptop1)
	// ioutil.WriteFile("laptop.json", jsonStr, 0644)
	fmt.Println(string(jsonStr))
	require.NoError(t, err)
}
