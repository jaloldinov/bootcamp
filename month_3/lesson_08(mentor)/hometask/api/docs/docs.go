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
        "/login": {
            "post": {
                "description": "api for create persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "create person",
                "parameters": [
                    {
                        "description": "data of person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            }
        },
        "/person": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "List persons",
                "parameters": [
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "limit for response",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "page of req",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "enum": [
                            "dev",
                            "backend",
                            "frontend"
                        ],
                        "type": "string",
                        "description": "filter by job",
                        "name": "job",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Person"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "api for create persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "create person",
                "parameters": [
                    {
                        "description": "data of person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            }
        },
        "/person/{id}": {
            "get": {
                "description": "get persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "List persons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            },
            "put": {
                "description": "api for update persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "update person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data of person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "api for create persons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "create person",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of person",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "data of person",
                        "name": "person",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreatePerson"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.ErrorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.CreatePerson": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "job": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LoginReq": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.LoginRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.Person": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "branch_id": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "job": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "response.CreateResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "response.ErrorResp": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
