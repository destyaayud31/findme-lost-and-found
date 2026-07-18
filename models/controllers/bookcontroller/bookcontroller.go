package bookcontroller

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
	var books []models.Book
	var bookResponse []models.BookResponse

	if err := config.DB.Joins("Author").Find(&books).Find(&bookResponse).Error; err != nil {
		helper.Response(w, 500, "Books Not Found", nil)
		return
	}
	helper.Response(w, 200, "List Books", bookResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var author models.Author
	err := config.DB.First(&author, book.AuthorID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Author Not Found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}

	helper.Response(w, 201, "Success Create Book", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	var bookResponse models.Book

	if err := config.DB.Joins("Author").First(&book, id).Find(&bookResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Book not found", nil)
			return
		}

		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, 200, "Detail Book", bookResponse)
	return

}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, http.StatusNotFound, "Book not found", nil)
			return
		}
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	var bookPayload models.Book
	if err := json.NewDecoder(r.Body).Decode(&bookPayload); err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	var author models.Author
	if bookPayload.AuthorID != 0 {
		if err := config.DB.First(&author, bookPayload.AuthorID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				helper.Response(w, 404, "Author not found", nil)
				return
			}
			helper.Response(w, 500, err.Error(), nil)
			return
		}
	}

	if err := config.DB.Model(&book).Updates(bookPayload).Error; err != nil {
		helper.Response(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.Response(w, http.StatusOK, "Success Update Book", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var book models.Book
	res := config.DB.Delete(&book, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0 {
		helper.Response(w, 404, "Book not found", nil)
		return
	}

	helper.Response(w, 200, "Success delete book", nil)
}

