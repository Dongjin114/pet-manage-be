# Pet Management API

Go로 구현된 대규모 프로젝트 표준 구조를 적용한 펫 관리 시스템 API입니다.

## 🏗️ 아키텍처

이 프로젝트는 **Clean Architecture**와 **Domain-Driven Design** 패턴을 적용하여 구현되었습니다.

### 📁 프로젝트 구조

```
pet-manage-be/
├── cmd/                    # 애플리케이션 진입점
│   └── server/
│       └── main.go
├── internal/               # 내부 패키지 (외부에서 import 불가)
│   ├── domain/            # 도메인 레이어
│   │   ├── entities/      # 엔티티
│   │   ├── repositories/  # 리포지토리 인터페이스
│   │   └── services/      # 도메인 서비스
│   ├── usecase/           # 유스케이스 레이어
│   │   ├── pet/
│   │   ├── owner/
│   │   └── meal/
│   ├── infrastructure/    # 인프라 레이어
│   │   ├── database/
│   │   ├── repository/
│   │   ├── config/
│   │   └── di/            # 의존성 주입
│   └── interface/         # 인터페이스 레이어
│       ├── http/
│       │   ├── handlers/
│       │   ├── middleware/
│       │   └── routes/
│       └── grpc/
├── pkg/                   # 외부에서 사용 가능한 라이브러리
│   ├── logger/
│   ├── validator/
│   └── utils/
├── api/                  # API 문서
├── migrations/           # DB 마이그레이션
├── docker/              # Docker 설정
└── scripts/             # 스크립트
```

## 🚀 시작하기

### 필수 요구사항

- Go 1.21 이상
- Git

### 설치 및 실행

1. **저장소 클론**
```bash
git clone <repository-url>
cd pet-manage-be
```

2. **의존성 설치**
```bash
go mod tidy
```

3. **서버 실행**
```bash
go run cmd/server/main.go
```

서버는 기본적으로 포트 8080에서 실행됩니다.

## 📚 API 엔드포인트

### 헬스체크
- `GET /health` - 서비스 상태 확인

### 소유자 관리
- `GET /api/v1/owners` - 소유자 목록 조회
- `GET /api/v1/owners/:id` - 특정 소유자 조회
- `POST /api/v1/owners` - 소유자 생성
- `PUT /api/v1/owners/:id` - 소유자 수정
- `DELETE /api/v1/owners/:id` - 소유자 삭제

### 펫 관리
- `GET /api/v1/pets` - 펫 목록 조회
- `GET /api/v1/pets/:id` - 특정 펫 조회
- `POST /api/v1/pets` - 펫 생성
- `PUT /api/v1/pets/:id` - 펫 수정
- `DELETE /api/v1/pets/:id` - 펫 삭제
- `GET /api/v1/pets/owner/:owner_id` - 소유자별 펫 목록 조회

### 급식 관리
- `GET /api/v1/meals/types` - 급식 타입 조회
- `GET /api/v1/meals/items` - 급식 아이템 목록 조회
- `GET /api/v1/meals/items/:id` - 특정 급식 아이템 조회
- `POST /api/v1/meals/items` - 급식 아이템 생성
- `PUT /api/v1/meals/items/:id` - 급식 아이템 수정
- `DELETE /api/v1/meals/items/:id` - 급식 아이템 삭제
- `GET /api/v1/meals/items/type/:type` - 타입별 급식 아이템 조회

## 🏛️ 아키텍처 설명

### 1. Domain Layer (도메인 레이어)
- **Entities**: 비즈니스 엔티티
- **Repositories**: 데이터 접근 인터페이스
- **Services**: 도메인 서비스

### 2. Use Case Layer (유스케이스 레이어)
- 비즈니스 로직 처리
- 도메인 규칙 적용
- 트랜잭션 관리

### 3. Infrastructure Layer (인프라 레이어)
- 데이터베이스 구현
- 외부 서비스 연동
- 설정 관리
- 의존성 주입

### 4. Interface Layer (인터페이스 레이어)
- HTTP 핸들러
- gRPC 서비스
- 미들웨어
- 라우팅

## 🔧 개발

### 새로운 기능 추가

1. **Domain Layer**: 엔티티 및 리포지토리 인터페이스 정의
2. **Use Case Layer**: 비즈니스 로직 구현
3. **Infrastructure Layer**: 데이터베이스 구현체 작성
4. **Interface Layer**: HTTP 핸들러 및 라우트 추가

### 데이터베이스 연동

현재는 목 데이터를 사용하고 있습니다. 실제 데이터베이스를 연동하려면:

1. `internal/infrastructure/repository/` 디렉토리의 구현체 수정
2. 데이터베이스 연결 설정 추가
3. 마이그레이션 스크립트 작성

## 🐳 Docker

```bash
# Docker 이미지 빌드
docker build -t pet-manage-be .

# 컨테이너 실행
docker run -p 8080:8080 pet-manage-be
```

## 📝 라이선스

이 프로젝트는 MIT 라이선스 하에 있습니다.