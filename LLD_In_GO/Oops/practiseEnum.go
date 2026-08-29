package main

// Exercise 2: HTTP Status Code
// Implement HTTP Status Code

// medium
// Problem: Create an HttpStatus enum where each status has a numeric code and a message string.

// Requirements:

// Values: OK(200, "OK"), BAD_REQUEST(400, "Bad Request"), NOT_FOUND(404, "Not Found"), INTERNAL_SERVER_ERROR(500, "Internal Server Error")
// isSuccess() method that returns true if the code is less than 400
// display() method that prints "CODE MESSAGE" (e.g. "200 OK")
// A static fromCode(int) method that returns the HttpStatus for a given code, or null/None if not found

type HttpStatus struct {
	Code int

	Message string
}
