package repos

import (
	"context"
	"errors"
	"github.com/cherrycutter/todo_app/internal/models"
	"github.com/jackc/pgx/v5"
	"github.com/pashagolub/pgxmock/v4"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetAllTodos(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer mockDB.Close(context.Background())

	r := NewTodoRepo(mockDB)

	tests := []struct {
		name    string
		mock    func()
		want    []models.TodoModel
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := pgxmock.NewRows([]string{"id", "title", "description", "completed", "created_at"}).
					AddRow(1, "title1", "description1", false, time.Now()).
					AddRow(2, "title2", "description2", true, time.Now())
				mockDB.ExpectQuery("SELECT \\* FROM todo").WillReturnRows(rows)
			},
			want: []models.TodoModel{
				{Id: 1, Title: "title1", Description: "description1", Completed: false, CreatedAt: time.Now()},
				{Id: 2, Title: "title2", Description: "description2", Completed: true, CreatedAt: time.Now()},
			},
			wantErr: false,
		},
		{
			name: "No Rows",
			mock: func() {
				rows := pgxmock.NewRows([]string{"id", "title", "description", "completed", "created_at"})
				mockDB.ExpectQuery("SELECT \\* FROM todo").WillReturnRows(rows)
			},
			want:    []models.TodoModel(nil),
			wantErr: false,
		},
		{
			name: "Query Error",
			mock: func() {
				mockDB.ExpectQuery("SELECT \\* FROM todo").WillReturnError(errors.New("query error"))
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.GetAllTodos(context.Background())
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mockDB.ExpectationsWereMet())
		})
	}
}

func TestGetTodoById(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer mockDB.Close(context.Background())

	r := NewTodoRepo(mockDB)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    models.TodoModel
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := pgxmock.NewRows([]string{"id", "title", "description", "completed", "created_at"}).
					AddRow(1, "title1", "description1", false, time.Now())
				mockDB.ExpectQuery("SELECT \\* FROM todo WHERE (.+)").WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				id: 1,
			},
			want:    models.TodoModel{Id: 1, Title: "title1", Description: "description1", CreatedAt: time.Now()},
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {
				mockDB.ExpectQuery("SELECT \\* FROM todo WHERE (.+)").WithArgs(404).WillReturnError(pgx.ErrNoRows)
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},
		{
			name: "Query Error",
			mock: func() {
				mockDB.ExpectQuery("SELECT \\* FROM todo WHERE (.+)").WithArgs(1).WillReturnError(errors.New("query error"))
			},
			input: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := r.GetTodoById(context.Background(), tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Equal(t, tt.want.Id, got.Id)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Description, got.Description)
				assert.Equal(t, tt.want.Completed, got.Completed)
				assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			}
			assert.NoError(t, mockDB.ExpectationsWereMet())
		})
	}
}

func TestCreateTodo(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer mockDB.Close(context.Background())

	r := NewTodoRepo(mockDB)

	tests := []struct {
		name    string
		mock    func()
		input   models.TodoModel
		want    models.TodoModel
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				rows := pgxmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now())
				mockDB.ExpectQuery("INSERT INTO todo").
					WithArgs("title", "description", false).
					WillReturnRows(rows)
			},
			input: models.TodoModel{
				Title:       "title",
				Description: "description",
				Completed:   false,
			},
			want: models.TodoModel{
				Id:          1,
				Title:       "title",
				Description: "description",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Empty Title",
			mock: func() {
			},
			input: models.TodoModel{
				Title:       "",
				Description: "description",
				Completed:   false,
			},
			want:    models.TodoModel{},
			wantErr: true,
		},
		{
			name: "Query Error",
			mock: func() {
				mockDB.ExpectQuery("INSERT INTO todo").
					WithArgs("title", "description", false).WillReturnError(errors.New("query error"))
			},
			input: models.TodoModel{
				Title:       "title",
				Description: "description",
				Completed:   false,
			},
			want:    models.TodoModel{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.CreateTodo(context.Background(), tt.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Id, got.Id)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Description, got.Description)
				assert.Equal(t, tt.want.Completed, got.Completed)
				assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			}
			assert.NoError(t, mockDB.ExpectationsWereMet())
		})
	}
}

