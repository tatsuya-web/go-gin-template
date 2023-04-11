package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/tatuya-web/go-gin-template/domain/model"
	"github.com/tatuya-web/go-gin-template/testutil"
	"github.com/tatuya-web/go-gin-template/testutil/fixture"
	"github.com/tatuya-web/go-gin-template/util"
	clock "github.com/tatuya-web/go-gin-template/util"
)

func prepareUser(ctx context.Context, t *testing.T, db Execer) model.UserID {
	t.Helper()
	u := fixture.User(nil)
	result, err := db.ExecContext(ctx, "INSERT INTO users (email, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?);", u.Email, u.Password, u.Role, u.CreatedAt, u.UpdatedAt)
	if err != nil {
		t.Fatalf("insert user: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatalf("got user_id: %v", err)
	}
	return model.UserID(id)
}

func preparePosts(ctx context.Context, t *testing.T, con Execer) (model.UserID, model.Posts) {
	t.Helper()
	userID := prepareUser(ctx, t, con)
	// otherUserID := prepareUser(ctx, t, con)
	c := clock.FixedClocker{}
	wants := model.Posts{
		{
			UserID: userID,
			Title:  "want post 1", Content: "test_content",
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
		{
			UserID: userID,
			Title:  "want post 2", Content: "test_content",
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
	}
	posts := model.Posts{
		wants[0],
		{
			UserID: 0,
			Title:  "not want post", Content: "test_content",
			CreatedAt: c.Now(), UpdatedAt: c.Now(),
		},
		wants[1],
	}
	result, err := con.ExecContext(ctx,
		`INSERT INTO posts (title, content, user_id, created_at, updated_at)
			VALUES
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?),
			    (?, ?, ?, ?, ?);`,
		posts[0].Title, posts[0].Content, posts[0].UserID, posts[0].CreatedAt, posts[0].UpdatedAt,
		posts[1].Title, posts[1].Content, posts[1].UserID, posts[1].CreatedAt, posts[1].UpdatedAt,
		posts[2].Title, posts[2].Content, posts[2].UserID, posts[2].CreatedAt, posts[2].UpdatedAt,
	)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	posts[0].ID = model.PostID(id)
	posts[1].ID = model.PostID(id + 1)
	posts[2].ID = model.PostID(id + 2)
	return userID, wants
}

func TestRepository_AddPost(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Execer
		p   *model.Post
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
			name: "ok post",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				p: &model.Post{
					Title:   "test_title",
					Content: "test_content",
					UserID:  1,
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
			if err := r.AddPost(tt.args.ctx, tt.args.db, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Repository.AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_UpdatePost(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Execer
		p   *model.Post
	}
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	_, wants := preparePosts(ctx, t, tx)
	wants[0].Title = "new_title"
	wants[0].Content = "new_content"
	wants[1].Title = "new_title"
	wants[1].Content = "new_content"

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok post",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				p:   wants[0],
			},
			wantErr: false,
		},
		{
			name: "ok post",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				p:   wants[1],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Clocker: tt.fields.Clocker,
			}
			if err := r.UpdatePost(tt.args.ctx, tt.args.db, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Repository.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_DeletePost(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Execer
		p   *model.Post
	}
	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	_, wants := preparePosts(ctx, t, tx)
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok post",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				p:   wants[0],
			},
			wantErr: false,
		},
		{
			name: "ok post",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: ctx,
				db:  tx,
				p:   wants[1],
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Clocker: tt.fields.Clocker,
			}
			if err := r.DeletePost(tt.args.ctx, tt.args.db, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Repository.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_ListPosts(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Queryer
		id  model.UserID
	}

	ctx := context.Background()
	tx, err := testutil.OpenDBForTest(t).BeginTxx(ctx, nil)
	// このテストケースが完了したらもとに戻す
	t.Cleanup(func() { _ = tx.Rollback() })
	if err != nil {
		t.Fatal(err)
	}
	wantUserID, wants := preparePosts(ctx, t, tx)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Posts
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "ok posts",
			fields: fields{
				Clocker: util.FixedClocker{},
			},
			args: args{
				ctx: context.Background(),
				db:  tx,
				id:  wantUserID,
			},
			want:    wants,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Clocker: tt.fields.Clocker,
			}
			got, err := r.ListPosts(tt.args.ctx, tt.args.db, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.ListPosts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%v : %v", got[0], tt.want[0])
				t.Errorf("Repository.ListPosts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_IsOwnPost(t *testing.T) {
	type fields struct {
		Clocker clock.Clocker
	}
	type args struct {
		ctx context.Context
		db  Queryer
		id  model.PostID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				Clocker: tt.fields.Clocker,
			}
			if got := r.IsOwnPost(tt.args.ctx, tt.args.db, tt.args.id); got != tt.want {
				t.Errorf("Repository.IsOwnPost() = %v, want %v", got, tt.want)
			}
		})
	}
}
