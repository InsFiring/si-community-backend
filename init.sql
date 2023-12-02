USE si_community;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
	register_number int auto_increment 	COMMENT '자동생성 등록번호'
	, email 		varchar(100) 		COMMENT '이메일 주소'
	, password		varchar(255)		COMMENT '패스워드'
	, nickname 		varchar(50) 		COMMENT '유저 닉네임'
	, company 		varchar(100)		COMMENT '회사'
	, is_active 	varchar(3)			COMMENT '계정활성화 상태 - 탈퇴 여부'
	, loggedin 		varchar(3)			COMMENT	'로그인 상태'
	, created_at 	datetime			COMMENT '가입 날짜'
	, updated_at 	datetime			COMMENT '정보 수정 날짜'
	, deleted_at 	datetime			COMMENT '삭제 날짜'
	, primary key(register_number)
	, unique(email)
	, unique(nickname)
) COMMENT '회원 정보 테이블'
;

DROP TABLE IF EXISTS articles;

-- 회사 평가 게시판
CREATE TABLE articles (
	article_id		int auto_increment 	COMMENT '자동생성 등록번호'
	, nickname 		varchar(50) 		COMMENT '글쓴이 닉네임'
	, company 		varchar(50) 		COMMENT '평가 대상 회사'
	, ratings 		int			 		COMMENT '회사 평가 점수'
	, title			varchar(255)		COMMENT '글 제목'
	, contents 		TEXT				COMMENT '글 내용'
	, view_counts 	int					COMMENT '글 조회수'
	, likes		 	int					COMMENT '글 좋아요'
	, unlikes		int					COMMENT '글 싫어요'
	, is_modified 	varchar(3)			COMMENT '글 수정 여부'
	, is_deleted	varchar(3)			COMMENT '글 삭제 여부'
	, created_at 	datetime			COMMENT '글 생성 날짜'
	, modified_at 	datetime			COMMENT '글 수정 날짜'
	, primary key(article_id)
) COMMENT '게시글 테이블'
;

DROP TABLE IF EXISTS article_replies;

CREATE TABLE article_replies (
	reply_id		int auto_increment 	COMMENT '자동생성 등록번호'
	, article_id	int					COMMENT '글 번호'
	, nickname 		varchar(50) 		COMMENT '댓글쓴이 닉네임'
	, contents 		TEXT				COMMENT '댓글 내용'
	, likes		 	int					COMMENT '댓글 좋아요'
	, unlikes		int					COMMENT '댓글 싫어요'
	, is_modified 	varchar(3)			COMMENT '댓글 수정 여부'
	, created_at 	datetime			COMMENT '댓글 생성 날짜'
	, modified_at 	datetime			COMMENT '댓글 수정 날짜'
	, primary key(reply_id)
) COMMENT '게시글 댓글 테이블'
; 

DROP TABLE IF EXISTS user_tokens;

CREATE TABLE user_tokens (
    token_id bigint auto_increment			COMMENT '자동생성 토큰번호'
    , register_number int					COMMENT '유저 등록 번호(pk)'
    , token TEXT NOT NULL					COMMENT '토큰'
    , expiration_time TIMESTAMP NOT NULL	COMMENT '토큰 만료 기간'
    , created_at 	datetime				COMMENT '토큰 생성 날짜'
	, updated_at 	datetime				COMMENT '토큰 수정 날짜'
	, deleted_at 	datetime				COMMENT '토큰 삭제 날짜'
    , primary key(token_id)
);
