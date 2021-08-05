package todo

import (
	"encoding/json"
	"net/http"
	"strconv"

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

func AddTask(rw http.ResponseWriter, r *http.Request) {
	// tokenString := r.Header.Get("Authorization")

	// tokenString = strings.ReplaceAll(tokenString, "Bearer ", "")
	// mySigningKey := []byte("password")
	// _, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
	// 	// Don't forget to validate the alg is what you expect:
	// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	// 		return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
	// 	}

	// 	// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
	// 	return mySigningKey, nil
	// })
	// if err != nil {
	// 	rw.WriteHeader(http.StatusUnauthorized)
	// 	return
	// }

	defer r.Body.Close()
	var task NewTaskTodo
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	New(task.Task)
}
func TaskDone(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	index := vars["index"]
	i, err := strconv.Atoi(index)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	// fmt.Println(tasks[i].Done)
	tasks[i].Done = true
	// task := tasks[i]
	// task.Done = true
}

func GetTask(rw http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(rw).Encode(tasks); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

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
