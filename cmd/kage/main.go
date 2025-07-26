package main

import (
	//i "github.com/thisismeamir/kage/internal/bootstrap"
	//"github.com/thisismeamir/kage/internal/server"
	//"log"
	"fmt"
	"github.com/thisismeamir/kage/internal/bootstrap/database"
	"os"
)

func main() {

	//config := i.LoadConfiguration(i.GetConfigPath())
	//i.SetGlobalConfig(config)
	//serverAddr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	art, _ := os.ReadFile("./util/asci-art")
	fmt.Println(string(art))
	var kageDB = &database.Database{}
	kageDB, err := kageDB.OpenDatabase("./data/kage.sqlite")
	if err != nil {
		fmt.Println("Failed to open database:", err)
		return
	}
	err = kageDB.CreateTable(database.TableSchema{
		TableName: "atoms",
		Schema: []database.ColumnSchema{
			{Name: "id", Type: "INTEGER PRIMARY KEY AUTOINCREMENT"},
			{Name: "name", Type: "TEXT NOT NULL"},
			{Name: "path", Type: "TEXT NOT NULL"},
		},
	})
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}
	fmt.Println("Database opened and table created successfully.")
	fmt.Println(kageDB.GetAllTables())
	a, _ := kageDB.GetTableSchema("atoms")
	for _, col := range a {
		fmt.Println(col.Name)
	}

	kageDB.DropTable("atoms")
	kageDB.CloseDatabase()
	//log.Println("Starting Server in", serverAddr)
	//srv := server.New()
	//if err := srv.Start(serverAddr); err != nil {
	//	log.Fatal("Failed to start server:", err)
	//}
}
