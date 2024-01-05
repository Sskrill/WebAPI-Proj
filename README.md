**This is a simple API application with a CRUD implementation with the creation of user data in a database (SQL).**

Example

```go
package main

import (
	webAPIUsers "github.com/Sskrill/WebAPI-Proj/internal/pkg"
	_ "github.com/lib/pq"
)

func main() {
cache := cache.NewCache(time.Second * 10)
	db := webAPIUsers.NewDB()

	handler := webAPIUsers.NewHandler(db,cache)

	webAPIUsers.NewRouting(handler)
}
```
**webAPIUsers.NewDB()** :This function returns a pointer to a **DataBase structure** that has a field of type ***sql.DB**, and the function connects to the database specified in the function.

**webAPIUsers.NewHandler(db)** :This method returns a structure with the ***Handler** type.Which has a field with the **CRUD** *interface* type (and the DataBase structure is implemented from it).And as an argument, it requires an object that is implemented from the **CRUD** *interface*.

**webAPIUsers.NewRouting(handler)** :This function starts the server with different routes.As an argument, it takes a ***Handler**, from which it will use handlers for the specified *routes*, and these handlers will use functions that are in the **CRUD** *interface*.

**cache.NewCache(time.Second * 10)** :This function creates a variable with a **cache structure** in which there are methods for caching data, and then we pass it as an argument to the handler creation function.

***This version uses net/http instead of gin. I just changed the handler file.go and also affected router.go .Now you need to write this "?id=(id)" after the users route in parentheses, you must specify its id.And I also added a small configuration with the file ".env"***
