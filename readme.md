
# FealtyX Backend Go Student API Assignment

This project is a RESTful API built with Go, designed to manage student data and generate summaries using the Ollama CLI. With endpoints to create, retrieve, update, and delete student records, this API can also generate concise summaries of student details through integration with the Ollama model.

## Requirements

- **Go (version used):** `go1.23.3.windows-amd64`
- **Ollama CLI**: Install the Ollama CLI by going to their website. Once installed, verify it by running:
  
  ```bash
  ollama --version
  ```

  Then, in a separate command prompt, run:

  ```bash
  ollama run llama3
  ```

  This will download the Llama3 model file and its dependencies into your local system. To start a server instance of Ollama, run:

  ```bash
  ollama start
  ```

  Finally, to pull the model into use, go to your terminal in your IDE and run:

  ```bash
  ollama pull llama3
  ```

## Set up the Project on a Local System

### Step 1: Clone the Repository

Clone the project repository to your local machine:

```bash
git clone https://github.com/Aman0-3-0Dixit/fealtyx-backend.git
cd fealtyx-backend
```

### Step 2: Install Go Modules

The project dependencies are managed with Go modules. Run the following command to download the necessary modules:

```bash
go mod tidy
```

### Step 3: Set Up Environment Variables

If your project uses environment variables, configure them accordingly.

## Running the API

To start the API server, run:

```bash
go run main.go
```

The server will start and listen on `http://localhost:8080` by default.

## API Endpoints

### 1. Create a New Student

- **Method**: POST
- **Endpoint**: `/students`
- **Description**: Adds a new student to the database.
- **Request Body**:
  ```json
  {
    "name": "John Doe",
    "age": 21,
    "email": "john.doe@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "age": 21,
    "email": "john.doe@example.com"
  }
  ```

### 2. Get All Students

- **Method**: GET
- **Endpoint**: `/students`
- **Description**: Retrieves a list of all students.
- **Response**:
  ```json
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
  ```

### 3. Get a Student by ID

- **Method**: GET
- **Endpoint**: `/students/{id}`
- **Description**: Retrieves details of a student by ID.
- **Response**:
  ```json
  {
    "id": 1,
    "name": "John Doe",
    "age": 21,
    "email": "john.doe@example.com"
  }
  ```

### 4. Update a Student by ID

- **Method**: PUT
- **Endpoint**: `/students/{id}`
- **Description**: Updates a student’s details by ID.
- **Request Body**:
  ```json
  {
    "name": "Jane Doe",
    "age": 22,
    "email": "jane.doe@example.com"
  }
  ```
- **Response**:
  ```json
  {
    "id": 1,
    "name": "Jane Doe",
    "age": 22,
    "email": "jane.doe@example.com"
  }
  ```

### 5. Delete a Student by ID

- **Method**: DELETE
- **Endpoint**: `/students/{id}`
- **Description**: Deletes a student by ID.
- **Response**:
  ```json
  {
    "message": "Student deleted successfully"
  }
  ```

### 6. Generate a Summary of a Student by ID Using Ollama

- **Method**: GET
- **Endpoint**: `/students/{id}/summary`
- **Description**: Generates a concise summary of a student’s information using the Ollama CLI.
- **Expected Response**:
  ```json
  {
    "summary": "John Doe is a 21-year-old student."
  }
  ```
  *Note*: In some cases, the response may include extra symbols due to terminal animations. An example:

  ```json
  {
    "summary": "Here is a concise summary of the student information:
ID: 1
Name: John Doe
Age: 21
Email: john@example.com⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ..."
  }
  ```

## Testing the API

To test API functionality, you can use `curl` commands from the terminal, as shown below, or a tool like Postman.

### Curl Commands to Test Endpoints

- **Add a Student**:
  ```bash
  curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "age": 21, "email": "john.doe@example.com"}'
  ```

- **Get All Students**:
  ```bash
  curl -X GET http://localhost:8080/students
  ```

- **Get a Student by ID**:
  ```bash
  curl -X GET http://localhost:8080/students/1
  ```

- **Get Student Summary**:
  ```bash
  curl -X GET http://localhost:8080/students/1/summary \
  -H "Content-Type: application/json"
  ```

- **Update a Student**:
  ```bash
  curl -X PUT http://localhost:8080/students/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Doe", "age": 22, "email": "jane.doe@example.com"}'
  ```

- **Delete a Student**:
  ```bash
  curl -X DELETE http://localhost:8080/students/1
  ```