package sdrangel

import (
	"encoding/json"
	"testing"
)

func TestDeviceActionsMarshalUsesPluginKey(t *testing.T) {
	da := DeviceActions{DeviceHwType: "FileInput", Direction: 0, ActionsKey: "fileInputActions", Actions: json.RawMessage(`{}`)}
	b, err := json.Marshal(da)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var m map[string]json.RawMessage
	if err := json.Unmarshal(b, &m); err != nil {
		t.Fatalf("unmarshal marshaled output: %v", err)
	}
	if _, ok := m["actions"]; ok {
		t.Errorf("marshaled output has generic actions key: %s", b)
	}
	if _, ok := m["fileInputActions"]; !ok {
		t.Errorf("marshaled output missing plugin key fileInputActions: %s", b)
	}
}
