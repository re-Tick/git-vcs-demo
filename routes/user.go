package routes

import (
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

func NewUserService() User {
	return User{
		Name: "John Doe",
		Age:  30,
	}
}

func (u *User) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Println("printing the request headers", req.Header)

	// read the request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "Error reading request body",
			http.StatusInternalServerError)
		return
	}
	fmt.Println("printing the request body", string(body))

	res.Write([]byte("Hello, World!"))
}
