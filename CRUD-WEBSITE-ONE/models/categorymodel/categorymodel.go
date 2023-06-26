package categorymodel

import (
	"go-web-native/config"
	"go-web-native/entities"

)

// jangan lupa mendefinisikan nilai apa yang akan di return
func GetAll() []entities.Category {
	rows, err := config.DB.Query("SELECT * FROM categories")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		if err := rows.Scan(&category.Id, &category.Name, &category.CreateAt, &category.UpdateAt); err != nil {
			panic(err)
		}

		categories = append(categories, category)
	}

	return categories

}

// mengembalikan dengan nilai boolean
func Create(category entities.Category) bool {
	result, err := config.DB.Exec("INSERT INTO categories (name, created_at, updated_at) VALUE (?,?,?)",
		category.Name, category.CreateAt, category.UpdateAt,
	)

	if err != nil {
		panic(err)
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	// karena nila yang di kembalikan adalah bool maka kita menggunakkan logika lebih dari 0
	// karena jika true nilainya adalah 1
	// dan jika false nilainya adalah 0
	return lastInsertId > 0
}

func Detail(id int) entities.Category {
	// QueryRow di gunakkan untuk mengambil salah satu data
	row := config.DB.QueryRow("SELECT id, name FROM categories WHERE id=?", id)

	var category entities.Category
	if err := row.Scan(&category.Id, &category.Name); err != nil {
		panic(err.Error())
	}

	return category
}

func Update(id int, category entities.Category) bool {
	// jangan lupa cek kode yang ingin di up ke database
	// nama ataupun semua kode yang akan di gunakkan untuk merubah mysql jangan sampai salah, itu membuat eror karena tidak bisa save data
	query, err := config.DB.Exec("UPDATE categories SET name = ? , updated_at = ? WHERE id = ?",
	category.Name, category.UpdateAt, id)
	if err != nil {
		panic(err)
	}

	// digunakkan untuk cek, apakah data sudah benar" ter update apa belum
	result, err := query.RowsAffected() 
		if err != nil {
			panic(err)
		}

		return result > 0

	
}


func Delete(id int) error {
	_, err := config.DB.Exec("DELETE FROM categories WHERE id = ?", id)
	return err
}