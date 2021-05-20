package lib

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
    "regexp"
	"strconv"
	"strings"
	"time"
)

type Page struct {
	Title        string    `json:"title"`
	CreatedAt    time.Time `json:"created_at"`
	URL          string    `json:"url"`
	EditingRoles string    `json:"editing_roles"`
	PageID       int       `json:"page_id"`
	LastEditedBy struct {
		ID             int    `json:"id"`
		DisplayName    string `json:"display_name"`
		AvatarImageURL string `json:"avatar_image_url"`
		HTMLURL        string `json:"html_url"`
		Pronouns       string `json:"pronouns"`
	} `json:"last_edited_by"`
	Published        bool        `json:"published"`
	HideFromStudents bool        `json:"hide_from_students"`
	FrontPage        bool        `json:"front_page"`
	HTMLURL          string      `json:"html_url"`
	TodoDate         interface{} `json:"todo_date"`
	UpdatedAt        time.Time   `json:"updated_at"`
	LockedForUser    bool        `json:"locked_for_user"`
	Body             string      `json:"body"`
}
// File contains information relating to an individual file on canvas
type File struct {
	ID                 int         `json:"id"`
	UUID               string      `json:"uuid"`
	FolderID           int         `json:"folder_id"`
	DisplayName        string      `json:"display_name"`
	Filename           string      `json:"filename"`
	UploadStatus       string      `json:"upload_status"`
	ContentType        string      `json:"content-type"`
	URL                string      `json:"url"`
	Size               int         `json:"size"`
	CreatedAt          time.Time   `json:"created_at"`
	UpdatedAt          time.Time   `json:"updated_at"`
	UnlockAt           interface{} `json:"unlock_at"`
	Locked             bool        `json:"locked"`
	Hidden             bool        `json:"hidden"`
	LockAt             interface{} `json:"lock_at"`
	HiddenForUser      bool        `json:"hidden_for_user"`
	ThumbnailURL       interface{} `json:"thumbnail_url"`
	ModifiedAt         time.Time   `json:"modified_at"`
	MimeClass          string      `json:"mime_class"`
	MediaEntryID       interface{} `json:"media_entry_id"`
	LockedForUser      bool        `json:"locked_for_user"`
	CanvadocSessionURL string      `json:"canvadoc_session_url"`
	CrocodocSessionURL interface{} `json:"crocodoc_session_url"`
}

// Folder contains information relating to folders on canvas
type Folder struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Position  int    `json:"position"`
	Indent    int    `json:"indent"`
	Type      string `json:"type"`
	ModuleID  int    `json:"module_id"`
	HTMLURL   string `json:"html_url"`
	ContentID int    `json:"content_id"`
	URL       string `json:"url"`
}

// Requester is a structure used in the http request to contain related data
type Requester struct {
	Context string
	Headers map[string]string
	BaseURL string
	Ignore  []string
}

// Status returned instead of structured response
type Status struct {
	Status string `json:"status"`
	Errors []struct {
		Message string `json:"message"`
	} `json:"errors"`
}

