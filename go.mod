module dbbook

go 1.13

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/kr/pretty v0.2.0 // indirect
	xorm.io/core v0.7.2
	xorm.io/xorm v0.8.2
)

replace xorm.io/xorm => gitea.com/istepheny/xorm v0.0.0-20200314074132-a917a437b40d
