package handlers

import (
    "encoding/json"
    "fealtyx-api/models"
    "fealtyx-api/util"
    "github.com/go-chi/chi"
    "fmt"
    "net/http"
    "strconv"
    "sync"
)

var (
    students = make(map[int]models.Student)
    nextID   = 1
    mu       sync.Mutex
)

// CreateStudent - POST /students
func CreateStudent(w http.ResponseWriter, r *http.Request) {
    var student models.Student
    if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }

    if student.Name == "" || student.Age <= 0 || student.Email == "" {
        http.Error(w, "Missing or invalid student data", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student.ID = nextID
    students[nextID] = student
    nextID++
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(student)
}

// GetAllStudents - GET /students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    var studentList []models.Student
    for _, student := range students {
        studentList = append(studentList, student)
    }
    mu.Unlock()

    json.NewEncoder(w).Encode(studentList)
}

// GetStudentByID - GET /students/{id}
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student, exists := students[id]
    mu.Unlock()

    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(student)
}

// UpdateStudentByID - PUT /students/{id}
func UpdateStudentByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    var updatedStudent models.Student
    if err := json.NewDecoder(r.Body).Decode(&updatedStudent); err != nil {
        http.Error(w, "Invalid input data", http.StatusBadRequest)
        return
    }

    if updatedStudent.Name == "" || updatedStudent.Age <= 0 || updatedStudent.Email == "" {
        http.Error(w, "Missing or invalid student data", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student, exists := students[id]
    if !exists {
        mu.Unlock()
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    updatedStudent.ID = student.ID
    students[id] = updatedStudent
    mu.Unlock()

    json.NewEncoder(w).Encode(updatedStudent)
}

// DeleteStudentByID - DELETE /students/{id}
func DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    if _, exists := students[id]; !exists {
        mu.Unlock()
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }
    delete(students, id)
    mu.Unlock()

    w.WriteHeader(http.StatusNoContent)
}

// GetStudentSummary - GET /students/{id}/summary
func GetStudentSummary(w http.ResponseWriter, r *http.Request) {
    idStr := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        http.Error(w, "Invalid student ID", http.StatusBadRequest)
        return
    }

    mu.Lock()
    student, exists := students[id]
    mu.Unlock()

    if !exists {
        http.Error(w, "Student not found", http.StatusNotFound)
        return
    }

    summary, err := util.GenerateOllamaSummary(student.ID, student.Name, student.Age, student.Email)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error generating summary: %v", err), http.StatusInternalServerError)
        return
    }

    response := map[string]string{"summary": summary}
    json.NewEncoder(w).Encode(response)
}


