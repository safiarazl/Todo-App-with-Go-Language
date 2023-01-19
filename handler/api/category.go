package api

import (
	"a21hc3NpZ25tZW50/entity"
	"a21hc3NpZ25tZW50/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type CategoryAPI interface {
	GetCategory(w http.ResponseWriter, r *http.Request)
	CreateNewCategory(w http.ResponseWriter, r *http.Request)
	DeleteCategory(w http.ResponseWriter, r *http.Request)
	GetCategoryWithTasks(w http.ResponseWriter, r *http.Request)
}

type categoryAPI struct {
	categoryService service.CategoryService
}

func NewCategoryAPI(categoryService service.CategoryService) *categoryAPI {
	return &categoryAPI{categoryService}
}

func (c *categoryAPI) GetCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userId, err := strconv.Atoi(r.Context().Value("id").(string))
	if userId == 0 {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category", "invalid user id")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	getCat, err := c.categoryService.GetCategories(r.Context(), userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println("log saya", getCat)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// JsonInByte, _ := json.Marshal(getCat)
	// w.Write(JsonInByte)
	json.NewEncoder(w).Encode(getCat)
	// end answer
}

func (c *categoryAPI) CreateNewCategory(w http.ResponseWriter, r *http.Request) {
	var category entity.CategoryRequest

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}

	// TODO: answer here
	if category.Type == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("create category", "invalid category request")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
		return
	}
	userID := r.Context().Value("id")
	if userID == nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("create category", "invalid user id")
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}
	intUserID , err := strconv.Atoi(userID.(string))
	newCategory := entity.Category{
		Type: category.Type,
		UserID: intUserID,
	}
	creatCat, err := c.categoryService.StoreCategory(r.Context(), &newCategory)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	msg := map[string]interface{}{
		"user_id": creatCat.UserID,
		"category_id": creatCat.ID,
		"message": "success create new category",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(msg)
	// end answer
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// TODO: answer here
	userId := r.Context().Value("id")
	categoryID, _ := strconv.Atoi(r.URL.Query().Get("category_id"))

	err := c.categoryService.DeleteCategory(r.Context(), categoryID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}
	msg := map[string]interface{}{
		"user_id": userId,
		"category_id": categoryID,
		"message": "success delete category",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
	// end answer
}

func (c *categoryAPI) GetCategoryWithTasks(w http.ResponseWriter, r *http.Request) {
	userId := r.Context().Value("id")

	idLogin, err := strconv.Atoi(userId.(string))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("get category task", err.Error())
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
		return
	}

	categories, err := c.categoryService.GetCategoriesWithTasks(r.Context(), int(idLogin))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categories)

}
