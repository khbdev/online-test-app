package adapter

import (
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/proto"
)


func ProtoGenerate(restBody []byte, protoMessage proto.Message) (proto.Message, error) {

	err := json.Unmarshal(restBody, protoMessage)
	if err != nil {
		return nil, fmt.Errorf(" ProtoGenerate error: %v", err)
	}
	return protoMessage, nil
}


func RestGenerate(protoMessage proto.Message) (map[string]interface{}, error) {

	data, err := json.Marshal(protoMessage)
	if err != nil {
		return nil, fmt.Errorf(" RestGenerate marshal error: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, fmt.Errorf(" RestGenerate unmarshal error: %v", err)
	}

	return result, nil
}
