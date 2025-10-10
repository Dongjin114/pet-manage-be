#!/bin/bash

# WSL에 프로젝트 복사 스크립트
echo "Copying project to WSL..."

# WSL 홈 디렉토리에 프로젝트 폴더 생성
wsl mkdir -p ~/pet-manage-be

# 현재 디렉토리의 모든 파일을 WSL로 복사
wsl cp -r . ~/pet-manage-be/

echo "Project copied to WSL at ~/pet-manage-be"
echo "Now you can connect to WSL and run:"
echo "cd ~/pet-manage-be"
echo "go mod tidy"
echo "go run main.go"


