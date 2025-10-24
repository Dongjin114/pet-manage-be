package main

import (
	"log"
	"pet-manage-be/internal/infrastructure/di"
)

func main() {
	// 의존성 주입 컨테이너 생성
	container := di.NewContainer()

	// 서버 시작
	port := ":" + container.Config.Server.Port
	log.Printf("서버가 포트 %s에서 시작됩니다", port)

	if err := container.Router.Run(port); err != nil {
		log.Fatal("서버 시작 실패:", err)
	}
}
