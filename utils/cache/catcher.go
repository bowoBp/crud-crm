package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func main() {
	firsRedis()
}

func chalange() {
	// Membuat koneksi ke Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Alamat dan port Redis server
		Password: "",               // Kata sandi Redis server (jika ada)
		DB:       0,                // Indeks database Redis
	})

	// Menutup koneksi Redis ketika program selesai
	defer client.Close()

	// Menyimpan data dalam Redis
	err := client.Set(context.Background(), "name", "John Doe", 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	// Mengambil data dari Redis
	val, err := client.Get(context.Background(), "name").Result()
	if err == redis.Nil {
		fmt.Println("Data tidak ditemukan dalam Redis")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Data dari Redis:", val)
	}

	// Menghapus data dari Redis
	err = client.Del(context.Background(), "name").Err()
	if err != nil {
		log.Fatal(err)
	}

	// Mengambil data setelah dihapus
	val, err = client.Get(context.Background(), "name").Result()
	if err == redis.Nil {
		fmt.Println("Data tidak ditemukan dalam Redis")
	} else if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Data dari Redis:", val)
	}
}

func firsRedis() {
	// Membuat koneksi ke Redis
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Alamat dan port Redis server
		Password: "",               // Kata sandi Redis server (jika ada)
		DB:       0,                // Indeks database Redis
	})

	// Menutup koneksi Redis ketika program selesai
	defer client.Close()

	// Membuat objek User
	user := User{
		ID:       1,
		Username: "john_doe",
		Email:    "johndoe@example.com",
	}

	// Mengubah objek menjadi bentuk JSON
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}

	// Menyimpan objek di Redis
	err = client.Set(context.Background(), "user:1", jsonData, 0).Err()
	if err != nil {
		log.Fatal(err)
	}

	//menghapus object dari redis
	err = client.Del(context.Background(), "user:1").Err()
	if err != nil {
		log.Fatal(err)
	}

	// Mengambil objek dari Redis
	//data, err := client.Get(context.Background(), "user:1").Bytes()
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Mendekode JSON menjadi objek User
	//var retrievedUser User
	//err = json.Unmarshal(data, &retrievedUser)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// Menampilkan data objek yang diambil dari Redis
	//fmt.Println("ID:", retrievedUser.ID)
	//fmt.Println("Username:", retrievedUser.Username)
	//fmt.Println("Email:", retrievedUser.Email)

	// Memeriksa apakah objek masih ada dalam Redis
	exists, err := client.Exists(context.Background(), "user:1").Result()
	if err != nil {
		log.Fatal(err)
	}

	if exists == 0 {
		fmt.Println("Data sudah dihapus dari Redis")
	} else {
		fmt.Println("Gagal menghapus data dari Redis")
	}
}