// Course is the toplevel struct containing all data related to an individual Course
type Course struct {
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

// Module is a struct containing subset of course information known as a 'module'
type Module struct {
	ID                        int           `json:"id"`
	Name                      string        `json:"name"`
	Position                  int           `json:"position"`
	UnlockAt                  interface{}   `json:"unlock_at"`
	RequireSequentialProgress bool          `json:"require_sequential_progress"`
	PublishFinalGrade         bool          `json:"publish_final_grade"`
	PrerequisiteModuleIds     []interface{} `json:"prerequisite_module_ids"`
	State                     string        `json:"state"`
	CompletedAt               time.Time     `json:"completed_at"`
	ItemsCount                int           `json:"items_count"`
	ItemsURL                  string        `json:"items_url"`
}

type NoModulesError struct {
	Course string
}

func (e *NoModulesError) Error() string {
	return fmt.Sprintf("%s, does not use the modules page\n", e.Course)
}

type NoFilesError struct {
	Course string
}

func (e *NoFilesError) Error() string {
	return fmt.Sprintf("course, %s, does not seem to have any files publicly available\n", strings.ReplaceAll(e.Course, " ", ""))
}

func GetCourses(r Requester, spec []string) ([]Course, error) {
	if len(r.Headers) == 0 {
		return nil, errors.New("empty headers")
	}
	req, err := http.NewRequest("GET", "https://"+r.BaseURL+r.Context, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", r.Headers["Authorization"])
	println("https://" + r.BaseURL + r.Context)
	println("Executing request")
	resp, err := http.DefaultClient.Do(req)
	println("Executed.")
	if err != nil {
		println("Error in request")
		return nil, err
	}
	defer resp.Body.Close()
	courses := make([]Course, 0)
	println("Reading Courses")
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &courses)

	if err != nil {
		fmt.Printf("%s, \n", body)
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}
	ret := make([]Course, 0)
	if len(spec) > 0 {
		println("filtering discovered courses")
		for _, course := range courses {
			for _, specifiedCourse := range spec {
				if strings.ToLower(strings.ReplaceAll(course.Name, " ", "")) == strings.ToLower(specifiedCourse) {
					ret = append(ret, course)
				}
			}
		}
	} else {
		ret = courses
	}
	return ret, nil
}

// Download downloads files to a given filepath from a given URL using data in a Requester Struct
func (file *File) Download(course Course, r Requester) {
	if file.URL == "" {
		log.Println(errors.New("no file URL"))
		return
	}
	filepath := strings.ReplaceAll("out/"+course.Name+"/"+file.Filename, " ", "")
	tmp := strings.Split(filepath, ".")
	fileExt := tmp[len(tmp)-1]
	for _, ext := range r.Ignore {
		if ext == fileExt {
			return
		}
	}
	// Get the data
	req, err := http.NewRequest("GET", file.URL, nil)
	if err != nil {
		log.Println(err)
		return
	}
	req.Header.Add("Authorization", r.Headers["Authorization"])
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)

	if err != nil {
		log.Println(err)
		return
	}
	defer out.Close()

	// // Write the body to file

	_, err = io.Copy(out, resp.Body)
	err = resp.Body.Close()
	if err != nil {
		return
		log.Println(err)
	}
	err = out.Close()
	if err != nil {
		return
		log.Println(err)
	}
	return
}

func (course *Course) GetFiles(r Requester) error {
	fmt.Printf("Looking for files in course, %s \n", course.Name)
	var unmarshalTypeError *json.UnmarshalTypeError
	if outputDir == "" {
		_ = os.MkdirAll("out/"+strings.ReplaceAll(course.Name, " ", ""), 0777)
	} else {
		_ = os.MkdirAll(outputDir+"/"+strings.ReplaceAll(course.Name, " ", ""), 0777)
	}
	req, err := http.NewRequest("GET", "https://"+r.BaseURL+r.Context+strconv.Itoa(course.ID)+"/files/?per_page=1000", nil)
	if err != nil {
		return err
	}
	files := make([]File, 0)

	req.Header.Add("Authorization", r.Headers["Authorization"])
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &files)
	if err != nil {
		switch {
		case errors.As(err, &unmarshalTypeError):
			status := Status{}
			e2 := json.Unmarshal(body, &status)
			if e2 != nil {
				println("Cannot unmarshal response")
				return e2
			} else {
				if status.Status == "unauthorised" {
					return &NoFilesError{course.Name}
				}
			}
		}
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return &NoFilesError{course.Name}
	}

	for _, file := range files {
		var filename string
		if outputDir == "" {
			filename = strings.ReplaceAll("out/"+course.Name+"/"+file.Filename, " ", "")
		} else {
			filename = strings.ReplaceAll(outputDir+"/"+course.Name+"/"+file.Filename, " ", "")
		}
		_, err = os.Stat(filename)
		if forceDownloadAll || os.IsNotExist(err) {
			file.Download(*course, r)
		}
	}

	return nil
}

func (course *Course) GetModules(r Requester) ([]Module, error) {
	if outputDir == "" {
		_ = os.Mkdir("out/"+strings.ReplaceAll(course.Name, " ", ""), 0777)
	} else {
		_ = os.Mkdir(outputDir+"/"+strings.ReplaceAll(course.Name, " ", ""), 0777)
	}
	req, err := http.NewRequest("GET", "https://"+r.BaseURL+r.Context+strconv.Itoa(course.ID)+"/modules/?per_page=1000", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", r.Headers["Authorization"])
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	modules := make([]Module, 0)
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &modules)
	if err != nil {
		return nil, err
	}
	if len(modules) == 0 {
		fmt.Printf("course, %s, does not use modules page \n ", strings.ReplaceAll(course.Name, " ", ""))
		return nil, &NoModulesError{course.Name}

	} else {
		return modules, nil
	}

}
func (module *Module) GetFolders(r Requester) ([]Folder, error) {

	req, _ := http.NewRequest("GET", module.ItemsURL, nil)
	req.Header.Add("Authorization", r.Headers["Authorization"])
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	folders := make([]Folder, 0)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &folders)
	if err != nil {
		return nil, err
	}
	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	return folders, nil

}

