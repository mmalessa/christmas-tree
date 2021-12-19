# Christmas tree lighting control tool - ws2811 strip  

## Requirements
- Raspberry PI (v2 min) with Raspbian
- docker with buidx package
- docker-compose

## Hardware
```
Strip +             --------> +5V
Strip GND           --------> GND
Strip DATA          --------> RPI PIN 12 (GPIO18)

RPI PIN 2           --------> +5V
RPI PIN 6           --------> GND
RPI PIN 12          --------> Strip DATA
```

## Use Docker buildx
You can download the latest buildx binary from the Docker buildx releases page on GitHub, 
`https://github.com/docker/buildx/releases/`
copy it to ~/.docker/cli-plugins folder with name docker-buildx and change the permission to execute:
```sh
chmod a+x ~/.docker/cli-plugins/docker-buildx
```

## Prepare
- `cp .env.dist .env` and change `RPI_IP` and `RPI_USER`
- Change (if needed) configuration `/config/*.yaml` files
- Run `make init` to init PC environment
- Run `make rpi-authorize` to copy ssh private key to RPI (login without password)
- Run `make` to compile code
- Run `make rpi-install` to copy binary file to RPI
- Run `make rpi-enable-service` to set and enable christmastree service on RPI
- Run `make rpi-start-service` to start christmastree service

## On every change
- Run `make`
- Run `make rpi-stop-service` (if started)
- Run `make rpi-install`
- Run `make rpi-start-service`



## Links
- https://docs.docker.com/buildx/working-with-buildx/
- https://www.docker.com/blog/getting-started-with-docker-for-arm-on-linux/
- https://github.com/rpi-ws281x/rpi-ws281x-go/
