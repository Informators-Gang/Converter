# File Conversion Service

## Overview

This is a simple and efficient file conversion service written in Go. The purpose of this service is to provide a convenient way to convert files from one format to another, supporting a variety of file types.

## Build and run

```bash
cd file-conversion-service
docker build -t file-conversion-service .
docker run -p 5353:5353 file-conversion-service
```

## Endpoints

### 1. File Uploading

- **URL**: `/upload`
- **Method**: `POST`
- **Body**: 
  - `file`: File to upload.
- **Response**: 
  - `file_id`: Unique identifier for the uploaded file.
- **Example Response**:
  ```json
  {
      "file_id": "3c6bd14a-3d9f-46aa-9764-dd1db8e768dd"
  }

### 2. File Information

- **Endpoint**: `/file_info?file_id=<file_id>`
- **Method**: `GET`
- **Description**: Retrieves information about an uploaded file.
- **Query Parameters**:
  - `file_id`: The unique identifier of the file.
- **Response**: JSON object containing file details.
- **Example Response**:
  ```json
  {
    "filename": "-2147483648_-214812.jpg",
    "size": 397455,
    "extension": "jpg",
    "convertible_formats": ["png", "tiff"]
  }

### 3. Convert File

- **Endpoint**: `/convert?file_id=<file_id>&convert_to=<format>`
- **Method**: `POST`
- **Description**: Converts an uploaded file to a specified format.
- **Query Parameters**:
  - `file_id`: The unique identifier of the file.
  - `convert_to`: The desired format for conversion (e.g., "png").
- **Response**: JSON object containing the new file ID.
- **Example Response**:
  ```json
  {
    "new_file_id": "16299075-9d78-4ff6-8a22-596a8982a173"
  }

### 4. Download File

- **Endpoint**: `/download?file_id=<file_id>`
- **Method**: `GET`
- **Description**: Downloads the file associated with the given file ID.
- **Query Parameters**:
  - `file_id`: The unique identifier of the file.
- **Response**: The file is downloaded.