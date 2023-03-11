spec:
	swagger generate spec -m -w ./cmd/server -o ./api/swagger.yaml
serve:
	swagger serve -F=swagger ./api/swagger.yaml
