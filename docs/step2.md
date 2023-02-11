# Database storage

## Status

This step uses a SQL backend storage, thanks to https://gorm.io

- Add a DB storage to store todo that have title and body (gorm.io)
- Create 5 swagger endpoints
  - List todos
  - Create a todo
  - Read a todo
  - Update a todo
  - Delete a todo


## Content

- Using gorm.io (https://gorm.io/docs/)
- Creating 1 DB entity (todo) and db.AutoMigrate() this entity
- Updating go-swagger to create different endpoints to manipulate these entities

## Code explanation

if you checkout step2 branch, you need to understand the code:

### gorm.io

Gorm is based on DB entity model/definition. In the `pkg/handler/db/todo.go`, there is a Todo definition:

```

type Todo struct {
	gorm.Model

	Title string
	Body  string `sql:"type:text"`
}
```

When we will use (in the `pkg/handler/db/db.go`) the 

```
...
var AutoMigrateTables = []interface{}{
	Todo{},
}
...

func NewDB() (*gorm.DB, error) {
  ...
	err = db.AutoMigrate(AutoMigrateTables...)
  ...
```

Gorm will connect to the database and create a Table called todo, with the
- title string column
- body string column

but also (because of the `gorm.Model`) 4 others columns automatically:
- ID uint (autoincrement)
- CreatedAt time.Time
- UpdatedAt time.Time
- DeletedAt DeletedAt

