codegen:
	oapi-codegen \
	-generate gin,spec \
	-package handler -o ./handler/handler.gen.go ./config/openapi.yaml

	oapi-codegen \
	-generate types \
	-package model -o ./model/model.gen.go ./config/openapi.yaml