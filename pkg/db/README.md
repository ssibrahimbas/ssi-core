## Ssi Core / Db

It is a package developed to establish database connections in microservice applications. Currently, only functions are written for `mongodb`, but you can also write for sql if you wish. It does not have any mandatory dependencies.

### Mongo

Here are some functions related to `mongodb`. This package uses the `mongo-driver` package. [Click to access](https://www.mongodb.com/docs/drivers/go/current/) the driver.

#### CalcMongoUri

This function is used to calculate the `mongodb` connection uri. It takes the `host`, `port`, `username`, `password`, `database` and `options` parameters. It returns the `uri` string.

Password and username not required

type:

```go
type UriParams struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}
```

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/db"
    "fmt"
)

func main() {
    uri := db.CalcMongoUri(db.UriParams{
        Host: "localhost",
        Port: "27017",
        User: "", // not required
        Pass: "", // not required
        Db:   "test-app",
	})
	fmt.Println(uri) // mongodb://localhost:27017/test-app
}
```

#### Connect

Connects to the database with the given parameters. Returns `Mongodb` struct.

type:

```go
type MongoDB struct {
	c   *mongo.Client
	db  *mongo.Database
	ctx context.Context
}
```

functions:

```go
func (m *MongoDB) GetCollection(n string) *mongo.Collection
func (m *MongoDB) TransformId(id string) primitive.ObjectID
```

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/db"
    "fmt"
)

func main() {
    uri := db.CalcMongoUri(db.UriParams{
        Host: "localhost",
        Port: "27017",
        User: "", // not required
        Pass: "", // not required
        Db:   "test-app",
	})
	d, err := db.NewMongo(uri, a.Cnf.Db.Name)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connected to database")
}
```

### GetCollection

Returns the collection with the given name.

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/db"
    "fmt"
)

func main() {
    uri := db.CalcMongoUri(db.UriParams{
        Host: "localhost",
        Port: "27017",
        User: "", // not required
        Pass: "", // not required
        Db:   "test-app",
	})
	d, err := db.NewMongo(uri, a.Cnf.Db.Name)
    if err != nil {
        panic(err)
    }
    c := d.GetCollection("users") // returns mongo.*Collection
    fmt.Println(c.Name())
}
```

### TransformId

Transforms the given string to `primitive.ObjectID`.

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/db"
    "fmt"
)

func main() {
    uri := db.CalcMongoUri(db.UriParams{
        Host: "localhost",
        Port: "27017",
        User: "", // not required
        Pass: "", // not required
        Db:   "test-app",
    })
    d, err := db.NewMongo(uri, a.Cnf.Db.Name)
    if err != nil {
        panic(err)
    }
    c := d.GetCollection("users")
    id := d.TransformId("5f9f1b9b9b9b9b9b9b9b9b9b")
    fmt.Println(id) // 5f9f1b9b9b9b9b9b9b9b9b9b
}
```