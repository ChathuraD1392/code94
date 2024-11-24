# code94

###Mini Social Media API
Core Functionality:
Create new posts (posts will just be a text message) > POST posts
Update posts > PUT posts/{id}
Retrieve all posts > GET posts
Like a specific post > POST posts/{id}/like
Add a comment to a specific post > POST posts/{id}/comment
Retrieve all details of a specific post, including all the comments and likes > GET posts/{id}
Getting Started
Step 1: Set the Configuration File
To configure the application, create a YAML configuration file, such as config.yaml, with the following structure:

server:
port: 8080
Next, set the MIN_SM_API_CONFIG_PATH environment variable to the path of your config.yaml file. This can be done in your terminal or programmatically.

For Linux/macOS:
export MIN_SM_API_CONFIG_PATH=/path/to/your/config.yaml
For Windows (Command Prompt):
set MIN_SM_API_CONFIG_PATH=C:\path\to\your\config.yaml
Step 2: Run the Application
Change the directory to cmd/min-sm-api:
cd cmd/min-sm-api
Run the application:
go run main.go
Note: You can find the API collection in the docs directory.
