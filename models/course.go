package models

//Model for course - file
type Course struct {
	CourseId   string `json:"courseid"`
	CourseName string `json:"coursename"`
	// CoursePrice int     `json:"-"`
	CoursePrice int    `json:"price"`
	Author      string `json:"author"`
}
