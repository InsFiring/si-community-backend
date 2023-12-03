// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/users": {
            "post": {
                "description": "회원 가입",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.Users"
                        }
                    }
                }
            }
        },
        "/v1/users/changePassword": {
            "post": {
                "description": "비밀번호 변경",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponseDto"
                        }
                    }
                }
            }
        },
        "/v1/users/signin": {
            "post": {
                "description": "로그인",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/models.UserResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gorm.DeletedAt": {
            "type": "object",
            "properties": {
                "time": {
                    "type": "string"
                },
                "valid": {
                    "description": "Valid is true if Time is not NULL",
                    "type": "boolean"
                }
            }
        },
        "models.UserResponseDto": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "expiration_time": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "register_number": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "token_id": {
                    "type": "integer"
                }
            }
        },
        "models.Users": {
            "type": "object",
            "properties": {
                "company": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "is_active": {
                    "type": "string"
                },
                "loggedin": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "register_number": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Your Gin API",
	Description:      "This is a sample Gin API with Swagger documentation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}