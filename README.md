# manage-template

## 실행

go run main.go

## 실행 환경

- abi, bin

  solcjs --abi .\contracts\Manage.sol --bin .\contracts\Manage.sol -o .\contracts\build

- abigen

  abigen --bin=.\contracts\build\contracts_Manage_sol_Manage.bin  --abi=.\contracts\build\contracts_Manage_sol_Manage.abi --pkg build --type Manage --out .\contracts\build\Manage.go


- swagger make openapi

  swag init --parseDependency --parseInternal --parseDepth 1

- convert openapi

  npx redoc-cli --output index.html bundle openapi.yaml

##

2022.09.20
