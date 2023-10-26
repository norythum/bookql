package db

import (
	"fmt"

	sqle "github.com/dolthub/go-mysql-server"
	"github.com/dolthub/go-mysql-server/memory"
	"github.com/dolthub/go-mysql-server/server"
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/dolthub/go-mysql-server/sql/types"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbName    = "bookrepo"
	TableName = "books"
	address   = "localhost"
	port      = 3306
)

// Init function creates in memory database and instantuates listener
func Init() {
	ctx := sql.NewEmptyContext()
	engine := sqle.NewDefault(
		memory.NewDBProvider(
			createTestDatabase(ctx),
		))

	config := server.Config{
		Protocol: "tcp",
		Address:  fmt.Sprintf("%s:%d", address, port),
	}
	s, err := server.NewDefaultServer(config, engine)
	if err != nil {
		panic(err)
	}
	if err = s.Start(); err != nil {
		panic(err)
	}
}

// createTestDatabase lays in a generic starter database and table for purposes of exercise
func createTestDatabase(ctx *sql.Context) *memory.Database {
	db := memory.NewDatabase(dbName)
	db.EnablePrimaryKeyIndexes()
	table := memory.NewTable(TableName, sql.NewPrimaryKeySchema(sql.Schema{
		{Name: "title", Type: types.Text, Nullable: false, Source: TableName, PrimaryKey: true},
		{Name: "author", Type: types.Text, Nullable: false, Source: TableName, PrimaryKey: true},
		{Name: "date_pub", Type: types.Text, Nullable: false, Source: TableName},
		{Name: "book_cvr_img", Type: types.Text, Nullable: false, Source: TableName},
	}), db.GetForeignKeyCollection())
	db.AddTable(TableName, table)

	_ = table.Insert(ctx, sql.NewRow("book 1", "Jane Deo", "01/05/2023", "image_here"))
	_ = table.Insert(ctx, sql.NewRow("book 5", "Jane Doe", "01/26/2006", "image_here"))
	_ = table.Insert(ctx, sql.NewRow("book 8", "John Doe", "10/5/1877", "image_here"))
	_ = table.Insert(ctx, sql.NewRow("Book 12", "John Doe", "11/12/1987", "image_here"))
	return db
}

// Connect is a helper func to allow swapping of database in one place
func Connect() *gorm.DB {
	db, err := gorm.Open(mysql.Open("tcp(localhost:3306)/"+dbName), &gorm.Config{})
	if err != nil {
		panic("failed to open db connection")
	}
	return db
}
