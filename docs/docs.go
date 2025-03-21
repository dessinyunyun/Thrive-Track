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
        "/Question": {
            "post": {
                "description": "Create Question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Create Question",
                "parameters": [
                    {
                        "description": "Payload Body for Create [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/question.Form"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/activated-client": {
            "patch": {
                "description": "ActivatedClient with token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "ActivatedClient",
                "parameters": [
                    {
                        "description": "Payload Body for Patch [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.ActivatedTokenForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "Payload Body for Login [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/refresh-token": {
            "patch": {
                "description": "RefreshToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "RefreshToken",
                "parameters": [
                    {
                        "description": "Payload Body for RefreshToken [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RefreshTokenForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "Payload Body for Register [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.RegisterForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/check-connection": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "responses": {}
            }
        },
        "/example": {
            "get": {
                "description": "Get All Example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Get All Example",
                "parameters": [
                    {
                        "type": "array",
                        "description": "ID in Array",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page of pagination",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create Example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Create Example",
                "parameters": [
                    {
                        "description": "Payload Body for Create [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/example.ExampleForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/example/{id}": {
            "get": {
                "description": "Get Detail Example by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Get Detail Example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update Example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Update Example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body for Update [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/example.ExampleForm"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete Example by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Example"
                ],
                "summary": "Delete Example",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/form-response": {
            "get": {
                "description": "Get All Form Response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form Response"
                ],
                "summary": "Get All Form Response",
                "parameters": [
                    {
                        "type": "array",
                        "description": "ID in Array",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by user id",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page of pagination",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create Form Response",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form Response"
                ],
                "summary": "Create Form Response",
                "parameters": [
                    {
                        "description": "Payload Body for Create [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/form_response.Form"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/form-response/{id}": {
            "get": {
                "description": "Get Detail Form Response by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Form Response"
                ],
                "summary": "Get Detail Form Response",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/question": {
            "get": {
                "description": "Get All Question",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get All Question",
                "parameters": [
                    {
                        "type": "array",
                        "description": "ID in Array",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by language",
                        "name": "language",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by category_id",
                        "name": "category_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page of pagination",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {}
            }
        },
        "/question/{language}/{order}": {
            "get": {
                "description": "Get Detail Question by language and order. lang(en/id)",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Question"
                ],
                "summary": "Get Detail Question",
                "parameters": [
                    {
                        "type": "string",
                        "description": "language (eng or id)",
                        "name": "language",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "order",
                        "name": "order",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/user": {
            "get": {
                "description": "Get All Example",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get All Users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer Token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "array",
                        "description": "ID in Array",
                        "name": "id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Username",
                        "name": "username",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by Email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Limit of pagination",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page of pagination",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Payload Body for Create [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserForm"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/user/{id}": {
            "get": {
                "description": "Get Detail User by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Detail User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Payload Body for Update [RAW]",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UserForm"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete User by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "auth.ActivatedTokenForm": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "auth.LoginForm": {
            "type": "object",
            "required": [
                "password"
            ],
            "properties": {
                "email": {
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
        "auth.RefreshTokenForm": {
            "type": "object",
            "required": [
                "refresh_token"
            ],
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "auth.RegisterForm": {
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
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "example.ExampleForm": {
            "type": "object",
            "required": [
                "email",
                "name",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "form_response.Form": {
            "type": "object",
            "properties": {
                "depression_level": {
                    "type": "string"
                },
                "total_score": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "question.Form": {
            "type": "object",
            "properties": {
                "category_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "example": {
                    "type": "string"
                },
                "language": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "user.UserForm": {
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
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
