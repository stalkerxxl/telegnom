package types

import (
	"encoding/json"
	"testing"
)

func TestMessageOrBool_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		wantIsMessage  bool
		wantSuccess    bool
		wantMessageNil bool
	}{
		{
			name:           "Valid True",
			input:          `true`,
			wantIsMessage:  false,
			wantSuccess:    true,
			wantMessageNil: true,
		},
		{
			name:           "Valid False",
			input:          `false`,
			wantIsMessage:  false,
			wantSuccess:    false,
			wantMessageNil: true,
		},
		{
			name:           "Valid Message Object",
			input:          `{"message_id": 123}`,
			wantIsMessage:  true,
			wantSuccess:    false,
			wantMessageNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mb MessageOrBool
			err := json.Unmarshal([]byte(tt.input), &mb)

			if err != nil {
				t.Fatalf("UnmarshalJSON() error = %v", err)
			}

			if mb.IsMessage() != tt.wantIsMessage {
				t.Errorf("IsMessage() = %v, want %v", mb.IsMessage(), tt.wantIsMessage)
			}

			if mb.Success() != tt.wantSuccess {
				t.Errorf("Success() = %v, want %v", mb.Success(), tt.wantSuccess)
			}

			if (mb.Message() == nil) != tt.wantMessageNil {
				t.Errorf("Message() nil status = %v, want %v", mb.Message() == nil, tt.wantMessageNil)
			}
		})
	}
}
