package web

import (
	"a21hc3NpZ25tZW50/client"
	"a21hc3NpZ25tZW50/entity"
	"embed"
	"encoding/json"
	"net/http"
	"path"
	"text/template"
)

type ModifyWeb interface {
	AddTask(w http.ResponseWriter, r *http.Request)
	AddTaskProcess(w http.ResponseWriter, r *http.Request)
	AddCategory(w http.ResponseWriter, r *http.Request)
	AddCategoryProcess(w http.ResponseWriter, r *http.Request)

	UpdateTask(w http.ResponseWriter, r *http.Request)
	UpdateTaskProcess(w http.ResponseWriter, r *http.Request)

	DeleteTask(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
}

type modifyWeb struct {
	taskClient     client.TaskClient
	categoryClient client.CategoryClient
	embed          embed.FS
}

func NewModifyWeb(tC client.TaskClient, cC client.CategoryClient, embed embed.FS) *modifyWeb {
	return &modifyWeb{tC, cC, embed}
}

func (a *modifyWeb) AddTask(w http.ResponseWriter, r *http.Request) {
	catId := r.URL.Query().Get("category")

	// ignore this
	_ = catId
	//
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "main", "add-task.html")
	tmpl, err := template.ParseFS(a.embed, loginPath, headerPath)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	if err := tmpl.Execute(w, catId); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	// TODO: answer here
}

func (a *modifyWeb) AddTaskProcess(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")
	title := r.FormValue("title")
	description := r.FormValue("description")
	category := r.URL.Query().Get("category")

	respCode, err := a.taskClient.CreateTask(title, description, category, userId.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ignore this
	_ = respCode
	//
	if respCode == 201 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/task/add?category="+category, http.StatusSeeOther)
	}
	// TODO: answer here
}

func (a *modifyWeb) AddCategory(w http.ResponseWriter, r *http.Request) {
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "main", "add-category.html")
	tmpl, err := template.ParseFS(a.embed, loginPath, headerPath)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	if err := tmpl.Execute(w, nil); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	// TODO: answer here
}

func (a *modifyWeb) AddCategoryProcess(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	category := r.FormValue("type")

	respCode, err := a.categoryClient.AddCategories(category, userId.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// ignore this
	_ = respCode

	if respCode == 201 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/category/add", http.StatusSeeOther)
	}

	// TODO: answer here
}

func (a *modifyWeb) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	task, err := a.taskClient.GetTaskById(taskId, r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// ignore this
	_ = task
	//
	headerPath := path.Join("views", "general", "header.html")
	loginPath := path.Join("views", "main", "update-task.html")
	tmpl, err := template.ParseFS(a.embed, loginPath, headerPath)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	if err := tmpl.Execute(w, task); err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	// TODO: answer here
}

func (a *modifyWeb) UpdateTaskProcess(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	categoryId := r.URL.Query().Get("category_id")
	if categoryId == "" {
		title := r.FormValue("title")
		description := r.FormValue("description")

		respCode, err := a.taskClient.UpdateTask(taskId, title, description, r.Context().Value("id").(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if respCode == 200 {
			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/task/update?task_id="+taskId, http.StatusSeeOther)
		}
	} else {
		_, err := a.taskClient.UpdateCategoryTask(taskId, categoryId, r.Context().Value("id").(string))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}

}

func (a *modifyWeb) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	_, err := a.taskClient.DeleteTask(taskId, r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)

	// TODO: answer here
}

func (a *modifyWeb) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.URL.Query().Get("category_id")

	_, err := a.categoryClient.DeleteCategory(categoryId, r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/dashboard", http.StatusFound)

	// TODO: answer here
}
