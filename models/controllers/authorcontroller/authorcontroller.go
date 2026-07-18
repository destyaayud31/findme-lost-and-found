package authorcontroller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"go_api_destya/config"
	"go_api_destya/helper"
	"go_api_destya/models"
	

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var authors []models.Author

	if err := config.DB.Find(&authors).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "List of authors", authors)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var author models.Author

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Create(&author).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusCreated, "Success create author", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Author not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Detail author", author)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var author models.Author

	if err := config.DB.First(&author, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Author not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&author); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&author).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusCreated, "Success Update Author", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var author models.Author
	res := config.DB.Delete(&author, id)

	if res.Error != nil {
		helper.Response(w, http.StatusInternalServerError, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, http.StatusNotFound, "Author not found", nil)
		return
	}

	helper.Response(w, http.StatusOK, "Success delete author", nil)
}
