package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/QRY91/wherewasi/internal/common"
	_ "modernc.org/sqlite"
)

type DB struct {
	*sql.DB
}

type ContextSession struct {
	ID          int64     `json:"id"`
	Project     string    `json:"project"`
	Timestamp   time.Time `json:"timestamp"`
	ContextData string    `json:"context_data"`
	SessionInfo string    `json:"session_info"`
	Keywords    string    `json:"keywords"`
	CreatedAt   time.Time `json:"created_at"`
}

// Initialize database connection
func NewDB(dbPath string) (*DB, error) {
	// If no path specified, use XDG-compliant default location
	if dbPath == "" {
		dataDir := common.GetDataDir()
		if err := os.MkdirAll(dataDir, 0755); err != nil {
			return nil, fmt.Errorf("failed to create data directory: %w", err)
		}
		dbPath = common.GetDefaultDBPath()
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	// Test connection
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Enable foreign keys and optimizations
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, fmt.Errorf("failed to enable foreign keys: %w", err)
	}

	if _, err := db.Exec("PRAGMA journal_mode = WAL"); err != nil {
		return nil, fmt.Errorf("failed to enable WAL mode: %w", err)
	}

	dbInstance := &DB{db}

	// Run migrations
	if err := dbInstance.migrate(); err != nil {
		return nil, fmt.Errorf("failed to run migrations: %w", err)
	}

	return dbInstance, nil
}

// Run database migrations
func (db *DB) migrate() error {
	// Check if migrations table exists
	var tableExists bool
	err := db.QueryRow(`
		SELECT COUNT(*) > 0 FROM sqlite_master 
		WHERE type='table' AND name='schema_migrations'
	`).Scan(&tableExists)

	if err != nil {
		return fmt.Errorf("failed to check migrations table: %w", err)
	}

	if !tableExists {
		// Create initial schema
		if err := db.createInitialSchema(); err != nil {
			return fmt.Errorf("failed to create initial schema: %w", err)
		}
	}

	return nil
}

// Create the initial database schema
func (db *DB) createInitialSchema() error {
	schema := `
	-- Context sessions for ripcord deployments
	CREATE TABLE context_sessions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		project TEXT NOT NULL,
		timestamp DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		context_data TEXT NOT NULL,
		session_info TEXT,
		keywords TEXT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Projects tracked by wherewasi
	CREATE TABLE projects (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT UNIQUE NOT NULL,
		path TEXT NOT NULL,
		last_activity DATETIME,
		git_repo BOOLEAN DEFAULT FALSE,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Cross-ecosystem tool communication (compatible with uroboro)
	CREATE TABLE tool_messages (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		from_tool TEXT NOT NULL DEFAULT 'wherewasi',
		to_tool TEXT NOT NULL,
		message_type TEXT NOT NULL,
		data TEXT NOT NULL,
		processed BOOLEAN DEFAULT FALSE,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		processed_at DATETIME
	);

	-- Usage and performance tracking
	CREATE TABLE usage_stats (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		command TEXT NOT NULL,
		project TEXT,
		duration_ms INTEGER,
		success BOOLEAN DEFAULT TRUE,
		error_message TEXT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Indexes for performance
	CREATE INDEX idx_context_sessions_project ON context_sessions(project);
	CREATE INDEX idx_context_sessions_timestamp ON context_sessions(timestamp);
	CREATE INDEX idx_context_sessions_keywords ON context_sessions(keywords);
	CREATE INDEX idx_projects_name ON projects(name);
	CREATE INDEX idx_tool_messages_to_tool ON tool_messages(to_tool, processed);
	CREATE INDEX idx_usage_stats_command ON usage_stats(command);

	-- Migration tracking
	CREATE TABLE schema_migrations (
		version INTEGER PRIMARY KEY,
		description TEXT NOT NULL,
		applied_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	);

	-- Initial migration record
	INSERT INTO schema_migrations (version, description) 
	VALUES (1, 'Initial wherewasi schema with context sessions and project tracking');
	`

	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %w", err)
	}

	return nil
}

