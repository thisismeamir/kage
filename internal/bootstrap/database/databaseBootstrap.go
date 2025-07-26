package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	db *sql.DB
}

type TableSchema struct {
	TableName string
	Schema    []ColumnSchema // Column name and type
}

type ColumnSchema struct {
	Name string
	Type string
}

func (db *Database) OpenDatabase(dbPath string) (*Database, error) {
	if db.db != nil {
		return db, nil // Database already open
	}

	var err error
	db.db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	// Optionally, you can ping the database to ensure it's reachable
	if err := db.db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Database) CloseDatabase() error {
	return db.db.Close()
}

func (db *Database) CreateTable(tableSchema TableSchema) error {
	if db.db == nil {
		return sql.ErrConnDone // Database not open
	}

	// Construct the CREATE TABLE SQL statement
	createTableSQL := "CREATE TABLE IF NOT EXISTS " + tableSchema.TableName + " ("
	for i, column := range tableSchema.Schema {
		createTableSQL += column.Name + " " + column.Type
		if i < len(tableSchema.Schema)-1 {
			createTableSQL += ", "
		}
	}
	createTableSQL += ");"

	// Execute the SQL statement
	_, err := db.db.Exec(createTableSQL)
	return err
}

func (db *Database) GetAllTables() ([]string, error) {
	if db.db == nil {
		return nil, sql.ErrConnDone // Database not open
	}

	rows, err := db.db.Query("SELECT name FROM sqlite_master WHERE type='table';")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tables []string
	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			return nil, err
		}
		tables = append(tables, tableName)
	}

	return tables, nil
}

func (db *Database) GetTableSchema(tableName string) ([]ColumnSchema, error) {
	if db.db == nil {
		return nil, sql.ErrConnDone // Database not open
	}

	rows, err := db.db.Query("PRAGMA table_info(" + tableName + ");")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schema []ColumnSchema
	for rows.Next() {
		var column ColumnSchema
		var cid int // Column ID, not needed here
		if err := rows.Scan(&cid, &column.Name, &column.Type, nil, nil, nil); err != nil {
			return nil, err
		}
		schema = append(schema, column)
	}

	return schema, nil
}

func (db *Database) DropTable(tableName string) error {
	if db.db == nil {
		return sql.ErrConnDone // Database not open
	}

	// Construct the DROP TABLE SQL statement
	dropTableSQL := "DROP TABLE IF EXISTS " + tableName + ";"

	// Execute the SQL statement
	_, err := db.db.Exec(dropTableSQL)
	return err
}
