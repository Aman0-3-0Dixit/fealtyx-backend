FealtyX Backend Go Student API Assignment
This project is a RESTful API built with Go, designed to manage student data and generate summaries using the Ollama CLI. With endpoints to create, retrieve, update, and delete student records, this API can also generate concise summaries of student details through integration with the Ollama model.


Go (version used): go1.23.3.windows-amd64
Ollama CLI: Install the Ollama CLI by going to their website and then running the following commands:
ollama --version   // Once installed, verify it by running
ollama run Llama3 // Run it in a separate command prompt. This will download the Llama3 model file and its dependency into the local system
ollama start //Run this command in command prompt or windows powershell to start a server instance of ollama in your system

Now go to your terminal in your IDE and run this command
ollama pull llama3 //This will pull the model into use 

Set up of project into a local system
Step 1: Clone the Repository
First, clone the project repository to your local machine:

git clone https://github.com/Aman0-3-0Dixit/fealtyx-backend.git
cd fealtyx-backend

Step 2: Install Go Modules
The project dependencies are managed with Go modules. Run the following command to download the necessary modules:

go mod tidy
This command will install any missing dependencies.

Step 3: Set Up Environment Variables
If your project uses environment variables

Running the API
To start the API server, run:

go run main.go
The server will start and listen on http://localhost:8080 by default.

API Endpoints
1. Create a New Student
Method: POST
Endpoint: /students
Description: Adds a new student to the database.
Request Body:
{
  "name": "John Doe",
  "age": 21,
  "email": "john.doe@example.com"
}
Response:
{
  "id": 1,
  "name": "John Doe",
  "age": 21,
  "email": "john.doe@example.com"
}

2. Get All Students
Method: GET
Endpoint: /students
Description: Retrieves a list of all students.
Response:
[
  {
    "id": 1,
    "name": "John Doe",
    "age": 21,
    "email": "john.doe@example.com"
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "age": 22,
    "email": "jane.smith@example.com"
  }
]

3. Get a Student by ID
Method: GET
Endpoint: /students/{id}
Description: Retrieves details of a student by ID.
Response:
{
  "id": 1,
  "name": "John Doe",
  "age": 21,
  "email": "john.doe@example.com"
}

4. Update a Student by ID
Method: PUT
Endpoint: /students/{id}
Description: Updates a student’s details by ID.
Request Body:
{
  "name": "Jane Doe",
  "age": 22,
  "email": "jane.doe@example.com"
}
Response:
{
  "id": 1,
  "name": "Jane Doe",
  "age": 22,
  "email": "jane.doe@example.com"
}

5. Delete a Student by ID
Method: DELETE
Endpoint: /students/{id}
Description: Deletes a student by ID.
Response:
{
  "message": "Student deleted successfully"
}

6. Generate a Summary of a Student by ID Using Ollama
Method: GET
Endpoint: /students/{id}/summary
Description: Generates a concise summary of a student’s information using the Ollama CLI.
Response:
{
  "summary": "John Doe is a 21-year-old student."
}
Here in the last endpoint, the response comes out to be a bit different which is flooded with spinning animations of the terminal with the summary that is provide by the ollama model which looks like this 
Response:
{"summary":"Here is a concise summary of the student information:ID: 1Name: John DoeAge: 21Email: john@example.com⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠦ ⠇ ⠇ ⠋ ⠙ ⠙ ⠸ ⠼ ⠴ ⠴ ⠦ ⠇ ⠇ ⠋ ⠙ ⠹ ⠹ ⠼ ⠴ ⠴ ⠧ ⠇ ⠇ ⠋ ⠋ ⠙ ⠸ ⠼ ⠴ ⠴ ⠧ ⠇ ⠏ ⠋ ⠋ ⠹ ⠹ ⠸ ⠴ ⠴ ⠧ ⠇ ⠏ ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠇ ⠋ ⠋ ⠙ ⠹ ⠼ ⠼ ⠦ ⠧ ⠧ ⠏ ⠏ ⠙ ⠹ ⠸ ⠼ ⠼ ⠦ ⠧ ⠇ ⠏ ⠏ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠇ ⠋ ⠙ ⠙ ⠸ ⠸ ⠴ ⠦ ⠧ ⠇ ⠇ ⠋ ⠋ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏ ⠋ ⠙ ⠙ ⠸ ⠼ ⠼ ⠦ ⠧ ⠇ ⠏ ⠋ ⠙ ⠙ ⠸ ⠼ ⠼ ⠦ ⠧ ⠧ ⠏ ⠋ ⠙ ⠹ ⠸ ⠼ ⠼ ⠦ ⠧ ⠇ ⠇ ⠏ ⠋ ⠹ ⠸ ⠸ ⠴ ⠦ ⠦ ⠇ ⠏ ⠋ ⠋ ⠹ ⠹ ⠼ ⠼ ⠦ ⠧ ⠇ ⠏ ⠏ ⠋ ⠹ ⠸ ⠼ ⠴ ⠴ ⠦ ⠇ ⠇ ⠋ ⠙ ⠹ ⠸ ⠸ ⠴ ⠦ ⠧ ⠧ ⠏ ⠋ ⠙ ⠹ ⠹ ⠸ ⠴ ⠦ ⠧ ⠇ ⠏ ⠋ ⠋ ⠹ ⠸ ⠼ ⠼ ⠦ ⠦ ⠇ ⠇ ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠧ ⠇ ⠏ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏ ⠋ ⠙ ⠹ ⠸ ⠼ ⠼ ⠦ ⠧ ⠧ ⠇ ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠴ ⠧ ⠧ ⠏ ⠋"}

Testing the API
To test API functionality, you can use a tool like curl or Postman.
Curl commands to test from the terminal

Add a Student:
curl -X POST http://localhost:8080/students
-H "Content-Type: application/json"
-d '{"name": "John Doe", "age": 21, "email": "john.doe@example.com"}'


Get All Students:
curl -X GET http://localhost:8080/students


Get a Student by ID:
curl -X GET http://localhost:8080/students/1


Get Student Summary:
curl -X GET http://localhost:8080/students/1/summary
-H "Content-Type: application/json"


Update a Student:
curl -X PUT http://localhost:8080/students/1
-H "Content-Type: application/json"
-d '{"name": "Jane Doe", "age": 22, "email": "jane.doe@example.com"}'


Delete a Student:
curl -X DELETE http://localhost:8080/students/1