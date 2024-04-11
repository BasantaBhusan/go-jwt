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
        "/user/kyc": {
            "post": {
                "description": "Create KYC (Know Your Customer) record.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KYC"
                ],
                "summary": "Create KYC",
                "parameters": [
                    {
                        "description": "KYC details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.CreateKYCRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "KYC created successfully",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    },
                    "400": {
                        "description": "Failed to read body or create KYC",
                        "schema": {
                            "$ref": "#/definitions/gin.H"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Log in a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "User login",
                "parameters": [
                    {
                        "description": "User credentials",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created Successfully",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "Create a new user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create new user",
                "parameters": [
                    {
                        "description": "User information",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User created Successfully",
                        "schema": {
                            "$ref": "#/definitions/controllers.SuccessResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.CreateKYCAddressRequest": {
            "type": "object",
            "required": [
                "district",
                "province",
                "ward_number"
            ],
            "properties": {
                "district": {
                    "type": "string"
                },
                "municipality": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "ward_number": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateKYCRequest": {
            "type": "object",
            "required": [
                "address",
                "full_name",
                "mobile_number",
                "service",
                "working_area"
            ],
            "properties": {
                "address": {
                    "$ref": "#/definitions/controllers.CreateKYCAddressRequest"
                },
                "firm_registered": {
                    "type": "boolean"
                },
                "full_name": {
                    "type": "string"
                },
                "is_kyc": {
                    "type": "boolean"
                },
                "mobile_number": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/controllers.CreateKYCServiceRequest"
                },
                "working_area": {
                    "$ref": "#/definitions/controllers.CreateKYCWorkingAreaRequest"
                }
            }
        },
        "controllers.CreateKYCServiceRequest": {
            "type": "object",
            "required": [
                "service_name"
            ],
            "properties": {
                "service_name": {
                    "type": "string"
                }
            }
        },
        "controllers.CreateKYCWorkingAreaRequest": {
            "type": "object",
            "required": [
                "activities",
                "area_name"
            ],
            "properties": {
                "activities": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "area_name": {
                    "type": "string"
                }
            }
        },
        "controllers.LoginRequest": {
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
                    "type": "string"
                }
            }
        },
        "controllers.SuccessResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.UserRequest": {
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
                    "type": "string"
                }
            }
        },
        "gin.H": {
            "type": "object",
            "additionalProperties": {}
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