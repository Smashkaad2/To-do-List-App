package configs

import (
	"log"
	"gin-proyect/models" // Asegúrate de que la ruta sea correcta
)

// SeedDB inicializa valores en la base de datos
func SeedDB() {

	err := DB.Exec("TRUNCATE TABLE persons").Error
	if err != nil {
		log.Fatalf("Failed to truncate persons table: %v", err)
	}

	err = DB.Exec("TRUNCATE TABLE tasks").Error
	if err != nil {
		log.Fatalf("Failed to truncate tasks table: %v", err)
	}
	
	persons := []models.Person{
		{Name: "John Doe", Address: "123 Elm St", Phone: 1234567890},
		{Name: "Jane Smith", Address: "456 Oak St", Phone: 9876543210},
	}

	for _, person := range persons {
		if err := DB.Create(&person).Error; err != nil {
			log.Fatalf("Failed to seed database in Table Person: %v", err)
		}
	}

	tasks := []models.Task{
		{TaskName: "C++ proyect", Status: "COMPLETED"},
		{TaskName: "Java proyect", Status: "PENDING"},
	}

	for _, task := range tasks {
		if err := DB.Create(&task).Error; err != nil {
			log.Fatalf("Failed to seed database in table Task: %v", err)
		}
	}

	log.Println("Database seeded successfully!")
}
