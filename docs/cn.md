# DBbook

📚 一键生成数据库表结构文档

# 特性

- 0️⃣ 零依赖
- 🔎 全文搜索
- ♻ 自动更新文档

# 下载

https://github.com/istepheny/dbbook/releases

# 配置

打开 `database.json` 填入数据库配置信息。

# 运行
```
$ ./dbbook
```

# 命令行参数
```
Usage: dbbook [-h] [-p port] [-t seconds]

Options:
  -h    显示帮助信息
  -p port
        监听端口 (默认 "3000")
  -t seconds
        每 t 秒更新一次文档 (默认 3600)
```

# 致谢

- https://github.com/docsifyjs/docsify
- https://github.com/go-xorm/xorm

# 许可证

[MIT](https://github.com/istepheny/dbbook/blob/master/LICENSE)