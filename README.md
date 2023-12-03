# si-community-backend
SI 커뮤니티 백엔드 저장소 - go로 개발

## 실행 환경
* docker-compose에 mysql db를 구동하여 로컬 테스트 가능
* init.sql 코드 내에 정의된 테이블들을 생성하며 필요에 따라 예제에 필요한 레코드를 생성하여 미리 저장할 수 있음

## 실행 방법
```
$ git clone https://github.com/InsFiring/si-community-backend.git
$ cd si-community-backend
$ docker-compose up -d
$ go run main.go
```

## API
