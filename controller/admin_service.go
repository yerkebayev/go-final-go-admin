package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	pb "github.com/yerkebayev/go-final-go/proto"
	"google.golang.org/grpc"
)

type AdminService struct {
	client pb.AdminServiceClient
}

func NewAdminService(client pb.AdminServiceClient) *AdminService {
	return &AdminService{client: client}
}

func (s *AdminService) AddCourse(w http.ResponseWriter, r *http.Request) {
	courseName := r.URL.Query().Get("name")
	if courseName == "" {
		http.Error(w, "Missing 'name' query parameter", http.StatusBadRequest)
		return
	}
	req := &pb.AddCourseRequest{Name: courseName}
	res, err := s.client.AddCourse(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteCourse(w http.ResponseWriter, r *http.Request) {
	courseIDStr := r.URL.Query().Get("id")
	if courseIDStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	courseID, err := strconv.ParseInt(courseIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}
	req := &pb.DeleteCourseRequest{Id: int32(courseID)}
	res, err := s.client.DeleteCourse(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	studentIDStr := r.URL.Query().Get("id")
	if studentIDStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	studentID, err := strconv.ParseInt(studentIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}
	req := &pb.DeleteStudentRequest{Id: int32(studentID)}
	res, err := s.client.DeleteStudent(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	teacherIDStr := r.URL.Query().Get("id")
	if teacherIDStr == "" {
		http.Error(w, "Missing 'id' query parameter", http.StatusBadRequest)
		return
	}
	teacherID, err := strconv.ParseInt(teacherIDStr, 10, 32)
	if err != nil {
		http.Error(w, "Invalid 'id' query parameter", http.StatusBadRequest)
		return
	}
	req := &pb.DeleteTeacherRequest{Id: int32(teacherID)}
	res, err := s.client.DeleteTeacher(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) GetCourses(w http.ResponseWriter, r *http.Request) {
	req := &pb.Empty{}
	res, err := s.client.GetCourses(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) GetStudents(w http.ResponseWriter, r *http.Request) {
	req := &pb.Empty{}
	res, err := s.client.GetStudents(context.Background(), req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func (s *AdminService) GetTeachers(w http.ResponseWriter, r *http.Request) {
	req := &pb.Empty{}
	res, err := s.client.GetTeachers(context.Background(), req)
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

	client := pb.NewAdminServiceClient(conn)
	adminService := NewAdminService(client)

	http.HandleFunc("/addCourse", adminService.AddCourse)
	http.HandleFunc("/deleteCourse", adminService.DeleteCourse)
	http.HandleFunc("/deleteStudent", adminService.DeleteStudent)
	http.HandleFunc("/deleteTeacher", adminService.DeleteTeacher)
	http.HandleFunc("/getCourses", adminService.GetCourses)
	http.HandleFunc("/getStudents", adminService.GetStudents)
	http.HandleFunc("/getTeachers", adminService.GetTeachers)

	log.Println("Admin service listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
