package main

import (
    "fealtyx-api/handlers"
    "github.com/go-chi/chi"
    "log"
    "net/http"
)

func main() {
    r := chi.NewRouter()

    
    r.Post("/students", handlers.CreateStudent)
    r.Get("/students", handlers.GetAllStudents)

    r.Route("/students/{id}", func(r chi.Router) {
        r.Get("/", handlers.GetStudentByID)
        r.Put("/", handlers.UpdateStudentByID)
        r.Delete("/", handlers.DeleteStudentByID)
        r.Get("/summary", handlers.GetStudentSummary)
    })

    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}

