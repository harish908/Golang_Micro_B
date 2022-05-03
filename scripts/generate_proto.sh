  
#!/bin/bash -x

goPath=$(go env GOPATH)

for file in proto/definitions/*.proto
do
  fileName=$(echo "$(basename "$file")")
  name=$(echo "$fileName" | cut -f 1 -d '.')
echo "procesing - $fileName"
  rm -rf proto/gen/$name
  # rm -rf lib/proto/gen/svc/$name

  mkdir -p proto/gen/$name
  # mkdir -p lib/proto/gen/svc/$name

  #protoc -I./lib/proto/definitions -I./lib/proto/third_party/googleapis --micro_out=./lib/proto/gen/svc/$name --go_out=./lib/proto/gen/svc/$name $fileName
  protoc -I./proto/definitions -I=$goPath/src/github.com/grpc-ecosystem/grpc-gateway --go-grpc_out=require_unimplemented_servers=false:./proto/gen/$name --go_out=./proto/gen/$name $fileName
  protoc -I./proto/definitions -I=$goPath/src/github.com/grpc-ecosystem/grpc-gateway --grpc-gateway_out=logtostderr=true:./proto/gen/$name $fileName
done