func TestUpdateTodo(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer mockDB.Close(context.Background())

	r := NewTodoRepo(mockDB)

	type args struct {
		id    int
		input models.TodoModel
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		want    models.TodoModel
		wantErr bool
	}{
		{
			name: "Ok_AllFields",
			mock: func() {
				rows := pgxmock.NewRows([]string{"id", "title", "description", "completed", "created_at"}).
					AddRow(1, "new title", "new description", false, time.Now())
				mockDB.ExpectQuery("UPDATE todo SET title = \\$1, description = \\$2, completed = \\$3 WHERE id = \\$4 RETURNING id, title, description, completed, created_at").
					WithArgs("new title", "new description", false, 1).
					WillReturnRows(rows)
			},
			input: args{
				id:    1,
				input: models.TodoModel{Id: 1, Title: "new title", Description: "new description", Completed: false},
			},
			want: models.TodoModel{
				Id:          1,
				Title:       "new title",
				Description: "new description",
				Completed:   false,
				CreatedAt:   time.Now(),
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {
				mockDB.ExpectQuery("UPDATE todo SET title = \\$1, description = \\$2, completed = \\$3 WHERE id = \\$4 RETURNING id, title, description, completed, created_at").
					WithArgs("new title", "new description", false, 404).
					WillReturnError(pgx.ErrNoRows)
			},
			input: args{
				id:    404,
				input: models.TodoModel{Id: 404, Title: "new title", Description: "new description", Completed: false},
			},
			want:    models.TodoModel{},
			wantErr: true,
		},
		{
			name: "Empty Title",
			mock: func() {
			},
			input: args{
				id:    1,
				input: models.TodoModel{Id: 1, Title: "", Description: "new description", Completed: false},
			},
			want:    models.TodoModel{},
			wantErr: true,
		},
		{
			name: "Query Error",
			mock: func() {
				mockDB.ExpectQuery("UPDATE todo SET title = \\$1, description = \\$2, completed = \\$3 WHERE id = \\$4 RETURNING id, title, description, completed, created_at").
					WithArgs("new title", "new description", false, 1).
					WillReturnError(errors.New("query error"))
			},
			input: args{
				id:    1,
				input: models.TodoModel{Id: 1, Title: "new title", Description: "new description", Completed: false},
			},
			want:    models.TodoModel{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.mock != nil {
				tt.mock()
			}

			got, err := r.UpdateTodo(context.Background(), tt.input.id, tt.input.input)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want.Id, got.Id)
				assert.Equal(t, tt.want.Title, got.Title)
				assert.Equal(t, tt.want.Description, got.Description)
				assert.Equal(t, tt.want.Completed, got.Completed)
				assert.WithinDuration(t, tt.want.CreatedAt, got.CreatedAt, time.Second)
			}
			assert.NoError(t, mockDB.ExpectationsWereMet())
		})
	}
}

func TestDeleteTodoById(t *testing.T) {
	mockDB, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("failed to create mock db: %v", err)
	}
	defer mockDB.Close(context.Background())

	r := NewTodoRepo(mockDB)

	type args struct {
		id int
	}
	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "Ok",
			mock: func() {
				mockDB.ExpectExec("DELETE FROM todo WHERE id = \\$1").
					WithArgs(1).
					WillReturnResult(pgxmock.NewResult("DELETE", 1))
			},
			input: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			mock: func() {
				mockDB.ExpectExec("DELETE FROM todo WHERE id = \\$1").
					WithArgs(404).
					WillReturnResult(pgxmock.NewResult("DELETE", 0))
			},
			input: args{
				id: 404,
			},
			wantErr: true,
		},
		{
			name: "Query Error",
			mock: func() {
				mockDB.ExpectExec("DELETE FROM todo WHERE id = \\$1").
					WithArgs(1).
					WillReturnError(errors.New("query error"))
			},
			input: args{
				id: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			err := r.DeleteTodoById(context.Background(), tt.input.id)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.NoError(t, mockDB.ExpectationsWereMet())
		})
	}
}
