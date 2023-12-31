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
**webAPIUsers.NewDB()** :This function returns a pointer to a **DataBase structure** that has a field of type ***sql.DB**, and the function connects to the database specified in the function.

**webAPIUsers.NewHandler(db)** :This method returns a structure with the ***Handler** type.Which has a field with the **CRUD** *interface* type (and the DataBase structure is implemented from it).And as an argument, it requires an object that is implemented from the **CRUD** *interface*.

**webAPIUsers.NewRouting(handler)** This function starts the server with different routes.As an argument, it takes a ***Handler**, from which it will use handlers for the specified *routes*, and these handlers will use functions that are in the **CRUD** *interface*.
