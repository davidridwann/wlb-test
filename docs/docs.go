// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/auth/login": {
            "post": {
                "description": "Login a user account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Login a user account",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userEntity.AuthRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userEntity.UserAccess"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register a new user account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Register a user account",
                "parameters": [
                    {
                        "description": "query params",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userEntity.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userEntity.UserAccess"
                        }
                    }
                }
            }
        },
        "/auth/user": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Responds with the data of user login.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Get user login data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userEntity.UserAccess"
                        }
                    }
                }
            }
        },
        "/auth/verification-account": {
            "post": {
                "description": "Verification user account",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Authentication"
                ],
                "summary": "Verification user account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Verification token",
                        "name": "verif",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/log": {
            "get": {
                "description": "Log activity",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Log"
                ],
                "summary": "Log activity",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Get all post",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/postEntity.PostShow"
                        }
                    }
                }
            }
        },
        "/post/comment": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Comment a post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Comment"
                ],
                "summary": "Comment a post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post Code",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Comment",
                        "name": "comment",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post/comment/reply": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Reply a comment",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reply"
                ],
                "summary": "Reply a comment",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Comment Code",
                        "name": "comment_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Reply",
                        "name": "reply",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post/create": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Create a new post",
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Create a new post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post Caption",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Post Enable/Disable",
                        "name": "is_comment",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Post Image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/postEntity.Post"
                        }
                    }
                }
            }
        },
        "/post/delete": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Soft Delete an existing post",
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Soft Delete an existing post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Existing Post Code",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post/like": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Like a post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Like"
                ],
                "summary": "Like a post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post Code",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post/show": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Show a specific post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Show a specific post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post Code",
                        "name": "code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/postEntity.PostShow"
                        }
                    }
                }
            }
        },
        "/post/unlike": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Unlike a post",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Like"
                ],
                "summary": "Unlike a post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Post Code",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/post/update": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update an existing post",
                "produces": [
                    "multipart/form-data"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Update an existing post",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Existing Post Code",
                        "name": "code",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Post Caption",
                        "name": "caption",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Post Enable/Disable",
                        "name": "is_comment",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Post Image",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/postEntity.Post"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "commentEntity.CommentWithReply": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "post_id": {
                    "type": "string"
                },
                "replies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/replyEntity.Reply"
                    }
                },
                "text": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "likeEntity.Like": {
            "type": "object",
            "required": [
                "user_id"
            ],
            "properties": {
                "user_id": {
                    "type": "string"
                }
            }
        },
        "postEntity.Post": {
            "type": "object",
            "required": [
                "caption"
            ],
            "properties": {
                "caption": {
                    "type": "string"
                },
                "is_comment": {
                    "type": "boolean"
                }
            }
        },
        "postEntity.PostShow": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "comments": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/commentEntity.CommentWithReply"
                    }
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "is_comment": {
                    "type": "boolean"
                },
                "likes": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/likeEntity.Like"
                    }
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "replyEntity.Reply": {
            "type": "object",
            "required": [
                "text"
            ],
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "userEntity.AuthRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                }
            }
        },
        "userEntity.User": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 6
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "userEntity.UserAccess": {
            "type": "object",
            "required": [
                "email",
                "name",
                "username"
            ],
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "emailIsVerified": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "WorkLife&Beyond BackEnd TEST API Docs",
	Description:      "BackEnd TEST API Documentations",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
