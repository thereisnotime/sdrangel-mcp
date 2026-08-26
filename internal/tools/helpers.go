package tools

import (
	"encoding/json"
	"fmt"
)

func marshalJSON(v any) (json.RawMessage, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("marshal: %w", err)
	}
	return json.RawMessage(b), nil
}
