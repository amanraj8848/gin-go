package database


import(
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB*	gorm.DB

func ConnectDatabase() {
	// Define DSN (Data Source Name)
	dsn := "host=localhost user=admin password=admin dbname=pb_container port=5432 sslmode=disable"
	var err error

	// Open connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// defer DB.Close() // This is correct, but do not use it because we want the database connection to stay open throughout the application lifetime.

	fmt.Println("Database connected successfully!")
}
func CloseDatabase() {
	db, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get sql db from gorm db: %v", err)
	}
	
	if err := db.Close(); err != nil {
		log.Fatalf("Failed to close database connection: %v", err)
	}
}
