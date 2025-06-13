package main

import (
    "fmt";
    "os";
    "log";
    "flag"
)

var db_path string = "../db/"

func fetch_value(key string) (string, error){
    path := db_path + key 
    value, err := os.ReadFile(path)
    return string(value), err
}

func delete_key(key string) error{
    path := db_path + key 
    err := os.Remove(path)
    return err
}

func create_kv(key string, value string) error{
    path := db_path + key 
    value_slice := []byte(value)
    err := os.WriteFile(path, value_slice, 0644)
    return err
}

func main(){
    err := os.Mkdir(db_path, 0750)
    if err != nil && !os.IsExist(err){
        log.Fatal(err)
    }
    switch os.Args[1]{
        case "create":
            create_flag := flag.NewFlagSet("create", flag.ExitOnError) 
            key_flag := create_flag.String("k", "", "key to write")
            value_flag := create_flag.String("v", "", "value to write")
            create_flag.Parse(os.Args[2:])
            if *key_flag == "" || *value_flag == ""{
                os.Exit(1)
            }
            err := create_kv(*key_flag, *value_flag)
            if err != nil{
                log.Fatal(err)
            }
        case "fetch":
            fetch_flag := flag.NewFlagSet("fetch", flag.ExitOnError) 
            key_flag := fetch_flag.String("k", "", "key to fetch")
            fetch_flag.Parse(os.Args[2:])
            if *key_flag == ""{
                os.Exit(1)
            }
            val, err := fetch_value(*key_flag)
            fmt.Println(val)
            if err != nil{
                log.Fatal(err)
            }
        default:
            fmt.Println("no command given!")
            os.Exit(1)
    }
}
