package test

import (
	"Gone/app/models"
	"go_service/models"
	"testing"
)

func TestPost(t *testing.T) {
	var dataInsertPost = []models.Post{
		models.Post{
			Slug:   "konten-uhuy",
			Title:  "Konten uhuy",
			Content: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
			Image: "Null",
		},
	}

	// db, err := database.ConnectApi()

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// t.Run("Testing insert mahasiswa", func(t *testing.T) {
	// 	for _, dataInsert := range dataInsertPost {
	// 		err := dataInsert.Insert(db)
	// 		if err != nil {
	// 			t.Fatal(err)
	// 		}
	// 	}
	// })

	// t.Run("Testing update  mahasiswa", func(t *testing.T) {
	// 	var updateData = map[string]interface{}{
	// 		"nama": "Abdi Teh Ayuna",
	// 	}

	// 	data := dataInsertMhs[0]

	// 	if err := data.Update(db, updateData); err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("Testing delete  mahasiswa", func(t *testing.T) {
	// 	data := dataInsertMhs[0]

	// 	if err := data.Delete(db); err != nil {
	// 		t.Fatal(err)
	// 	}
	// })

	// t.Run("Testing get  mahasiswa", func(t *testing.T) {
	// 	value, err := model.GetMahasiswa(db, "4444444")

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	fmt.Printf("%v", value)
	// })

	// t.Run("Testing get all mahasiswa", func(t *testing.T) {
	// 	value, err := model.GetAllMahasiswa(db)

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	fmt.Printf("%v", value)
	// })
}