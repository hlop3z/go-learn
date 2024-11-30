# GORM - CRUD(S) Operations

```go
package main

import (
 "fmt"
 "time"

 "gorm.io/driver/sqlite"
 "gorm.io/gorm"
 _ "modernc.org/sqlite" // Import the modernc driver
)

/*
type Model struct {
 ID        uint `gorm:"primarykey"`
 CreatedAt time.Time
 UpdatedAt time.Time
 DeletedAt time.Time `gorm:"index" default:"NULL"`
}
*/

// Base model with common fields
type Product struct {
 gorm.Model
 Code  string
 Price uint
}

// Always use UTC
func DateTimeUTC() time.Time {
 return time.Now().UTC()
}

// Error Wrap
func check(err error) {
 if err != nil {
  panic(err)
 }
}

func main() {
 // Initialize the database with the "test.db" file and auto-migrate models
 db, err := InitDB("file:test.db", &Product{})
 check(err)

 // Test Database
 TestDB(db)

}

// InitDB initializes a GORM SQLite database connection
func InitDB(dsn string, models ...interface{}) (*gorm.DB, error) {
 // Open the database connection
 db, err := gorm.Open(sqlite.Dialector{
  DriverName: "sqlite",
  DSN:        dsn,
 }, &gorm.Config{})

 // Check for Errors
 if err != nil {
  return nil, fmt.Errorf("failed to connect to database: %w", err)
 }

 // Migrate the schema
 if len(models) > 0 {
  if err := db.AutoMigrate(models...); err != nil {
   return nil, fmt.Errorf("failed to migrate database schema: %w", err)
  }
 }

 return db, nil
}

// Examples for GORM SQLite Database
// TestDB demonstrates CRUD operations on the "products" table using GORM's db.Table method.
func TestDB(db *gorm.DB) {

 // C
 dbCreate((db))

 // R
 dbDetail((db))

 // U
 time.Sleep(1 * time.Second) // Pause for demonstration purposes
 dbUpdate((db))

 // D
 time.Sleep(1 * time.Second) // Pause for demonstration purposes
 dbDelete((db))

 // S ~ R
 dbSearch((db))

}

// ----------- Create Operation -----------
func dbCreate(db *gorm.DB) {
 // Insert a new product into the "products" table
 db.Table("products").Create(map[string]interface{}{
  "code":       "F42",         // Product code
  "price":      100,           // Product price
  "created_at": DateTimeUTC(), // UTC timestamp for created_at
  "updated_at": DateTimeUTC(), // UTC timestamp for updated_at
 })

}

// ----------- Update Operation -----------
func dbUpdate(db *gorm.DB) {
 // Update the product with ID=1
 db.Table("products").
  Where("id IN ?", []int{1}). // Match product with ID=1
  Updates(map[string]interface{}{
   "code":       "D42",         // Update product code
   "price":      500,           // Update product price
   "updated_at": DateTimeUTC(), // Update timestamp
  })

}

// ----------- Delete Operation -----------
func dbDelete(db *gorm.DB) {
 // Soft-delete the product by setting the "deleted_at" timestamp
 where_ids := db.Where("id IN ?", []int{1})
 db.Table("products").
  Where(where_ids). // Match product with ID=1
  Updates(map[string]interface{}{
   "deleted_at": DateTimeUTC(), // Mark as deleted
  })

}

// ----------- Read Operation -----------
func dbDetail(db *gorm.DB) {
 // Retrieve a product with ID=1, ensuring it's not marked as deleted
 record := map[string]interface{}{}
 db.Table("products").
  Select("*").                 // Select all fields
  Where("id = ?", 1).          // Match product with ID=1
  Where("deleted_at IS NULL"). // Exclude logically deleted records
  Take(&record)

 // Print the retrieved record
 fmt.Println("Retrieved Record:", record)
}

// ----------- List Operation -----------
func dbSearch(db *gorm.DB) {
 // Retrieve all products that are not marked as deleted
 var results []map[string]interface{}
 db.Table("products").
  Select("*").                 // Select all fields
  Where("id > ?", 0).          // Match all products with ID > 0
  Where("deleted_at IS NULL"). // Exclude logically deleted records
  Find(&results)

 // Print the list of active products
 fmt.Println("Active Products:", results)
}
```
