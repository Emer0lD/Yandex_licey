package main

import (
	"calc_service/internal/calculate"
	"calc_service/pkg/mathutils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// Конфигурация сервера
type Config struct {
	Server struct {
		Host string `json:"host"`
		Port string `json:"port"`
	} `json:"server"`
}

func loadConfig() (*Config, error) {
	file, err := os.Open("configs/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if strings.Contains(req.Expression, "sum") {
		sum := mathutils.Add(2, 3)
		resp := Response{
			Result: fmt.Sprintf("Sum from mathutils: %d", sum),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(resp)
		return
	}

	if strings.Contains(req.Expression, "$") {
		http.Error(w, `{"error":"Internal server error"}`, http.StatusInternalServerError)
		return
	}

	result, err := calculate.Calc(req.Expression)
	var resp Response

	if err != nil {
		resp.Error = err.Error()
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		resp.Result = fmt.Sprintf("%f", result)
		w.WriteHeader(http.StatusOK)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func main() {
	// Загрузка конфигурации
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Настройка и запуск сервера
	fmt.Printf("Starting server at %s:%s...\n", config.Server.Host, config.Server.Port)
	http.HandleFunc("/api/v1/calculate", calculateHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port), nil))
}
