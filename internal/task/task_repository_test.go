package task

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Dunkansdk/kanban-dunkan/internal/database"
	"github.com/stretchr/testify/assert"
)

// MockDB implements IConnection interface for testing
type MockDB struct {
	db *sql.DB
}

func (m *MockDB) GetConnection() *sql.DB {
	return m.db
}

func setupMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock, TaskRepository) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}

	mockDB := &MockDB{db: db}
	handler := database.CreateConnection(mockDB)
	repo := NewTaskRepository(handler)

	return db, mock, repo
}

func TestGetById(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "code", "name", "description"}).
		AddRow(int64(1), "TSK-001", "Test Task", "Test Description") // Explicitly use int64

	mock.ExpectQuery("SELECT id, code, name, description FROM task WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	task, err := repo.GetById(1)

	assert.NoError(t, err)
	assert.Equal(t, int64(1), task.ID) // Compare with int64
	assert.Equal(t, "TSK-001", task.Code)
	assert.Equal(t, "Test Task", task.Name)
	assert.Equal(t, "Test Description", task.Content)
}

func TestInsert(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	task := &Task{
		Code:    "TSK-002",
		Name:    "New Task",
		Content: "New Description",
	}

	mock.ExpectPrepare("INSERT INTO task").
		ExpectExec().
		WithArgs(task.Code, task.Name, task.Content).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err := repo.Insert(task)
	assert.NoError(t, err)
}

func TestGetAllStatuses(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "TODO").
		AddRow(2, "IN_PROGRESS").
		AddRow(3, "DONE")

	mock.ExpectQuery("SELECT id, name FROM task_status").
		WillReturnRows(rows)

	statuses := repo.GetAllStatuses()

	assert.Len(t, statuses, 3)
	assert.Equal(t, "TODO", statuses[0].Name)
	assert.Equal(t, "IN_PROGRESS", statuses[1].Name)
	assert.Equal(t, "DONE", statuses[2].Name)
}

func TestGetStatusById(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "TODO")

	mock.ExpectQuery("SELECT id, name FROM task_status WHERE id = \\$1").
		WithArgs(1).
		WillReturnRows(rows)

	status, err := repo.GetStatusById(1)

	assert.NoError(t, err)
	assert.Equal(t, 1, status.ID)
	assert.Equal(t, "TODO", status.Name)
}

func TestGetAllByStatus(t *testing.T) {
	db, mock, repo := setupMockDB(t)
	defer db.Close()

	status := TaskStatus{ID: 1, Name: "TODO"}

	rows := sqlmock.NewRows([]string{"id", "code", "name", "description", "status_id", "status_name"}).
		AddRow(1, "TSK-001", "Task 1", "Description 1", 1, "TODO").
		AddRow(2, "TSK-002", "Task 2", "Description 2", 1, "TODO")

	mock.ExpectQuery("SELECT task.id, task.code, task.name, task.description, status.id, status.name FROM task AS task JOIN task_status AS status").
		WithArgs(status.ID).
		WillReturnRows(rows)

	tasks, err := repo.GetAllByStatus(status)

	assert.NoError(t, err)
	assert.Len(t, tasks, 2)
	assert.Equal(t, "TSK-001", tasks[0].Code)
	assert.Equal(t, "TSK-002", tasks[1].Code)
	assert.Equal(t, "TODO", tasks[0].Status.Name)
}
