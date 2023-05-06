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
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
	}
	userid, _ := strconv.Atoi(id)
	ctg, err := c.categoryService.GetCategories(r.Context(), userid)
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(ctg)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
	}

	// TODO: answer here
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
	if category.Type == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid category request"))
	}
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
	}
	userid, _ := strconv.Atoi(id)
	ctg, err := c.categoryService.StoreCategory(r.Context(), &entity.Category{
		Type:   category.Type,
		UserID: userid,
	})
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":     userid,
		"category_id": ctg.ID,
		"message":     "success create new category",
	})
	// TODO: answer here
}

func (c *categoryAPI) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("id").(string)
	if id == "" {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("invalid user id"))
	}
	categoryID := r.URL.Query().Get("category_id")
	catID, _ := strconv.Atoi(categoryID)
	err := c.categoryService.DeleteCategory(r.Context(), catID)
	w.WriteHeader(200)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(entity.NewErrorResponse("error internal server"))
	}

	w.WriteHeader(200)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":     id,
		"category_id": catID,
		"message":     "success delete category",
	})
	// TODO: answer here
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
