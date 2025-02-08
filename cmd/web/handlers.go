package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

type Classification struct {
	Number     int      `json:"number"`
	IsPrime    bool     `json:"is_prime"`
	IsPerfect  bool     `json:"is_perfect"`
	Properties []string `json:"properties"`
	DigitSum   int      `json:"digit_sum"`
	Funfact    string   `json:"fun_fact"`
}

func getNumber(r *http.Request) (int, error) {
	numberStr := r.URL.Query().Get("number")
	if numberStr == "" {
		return 0, fmt.Errorf("no number provided")
	}

	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return 0, fmt.Errorf("invalid number provided")
	}

	return number, nil
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func isPerfect(n int) bool {
	sum := 1

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			if i*i != n {
				sum = sum + i + n/i
			} else {
				sum = sum + i
			}
		}
	}
	if sum == n && n != 1 {
		return true
	}
	return false
}

func isArmstrong(n int) bool {
	originalNumber := n
	numDigits := countDigits(n)
	sum := 0

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
		n /= 10
	}

	return sum == originalNumber
}

func countDigits(n int) int {
	count := 0
	for n != 0 {
		n /= 10
		count++
	}
	return count
}

func getProperties(n int) []string {
	properties := []string{}
	if isArmstrong(n) {
		properties = append(properties, "armstrong")
	}
	if n%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	return properties
}

func digitSum(n int) int {
	sum := 0
	for n != 0 {
		digit := n % 10
		sum += digit
		n /= 10
	}

	return sum
}

func getFunfact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d/math", n)

	resp, err := http.Get(url)
	if err != nil {
		return "No fun fact available"
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "No fun fact available"
	}

	return string(body)
}

func classifyNumbers(w http.ResponseWriter, r *http.Request) {
	number, err := getNumber(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"number": "alphabet",
			"error":  true,
		})
		return
	}

	numberClass := Classification{
		Number:     number,
		IsPrime:    isPrime(number),
		IsPerfect:  isPerfect(number),
		Properties: getProperties(number),
		DigitSum:   digitSum(number),
		Funfact:    getFunfact(number),
	}

	//w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(numberClass); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Encoded JSON: %+v\n", numberClass)

}
