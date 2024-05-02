package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type DataPoint struct {
	Size      float64
	Rooms     int
	Age       int
	Price     float64
	Predicted float64
	Residual  float64
}

func generateData(seed int64, size int) []DataPoint {
	data := make([]DataPoint, size)
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < size; i++ {
		size := r.Float64()*(180-70) + 70
		price := r.Float64()*(150000-90000) + 90000
		data[i] = DataPoint{
			Size:  size,
			Rooms: r.Intn(5) + 3,
			Age:   r.Intn(15) + 1,
			Price: price,
		}
	}
	return data
}

func linearRegression(data []DataPoint) (float64, float64) {
	var sumX, sumY, sumXY, sumXSquare float64
	for _, point := range data {
		sumX += point.Size
		sumY += point.Price
		sumXY += point.Size * point.Price
		sumXSquare += point.Size * point.Size
	}
	n := float64(len(data))
	slope := (n*sumXY - sumX*sumY) / (n*sumXSquare - sumX*sumX)
	intercept := (sumY - slope*sumX) / n
	return slope, intercept
}

func predictPrice(size float64, slope float64, intercept float64) float64 {
	return slope*size + intercept
}

func main() {
	const (
		tests    = 1000
		dataSize = 1000000
		trimPct  = 0.05
	)
	seed := time.Now().UnixNano()
	start := time.Now()

	var totalPredictedPrice, totalSize float64
	predictedPrices := make([]float64, tests)

	for i := 0; i < tests; i++ {
		data := generateData(seed+int64(i), dataSize)
		slope, intercept := linearRegression(data)

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		size := r.Float64()*(120-70) + 70
		predictedPrice := predictPrice(size, slope, intercept)
		predictedPrices[i] = predictedPrice
		totalPredictedPrice += predictedPrice
		totalSize += size
	}

	sort.Float64s(predictedPrices)
	trimIndex := int(float64(tests) * trimPct)
	trimmedPrices := predictedPrices[trimIndex : tests-trimIndex]

	var sumTrimmed float64
	for _, price := range trimmedPrices {
		sumTrimmed += price
	}
	trimmedMean := sumTrimmed / float64(len(trimmedPrices))

	averagePredictedPrice := totalPredictedPrice / float64(tests)
	averageSize := totalSize / float64(tests)
	fmt.Printf("Tamaño promedio predicho: %.2f metros cuadrados\n", averageSize)
	fmt.Printf("Precio promedio predicho: %.2f\n", averagePredictedPrice)
	fmt.Printf("Media recortada del precio (media recortada al %d%%): %.2f\n", int(trimPct*100), trimmedMean)
	fmt.Printf("Tiempo total de ejecución: %s\n", time.Since(start))
}
