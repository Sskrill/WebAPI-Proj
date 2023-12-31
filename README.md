**This is a simple API application with a CRUD implementation with the creation of user data in a database (SQL).**

Example

```go
package main

import (
	webAPIUsers "github.com/Sskrill/WebAPI-Proj/internal/pkg"
	_ "github.com/lib/pq"
)

func main() {

	db := webAPIUsers.NewDB()

	handler := webAPIUsers.NewHandler(db)

	webAPIUsers.NewRouting(handler)
}
```
**webAPIUsers.NewDB()** This function returns a pointer to a **DataBase structure** that has a field of type ***sql.DB**, and the function connects to the database specified in the function.