func (folder *Folder) GetFiles(r Requester, course Course) error {
	if (folder.URL != "") && !strings.Contains(folder.URL, "/quizzes/") {
		req, _ := http.NewRequest("GET", folder.URL, nil)
		req.Header.Add("Authorization", r.Headers["Authorization"])
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		var file File
		body, err := ioutil.ReadAll(resp.Body)
        if strings.Contains(folder.URL, "/pages/"){
            var page Page
            err = json.Unmarshal(body, &page)
            //fmt.Printf("Page: %v\n",page)
            body_list := strings.Split(page.Body," ")
            urls := make([]string,1)
            for _,each := range body_list{
                if strings.Contains(each,"canvas.bham.ac.uk"){
                    urls = append(urls, each)
                }
            }
            //fmt.Printf("DETECTED URLS: %v\n", urls)
            for _, url := range urls{
                if strings.Contains(url, "files") && strings.Contains(url, "api"){
                    url_re := regexp.MustCompile(`https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
                    file_url := url_re.FindString(url)

                    req, err := http.NewRequest("GET", file_url, nil)
                    if err != nil {
                        log.Fatal(err)
                    }
                    req.Header.Add("Authorization", r.Headers["Authorization"])
                    resp, err := http.DefaultClient.Do(req)
                    if err != nil {
                        return err
                    }
		            body, err := ioutil.ReadAll(resp.Body)
                    if err != nil {
                        return err
                    }
                    json.Unmarshal(body, &file)
                    err = resp.Body.Close()
                    if err != nil {
                        return err
                    }
                    fmt.Printf("Downloading file: %v\n", file.DisplayName)
                    filename := strings.ReplaceAll("out/"+course.Name+"/"+file.Filename, " ", "")
                    _, err = os.Stat(filename)
                    if forceDownloadAll || os.IsNotExist(err) {
                        go file.Download(course, r)
                        if err != nil {
                            return err
                        }
                    }



                }

            }

        }else{
            if err != nil {
                return err
            }
            err = json.Unmarshal(body, &file)
            if err != nil {
                fmt.Printf("%s cannot be unmarshalled from folder", body)
                return err
            }
            err = resp.Body.Close()
            if err != nil {
                return err
            }
            filename := strings.ReplaceAll("out/"+course.Name+"/"+file.Filename, " ", "")
            _, err = os.Stat(filename)
            if forceDownloadAll || os.IsNotExist(err) {
                go file.Download(course, r)
                if err != nil {
                    return err
                }
            }
        }
    }
	return nil
}

var (
	forceDownloadAll bool
	outputDir        string
)

func ReadConfig() (*viper.Viper, error) {
	fmt.Println("Reading config file")

	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetDefault("AuthToken", "token")
	err := v.ReadInConfig()
	return v, err
}

func GetRequester() (Requester, error) {
	fmt.Println("listing modules")
	baseUrl := "canvas.bham.ac.uk"
	config, err := ReadConfig()
	if err != nil {
		return Requester{}, err
	}
	authToken := config.GetString("AuthToken")

	headers := make(map[string]string)
	fmt.Println("Using lib.AuthToken, %s", authToken)
	fmt.Println("", authToken)

	headers["Authorization"] = "Bearer " + authToken

	dat, err := ioutil.ReadFile(".scrapeignore")
	if err != nil {
		log.Fatal(err)
	}
	strData := string(dat)
	ignore := strings.Split(strData, "\n")
	ignore = ignore[:len(ignore)-1] // remove last empty value
	println("Ignoring the following extensions:\n " + strData)

	requester := Requester{
		Context: "/api/v1/courses?per_page=1000",
		Headers: headers,
		BaseURL: baseUrl,
		Ignore:  ignore,
	}

	return requester, nil
}
