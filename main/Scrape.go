package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// File contains information relating to an individual file on canvas
type File struct {
	ID          int
	URL         string
	DisplayName string
}

// Folder contains information relating to folders on canvas
type Folder struct {
	ID              int
	Files           []File
	SubDirectories  *Folder
	Name            string
	ParentDirectory *Folder
}

// Requester is a structure used in the http request to contain related data
type Requester struct {
	Context string
	Headers map[string]string
	BaseURL string
}

// Module is the toplevel struct containing all data related to an individual module
type Module []struct {
	ID                          int         `json:"id"`
	Name                        string      `json:"name"`
	AccountID                   int         `json:"account_id"`
	UUID                        string      `json:"uuid"`
	StartAt                     time.Time   `json:"start_at"`
	GradingStandardID           interface{} `json:"grading_standard_id"`
	IsPublic                    bool        `json:"is_public"`
	CreatedAt                   time.Time   `json:"created_at"`
	CourseCode                  string      `json:"course_code"`
	DefaultView                 string      `json:"default_view"`
	RootAccountID               int         `json:"root_account_id"`
	EnrollmentTermID            int         `json:"enrollment_term_id"`
	License                     string      `json:"license"`
	GradePassbackSetting        interface{} `json:"grade_passback_setting"`
	EndAt                       interface{} `json:"end_at"`
	PublicSyllabus              bool        `json:"public_syllabus"`
	PublicSyllabusToAuth        bool        `json:"public_syllabus_to_auth"`
	StorageQuotaMb              int         `json:"storage_quota_mb"`
	IsPublicToAuthUsers         bool        `json:"is_public_to_auth_users"`
	ApplyAssignmentGroupWeights bool        `json:"apply_assignment_group_weights"`
	Calendar                    struct {
		Ics string `json:"ics"`
	} `json:"calendar"`
	TimeZone    string `json:"time_zone"`
	Blueprint   bool   `json:"blueprint"`
	Enrollments []struct {
		Type                           string `json:"type"`
		Role                           string `json:"role"`
		RoleID                         int    `json:"role_id"`
		UserID                         int    `json:"user_id"`
		EnrollmentState                string `json:"enrollment_state"`
		LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
	} `json:"enrollments"`
	HideFinalGrades                  bool   `json:"hide_final_grades"`
	WorkflowState                    string `json:"workflow_state"`
	RestrictEnrollmentsToCourseDates bool   `json:"restrict_enrollments_to_course_dates"`
	OverriddenCourseVisibility       string `json:"overridden_course_visibility,omitempty"`
	Locale                           string `json:"locale,omitempty"`
}

func getModules(r Requester) ([]Module, error) {
	if len(r.Headers) == 0 {
		return nil, errors.New("Empty headers")
	}
	println("Creating client")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://"+r.BaseURL+r.Context+"/courses", nil)
	req.Header.Add("Authorization", r.Headers["Authorization"])
	println("https://" + r.BaseURL + r.Context)
	println("Executing request")
	resp, err := client.Do(req)
	println("Executed.")
	// return modules, nil
	if err != nil {
		println("Error in request")
		return nil, err
	}
	defer resp.Body.Close()
	println("creating modules array")
	modules := make([]Module, 0)
	println("Reading modules")
	body, err := ioutil.ReadAll(resp.Body)
	println(string(body))
	println("Unmarshalling modules")
	json.Unmarshal(body, &modules)
	return modules, nil
}

func main() {

	baseURLPtr := flag.String("baseUrl", "canvas.bham.ac.uk", "baseUrl for canvas curl, default canvas.bham.ac.uk")
	authorisationTokenPtr := flag.String("auth", "", "Authorisation key from canvas")

	flag.Parse()
	if *authorisationTokenPtr == "" {
		println("ERROR: Please enter an authorisation token!")
		os.Exit(0)
	}
	headers := make(map[string]string)
	headers["Authorization"] = "Bearer " + *authorisationTokenPtr

	requester := Requester{
		Context: "/api/v1",
		Headers: headers,
		BaseURL: *baseURLPtr,
	}

	modules, err := getModules(requester)
	if err != nil {
		log.Fatal(err)
	} else {
		println("Scraped Modules:")
		fmt.Printf("%v", modules)
	}

}
