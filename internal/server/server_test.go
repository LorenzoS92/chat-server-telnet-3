package server

import "testing"

const errorPort = "10000"

// As server will enter in for loop to hear tcp connection, I won't test the initializeConnection happy path
// This is just for show that I will unit test written code.
func TestServer_RunChatServer(t *testing.T) {
	type fields struct {
		Port string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name:    "error defining port",
			fields:  fields{Port: errorPort},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				Port: tt.fields.Port,
			}
			if err := s.RunChatServer(); (err != nil) != tt.wantErr {
				t.Errorf("RunChatServer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
