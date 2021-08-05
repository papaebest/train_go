package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"todo/todo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

var tasks map[int]*Task = make(map[int]*Task)
var index int

type Task struct {
	Title string
	Done  bool
}

type NewTaskTodo struct {
	Task string `json:"task"`
}

func main() {
	r := mux.NewRouter()
	r.Use(authMiddleware)

	r.HandleFunc("/auth", func(rw http.ResponseWriter, r *http.Request) {
		mySigningKey := []byte("password")

		// Create the Claims
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Minute).Unix(),
			Issuer:    "test",
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		ss, err := token.SignedString(mySigningKey)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		json.NewEncoder(rw).Encode(map[string]string{
			"token": ss,
		})
	})

	api := r.NewRoute().Subrouter()
	api.Use(authMiddleware)
	// r.HandleFunc("/todos", func(rw http.ResponseWriter, r *http.Request) {
	api.HandleFunc("/todos", todo.AddTask).Methods(http.MethodPut)

	api.HandleFunc("/todos/{index}", todo.TaskDone).Methods(http.MethodPut)

	api.HandleFunc("/todos", todo.GetTask).Methods(http.MethodGet)

	http.ListenAndServe(":9090", r)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
		mySigningKey := []byte("password")
		_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return mySigningKey, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func addNum(num1 int, num2 int) (result int) {
	result = num1 + num2
	return result
}

func New(task string) {
	defer func() {
		index++
	}()

	tasks[index] = &Task{
		Title: task,
		Done:  false,
	}
	// return task
}

func List() map[int]*Task {
	return tasks
}
