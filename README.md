
### Build Static Web
```bash
go-bindata-assetfs.exe -o pkg/server/browser/webroot.go -pkg browser -nocompress=true website/build/...
```