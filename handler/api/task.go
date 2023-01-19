package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type TaskAPI interface {
	GetTask(w http.ResponseWriter, r *http.Request)
	CreateNewTask(w http.ResponseWriter, r *http.Request)
	UpdateTask(w http.ResponseWriter, r *http.Request)
	DeleteTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskCategory(w http.ResponseWriter, r *http.Request)
}

type taskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) *taskAPI {
	return &taskAPI{taskService}
}

func (t *taskAPI) GetTask(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userId := r.Context().Value("id")
	if userId == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get task", "invalid user id")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		idUser, _ := strconv.Atoi(userId.(string))
		getTask , err := t.taskService.GetTasks(r.Context(), idUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getTask)
	} else {
		idTask, _ := strconv.Atoi(taskID)
		getTask, err := t.taskService.GetTaskByID( r.Context(), idTask)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(getTask)
	}
	// end answer
}

func (t *taskAPI) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}

	// TODO: answer here
	if task.Title == "" || task.Description == "" || task.CategoryID == 0 || task == (entity.TaskRequest{}) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}
	userId := r.Context().Value("id")
	if userId == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get task", "invalid user id")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	intUserID, _ := strconv.Atoi(userId.(string))
	taskET := entity.Task{
		ID: task.ID,
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID:      intUserID,
	}
	storeTask, err := t.taskService.StoreTask(r.Context(), &taskET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	msg := map[string]interface{}{
		"user_id": storeTask.UserID,
		"task_id": storeTask.ID,
		"message": "success create new task",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
	// end answer
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userID := r.Context().Value("id")
	taskID, _ := strconv.Atoi(r.URL.Query().Get("task_id"))
	err := t.taskService.DeleteTask(r.Context(), taskID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	msg := map[string]interface{}{
		"user_id": userID,
		"task_id": taskID,
		"message": "success delete task",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
	// end answer
}

func (t *taskAPI) UpdateTask(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	// TODO: answer here
	userID := r.Context().Value("id")
	if userID == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get task", "invalid user id")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	idUser, _ := strconv.Atoi(userID.(string))
	taskET := entity.Task{
		ID : task.ID,
		Title:       task.Title,
		Description: task.Description,
		CategoryID:  task.CategoryID,
		UserID: 	idUser,
	}
	updateTas,err := t.taskService.UpdateTask(r.Context(), &taskET)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	msg := map[string]interface{}{
		"user_id": updateTas.UserID,
		"task_id": updateTas.ID,
		"message": "success update task",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
	// end answer
}

func (t *taskAPI) UpdateTaskCategory(w http.ResponseWriter, r *http.Request) {
	var task entity.TaskCategoryRequest

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid decode json"))
		return
	}

	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	var updateTask = entity.Task{
		ID:         task.ID,
		CategoryID: task.CategoryID,
		UserID:     idLogin,
	}

	_, err = t.taskService.UpdateTask(r.Context(), &updateTask)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userId,
		"task_id": task.ID,
		"message": "success update task category",
	})
}
