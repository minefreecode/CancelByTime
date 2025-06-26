package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("=== Отмена со временем ===")

	//Пример 1.
	ctx, cancel := context.WithTimeout(context.Background(), 7*time.Second)
	defer cancel()

	result := make(chan string, 1)
	go longRunningTask(ctx, result)

	select {
	case <-result:
		fmt.Printf("Таск завершён")
	case <-ctx.Done():
		fmt.Printf("Таск выполняется слишком долго: %v", ctx.Err())
	}
}

func longRunningTask(ctx context.Context, result chan<- string) {
	workTime := time.Duration(rand.Intn(3000)+1000) * time.Millisecond
	fmt.Printf("Старт длинного таска %v...\n", workTime)

	select {
	case <-time.After(workTime):
		result <- "Длинный таск завершён успешно"
	case <-ctx.Done():
		fmt.Printf("Длинный таск отменён: %v\n", ctx.Err())
		return
	}
}
