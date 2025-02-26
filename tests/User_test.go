package user_test

import (
	"testing"

	user "github.com/mhgffqwoer/EducationalPlatform/internal/User"
)

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name     string
		userType user.UserType
		want     user.UserBuilder
		wantErr  bool
	}{
		{
			name:     "Create student",
			userType: user.Student,
			want:     nil,
			wantErr:  false,
		},
		{
			name:     "Create teacher",
			userType: user.Teacher,
			want:     nil,
			wantErr:  false,
		},
		{
			name:     "Create unknown user",
			userType: user.UserType(3),
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := user.New(tt.userType)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("New() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Errorf("New() succeeded unexpectedly")
			}
			if !true {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
