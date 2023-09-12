package main

import (
	"testing"
	"log"
)

func TestAdd(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		result := Add(2, 3)
		if result != 5 {
			t.Errorf("Expected 5, but got %d", result)
		} else {
			log.Println("Success: TestAdd for positive numbers passed.")
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		result := Add(-2, -3)
		if result != -5 {
			t.Errorf("Expected -5, but got %d", result)
		} else {
			log.Println("Success: TestAdd for negative numbers passed.")
		}
	})

	t.Run("zero", func(t *testing.T) {
		result := Add(0, 0)
		if result != 0 {
			t.Errorf("Expected 0, but got %d", result)
		} else {
			log.Println("Success: TestAdd for zero passed.")
		}
	})

	t.Run("positive and negative", func(t *testing.T) {
		result := Add(2, -3)
		if result != -1 {
			t.Errorf("Expected -1, but got %d", result)
		} else {
			log.Println("Success: TestAdd for positive and negative passed.")
		}
	})
}

func Add(a, b int) int {
	return a + b
}
