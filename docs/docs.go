// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://developers.ctd.com.ar/es_ar/terminos-y-condiciones",
        "contact": {
            "name": "API Support",
            "url": "https://developers.ctd.com.ar/support"
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
        "/appointments/": {
            "post": {
                "description": "Crea un nuevo turno en el sistema",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "POST: agregar turno.",
                "parameters": [
                    {
                        "description": "Appointment object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    }
                }
            }
        },
        "/appointments/create-by-dni-and-license": {
            "post": {
                "description": "Crea un nuevo turno asociado a un paciente identificado por su DNI y a un dentista identificado por su matrícula",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "POST: agregar turno por DNI del paciente y matrícula del dentista.",
                "parameters": [
                    {
                        "description": "Appointment data",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appointments.AppointmentCreateWithDNIAndLicense"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    }
                }
            }
        },
        "/appointments/patient": {
            "get": {
                "description": "Obtiene los turnos de un paciente por su DNI con el detalle de las entidades de paciente y dentista desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "GET: traer turno por DNI del paciente. Debe traer el detalle del turno (Fecha-Hora, descripción, Paciente y Dentista) y el dni deberá ser recibido por QueryParams.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Patient's DNI",
                        "name": "dni",
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
                                "$ref": "#/definitions/appointments.AppointmentGetWithAllEntities"
                            }
                        }
                    }
                }
            }
        },
        "/appointments/{id}": {
            "get": {
                "description": "Obtiene un turno por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "GET: traer turno por ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualiza un turno por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "PUT: actualizar turno.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Appointment object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    }
                }
            },
            "delete": {
                "description": "Elimina un turno por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "DELETE: eliminar turno.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Appointment deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Actualiza un turno parcialmente por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "PATCH: actualizar un turno por alguno de sus campos.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Appointment patch object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/appointments.AppointmentPatch"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appointments.Appointment"
                        }
                    }
                }
            }
        },
        "/appointments/{id}/details": {
            "get": {
                "description": "Obtiene un turno por ID con el detalle de las entidades de paciente y dentista desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "turnos"
                ],
                "summary": "GET: traer turno por ID pero con el detalle de las entidades",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Appointment ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/appointments.AppointmentGetWithAllEntities"
                        }
                    }
                }
            }
        },
        "/dentists/": {
            "post": {
                "description": "Crear dentista en el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "POST: agregar dentista.",
                "parameters": [
                    {
                        "description": "Dentist object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    }
                }
            }
        },
        "/dentists/{id}": {
            "get": {
                "description": "Obtener dentista por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "GET: traer dentista por ID.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualizar dentista por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "PUT: actualizar dentista.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dentist object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    }
                }
            },
            "delete": {
                "description": "Eliminar dentista por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "DELETE: eliminar el dentista.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Dentist deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Actualizar parcialmente un dentista por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "odontologos"
                ],
                "summary": "PATCH: actualizar un dentista por alguno de sus campos.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Dentist ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Dentist patch object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dentists.DentistPatch"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dentists.Dentist"
                        }
                    }
                }
            }
        },
        "/patients/": {
            "post": {
                "description": "Crea un nuevo paciente en el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "POST: agregar paciente.",
                "parameters": [
                    {
                        "description": "Patient object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    }
                }
            }
        },
        "/patients/list": {
            "get": {
                "description": "Obtener listado de paciente desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "GET: lista todos los paciente.",
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/[]patients.Patient"
                        }
                    }
                }
            }
        },
        "/patients/{id}": {
            "get": {
                "description": "Obtener paciente por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "GET: traer paciente por ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    }
                }
            },
            "put": {
                "description": "Actualizar paciente por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "PUT: actualizar paciente.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patient object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    }
                }
            },
            "delete": {
                "description": "Elimina un paciente por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "DELETE: eliminar al paciente.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Patient deleted successfully",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "patch": {
                "description": "Actualiza parcialmente un paciente por ID desde el repositorio",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "pacientes"
                ],
                "summary": "PATCH: actualizar un paciente por alguno de sus campos.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Patient ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Patient patch object",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/patients.PatientPatch"
                        }
                    },
                    {
                        "type": "string",
                        "description": "PRIVATE-KEY",
                        "name": "PRIVATE-KEY",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "PUBLIC-KEY",
                        "name": "PUBLIC-KEY",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/patients.Patient"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "appointments.Appointment": {
            "type": "object",
            "required": [
                "date",
                "dentist_id",
                "hour",
                "patient_id"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "dentist_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "patient_id": {
                    "type": "integer"
                }
            }
        },
        "appointments.AppointmentCreateWithDNIAndLicense": {
            "type": "object",
            "required": [
                "date",
                "dentist_license",
                "hour",
                "patient_dni"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "dentist_license": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "patient_dni": {
                    "type": "string"
                }
            }
        },
        "appointments.AppointmentGetWithAllEntities": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "dentist": {
                    "$ref": "#/definitions/appointments.DentistInfo"
                },
                "description": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "patient": {
                    "$ref": "#/definitions/appointments.PatientInfo"
                }
            }
        },
        "appointments.AppointmentPatch": {
            "type": "object",
            "properties": {
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "hour": {
                    "type": "string"
                }
            }
        },
        "appointments.DentistInfo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "license": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "appointments.PatientInfo": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "discharge_date": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dentists.Dentist": {
            "type": "object",
            "required": [
                "lastname",
                "license",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "license": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "dentists.DentistPatch": {
            "type": "object",
            "properties": {
                "lastname": {
                    "type": "string"
                },
                "license": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "patients.Patient": {
            "type": "object",
            "required": [
                "address",
                "discharge_date",
                "dni",
                "lastname",
                "name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "discharge_date": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "patients.PatientPatch": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "discharge_date": {
                    "type": "string"
                },
                "dni": {
                    "type": "string"
                },
                "lastname": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Grupo 8 Final Backend",
	Description:      "Solución al final de backend de Digital House - Grupo 8 por Felipe Monterrosa, Javier Triana, Fabricio Montivero, Gaston Diaz",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
