package admin

import (
	"context"
	"log"
	"net/http"

	"encoding/json"
	pb "github.com/yerkebayev/go-final-go/proto"
	"google.golang.org/grpc"
)

type AdminService struct {
	client pb.AdminServiceClient
}

func NewAdminService(client pb.MainServiceClient) *AdminService {
	return &AdminService{client: client}
}

func (s *AdminService) AddCourse(w http.ResponseWriter, r *http.Request) {
	var req pb.AddCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.client.AddCourse(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	var req pb.DeleteCourseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.client.DeleteCourse(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var req pb.DeleteStudentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.client.DeleteStudent(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	var req pb.DeleteTeacherRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := s.client.DeleteTeacher(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewMainServiceClient(conn)
	adminService := NewAdminService(client)

	http.HandleFunc("/addCourse", adminService.AddCourse)
	http.HandleFunc("/deleteCourse", adminService.DeleteCourse)
	http.HandleFunc("/deleteStudent", adminService.DeleteStudent)
	http.HandleFunc("/deleteTeacher", adminService.DeleteTeacher)

	log.Println("Admin service listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
