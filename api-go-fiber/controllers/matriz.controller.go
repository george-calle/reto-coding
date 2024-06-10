package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type MatrizRequest struct {
	Matriz [][]int `json:"matriz"`
}

type MatrizResponse struct {
	Matriz [][]int `json:"matriz"`
}

type NodeAPIResponse struct {
	Val_valor_maximo    float32 `json:"val_valor_maximo"`
	Val_valor_minimo    float32 `json:"val_valor_minimo"`
	Val_promedio        float32 `json:"val_promedio"`
	Val_suma_total      float32 `json:"val_suma_total"`
	Val_matriz_diagonal float32 `json:"val_matriz_diagonal"`
}

func RotateMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	m := len(matrix[0])
	rotated := make([][]int, m)
	for i := range rotated {
		rotated[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			rotated[j][n-i-1] = matrix[i][j]
		}
	}
	return rotated
}

func PostMatriz(c *fiber.Ctx) error {
	var request MatrizRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if len(request.Matriz) == 0 || len(request.Matriz[0]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "The matrix cannot be empty"})
	}

	cols := len(request.Matriz[0])

	for _, row := range request.Matriz {
		if len(row) != cols {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "The matrix must be rectangular"})
		}
	}

	rotatedMatrix := RotateMatrix(request.Matriz)
	response := MatrizResponse{Matriz: rotatedMatrix}

	//Call external API
	responseAPI, err := notifyNodeAPI(response)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error calling external API"})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"Matriz_Rotada": response, "MatrizAPI_Estadisticas": responseAPI})
}

func notifyNodeAPI(response MatrizResponse) (*NodeAPIResponse, error) {
	fmt.Println("notifyNodeAPI_v2")
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	nodeAPIURL := os.Getenv("NODE_API_URL")
	jsonValue, _ := json.Marshal(response)

	fmt.Println(nodeAPIURL)

	resp, err := http.Post(nodeAPIURL, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("post")
		fmt.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	fmt.Println("resultado API Node")
	fmt.Println(resp.Body)

	var responseAPI NodeAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&responseAPI); err != nil {
		fmt.Println("decode")
		fmt.Println(err)
		return nil, err
	}

	return &responseAPI, nil
}

func PostMatriz_SinAPI(c *fiber.Ctx) error {
	var request MatrizRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	if len(request.Matriz) == 0 || len(request.Matriz[0]) == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "The matrix cannot be empty"})
	}

	cols := len(request.Matriz[0])

	for _, row := range request.Matriz {
		if len(row) != cols {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "The matrix must be rectangular"})
		}
	}

	rotatedMatrix := RotateMatrix(request.Matriz)
	response := MatrizResponse{Matriz: rotatedMatrix}

	return c.Status(http.StatusOK).JSON(response)
}
