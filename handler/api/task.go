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
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
	}
	userid, _ := strconv.Atoi(id)
	taskID := r.URL.Query().Get("task_id")
	if taskID == "" {
		tsk, err := t.taskService.GetTasks(r.Context(), userid)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		} else {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(tsk)
			return
		}
	} else {
		taskID2, _ := strconv.Atoi(taskID)
		tsk, err := t.taskService.GetTaskByID(r.Context(), taskID2)
		if err != nil {
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
			return
		} else {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(tsk)
			return
		}
	}

	// TODO: answer here
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
	if task.Title == "" || task.Description == "" || task.CategoryID == 0 {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid task request"))
		return
	}
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	userid, _ := strconv.Atoi(id)
	tsk, err := t.taskService.StoreTask(r.Context(), &entity.Task{
		ID:          task.ID,
		UserID:      userid,
		CategoryID:  task.CategoryID,
		Title:       task.Title,
		Description: task.Description,
	})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": userid,
		"task_id": tsk.ID,
		"message": "success create new task",
	})

	// TODO: answer hsere
}

func (t *taskAPI) DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	taskID := r.URL.Query().Get("task_id")
	tskID, _ := strconv.Atoi(taskID)
	err := t.taskService.DeleteTask(r.Context(), tskID)
	w.WriteHeader(200)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": id,
		"task_id": tskID,
		"message": "success delete task",
	})

	// TODO: answer here
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
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	taskid, _ := strconv.Atoi(id)
	tsk, err := t.taskService.UpdateTask(r.Context(), &entity.Task{
		ID:          task.ID,
		UserID:      taskid,
		CategoryID:  task.CategoryID,
		Title:       task.Title,
		Description: task.Description,
	})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id": tsk.UserID,
		"task_id": taskid,
		"message": "success update task",
	})
	// TODO: answer here
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
		UserID:     int(idLogin),
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
