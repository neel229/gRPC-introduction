package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobuToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("There was an error marshaling the message: %w", err)
	}

	if err := ioutil.WriteFile(filename, []byte(data), 0644); err != nil {
		return fmt.Errorf("There was an error writing data to the file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile write protocol buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("There was an error while serializing the message: %w", err)
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("Cannot write binary data to file %s: %w", filename, err)
	}

	return nil
}

// ReadProtobufFromFile reads protocol buffer message from binary file
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("There was an error reading the file: %w", err)
	}

	if err := proto.Unmarshal(data, message); err != nil {
		return fmt.Errorf("There was a problem decoding the data from the given file: %w", err)
	}

	return nil
}
