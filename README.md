# Arkose Fetch

```
docker run -d -p 3610:3610 ghcr.io/imtms/funcaptcha:latest
```
或
```
version: "3"
services:
  arkose-token:
     container_name: arkose-token
     image: ghcr.io/imtms/funcaptcha:latest
     ports:
       - 3610:3610
```

如果你有自己的bx发布url，请使用BX_URL环境变量传入

```
docker run -d -p 3610:3610 -e BX_URL=https://bx.tms.im ghcr.io/imtms/funcaptcha:latest
```
或
```
version: "3"
services:
  arkose-token:
     container_name: arkose-token
     image: ghcr.io/imtms/funcaptcha:latest
     environment:
       - BX_URL=https://bx.tms.im
     ports:
       - 3610:3610
```
Forked from: [acheong08/funcaptcha](https://github.com/acheong08/funcaptcha)
