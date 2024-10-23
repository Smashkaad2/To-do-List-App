package configs

import (
	"log"
	"gin-proyect/models" // Asegúrate de que la ruta sea correcta
)

// SeedDB inicializa valores en la base de datos
func SeedDB() {

	DB.Exec("SET FOREIGN_KEY_CHECKS = 0")

	// Truncar las tablas
	if err := DB.Exec("TRUNCATE TABLE tasks").Error; err != nil {
		log.Fatalf("Failed to truncate tasks table: %v", err)
	}
	if err := DB.Exec("TRUNCATE TABLE persons").Error; err != nil {
		log.Fatalf("Failed to truncate persons table: %v", err)
	}

	DB.Exec("SET FOREIGN_KEY_CHECKS = 1")
	log.Println("Tables truncated successfully!")

	// Inicializar datos para Person con Tasks
	persons := []models.Person{
		{
			Name:    "John Doe", 
			Address: "123 Elm St", 
			Phone:   1234567890,
			Tasks: []models.Task{
				{TaskName: "C++ proyect", Status: "COMPLETED"},
				{TaskName: "Go project", Status: "PENDING"},
			},
		},
		{
			Name:    "Jane Smith", 
			Address: "456 Oak St", 
			Phone:   9876543210,
			Tasks: []models.Task{
				{TaskName: "Java project", Status: "PENDING"},
				{TaskName: "Python project", Status: "COMPLETED"},
			},
		},
	}

	// Insertar personas con sus tareas asociadas
	for _, person := range persons {
		if err := DB.Create(&person).Error; err != nil {
			log.Fatalf("Failed to seed database in Table Person with Tasks: %v", err)
		}
	}

	log.Println("Database seeded successfully!")
}
