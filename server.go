package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

type House struct {
	Size     float64
	Bedrooms float64
	Age      float64
	Location string
}

type PredictionRequest struct {
	House House
}

type PredictionResponse struct {
	Price float64
}

func calculatePrice(h House) float64 {
	// Calcular el precio base sin aplicar la variabilidad
	basePrice := 1200*h.Size + 500*h.Bedrooms
	ageMultiplier := 1 + 0.08*h.Age
	locationMultiplier := 1.0

	switch h.Location {
	case "A":
		locationMultiplier = 1.06
	case "B":
		locationMultiplier = 1.02
	case "D":
		locationMultiplier = 0.98
	}

	// Introducir variabilidad en el número de habitaciones y antigüedad
	rand.Seed(time.Now().UnixNano())

	// Variabilidad para el número de habitaciones (±1 o ±2 habitaciones)
	bedroomsVariability := float64(rand.Intn(3) - 1) // Entre -1 y 1 habitaciones
	bedroomsAdjusted := h.Bedrooms + bedroomsVariability

	// Variabilidad para la antigüedad (±1 o ±2 años)
	ageVariability := float64(rand.Intn(3) - 1) // Entre -1 y 1 años
	ageAdjusted := h.Age + ageVariability

	// Calcular el precio final aplicando la variabilidad
	finalPrice := basePrice * ageMultiplier * locationMultiplier

	// Utilizar las variables bedroomsAdjusted y ageAdjusted en el cálculo final
	finalPriceAdjusted := finalPrice + bedroomsAdjusted + ageAdjusted

	return finalPriceAdjusted
}

func calculateMAE(data [][]string) float64 {
	var totalError float64
	var count int

	for _, row := range data {
		size, _ := strconv.ParseFloat(row[0], 64)
		bedrooms, _ := strconv.ParseFloat(row[1], 64)
		age, _ := strconv.ParseFloat(row[2], 64)
		location := row[3]
		actualPrice, _ := strconv.ParseFloat(row[4], 64)

		predictedPrice := calculatePrice(House{Size: size, Bedrooms: bedrooms, Age: age, Location: location})
		error := math.Abs(predictedPrice - actualPrice)
		totalError += error
		count++
	}

	return totalError / float64(count)
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	decoder := json.NewDecoder(conn)
	var request PredictionRequest
	if err := decoder.Decode(&request); err != nil {
		log.Println("Error decodificando la solicitud:", err)
		return
	}

	price := calculatePrice(request.House)
	fmt.Printf("Predicción para casa con características %v: %.2f dólares\n", request.House, price)

	response := PredictionResponse{Price: price}
	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&response); err != nil {
		log.Println("Error codificando la respuesta:", err)
		return
	}
}

func main() {
	// Leer el dataset para calcular la precisión
	file, err := os.Open("house_prices.csv")
	if err != nil {
		log.Fatal("No se puede abrir el archivo:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error leyendo el archivo CSV:", err)
	}

	// Calcular la precisión
	mae := calculateMAE(data[1:]) // Excluir encabezado
	fmt.Printf("Error absoluto medio (MAE) del modelo: %.2f dólares\n", mae)

	// Iniciar el servidor
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	log.Println("El servidor está escuchando en el puerto 8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error aceptando la conexión:", err)
			continue
		}
		go handleConnection(conn)
	}
}
