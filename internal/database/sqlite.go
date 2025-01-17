package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	gap "github.com/muesli/go-app-paths"
)

type SQLite3DB struct {
	instance *sql.DB
}

// https://go.dev/doc/database/sql-injection

func initalizePath(path string) error {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return os.Mkdir(path, 0o770)
		}
		return err
	}
	return nil
}

func path() string {
	scope := gap.NewScope(gap.User, "tasks")
	dirs, err := scope.DataDirs()
	if err != nil {
		log.Fatal(err)
	}
	var dir string
	if len(dirs) > 0 {
		dir = dirs[0]
	} else {
		dir, _ = os.UserHomeDir()
	}
	if err := initalizePath(dir); err != nil {
		log.Fatal(err)
	}
	return filepath.Join(dir, "kandun.db")
}

func (db *SQLite3DB) GetConnection() *sql.DB {
	if db.instance != nil {
		return db.instance
	}

	conn, err := sql.Open("sqlite3", path())

	if err != nil {
		log.Fatal(err)
	}

	if _, err := conn.Exec(initialData()); err != nil {
		panic(err)
	}

	db.instance = conn

	return db.instance
}

func initialData() string {
	return `-- Crear la tabla task_status si no existe
CREATE TABLE IF NOT EXISTS task_status (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE
);

-- Crear la tabla project si no existe
CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    start_date TIMESTAMP,
    end_date TIMESTAMP
);

-- Crear la tabla task si no existe
CREATE TABLE IF NOT EXISTS task (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	code VARCHAR(15) NOT NULL, 
    name VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    end_time TIMESTAMP,
    task_status_id INTEGER NOT NULL,
    project_id INTEGER,
    CONSTRAINT fk_task_status
        FOREIGN KEY(task_status_id)
	    REFERENCES task_status(id),
    CONSTRAINT fk_project
        FOREIGN KEY(project_id)
	    REFERENCES project(id),
    CONSTRAINT ck_end_time
        CHECK (end_time > created_at)
);

-- Crear índices para mejorar el rendimiento si no existen
CREATE INDEX IF NOT EXISTS idx_task_task_status_id ON task (task_status_id);
CREATE INDEX IF NOT EXISTS idx_task_project_id ON task (project_id);

INSERT INTO task_status (id, name) VALUES
(0, 'TODO'),
(1, 'IN_PROGRESS'),
(2, 'DONE')
ON CONFLICT (id) DO NOTHING;

INSERT INTO project (id, name, description) VALUES
(0, 'Sample Project', 'Sample project, for testing purpose only')
ON CONFLICT (id) DO NOTHING;

INSERT INTO task (id, code, name, description, task_status_id, project_id) VALUES
(0, 'TK-0011', 'Plan Product Launch', '## Develop a detailed plan for the new product launch', 0, 0),
(1, 'TK-0012', 'Design Marketing Campaign', 'Create and design an effective marketing campaign to attract new customers', 1, 0),
(2, 'TK-0013', 'Conduct Market Analysis', 'Research and analyze current market trends and competitors', 0, 0),
(3, 'TK-0014', 'Develop Prototype', 'Create a functional prototype of the new technological device', 0, 0),
(4, 'TK-0015', 'Write Research Article', 'Draft a detailed article on recent findings in the field of biotechnology', 1, 0),
(5, 'TK-0016', 'Organize International Conference', 'Plan and coordinate all aspects of an international conference on artificial intelligence', 2, 0),
(6, 'TK-0017', 'Review Source Code', 'Conduct a thorough review of the project''s source code', 0, 0),
(7, 'TK-0018', 'Lead Brainstorming Session', 'Conduct a brainstorming session to generate innovative ideas for the new project', 1, 0),
(8, 'TK-0019', 'Prepare Investor Presentation', 'Create and rehearse an engaging presentation for potential project investors', 0, 0),
(9, 'TK-0020', 'Implement SEO Strategies', 'Optimize website content to improve search engine rankings', 0, 0),
(10, 'TK-0021', 'Develop Mobile App', 'Design and develop a new mobile application for e-commerce', 1, 0),
(11, 'TK-0022', 'Conduct User Testing', 'Perform user testing to gather feedback on the new software', 2, 0),
(12, 'TK-0023', 'Optimize Database Performance', 'Improve the performance of the database for faster query response times', 0, 0),
(13, 'TK-0024', 'Create Content Strategy', 'Develop a comprehensive content strategy for the upcoming quarter', 1, 0),
(14, 'TK-0025', 'Manage Social Media Accounts', 'Oversee the daily operations and content posting for all social media accounts', 2, 0),
(15, 'TK-0026', 'Research Emerging Technologies', 'Investigate new and emerging technologies that could benefit the company', 0, 0),
(16, 'TK-0027', 'Update Website Design', 'Redesign the company website to improve user experience and visual appeal', 0, 0),
(17, 'TK-0028', 'Train New Employees', 'Conduct training sessions for newly hired employees', 2, 0),
(18, 'TK-0029', 'Analyze Sales Data', 'Examine and interpret sales data to identify trends and insights', 0, 0),
(19, 'TK-0030', 'Develop Customer Loyalty Program', 'Create a program to reward loyal customers and encourage repeat business', 1, 0),
(20, 'TK-0031', 'Small Content', 'Small Content', 0, 0)
ON CONFLICT (id) DO NOTHING;`

}
