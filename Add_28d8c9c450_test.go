package main

import (
	"testing"
)

func TestAdd_28d8c9c450(t *testing.T) {
	t.Run("positive numbers", func(t *testing.T) {
		result := Add_28d8c9c450(2, 3)
		if result != 5 {
			t.Errorf("Expected 5, but got %d", result)
		} else {
			t.Log("Success: TestAdd_28d8c9c450 for positive numbers passed.")
		}
	})

	t.Run("negative numbers", func(t *testing.T) {
		result := Add_28d8c9c450(-2, -3)
		if result != -5 {
			t.Errorf("Expected -5, but got %d", result)
		} else {
			t.Log("Success: TestAdd_28d8c9c450 for negative numbers passed.")
		}
	})

	t.Run("zero", func(t *testing.T) {
		result := Add_28d8c9c450(0, 0)
		if result != 0 {
			t.Errorf("Expected 0, but got %d", result)
		} else {
			t.Log("Success: TestAdd_28d8c9c450 for zero passed.")
		}
	})

	t.Run("positive and negative", func(t *testing.T) {
		result := Add_28d8c9c450(2, -3)
		if result != -1 {
			t.Errorf("Expected -1, but got %d", result)
		} else {
			t.Log("Success: TestAdd_28d8c9c450 for positive and negative passed.")
		}
	})
}

func Add_28d8c9c450(a, b int) int {
	return a + b
}
