package schema

import (
	"fmt"

	"gorm.io/gorm"
)

// type update struct {
// 	as string
// 	to string
// }

func GetAll() []string {
	return []string{
		Security(),
	}
}

func Create(db *gorm.DB) {
	for _, name := range GetAll() {
		query := fmt.Sprintf("IF NOT EXISTS (SELECT * FROM sys.schemas WHERE name = '%s') EXEC('CREATE SCHEMA %s')", name, name)
		db.Exec(query)
	}
}

// func Update() update {
// 	// 	err := db.Exec("ALTER SCHEMA Profile TRANSFER Security.Users").Error
// 	// if err != nil {
// 	//     log.Fatal("Failed to move table to new schema: ", err)
// 	// }
// 	new := update{as: "Security", to: "Profile"}
// 	return new
// }

func Security() string {
	return "Security"
}
