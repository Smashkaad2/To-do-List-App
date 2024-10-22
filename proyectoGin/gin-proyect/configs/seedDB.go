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
	
	persons := []models.Person{
		{Name: "John Doe", Address: "123 Elm St", Phone: 1234567890},
		{Name: "Jane Smith", Address: "456 Oak St", Phone: 9876543210},
	}

	// Insertar datos iniciales
	for _, person := range persons {
		if err := DB.Create(&person).Error; err != nil {
			log.Fatalf("Failed to seed database: %v", err)
		}
	}
	log.Println("Database seeded successfully!")
}
