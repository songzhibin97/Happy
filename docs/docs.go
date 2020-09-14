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
        "termsOfService": "Open source",
        "contact": {
            "name": "SongZhiBin",
            "email": "718428482@qq.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/CreatePost": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于新建帖子接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关"
                ],
                "summary": "新建帖子",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "社区ID*",
                        "name": "community_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "内容",
                        "name": "content",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "标题*",
                        "name": "title",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/GetPostDetail": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于获取帖子详情接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关"
                ],
                "summary": "获取帖子详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "帖子ID",
                        "name": "ID",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/Login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户登录的接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "登录",
                "parameters": [
                    {
                        "type": "string",
                        "description": "密码*",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名*",
                        "name": "username",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/PostList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于获取帖子列表接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "帖子相关"
                ],
                "summary": "获取帖子列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "integer",
                        "description": "获取帖子的模式",
                        "name": "Mode",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "社区ID/用户ID",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "分页页码",
                        "name": "Page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页最大数量",
                        "name": "Max",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/SignUp": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户注册的接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "注册",
                "parameters": [
                    {
                        "type": "string",
                        "description": "re密码*",
                        "name": "confirm_password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "邮箱*",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "密码*",
                        "name": "password",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "用户名*",
                        "name": "username",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "验证码*",
                        "name": "verification_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/VerificationCode": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于发送验证码的接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户相关"
                ],
                "summary": "发送验证码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "发送地址*",
                        "name": "email",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/Vote": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于用户投票接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "投票相关"
                ],
                "summary": "用户投票接口",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "example": "0",
                        "description": "投票 反对 取消 赞成",
                        "name": "direction",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "0",
                        "description": "帖子id",
                        "name": "post_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/communityDetail": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于获取社区详情接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关"
                ],
                "summary": "获取社区详情",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    },
                    {
                        "type": "string",
                        "example": "0",
                        "description": "社区ID*",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        },
        "/communityList": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用于获取获取社区列表接口 内部调用grpc接口",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "社区相关"
                ],
                "summary": "获取社区列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "Authorization",
                        "in": "header"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ResponseStruct"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.ResponseStruct": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "数据",
                    "type": "object"
                },
                "msg": {
                    "description": "提示信息",
                    "type": "object"
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
	Version:     "0.01",
	Host:        "Gin:8080 Grpc: 8082/8083",
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Happy",
	Description: "基于Grpc/Gin 扩展的一个社区系统",
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
