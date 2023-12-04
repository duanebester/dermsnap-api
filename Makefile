tidy:
	go mod tidy

codegen-api:
	oapi-codegen \
	-generate fiber,types,spec \
	-package http -o api/http/http.gen.go openapi/api.yaml

codegen-public:
	oapi-codegen \
	-generate fiber,types,spec \
	-package public -o api/public/public.gen.go openapi/public.yaml

codegen: codegen-api codegen-public