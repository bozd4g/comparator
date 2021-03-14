// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "me@furkanbozdag.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "description": "This method redirects to swagger ui",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "IndexController"
                ],
                "summary": "redirectToSwaggerUi",
                "responses": {
                    "308": {
                        "description": "Redirect",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/configs": {
            "get": {
                "description": "This method returns all configs",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configs"
                ],
                "summary": "Get all configs",
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/configs.Dto"
                            }
                        }
                    }
                }
            }
        },
        "/api/configs/{name}": {
            "get": {
                "description": "This method returns spesific config",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Configs"
                ],
                "summary": "Get configs by name",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Config name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "$ref": "#/definitions/configs.Dto"
                        }
                    }
                }
            }
        },
        "/api/products": {
            "post": {
                "description": "This method returns all prices with their total price",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all prices with their total",
                "parameters": [
                    {
                        "description": "Products",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "type": "string"
                            }
                        }
                    },
                    {
                        "enum": [
                            "Books"
                        ],
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/products.MultipleDto"
                            }
                        }
                    }
                }
            }
        },
        "/api/products/{name}": {
            "get": {
                "description": "This method returns all prices",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all prices",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product name",
                        "name": "name",
                        "in": "path",
                        "required": true
                    },
                    {
                        "enum": [
                            "Books"
                        ],
                        "type": "string",
                        "description": "Category",
                        "name": "category",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/products.Dto"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "configs.Dto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "sites": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/configs.SiteDto"
                    }
                }
            }
        },
        "configs.SiteDto": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "steps": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/configs.StepDto"
                    }
                }
            }
        },
        "configs.StepDto": {
            "type": "object",
            "properties": {
                "action": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "selector": {
                    "type": "string"
                }
            }
        },
        "products.Dto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "product": {
                    "type": "object",
                    "$ref": "#/definitions/products.ProductDto"
                }
            }
        },
        "products.MultipleDto": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "products": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/products.ProductDto"
                    }
                },
                "total": {
                    "type": "number"
                }
            }
        },
        "products.ProductDto": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Comparator API",
	Description: "This is a price comparator application.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
