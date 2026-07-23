package main

import "fmt"

// ============================================================
// 1. STRUCT = CLASS REPLACEMENT
// ============================================================
//
// C++:
//
// class User {
//     private:
//         string name;
//         int age;
// };
//
// Go does not have classes.
// A struct stores data.
//
// Fields starting with lowercase are PRIVATE.
// Fields starting with uppercase are PUBLIC.
//

type User struct {

	// private fields
	// Only code inside this package can access them directly.
	name  string
	email string
	age   int

	// public field
	// Anyone can access this.
	ID int
}

// ============================================================
// 2. CONSTRUCTOR
// ============================================================
//
// C++:
//
// User(string name){
//     this->name = name;
// }
//
// Go has no constructor keyword.
//
// Convention:
//
// New<Type>()
//
// This function creates and returns a User object.
//

func NewUser(
	name string,
	email string,
	age int,
	id int,

) *User {

	// &User creates a pointer to User.
	// Similar to:
	//
	// User* u = new User();
	//

	return &User{

		name:  name,
		email: email,
		age:   age,
		ID:    id,
	}

}

// ============================================================
// 3. GETTERS
// ============================================================
//
// C++:
//
// string GetName(){
//     return name;
// }
//
// Go convention:
// Instead of GetName()
// use Name()
//
// Because the receiver already tells us it belongs to User.
//

func (u User) Name() string {

	return u.name

}

func (u User) Email() string {

	return u.email

}

func (u User) Age() int {

	return u.age

}

// ============================================================
// 4. SETTERS
// ============================================================
//
// C++:
//
// void SetAge(int age){
//
// }
//
//
//
// In Go setters usually use pointer receivers.
//
// Why?
//
// Because we need to modify the original object.
//
//

func (u *User) SetName(name string) {

	u.name = name

}

func (u *User) SetAge(age int) {

	// validation logic
	// This is encapsulation.
	//
	// Nobody can put invalid age directly.

	if age > 0 {

		u.age = age

	}

}

// ============================================================
// 5. ENCAPSULATION
// ============================================================
//
// We hide data:
//
// name
// email
// age
//
// They cannot be modified directly:
//
// user.age = -100   ❌
//
// They must go through methods:
//
// user.SetAge(25)   ✅
//
// This is similar to:
//
// private variables + public methods
//

func (u User) Display() {

	fmt.Println(
		"User:",
		u.name,
		u.email,
		u.age,
	)

}

// ============================================================
// 6. POINTER RECEIVER VS VALUE RECEIVER
// ============================================================
//
// VALUE RECEIVER:
//
// func (u User)
//
// receives a COPY.
//
// Good for reading data.
//
//
//
// POINTER RECEIVER:
//
// func (u *User)
//
// receives original memory address.
//
// Used for modifications.
//

func (u User) PrintInfo() {

	fmt.Println(
		u.name,
	)

}

// ============================================================
// 7. STATIC VARIABLE EQUIVALENT
// ============================================================
//
// C++:
//
// class User {
//
// static int count;
//
// }
//
//
//
// Go uses package variables.
//
//

var totalUsers int

func CreateUser(
	name string,
	email string,
	age int,

) *User {

	totalUsers++

	return NewUser(
		name,
		email,
		age,
		totalUsers,
	)

}

// ============================================================
// 8. INTERFACE = ABSTRACT CLASS / POLYMORPHISM
// ============================================================
//
// C++:
//
// class Auth {
//     virtual void Login() = 0;
// };
//
//
//
// Go:
//
// Interface defines behavior.
//
// No inheritance required.
//

type Authenticator interface {
	Login()
}

// ============================================================
// 9. USER IMPLEMENTS INTERFACE
// ============================================================
//
// In Go there is no:
//
// implements Authenticator
//
// keyword.
//
// If a type has Login(),
// it automatically satisfies the interface.
//

func (u User) Login() {

	fmt.Println(
		"User login:",
		u.name,
	)

}

// ============================================================
// 10. ANOTHER TYPE IMPLEMENTING SAME INTERFACE
// ============================================================
//
// This is polymorphism.
//
// Different objects.
// Same behavior.
//

type Admin struct {
	User

	permissions []string
}

func (a Admin) Login() {

	fmt.Println(
		"Admin login:",
		a.name,
	)

}

// ============================================================
// 11. POLYMORPHIC FUNCTION
// ============================================================
//
// This function does not care whether it receives:
//
// User
//
// or
//
// Admin
//
// It only cares that Login() exists.
//

func Authenticate(
	a Authenticator,

) {

	a.Login()

}

// ============================================================
// 12. COMPOSITION INSTEAD OF INHERITANCE
// ============================================================
//
// C++:
//
// class Admin : public User
//
// Go:
//
// Embed User inside Admin.
//
//
//
// Admin gets User behavior.
//

func CreateAdmin() Admin {

	user := NewUser(
		"Admin",
		"admin@test.com",
		40,
		999,
	)

	return Admin{

		User: *user,

		permissions: []string{

			"delete",
			"create",
		},
	}

}

// ============================================================
// 13. MAIN FUNCTION
// ============================================================

func main() {

	// Object creation
	//
	// Similar to:
	//
	// User user = new User();

	user := CreateUser(
		"John",
		"john@test.com",
		30,
	)

	fmt.Println(
		"Total Users:",
		totalUsers,
	)

	// Getter

	fmt.Println(
		user.Name(),
	)

	// Setter

	user.SetAge(35)

	user.Display()

	// Public field access

	fmt.Println(
		user.ID,
	)

	// Interface polymorphism

	Authenticate(user)

	// Composition example

	admin := CreateAdmin()

	Authenticate(admin)

	fmt.Println(
		admin.permissions,
	)

}
