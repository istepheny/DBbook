package schema

import (
	"dbbook/pkg/config"
	"dbbook/pkg/database"
	"dbbook/pkg/document"
	"log"
	"strings"
	"xorm.io/core"
	"xorm.io/xorm/schemas"
)

var (
	indexTypes = map[int]string{
		core.IndexType:  "index",
		core.UniqueType: "unique",
	}
)

func Query(configs []config.Database) (book document.Book) {
	for _, config := range configs {
		tables := getTables(config)
		sidebar, pages := parseTables(config, tables)
		book.Sidebars = append(book.Sidebars, sidebar)
		book.Pages = append(book.Pages, pages...)
	}

	return book
}

func getTables(config config.Database) []*schemas.Table {
	DB := database.Connect(config)
	tables, e := DB.DBMetas()
	if e != nil {
		log.Fatalf("Failed to connect to database: %s", e)
	}
	return tables
}

func parseTables(config config.Database, tables []*schemas.Table) (sidebar document.Sidebar, pages []document.Page) {
	sidebar.DataBase = config.Database
	sidebar.CoverTable = tables[0].Name

	for _, table := range tables {
		sidebar.Tables = append(sidebar.Tables, table.Name)
		page := parseTable(config, table)
		pages = append(pages, page)
	}

	return sidebar, pages
}

func parseTable(config config.Database, table *schemas.Table) (page document.Page) {
	page.DataBase = config.Database
	page.Table = table.Name
	page.Engine = table.StoreEngine
	page.Comment = table.Comment

	for _, column := range table.Columns() {
		page.Columns = append(page.Columns, document.Column{
			Name:    column.Name,
			Type:    column.SQLType.Name,
			Length:  column.SQLType.DefaultLength,
			Default: column.Default,
			Comment: column.Comment,
		})
	}

	for _, pk := range table.PrimaryKeys {
		page.Indexes = append(page.Indexes, document.Index{
			IndexName:  pk,
			ColumnName: pk,
			Type:       "primary",
		})
	}

	for _, index := range table.Indexes {
		page.Indexes = append(page.Indexes, document.Index{
			IndexName:  index.Name,
			ColumnName: strings.Join(index.Cols, ","),
			Type:       indexTypes[index.Type],
		})
	}

	return page
}
