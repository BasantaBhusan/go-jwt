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
        "/search": {
            "get": {
                "description": "Search for users by email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "Perform a search",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users matching the search query",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.User"
                            }
                        }
                    }
                }
            }
        },
        "/search/address": {
            "get": {
                "description": "Search based on the address model and return associated working area, activities, and services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "Perform a search based on the address model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "query",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search result",
                        "schema": {
                            "$ref": "#/definitions/controllers.SearchResult"
                        }
                    }
                }
            }
        },
        "/search/advanced": {
            "get": {
                "description": "Perform an advanced global search across all models based on the provided query string",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "Perform an advanced global search across all models",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of results matching the advanced search query",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Kyc"
                            }
                        }
                    }
                }
            }
        },
        "/search/all": {
            "get": {
                "description": "Perform a global search across all models",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "Perform a global search across all models",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/search/all/address/{province}/{district}/{municipality}/{ward_number}": {
            "get": {
                "description": "Search based on the address model and return associated working area, activities, and services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Search"
                ],
                "summary": "Perform a search based on the address model",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Province",
                        "name": "province",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "District",
                        "name": "district",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Municipality",
                        "name": "municipality",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Ward Number",
                        "name": "ward_number",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Search result",
                        "schema": {
                            "$ref": "#/definitions/controllers.SearchResult"
                        }
                    }
                }
            }
        },
        "/user/all": {
            "get": {
                "description": "Retrieve all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controllers.UserResponse"
                            }
                        }
                    }
                }
            }
        },
        "/user/kyc/create": {
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
                        "description": "KYC created successfully"
                    },
                    "400": {
                        "description": "Failed to read body or create KYC"
                    }
                }
            }
        },
        "/user/kyc/update/{id}": {
            "put": {
                "description": "Update KYC (Know Your Customer) record by User ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KYC"
                ],
                "summary": "Update KYC by User ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "KYC details",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.UpdateKYCRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "KYC updated successfully"
                    },
                    "400": {
                        "description": "Invalid user ID or failed to read body"
                    },
                    "404": {
                        "description": "KYC not found for the given user ID"
                    }
                }
            }
        },
        "/user/kyc/{id}": {
            "get": {
                "description": "Retrieve KYC (Know Your Customer) record by User ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "KYC"
                ],
                "summary": "Get KYC by User ID",
                "parameters": [
                    {
                        "type": "integer",
                        "format": "int64",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "KYC information",
                        "schema": {
                            "$ref": "#/definitions/models.Kyc"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID"
                    },
                    "404": {
                        "description": "KYC not found for the given user ID"
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
        "/user/logout": {
            "get": {
                "description": "Clear Cookie.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Logout user",
                "responses": {
                    "200": {
                        "description": "Sucessfully logged out."
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
        },
        "/user/validate": {
            "get": {
                "description": "Validate User.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Validate user",
                "responses": {
                    "200": {
                        "description": "Ok"
                    },
                    "401": {
                        "description": "Unauthorized"
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
                "latitude": {
                    "type": "string"
                },
                "longitude": {
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
        "controllers.SearchResult": {
            "type": "object",
            "properties": {
                "activities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Activity"
                    }
                },
                "address": {
                    "$ref": "#/definitions/models.Address"
                },
                "associated_services": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Service"
                    }
                },
                "working_area": {
                    "$ref": "#/definitions/models.WorkingArea"
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
        "controllers.UpdateKYCAddressRequest": {
            "type": "object",
            "properties": {
                "district": {
                    "type": "string"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
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
        "controllers.UpdateKYCRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/controllers.UpdateKYCAddressRequest"
                },
                "firm_registered": {
                    "type": "boolean"
                },
                "full_name": {
                    "type": "string"
                },
                "mobile_number": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/controllers.UpdateKYCServiceRequest"
                },
                "working_area": {
                    "$ref": "#/definitions/controllers.UpdateKYCWorkingAreaRequest"
                }
            }
        },
        "controllers.UpdateKYCServiceRequest": {
            "type": "object",
            "properties": {
                "service_name": {
                    "type": "string"
                }
            }
        },
        "controllers.UpdateKYCWorkingAreaRequest": {
            "type": "object",
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
        "controllers.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                }
            }
        },
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
        "models.Activity": {
            "type": "object",
            "properties": {
                "activityName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "kycID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "workingAreaID": {
                    "type": "integer"
                }
            }
        },
        "models.Address": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "district": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "kycID": {
                    "type": "integer"
                },
                "latitude": {
                    "type": "string"
                },
                "longitude": {
                    "type": "string"
                },
                "municipality": {
                    "type": "string"
                },
                "province": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "wardNumber": {
                    "type": "string"
                }
            }
        },
        "models.InvestmentOption": {
            "type": "string",
            "enum": [
                "up to 5 Lakhs",
                "up to 10 Lakhs",
                "up to 25 Lakhs",
                "up to 50 Lakhs",
                "up to 1 Crore",
                "above 1 Crore"
            ],
            "x-enum-varnames": [
                "UpTo5LAKHS",
                "UpTo10LAKHS",
                "UpTo25LAKHS",
                "UpTo50LAKHS",
                "UpTo1CRORE",
                "Above1CRORE"
            ]
        },
        "models.Kyc": {
            "type": "object",
            "properties": {
                "address": {
                    "$ref": "#/definitions/models.Address"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "firmRegistered": {
                    "type": "boolean"
                },
                "fullName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "mobileNumber": {
                    "type": "string"
                },
                "service": {
                    "$ref": "#/definitions/models.Service"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "workingArea": {
                    "$ref": "#/definitions/models.WorkingArea"
                }
            }
        },
        "models.Service": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "investment": {
                    "$ref": "#/definitions/models.InvestmentOption"
                },
                "kycID": {
                    "type": "integer"
                },
                "serviceName": {
                    "$ref": "#/definitions/models.ServiceType"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                }
            }
        },
        "models.ServiceType": {
            "type": "string",
            "enum": [
                "Expert Advice",
                "Business Partnership",
                "Bank Loan Facilitation",
                "Training and Coaching",
                "Cold Store Construction",
                "Assistance in Marketing",
                "Investment"
            ],
            "x-enum-varnames": [
                "ExpertAdvice",
                "BusinessPartnership",
                "BankLoanFacilitation",
                "TrainingAndCoaching",
                "ColdStoreConstruction",
                "AssistanceInMarketing",
                "InvestmentService"
            ]
        },
        "models.User": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
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
                "is_kyc": {
                    "type": "boolean"
                },
                "kyc": {
                    "$ref": "#/definitions/models.Kyc"
                },
                "password": {
                    "type": "string"
                },
                "profile_picture": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "models.WorkingArea": {
            "type": "object",
            "properties": {
                "activities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Activity"
                    }
                },
                "areaName": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "$ref": "#/definitions/gorm.DeletedAt"
                },
                "id": {
                    "type": "integer"
                },
                "kycID": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
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
