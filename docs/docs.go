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
        "/token/balanceToken": {
            "get": {
                "description": "do BalanceToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "BalanceToken",
                "parameters": [
                    {
                        "type": "string",
                        "example": "0x46C65D87bE47255882561bcc7CFf3bBA186F0848",
                        "description": "Account Address",
                        "name": "address",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ResponseObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/token/transactionReceipt": {
            "get": {
                "description": "do TransactionReceipt",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TransactionReceipt",
                "parameters": [
                    {
                        "type": "string",
                        "example": "0x81f05c7087dc3fc1ff6576f6303694b803122366e5fba54afd1f526a2f567d56",
                        "description": "Hash Address",
                        "name": "hash",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ResponseObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/token/transferToken": {
            "post": {
                "description": "do TransferToken",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TransferToken",
                "parameters": [
                    {
                        "description": "To Address, From Address, Value Token",
                        "name": "RequestParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RequestParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ResponseObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/token/transferTokenCancel": {
            "post": {
                "description": "do TransferTokenCancel",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TransferTokenCancel",
                "parameters": [
                    {
                        "description": "To Address, From Address, Value Token",
                        "name": "RequestParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RequestParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ResponseObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        },
        "/token/transferTokenPolling": {
            "post": {
                "description": "do TransferTokenPolling",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "TransferTokenPolling",
                "parameters": [
                    {
                        "description": "To Address, From Address, Value Token",
                        "name": "RequestParam",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.RequestParam"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/config.ResponseObject"
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    }
                }
            }
        }
    },
    "definitions": {
        "big.Int": {
            "type": "object"
        },
        "config.ResponseObject": {
            "type": "object",
            "properties": {
                "body": {},
                "isSuccess": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "statusCode": {
                    "type": "integer"
                },
                "timeStamp": {
                    "type": "string"
                }
            }
        },
        "controller.RequestParam": {
            "type": "object",
            "properties": {
                "From": {
                    "type": "string",
                    "example": "0x46C65D87bE47255882561bcc7CFf3bBA186F0848"
                },
                "To": {
                    "type": "string",
                    "example": "0x5EFC0751759b01759BEc45a06930A2d338a5E234"
                },
                "Value": {
                    "$ref": "#/definitions/big.Int"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample mange-template server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}