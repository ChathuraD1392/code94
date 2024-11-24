# Mini Social Media API

##1. Solution:
While maintaining the guidelines for good coding practice in Golang, I created and implemented a clean API responsible for the posts of a mini social media application. The implementation respects the standard practices in Golang programming and makes use of a clean architecture of a modular package design.

The critical aspects of the solution include:

Unit Tests: Unit tests have been implemented for the business logic in order to validate the features of the platform. Other tests can be added for other packages for higher coverage.

Golang Generics: The application architecture is heavily dependent on the use of Golang generics that were introduced in earlier versions, hence the need for flexibility and reusability.

In-Memory Storage: In order to demonstrate the application programming interface (API) in the simplest manner possible, in-memory data storage is employed. However, the modular design permits easy switching to a database if necessary.

Traceable Logger: A traceable logger has been achieved, thus making it possible to trace errors through unique UMID of requests, for easier and better debugging and monitoring.

Configuration Management: All application settings are read from the environment variables, providing more responsiveness.

This design shows a high level of scale, sound maintenance, and contemporary trends in the Golang development environment.


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

##3.Assumptions
Authentication level: Use of a login and password is not required on the platform when creating posts.
Anonymous Posts: The feature allows creating posts anonymously not tagging them to a specific member.

##4.Possible Improvments

Adding security features to authenticate posts with registered users is important as it prevents data manipulation.
Use technologies such as OAuth2, JWT or session-based authentication which best suit the application.
MongoDB, PostgreSQL, MySQL – persistent storage of this type is reasonable and able to import and work with pretty much large amount of data sets.


 Search functionality: Elasticsearch for posts and efficient search for content.

Improving the system performance: Apply cache (to Redis for example) for frequently used information. Restrict request frequency to discourage misuse and improve service quality.

Better visibility and controls: Increase the logfile coverage to include response to request/response timmings and system global performance. Use monitoring  for real time statistics, customisation can also be done.
