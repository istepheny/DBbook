package document

import (
	"dbbook/pkg/helper"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"text/template"
	"time"
)

type Book struct {
	Sidebars []Sidebar
	Pages    []Page
}

type Sidebar struct {
	DataBase   string
	CoverTable string
	Tables     []string
}

type Page struct {
	DataBase string
	Engine   string
	Table    string
	Comment  string
	Columns  []Column
	Indexes  []Index
}

type Column struct {
	Name    string
	Type    string
	Length  int
	Default string
	Comment string
}

type Index struct {
	IndexName  string
	ColumnName string
	Type       string
}

func Write(book Book) {
	mkDocsDir()

	for _, sidebar := range book.Sidebars {
		SidebarRender := renderSidebar(sidebar)
		writeSidebar(SidebarRender)
	}

	for _, page := range book.Pages {
		databaseDir := mkDatabaseDir(page.DataBase)
		pageRendered := renderPage(page)
		writePage(databaseDir, page.Table, pageRendered)
	}

	writeReadMe(renderReadMe())
}

func renderReadMe() []byte {
	templateReadMe := getTemplateReadMe()

	tmpl, e := template.New("readme").Parse(templateReadMe)

	if e != nil {
		log.Fatal(e)
	}

	time := time.Now().Format("2006-01-02 15:04:05")

	return render(tmpl, time)
}

func writeReadMe(readMeRendered []byte) {
	readMePath := helper.BookPath() + "README.md"
	ioutil.WriteFile(readMePath, readMeRendered, 0644)
}

func renderPage(page Page) []byte {
	templatePage := getTemplatePage()

	tmpl, e := template.New("page").Parse(templatePage)

	if e != nil {
		log.Fatal(e)
	}

	return render(tmpl, page)
}

func writePage(databaseDir, tableName string, PageRendered []byte) {
	pagePath := databaseDir + tableName + ".md"
	ioutil.WriteFile(pagePath, PageRendered, 0644)
}

func renderSidebar(sidebar Sidebar) []byte {
	templateSidebar := getTemplateSidebar()
	tmpl, e := template.New("sidebar").Parse(templateSidebar)

	if e != nil {
		log.Fatal(e)
	}

	return render(tmpl, sidebar)
}

func writeSidebar(sidebarRendered []byte) {
	sidebarPath := helper.BookPath() + "_sidebar.md"

	f, e := os.OpenFile(sidebarPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if e != nil {
		log.Fatal(e)
	}

	defer f.Close()

	if _, e = f.WriteString(string(sidebarRendered)); e != nil {
		log.Fatal(e)
	}
}

func render(tmpl *template.Template, data interface{}) []byte {

	buffer := new(bytes.Buffer)
	e := tmpl.Execute(buffer, data)

	if e != nil {
		log.Fatal(e)
	}

	return buffer.Bytes()
}

func mkDocsDir() {
	docsPath := helper.BookPath()
	_ = os.RemoveAll(docsPath)
	helper.Mkdir(docsPath)
}

func mkDatabaseDir(databaseName string) string {
	databaseDir := helper.BookPath() + databaseName + string(os.PathSeparator)

	helper.Mkdir(databaseDir)

	return databaseDir
}

func getTemplateReadMe() string {
	return getTemplate("readme")
}
func getTemplatePage() string {
	return getTemplate("page")
}

func getTemplateSidebar() string {
	return getTemplate("sidebar")
}

func getTemplate(template string) string {
	templatePath := helper.TemplatePath() + template + ".tmpl"
	templateBytes, e := ioutil.ReadFile(templatePath)

	if e != nil {
		log.Fatalln(e)
	}

	return string(templateBytes)
}
