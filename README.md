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
Forked from: [acheong08/funcaptcha](https://github.com/acheong08/funcaptcha)
