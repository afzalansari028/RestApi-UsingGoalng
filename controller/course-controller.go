package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/afzal/go-course/helper"
	"github.com/afzal/go-course/models"
	"github.com/gorilla/mux"
)

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")

	db := helper.SetupDB()
	rows, err := db.Query("select * from courses")
	if err != nil {
		log.Fatal(err)
	}

	var courses []models.Course

	for rows.Next() {
		var courseId string
		var courseName string
		var price int
		var author string

		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	// fmt.Println(courses)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(courses)
}
func GetOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	id := params["id"]

	db := helper.SetupDB()
	rows, err := db.Query("select * from courses where course_id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	var courses []models.Course
	for rows.Next() {
		var courseId string
		var courseName string
		var price int
		var author string

		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	// log.Print(courses)
	if courses != nil {
		json.NewEncoder(w).Encode(courses)
	} else {
		json.NewEncoder(w).Encode("No data is available of this id")
	}
}

func CreateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	//what about - {}
	var course models.Course
	var courses []models.Course
	json.NewDecoder(r.Body).Decode(&course)
	if course.CourseName == "" {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	if course.CourseId == "" {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	db := helper.SetupDB()
	var courseId = course.CourseId
	var courseName = course.CourseName
	var price = course.CoursePrice
	var author = course.Author
	// fmt.Println(courseId, courseName, price, author)

	//validate duplicate data
	rows, _ := db.Query("SELECT * FROM courses")
	for rows.Next() {
		var (
			courseId   string
			courseName string
			price      int
			author     string
		)
		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}
	for _, dbCourses := range courses {
		if courseId == dbCourses.CourseId {
			json.NewEncoder(w).Encode("This course is already available, Please try with other...!")
			return
		}
	}

	db.Query("INSERT INTO courses(course_id,course_name,course_price,author) values(?,?,?,?)", courseId, courseName, price, author)
	json.NewEncoder(w).Encode("Course added...")
	w.WriteHeader(http.StatusOK)
}

func UpdateOnecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	// w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	var updatecourse models.Course
	json.NewDecoder(r.Body).Decode(&updatecourse)
	var (
		courseName   = updatecourse.CourseName
		coursePrice  = updatecourse.CoursePrice
		courseAuthor = updatecourse.Author
	)
	fmt.Println(courseName, coursePrice, courseAuthor)
	db := helper.SetupDB()
	db.Query("UPDATE courses SET course_name=?,course_price=?,author=? where course_id=?", courseName, coursePrice, courseAuthor, id)
	json.NewEncoder(w).Encode("Course updated")
	w.WriteHeader(http.StatusOK)
}

func DeleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course from db")
	w.Header().Set("Content-Type,", "application/json")

	params := mux.Vars(r)
	id := params["id"]

	db := helper.SetupDB()

	rows, _ := db.Query("SELECT * FROM courses")
	var courses []models.Course
	for rows.Next() {
		var (
			courseId   string
			courseName string
			price      int
			author     string
		)
		err := rows.Scan(&courseId, &courseName, &price, &author)
		if err != nil {
			log.Fatal(err)
		}
		courses = append(courses, models.Course{CourseId: courseId, CourseName: courseName, CoursePrice: price, Author: author})
	}

	for _, courseDB := range courses {
		if courseDB.CourseId == id {
			db.Query("DELETE FROM courses where course_id = ? ", id)
			json.NewEncoder(w).Encode("Course is deleted of given id")
			return
		}
	}
	json.NewEncoder(w).Encode("Course not found by given id")
	w.WriteHeader(http.StatusNotFound)
}
