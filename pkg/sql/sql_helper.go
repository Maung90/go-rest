package sql

import (
	"database/sql"
	"fmt"
	"strings"
)

// SQLBuilder struct utama untuk membuat query dinamis
type SQLBuilder struct {
	db        *sql.DB
	table     string
	selectCols []string
	whereClauses []string
	orderBy   string
	limit     string
	args      []interface{}
}

// Constructor
func NewSQLBuilder(db *sql.DB, table string) *SQLBuilder {
	return &SQLBuilder{
		db:    db,
		table: table,
		selectCols: []string{"id", "created_at", "updated_at"}, // default
	}
}

// Select kolom tertentu
func (b *SQLBuilder) Select(columns ...string) *SQLBuilder {
	if len(columns) > 0 {
		b.selectCols = append(b.selectCols, columns...)
	}
	return b
}

// Where dengan parameter dinamis
func (b *SQLBuilder) Where(condition string, args ...interface{}) *SQLBuilder {
	b.whereClauses = append(b.whereClauses, condition)
	b.args = append(b.args, args...)
	return b
}

// OrderBy clause
func (b *SQLBuilder) OrderBy(order string) *SQLBuilder {
	b.orderBy = order
	return b
}

// Limit clause
func (b *SQLBuilder) Limit(limit int) *SQLBuilder {
	b.limit = fmt.Sprintf("LIMIT %d", limit)
	return b
}

// Get → SELECT * FROM table WHERE ... ORDER BY ... LIMIT ...
func (b *SQLBuilder) Get() (*sql.Rows, error) {
	query := fmt.Sprintf("SELECT %s FROM %s", strings.Join(b.selectCols, ", "), b.table)

	if len(b.whereClauses) > 0 {
		query += " WHERE " + strings.Join(b.whereClauses, " AND ")
	}
	if b.orderBy != "" {
		query += " ORDER BY " + b.orderBy
	}
	if b.limit != "" {
		query += " " + b.limit
	}

	stmt, err := b.db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(b.args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

// Insert → INSERT INTO table (columns...) VALUES (?, ?, ...)
func Insert(db *sql.DB, table string, data map[string]interface{}) (sql.Result, error) {
	columns := []string{}
	placeholders := []string{}
	args := []interface{}{}

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		args = append(args, val)
	}

	query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// Update → UPDATE table SET col1=?, col2=? WHERE condition
func Update(db *sql.DB, table string, data map[string]interface{}, where string, whereArgs ...interface{}) (sql.Result, error) {
	setClauses := []string{}
	args := []interface{}{}

	for col, val := range data {
		setClauses = append(setClauses, fmt.Sprintf("%s = ?", col))
		args = append(args, val)
	}

	args = append(args, whereArgs...)

	query := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s",
		table,
		strings.Join(setClauses, ", "),
		where,
	)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	return stmt.Exec(args...)
}

// Delete → DELETE FROM table WHERE condition
func Delete(db *sql.DB, table string, where string, whereArgs ...interface{}) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", table, where)

	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	return stmt.Exec(whereArgs...)
}

// RawQuery menjalankan custom SELECT query dengan placeholder aman.
func RawQuery(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	return rows, nil
}

// RawExec menjalankan custom INSERT, UPDATE, DELETE query dengan placeholder aman.
func RawExec(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return nil, fmt.Errorf("exec failed: %w", err)
	}
	return result, nil
}


// 💡 Contoh Penggunaan

//  🔍 Select
// rows, err := helper.NewSQLBuilder(db, "daily_stories").
// 	Select("user_id", "story_text", "mood").
// 	Where("user_id = ?", 1).
// 	OrderBy("created_at DESC").
// 	Limit(5).
// 	Get()


//  ➕ Insert
// _, err := helper.Insert(db, "daily_stories", map[string]interface{}{
// 	"user_id":     1,
// 	"story_date":  "2025-11-03",
// 	"story_text":  "Hari ini belajar Golang SQL builder",
// 	"mood":        "excited",
// 	"created_at":  time.Now(),
// 	"updated_at":  time.Now(),
// })


//  ✏️ Update
// _, err := helper.Update(db, "daily_stories", map[string]interface{}{
// 	"mood": "happy",
// }, "id = ? AND user_id = ?", 5, 1)


//  ❌ Delete 
// _, err := helper.Delete(db, "daily_stories", "id = ? AND user_id = ?", 5, 1)


//  ✅ Custom SELECT
// rows, err := helper.RawQuery(db, `
// 	SELECT user_id, COUNT(*) as total_story, mood 
// 	FROM daily_stories 
// 	WHERE created_at BETWEEN ? AND ? 
// 	GROUP BY mood, user_id
// `, "2025-10-01", "2025-10-31")

// if err != nil {
// 	log.Fatal(err)
// }
// defer rows.Close()

// for rows.Next() {
// 	var userID int
// 	var total int
// 	var mood string
// 	rows.Scan(&userID, &total, &mood)
// 	fmt.Println(userID, mood, total)
// }

//  ✅ Custom UPDATE / DELETE
// _, err := helper.RawExec(db, `
// 	UPDATE daily_stories 
// 	SET mood = ? 
// 	WHERE id = ? AND user_id = ?
// `, "happy", 12, 1)

// if err != nil {
// 	log.Fatal(err)
// }

