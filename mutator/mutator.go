package mutator

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func MutateConfig() {
    // Read the environment variable
    // DB Password
    db_password := os.Getenv("db_password")
    if db_password == "" {
        fmt.Println("db_password environment variable is not set.")
        return
    }

    // Redis Password
    redis_password := os.Getenv("redis_password")
    if redis_password == "" {
        fmt.Println("redis_password environment variable is not set.")
        return
    }

    // AWS Secret
    aws_secret := os.Getenv("aws_secret")
    if aws_secret == "" {
        fmt.Println("aws_secret environment variable is not set.")
        return
    }

    // Read the JSON file
    file, err := os.Open("config/app.prod.json")
    if err != nil {
        log.Fatalf("Error opening file: %v", err)
    }
    defer file.Close()

    // Decode the JSON data
    byteValue, _ := ioutil.ReadAll(file)
    var config map[string]interface{}
    if err := json.Unmarshal(byteValue, &config); err != nil {
        log.Fatalf("Error decoding JSON: %v", err)
    }

    // Update Database Password value
    if dbUser, ok := config["DBUser"].(map[string]interface{}); ok {
        if write, ok := dbUser["Write"].(map[string]interface{}); ok {
            write["Password"] = db_password
        }
    }
    // Update Redis Password value
    if redis, ok := config["Redis"].(map[string]interface{}); ok {
        redis["Password"] = redis_password
    }
    // Update AWS Secret
    if codepush, ok := config["CodePush"].(map[string]interface{}); ok {
        if aws, ok := codepush["Aws"].(map[string]interface{}); ok {
            aws["Secret"] = aws_secret
        }
    }


    // Encode back to JSON
    updatedJSON, err := json.MarshalIndent(config, "", "    ")
    if err != nil {
        log.Fatalf("Error encoding JSON: %v", err)
    }

    // Write the updated JSON back to the file
    if err := ioutil.WriteFile("config/app.prod.json", updatedJSON, 0644); err != nil {
        log.Fatalf("Error writing to file: %v", err)
    }

    fmt.Println("Credentials written to config/app.prod.json successfully.")
}