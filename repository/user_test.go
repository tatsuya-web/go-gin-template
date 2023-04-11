package repository

import (
	"context"
	"testing"

	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/testutil"
	"github.com/tatuya-web/go-gin-template/util"
	clock "github.com/tatuya-web/go-gin-template/util"
)

func TestRepository_RegisterUser(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Execer
		u   *model.User
	}
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok user",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				u: &model.User{
					Email:    "test@example.com",
					Password: "password",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Clocker: tt.fields.Clocker,
			}
			if err := r.RegisterUser(tt.args.ctx, tt.args.db, tt.args.u); (err != nil) != tt.wantErr {
				t.Errorf("Repository.RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
