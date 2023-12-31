// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/data/id": {
            "get": {
                "description": "get dataset with specific id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "data"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/data/idrange": {
            "get": {
                "description": "get a range of ids in a given time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "data"
                ],
                "summary": "id range",
                "parameters": [
                    {
                        "type": "string",
                        "description": "start date",
                        "name": "start",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "end date",
                        "name": "end",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Dataset"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorMsg"
                        }
                    }
                }
            }
        },
        "/data/timerange": {
            "get": {
                "description": "get a range of ids in a given time range",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "data"
                ],
                "summary": "get time range",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": ""
                        }
                    }
                }
            }
        },
        "/example/helloworld": {
            "get": {
                "description": "do ping hello world",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "example"
                ],
                "summary": "ping example",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Dataset": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "Actual data\nexample: { \"key\": \"dataset1\" }",
                    "type": "integer"
                },
                "id": {
                    "description": "The ID of a dataaset\nexample: 130670",
                    "type": "string"
                }
            }
        },
        "models.ErrorMsg": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "status": {
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
