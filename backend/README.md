# File Conversion Service

## Overview

This is a simple and efficient backend server in Python. The purpose of this service is to provide a convenient way to communicate between front and converter service.

## Build and run

```bash
docker-compose up
```

## Endpoints

### 1. File Uploading

- **URL**: `/convert`
- **Method**: `POST`
- **Body**: 
  - `file`: File to upload.
  - `convert_to`: Format to convert to
- **Response**: The file is downloaded.
