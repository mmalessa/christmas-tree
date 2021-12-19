Christmas tree lighting control tool - ws2811 strip  

# Requirements
- Raspberry PI with Raspbian
- docker with buidx package
- docker-compose

# Hardware
```
Strip +             --------> +5V
Strip GND           --------> GND
Strip DATA          --------> RPI PIN 12 (GPIO18)

RPI PIN 2           --------> +5V
RPI PIN 6           --------> GND
RPI PIN 12          --------> Strip DATA
```

# Docker buildx
You can download the latest buildx binary from the Docker buildx releases page on GitHub, 
`https://github.com/docker/buildx/releases/`
copy it to ~/.docker/cli-plugins folder with name docker-buildx and change the permission to execute:
```
chmod a+x ~/.docker/cli-plugins/docker-buildx
```
then
```
docker run --rm --privileged docker/binfmt:820fdd95a9972a5308930a2bdfb8573dd4447ad3
docker buildx ls

## Result
# NAME/NODE DRIVER/ENDPOINT STATUS  PLATFORMS
# default * docker                  
#  default default         running linux/amd64, linux/386, linux/arm64, linux/ppc64le, linux/s390x, linux/arm/v7, linux/arm/v6
```


## Links
- https://docs.docker.com/buildx/working-with-buildx/
- https://www.docker.com/blog/getting-started-with-docker-for-arm-on-linux/
- https://github.com/rpi-ws281x/rpi-ws281x-go/
