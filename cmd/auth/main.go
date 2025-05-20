package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"final/internal/app/final/v1"
	"final/internal/config"
	"final/internal/repository/postgres"
	"final/internal/service"
	"final/internal/utils/observability/log"
)

func main() {
	// Инициализируем логгер
	logger := log.NewLogger(log.LevelDebug)

	// Контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Загружаем конфиг
	cfg, err := config.Load()
	if err != nil {
		logger.Error("ошибка загрузки конфига", err)
		os.Exit(1)
	}

	// Инициализируем базу данных
	db := postgres.NewPostgres(ctx, logger, cfg.Postgres)
	if db == nil {
		logger.Error("не удалось инициализировать базу данных")
		os.Exit(1)
	}
	defer db.Close()

	// Инициализируем сервис
	svc := service.NewService(logger, db)

	// Создаём gRPC-сервер
	server := final.NewServer(cfg, logger, svc)

	// Запуск сервера в отдельной горутине
	go func() {
		if err := server.Listen(); err != nil {
			logger.Error("ошибка запуска gRPC-сервера", err)
			os.Exit(1)
		}
	}()

	logger.Info("сервер запущен")

	// Ожидаем завершения по сигналу
	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, syscall.SIGINT, syscall.SIGTERM)
	<-quitCh

	// Грейсфул-шатдаун
	time.Sleep(100 * time.Millisecond)

	if err := server.Stop(ctx); err != nil {
		logger.Error("ошибка при остановке сервера", err)
	}

	logger.Info("приложение завершено")
}