// Save a context session
func (db *DB) SaveContext(project, contextData, sessionInfo, keywords string) (*ContextSession, error) {
	timestamp := time.Now()

	query := `
		INSERT INTO context_sessions (project, context_data, session_info, keywords, timestamp)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := db.Exec(query, project, contextData, sessionInfo, keywords, timestamp)
	if err != nil {
		return nil, fmt.Errorf("failed to save context: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get insert ID: %w", err)
	}

	session := &ContextSession{
		ID:          id,
		Project:     project,
		Timestamp:   timestamp,
		ContextData: contextData,
		SessionInfo: sessionInfo,
		Keywords:    keywords,
		CreatedAt:   timestamp,
	}

	return session, nil
}

// Get recent context sessions for a project
func (db *DB) GetRecentContexts(project string, limit int) ([]ContextSession, error) {
	query := `
		SELECT id, project, timestamp, context_data, session_info, keywords, created_at
		FROM context_sessions 
		WHERE project = ?
		ORDER BY timestamp DESC 
		LIMIT ?
	`

	rows, err := db.Query(query, project, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to query recent contexts: %w", err)
	}
	defer rows.Close()

	var sessions []ContextSession
	for rows.Next() {
		var session ContextSession
		err := rows.Scan(
			&session.ID,
			&session.Project,
			&session.Timestamp,
			&session.ContextData,
			&session.SessionInfo,
			&session.Keywords,
			&session.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan context session: %w", err)
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

// Search stored contexts by keyword
func (db *DB) SearchStoredContexts(keyword string) ([]ContextSession, error) {
	query := `
		SELECT id, project, timestamp, context_data, session_info, keywords, created_at
		FROM context_sessions 
		WHERE context_data LIKE ? OR keywords LIKE ? OR session_info LIKE ?
		ORDER BY timestamp DESC 
		LIMIT 10
	`

	searchTerm := "%" + keyword + "%"
	rows, err := db.Query(query, searchTerm, searchTerm, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("failed to search contexts: %w", err)
	}
	defer rows.Close()

	var sessions []ContextSession
	for rows.Next() {
		var session ContextSession
		err := rows.Scan(
			&session.ID,
			&session.Project,
			&session.Timestamp,
			&session.ContextData,
			&session.SessionInfo,
			&session.Keywords,
			&session.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan context session: %w", err)
		}
		sessions = append(sessions, session)
	}

	return sessions, nil
}

// Track project activity for ecosystem intelligence
func (db *DB) TrackProject(name, path string, isGitRepo bool) error {
	query := `
		INSERT INTO projects (name, path, git_repo, last_activity)
		VALUES (?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(name) DO UPDATE SET
			last_activity = CURRENT_TIMESTAMP,
			path = excluded.path,
			git_repo = excluded.git_repo
	`

	_, err := db.Exec(query, name, path, isGitRepo)
	if err != nil {
		return fmt.Errorf("failed to track project: %w", err)
	}

	return nil
}

// Send message to another ecosystem tool (uroboro compatibility)
func (db *DB) SendToolMessage(toTool, messageType, data string) error {
	query := `
		INSERT INTO tool_messages (to_tool, message_type, data)
		VALUES (?, ?, ?)
	`

	_, err := db.Exec(query, toTool, messageType, data)
	if err != nil {
		return fmt.Errorf("failed to send tool message: %w", err)
	}

	return nil
}

// Track command usage for analytics
func (db *DB) TrackUsage(command, project string, durationMs int, success bool, errorMsg string) error {
	query := `
		INSERT INTO usage_stats (command, project, duration_ms, success, error_message)
		VALUES (?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query, command, project, durationMs, success, errorMsg)
	if err != nil {
		return fmt.Errorf("failed to track usage: %w", err)
	}

	return nil
}
