package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
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

func main() {

	url := "https://raw.githubusercontent.com/parzzd/ta4/main/house_price_8.csv?token=GHSAT0AAAAAACTRFWPRZ4BPMJ6ASHCND45WZTLQDSA"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error al obtener el archivo CSV:", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: estado HTTP %d\n", resp.StatusCode)
	}

	reader := csv.NewReader(resp.Body)

	header, err := reader.Read()
	if err != nil {
		log.Fatal("Error al leer la cabecera del CSV:", err)
	}

	fmt.Println("Cabecera del archivo CSV:")
	fmt.Println(header)

	row, err := reader.Read()
	if err != nil {
		log.Fatal("Error al leer una fila del CSV:", err)
	}

	size, _ := strconv.ParseFloat(row[0], 64)
	bedrooms, _ := strconv.ParseFloat(row[1], 64)
	age, _ := strconv.ParseFloat(row[2], 64)
	location := row[3]

	request := PredictionRequest{
		House: House{
			Size:     size,
			Bedrooms: bedrooms,
			Age:      age,
			Location: location,
		},
	}

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("Error al conectarse al servidor:", err)
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	if err := encoder.Encode(&request); err != nil {
		log.Fatal("Error codificando la solicitud:", err)
	}

	var response PredictionResponse
	decoder := json.NewDecoder(conn)
	if err := decoder.Decode(&response); err != nil {
		log.Fatal("Error decodificando la respuesta:", err)
	}

	fmt.Printf("El precio estimado de la casa es: %.2f dólares\n", response.Price)
}
