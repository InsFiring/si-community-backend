# si-community-backend
SI 커뮤니티 백엔드 저장소 - go로 개발

## 실행 환경
* docker-compose에 mysql db를 구동하여 로컬 테스트 가능
* init.sql 코드 내에 정의된 테이블들을 생성하며 필요에 따라 예제에 필요한 레코드를 생성하여 미리 저장할 수 있음

## 실행 방법
* mysql docker 실행
```
$ git clone https://github.com/InsFiring/si-community-backend.git
$ cd si-community-backend
$ docker-compose up -d
```
* go언어로 빌드된 backend 실행
```
$ bin/si-community-backend -config ./config/configuration_test.toml
```

## API
구현한 API 리스트

### 회원 관련 API
|Method|URL|설명|
|---|---|---|
|POST|http://127.0.0.1:8000/v1/users|회원 가입|
|POST|http://127.0.0.1:8000/v1/users/signin|로그인|
|POST|http://127.0.0.1:8000/v1/users/changePassword|비밀번호 변경|

### 게시글 관련 API
|Method|URL|설명|
|---|---|---|
|POST|http://127.0.0.1:8000/v1/article|게시글 등록|
|GET|http://127.0.0.1:8000/v1/article|게시글 리스트 조회|
|GET|http://127.0.0.1:8000/v1/article/:id|단일 게시글 조회|
|PUT|http://127.0.0.1:8000/v1/article/:id|단일 게시글 수정|
|GET|http://127.0.0.1:8000/v1/article/:id/like|게시글 좋아요 추가|
|GET|http://127.0.0.1:8000/v1/article/:id/cancel_like|게시글 좋아요 취소|
|GET|http://127.0.0.1:8000/v1/article/:id/unlike|게시글 싫어요 추가|
|GET|http://127.0.0.1:8000/v1/article/:id/cancel_unlike|게시글 싫어요 취소|
|DELETE|http://127.0.0.1:8000/v1/article/:id|게시글 삭제|

### 게시글 댓글 관련 API
|Method|URL|설명|
|---|---|---|
|POST|http://127.0.0.1:8000/v1/article_reply|게시글 등록|
|GET|http://127.0.0.1:8000/v1/article/:id/article_replies|게시글 리스트 조회|
|PUT|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id|단일 게시글 수정|
|GET|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id/like|게시글 좋아요 추가|
|GET|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id/cancel_like|게시글 좋아요 취소|
|GET|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id/unlike|게시글 싫어요 추가|
|GET|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id/cancel_unlike|게시글 싫어요 취소|
|DELETE|http://127.0.0.1:8000/v1/article/:id/article_replies/:reply_id|게시글 삭제|


* [swagger](http://localhost:8000/swagger/index.html) 링크에서 일부 확인 가능