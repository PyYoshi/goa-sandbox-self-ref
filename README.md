```bash
export GO111MODULE=on

go get -u goa.design/goa/v3/...@v3
go get -u github.com/golang/protobuf/protoc-gen-go

goa gen github.com/PyYoshi/goa-sandbox-self-ref/design
goa example github.com/PyYoshi/goa-sandbox-self-ref/design
```

#

v3:
