# Image Resize Service

### Instruction

**Prerequisite**:
Make sure you have [Docker](https://www.docker.com/get-started)
Also make sure you have docker-compose. It should come with Docker Mac/Windows version but might need to be installed separately with package managers with different version of Unix.

* If you have go installed, you can download this by `go get github.com/funfoolsuzi/imgresize`. Or just clone it.

To build the Docker image:
```
make build
```

To start the entire service via docker-compose:
```
make up
```

To tear down the entire service via docker-compose:
```
make down
```

Once the docker-compose is up, try a few url in your browser:
"http://127.0.0.1:8080/originals/gophers/fancygopher.jpg"
"http://127.0.0.1:8080/resized/gophers/fancygopher_h100_w100.jpg"

### Description

This app relies on another go repo: [github.com/h2non/imaginary](https://github.com/h2non/imaginary)

The app uses _imaginary_ as a service in docker-compose to resize image. The service itself is pre-loaded with a few images. The original images can be accessed at "http://127.0.0.1:8080/originals/". It is built with go's default http file server. The resized images can be accessed from "http://127.0.0.1:8080/resized/". The file path should be the same with the originals. But the image name is different. It follows a convention. For example, if the original is saved at "/originals/this/image.jpg", then you can access the resized image at "/resized/this/image_h100_w100.jpg". So to sum up, for resized image `/resized/{path_to_image}/{image_name}_h{height}_w{width}.{file_extention}`.

If a resized image doesn't exist at first, the main service will contact an instance of of imaginary for a resized image and then store it at certain location for future access.

log for the main service is persisted at `./log/app.log`. Once you run the service, you can get most updated log entries withouth `docker exec -it container_name /bin/sh`.

