# DBbook

📚 Generate database schema documentation on the fly.

# Features

- 0️⃣ Zero dependency
- 🔎 Full text search
- ♻ Automatically update your documentation

# Download

https://github.com/istepheny/dbbook/releases

# Configuration

Open `database.json` and put in your database configuration.

# Run and enjoy
```
$ ./dbbook
```

# Command Usage
```
Usage: dbbook [-h] [-p port] [-t seconds]

Options:
  -h    Show this help.
  -p port
        Listening port. (default "3000")
  -t seconds
        Update documentation every t seconds. (default 3600)
```

# Credits

- https://github.com/docsifyjs/docsify
- https://github.com/go-xorm/xorm

# License

[MIT](https://github.com/istepheny/dbbook/blob/master/LICENSE)