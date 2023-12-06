# File Conversion Service

## Overview

This is a great conversion service with service written on go, server written on python and frontend on react.js.

## Build and run

### Prerequisities


In order to run this container you'll need docker installed.

* [Windows](https://docs.docker.com/windows/started)
* [OS X](https://docs.docker.com/mac/started/)
* [Linux](https://docs.docker.com/linux/started/)

After that, check the Dockerfile in all projects and check the env setup. In the frontend/.env file you can set up the backend environment and port for web.

All you need to do is to run 

```bash
docker-compose -f docker-compose.yml up
```

And then view frontend in the browser by the link provided in the logs.