package structutil

import (
	"fmt"

	"google.golang.org/protobuf/types/known/structpb"
)

func ToProtoStruct(m map[string]string) (*structpb.Struct, error) {
	mapAny := make(map[string]any, len(m))
	for k, v := range m {
		mapAny[k] = v
	}

	return structpb.NewStruct(mapAny)
}

func ToStringMap(s *structpb.Struct) map[string]string {
	if s == nil {
		return nil
	}

	result := make(map[string]string, len(s.Fields))
	for k, v := range s.AsMap() {
		if strVal, ok := v.(string); ok {
			result[k] = strVal
		} else {
			result[k] = fmt.Sprintf("%v", v)
		}
	}

	return result
}
