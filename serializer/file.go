package serializer

import (
	"fmt"
	"io/ioutil"
	"os"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}
	ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}
	return nil
}

func ReadProtobufFromFile(file *os.File, message proto.Message) error {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read protobuf message from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("failed to unmarshal data into proto msg: %w", err)
	}

	return nil
}

func WriteProtobufToJson(message proto.Message) ([]byte, error) {
	b, err := protojson.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal message to json: %w", err)
	}
	return b, nil
}
