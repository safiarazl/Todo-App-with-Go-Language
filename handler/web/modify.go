package web

import (
	"a21hc3NpZ25tZW50/client"
	"embed"
	"log"
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

	// TODO: answer here
	addTask := path.Join("views", "main", "add-task.html")
	header := path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(a.embed, addTask, header))
	exc := tmpl.Execute(w, catId)
	if exc != nil {		
		http.Error(w, exc.Error(), http.StatusInternalServerError)
		return
	}
	// end of answer
}

func (a *modifyWeb) AddTaskProcess(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	title := r.FormValue("title")
	description := r.FormValue("description")
	// category := r.URL.Query().Get("category")
	category := r.FormValue("category")
	// log.Println("hasil cat addtaskprocess: ", category)
	respCode, err := a.taskClient.CreateTask(title, description, category, userId.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		// log.Println("error in add task process: ", err.Error())
		return
	}

	// ignore this
	_ = respCode
	//

	// TODO: answer here
	link := "/task/add?category=" + category
	if respCode == 201 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	} else {
		http.Redirect(w, r, link, http.StatusSeeOther)
	}
	// ragu
	// end of answer
}

func (a *modifyWeb) AddCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	addCat := path.Join("views", "main", "add-category.html")
	header := path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(a.embed, addCat, header))
	exc := tmpl.Execute(w, nil)
	if exc != nil {		
		http.Error(w, exc.Error(), http.StatusInternalServerError)
		return
	}
	// end of answer
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
	//

	// TODO: answer here
	if respCode == 201 {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/category/add", http.StatusSeeOther)
	}
	// end of answer
}

func (a *modifyWeb) UpdateTask(w http.ResponseWriter, r *http.Request) {
	// taskId := r.URL.Query().Get("task_id")
	taskId := r.FormValue("task_id")

	task, err := a.taskClient.GetTaskById(taskId, r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("error di addtask: ", err.Error())
		return
	}

	// ignore this
	_ = task
	//

	// TODO: answer here
	updtTask := path.Join("views", "main", "update-task.html")
	header := path.Join("views", "general", "header.html")

	var tmpl = template.Must(template.ParseFS(a.embed, updtTask, header))
	exc := tmpl.Execute(w, task)
	if exc != nil {		
		http.Error(w, exc.Error(), http.StatusInternalServerError)
		return
	}
	// end of answer
}

func (a *modifyWeb) UpdateTaskProcess(w http.ResponseWriter, r *http.Request) {
	// taskId := r.URL.Query().Get("task_id")
	taskId := r.FormValue("task_id")
	// categoryId := r.FormValue("category_id")
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

	// TODO: answer here
	if err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
	// end of answer
}

func (a *modifyWeb) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	categoryId := r.URL.Query().Get("category_id")

	_, err := a.categoryClient.DeleteCategory(categoryId, r.Context().Value("id").(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: answer here
	if err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
	// end of answer
}
