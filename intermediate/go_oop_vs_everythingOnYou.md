Coming from C++, the biggest mental shift in Go is this:

> Go does not replace classes with a different class syntax. Instead, it combines **structs + methods + interfaces + packages** to achieve most OOP goals, but with a different philosophy.

In C++ you might think:

```
Class = Data + Methods + Access Control + Inheritance
```

In Go:

```
Struct = Data
Methods = Behavior
Packages = Access Control
Interfaces = Polymorphism
Composition = Code Reuse (instead of inheritance)
```

Let's build a production-style `User` example and map every concept to OOP.

---

# 1. C++ Class vs Go Struct

### C++

```cpp
class User {
private:
    string name;
    int age;

public:
    void SetName(string name);
    string GetName();
};
```

### Go equivalent

```go
package user

type User struct {
    name string
    age  int
}
```

Here:

* `User` is your class-like structure.
* `name` and `age` are private fields.

But wait:

How does Go know private/public?

---

# 2. Access Modifiers (public/private)

In C++:

```cpp
private:
    string name;

public:
    string GetName();
```

In Go, access is controlled by **capitalization**.

## Public

Starts with uppercase:

```go
type User struct {
    Name string
}
```

Anyone can access:

```go
u.Name = "John"
```

---

## Private

Starts with lowercase:

```go
type User struct {
    name string
}
```

Only the same package can access:

```go
u.name = "John"
```

Outside package:

```go
user.name // ❌ Error
```

This is Go's replacement for access modifiers.

---

# 3. Encapsulation

In OOP:

> Hide internal data and expose controlled operations.

Example:

C++:

```cpp
class User {

private:
    int age;

public:

    void SetAge(int age){
        if(age > 0)
            this->age = age;
    }

};
```

Go:

```go
package user


type User struct {
    name string
    age  int
}


func (u *User) SetAge(age int) {

    if age > 0 {
        u.age = age
    }

}
```

Now nobody can directly do:

```go
u.age = -10
```

because it is private.

They must use:

```go
u.SetAge(25)
```

This is encapsulation.

---

# 4. Constructor

Go has no constructor keyword.

C++:

```cpp
User(string name){
    this->name=name;
}
```

Go uses a factory function.

```go
func NewUser(name string, age int) *User {

    return &User{
        name: name,
        age: age,
    }
}
```

Usage:

```go
u := user.NewUser("Alex",25)
```

Convention:

```
New<Type>()
```

is the Go equivalent of a constructor.

---

# 5. Complete User Class Equivalent

A realistic Go user package:

```
user/
 |
 |-- user.go
```

## user.go

```go
package user


type User struct {

    name string
    email string
    age int

}


func NewUser(
    name string,
    email string,
    age int,
) *User {


    return &User{
        name:name,
        email:email,
        age:age,
    }

}
```

---

# 6. Getters

C++:

```cpp
string GetName(){
    return name;
}
```

Go:

```go
func (u User) Name() string {

    return u.name

}
```

Notice:

Go does not usually write:

```go
GetName()
```

Instead:

```go
Name()
```

Usage:

```go
fmt.Println(u.Name())
```

---

# 7. Setters

C++:

```cpp
void SetName(string name)
{
    this->name=name;
}
```

Go:

```go
func (u *User) SetName(name string){

    u.name=name

}
```

Usage:

```go
u.SetName("Bob")
```

---

# 8. Why Pointer Receiver in Setters?

Look:

```go
func (u User) SetName(name string)
```

This receives a copy.

Example:

```
Original User

name="Alex"


Copy User

name="Alex"
```

You modify the copy.

The original stays unchanged.

---

Pointer:

```go
func (u *User) SetName(name string)
```

Now:

```
u -----> Original User
```

Changes affect the real object.

Therefore:

* Getters → value receiver
* Setters → pointer receiver

Common Go style.

---

# 9. Full Production Style User

```go
package user


type User struct {

    name string
    email string
    age int

}


func NewUser(
    name string,
    email string,
    age int,
) *User {


    return &User{
        name:name,
        email:email,
        age:age,
    }

}


func (u User) Name() string {

    return u.name

}


func (u User) Email() string {

    return u.email

}


func (u User) Age() int {

    return u.age

}


func (u *User) SetName(name string){

    u.name=name

}


func (u *User) SetEmail(email string){

    u.email=email

}


func (u *User) SetAge(age int){

    if age > 0 {
        u.age=age
    }

}
```

Usage:

```go
package main


import (
    "fmt"
    "myapp/user"
)


func main(){

    u := user.NewUser(
        "John",
        "john@gmail.com",
        30,
    )


    fmt.Println(u.Name())


    u.SetAge(35)


    fmt.Println(u.Age())

}
```

---

# 10. Static Members (C++)

C++:

```cpp
class User{

static int count;

};
```

Go has no static fields.

Usually:

```go
var userCount int
```

package-level variable:

```go
package user


var totalUsers int


func NewUser(...) *User {

    totalUsers++

    return &User{}

}
```

---

# 11. Inheritance

C++:

```cpp
class Admin : public User
{

};
```

Go does NOT do inheritance.

Instead:

## Composition

C++ inheritance:

```
Admin
 |
 User
```

Go:

```go
type Admin struct {

    User

    permissions []string

}
```

Now:

```go
admin.Name()
```

works.

This is called embedding.

---

# 12. Polymorphism

C++:

```cpp
virtual void Login();
```

Go uses interfaces.

Example:

```go
type Authenticator interface {

    Login()

}
```

Any type with Login() satisfies it.

User:

```go
func (u User) Login(){

    fmt.Println("User login")

}
```

Admin:

```go
func (a Admin) Login(){

    fmt.Println("Admin login")

}
```

Now:

```go
func Authenticate(a Authenticator){

    a.Login()

}
```

Both work.

No `implements` keyword.

---

# 13. Final OOP Mapping

| OOP Concept     | C++               | Go                        |
| --------------- | ----------------- | ------------------------- |
| Class           | class             | struct                    |
| Object          | object            | struct instance           |
| Constructor     | constructor       | NewFunction               |
| Private         | private           | lowercase fields          |
| Public          | public            | uppercase fields          |
| Getter          | GetName()         | Name()                    |
| Setter          | SetName()         | SetName()                 |
| Encapsulation   | access modifiers  | packages + private fields |
| Inheritance     | class inheritance | composition               |
| Virtual methods | virtual           | interfaces                |
| Abstract class  | abstract class    | interface                 |
| Static          | static            | package variables         |
| Method          | member function   | receiver method           |

---

# How Production Go Code Usually Looks

A common backend structure:

```
project/

cmd/
   api/
      main.go

internal/

   user/
      user.go
      service.go
      repository.go

   database/
      postgres.go

```

The `User` struct is usually just the data model.

Business logic goes into services:

```go
type UserService struct {
    repo UserRepository
}


func (s UserService) CreateUser(){

}
```

This is more common in Go than putting everything inside a giant "User class".

---

The key mindset change from C++:

**C++ asks: "What object hierarchy should I create?"**

**Go asks: "What data exists, what behavior is needed, and what interfaces connect them?"**

Go can do OOP, but it favors **composition over inheritance and small focused types over large classes.**
