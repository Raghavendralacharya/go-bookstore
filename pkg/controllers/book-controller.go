package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/raghavendra/go-bookstore/pkg/utils"
	"github.com/raghavendra/go-bookstore/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks :=models.GetAllBooks();
	res, _ :=json.Marshal(newBooks)
	w.Header().Set("Content-Type","pkglication/json");
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func GetBookById(w http.ResponseWriter, r *http.Request){
	vars :=mux.Vars(r);
	bookId := vars["bookId"];
	ID,err := strconv.ParseInt(bookId,0,0);
	if err != nil {
		fmt.Println("errors while parsing")
	}
	bookDetails, _ :=models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res);
}


func CreateBook(w http.ResponseWriter, r *http.Request){
	fmt.Println("creating book", r.Body,"--");
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	fmt.Println("creating book", CreateBook.Name,"--",CreateBook.Author);
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b);
	// w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r);
	bookId := vars["bookId"]
	fmt.Println(bookId)
	ID, err := strconv.ParseInt(bookId, 0,0);
	if err != nil {
		fmt.Println("error while parsing");
	}
	book := models.DeleteBook(ID);
	res, _ := json.Marshal(book);
	w.Header().Set("Content-Type", "pkglication/json");
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}


func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r);
	bookId := vars["bookId"]
	fmt.Println(bookId)
	ID, err := strconv.ParseInt(bookId, 0,0)
	if err != nil {
		fmt.Println("error while parsing");
	}
	bookDetails, db := models.GetBookById(ID);
	if bookDetails.Name != "" {
		bookDetails.Name = updateBook.Name
	}

	if bookDetails.Author != "" {
		bookDetails.Author = updateBook.Author
	}

	if bookDetails.Publication != "" {
		bookDetails.Publication = updateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails);
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK);
	w.Write(res);
}



