package main

import (
    "github.com/kamva/mgm/v3"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func initDB() {
    // Replace with your MongoDB URI
    err := mgm.SetDefaultConfig(nil, "your-database-name", options.Client().ApplyURI("mongodb+srv://bollzhan4:Z1h2a3n4b5o6l7n8@untitled.3hbsy.mongodb.net/?retryWrites=true&w=majority&appName=untitled"))
    if err != nil {
        panic(err)
    }
}