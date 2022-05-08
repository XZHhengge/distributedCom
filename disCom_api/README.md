# distributedCom
## disCom_api
**使用tornado作为web通过grpc与disCom_srv通讯**

在proto目录下
`python -m grpc_tools.protoc --python_out=. --grp
c_python_out=. -I. code.proto`