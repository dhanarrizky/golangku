package categorycontroller

import (
	"go-web-native/entities"
	"go-web-native/models/categorymodel"
	"html/template"
	"net/http"
	"strconv"
	"time"

)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodel.GetAll()
	data := map[string]any{
		"categories": categories,
	}

	// mengambil template apakah ada eror atau tidak
	temp, err := template.ParseFiles("views/category/index.html")
	if err != nil {
		panic(err)
	}
	// w di gunakkan untuk menulis, dan nilai di terakhir di gunakkan untuk deklarasi, nilai apa yang akan di berikan
	temp.Execute(w, data)

}

func Add(w http.ResponseWriter, r *http.Request) {
	// maksud dari r adalah request
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/create.html")
		if err != nil {
			panic(err)
		}

		temp.Execute(w, nil)
	}

	if r.Method == "POST" {
		var category entities.Category

		category.Name = r.FormValue("name")
		category.CreateAt = time.Now()
		category.UpdateAt = time.Now()

		if ok := categorymodel.Create(category); !ok {
			temp, _ := template.ParseFiles("views/category/create.html")
			temp.Execute(w, nil)
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}
}

func Edit(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		temp, err := template.ParseFiles("views/category/edit.html")
		if err != nil {
			panic(err)
		}

		// kenapa pakai name idsting karena default yang akan di kirim oleh go dalam bentuk string,
		// sehingga jika kita ingin menggunakkannuya untukid apa untuk oprasi jangan lupa untuk di convert terlebihdahulu
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)

		if err != nil {
			panic(err)
		}

		category := categorymodel.Detail(id)
		data := map[string]any{
			"category":category,
		}

		temp.Execute(w, data)
	}

	// jika edit.html mengirim suatu data
	if r.Method == "POST"{
		var category entities.Category

		// mengambil nilai id yang ada di form (hidden)
		idString := r.FormValue("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		category.Name = r.FormValue("name")
		category.UpdateAt = time.Now()

		if ok := categorymodel.Update(id, category); !ok {
			http.Redirect(w, r, r.Header.Get("Referer"), http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/categories", http.StatusSeeOther)
	}

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		panic(err)
	}

	// panggil category model untuk delete, dengan fungsi di bawah dan sekalian cek adakah eror atau tidak
	if err := categorymodel.Delete(id); err != nil {
		panic(err)
	}

	http.Redirect(w, r, "categories", http.StatusSeeOther)
}
