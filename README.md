# Mini Social Media API

## Core Functionality:
1. Create new posts (posts will just be a text message) > `POST posts`
2. Update posts > `PUT posts/{id}`
3. Retrieve all posts > `GET posts`
4. Like a specific post > `POST posts/{id}/like`
5. Add a comment to a specific post > `POST posts/{id}/comment`
6. Retrieve all details of a specific post, including all the comments and likes >  `GET posts/{id}`


## Getting Started

### Step 1: Set the Configuration File

To configure the application, create a YAML configuration file, such as `config.yaml`, with the following structure:

```yaml
server:
  port: 8080
```

Next, set the `MIN_SM_API_CONFIG_PATH` environment variable to the path of your `config.yaml` file. This can be done in your terminal or programmatically.

- **For Linux/macOS**:
```bash
export MIN_SM_API_CONFIG_PATH=/path/to/your/config.yaml
```

- **For Windows (Command Prompt)**:
```cmd
set MIN_SM_API_CONFIG_PATH=C:\path\to\your\config.yaml
```

### Step 2: Run the Application

1. Change the directory to `cmd/min-sm-api`:

```bash
cd cmd/min-sm-api
```

2. Run the application:

```bash
go run main.go
```

> Note: You can find the API collection in the `docs` directory.
