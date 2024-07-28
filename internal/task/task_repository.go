package task

import (
	"database/sql"
	"log"

	"github.com/Dunkansdk/kanban-dunkan/internal/database"
)

type TaskRepository interface {
	Insert(in *Task) error
	GetById(id int) (Task, error)
	GetAllByStatus(status TaskStatus) ([]Task, error)
	GetAllStatuses() []TaskStatus
	GetStatusById(id int) (TaskStatus, error)
}

type pTaskRepository struct {
	connection *sql.DB
}

type TaskConnection struct {
	connection *database.ConnectionHandler
}

func NewTaskRepository(handler *database.ConnectionHandler) TaskRepository {
	return &TaskConnection{connection: handler}
}

// Get implements TaskConnection.
func (tc *TaskConnection) GetById(id int) (Task, error) {
	var task Task
	row := tc.connection.Connection().QueryRow("SELECT id, code, name, description FROM task WHERE id = $1", id)
	if err := row.Scan(&task.ID, &task.Code, &task.Name, &task.Content); err != nil {
		return task, err
	}
	return task, nil
}

// InsertTask implements TaskConnection.
func (tc *TaskConnection) Insert(in *Task) error {
	sqlStatement := `
		INSERT INTO task (code, name, description, status, task_status_id)
		VALUES ($1, $2, $3, 0)
	`
	_, err := tc.connection.Connection().Exec(sqlStatement, in.Code, in.Name, in.Content)

	return err
}

func (tc *TaskConnection) GetAllStatuses() []TaskStatus {
	rows, err := tc.connection.Connection().Query("SELECT id, name FROM task_status")
	if err != nil {
		log.Fatal(err)
	}

	var status_list []TaskStatus
	for rows.Next() {
		var ts TaskStatus
		err := rows.Scan(&ts.ID, &ts.Name)
		if err != nil {
			log.Fatal(err)
		}
		status_list = append(status_list, ts)
	}

	return status_list
}

func (tc *TaskConnection) GetStatusById(id int) (TaskStatus, error) {
	var status TaskStatus
	row := tc.connection.Connection().QueryRow("SELECT id, name FROM task_status WHERE id = $1", id)
	if err := row.Scan(&status.ID, &status.Name); err != nil {
		return status, err
	}
	return status, nil
}

// GetByStatus implements TaskRepository.
func (tc *TaskConnection) GetAllByStatus(status TaskStatus) ([]Task, error) {
	rows, err := tc.connection.Connection().Query(`SELECT task.id, task.code, task.name, task.description, status.id, status.name 
		FROM task AS task
		JOIN task_status AS status ON status.id = task.task_status_id
		WHERE task_status_id = $1`, status.ID)
	if err != nil {
		log.Fatal(err)
	}

	var tasks []Task
	for rows.Next() {
		var t Task
		var ts TaskStatus
		err := rows.Scan(&t.ID, &t.Code, &t.Name, &t.Content, &ts.ID, &ts.Name)
		if err != nil {
			log.Fatal(err)
		}
		t.Status = ts
		tasks = append(tasks, t)
	}

	return tasks, nil
}
