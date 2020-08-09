//+build integration

package mysql

import (
	"testing"

	"github.com/jinzhu/gorm"
)

func TestDatabase_FindUserWithUsername(t *testing.T) {
	type fields struct {
		DB *gorm.DB
	}
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantUser Customer
		wantErr  bool
	}{
		{
			name: "find the taiwan",
			fields: fields{DB: db.DB},
			args: args{name: "taiwan"},
			wantUser: Customer{Username: "taiwan"},
			wantErr: false,
		},
		{
			name: "can't find the ghost",
			fields: fields{DB: db.DB},
			args: args{name: "ghost"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Database{
				DB: tt.fields.DB,
			}
			gotUser, err := d.FindUserWithUsername(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Database.FindUserWithUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotUser.Username != tt.wantUser.Username {
				t.Errorf("Database.FindUserWithUsername() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}
