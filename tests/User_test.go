package tests

import (
	"testing"

	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

func TestBuildUsers(t *testing.T) {
	tests := []struct {
		name        string
		userBuilder user.UserBuilder
		userName    string
		wantName    string
		wantId      int
	}{
		{
			name:        "Create student",
			userBuilder: user.New().Student(),
			userName:    "John Doe",
			wantName:    "John Doe",
			wantId:      5,
		},
		{
			name:        "Create teacher",
			userBuilder: user.New().Student(),
			userName:    "Jane Doe",
			wantName:    "Jane Doe",
			wantId:      6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, _ := tt.userBuilder.SetName(tt.userName).Build()
			if user.GetName() != tt.wantName {
				t.Errorf("GetName() = %v, want %v", user.GetName(), tt.wantName)
			}
			if user.GetID().Int() != tt.wantId {
				t.Errorf("GetID() = %v, want %v", user.GetID().Int(), tt.wantId)
			}
		})
	}
}
