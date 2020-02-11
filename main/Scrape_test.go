package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCourse_getFiles(t *testing.T) {
	type fields struct {
		ID                          int
		Name                        string
		AccountID                   int
		UUID                        string
		StartAt                     time.Time
		GradingStandardID           interface{}
		IsPublic                    bool
		CreatedAt                   time.Time
		CourseCode                  string
		DefaultView                 string
		RootAccountID               int
		EnrollmentTermID            int
		License                     string
		GradePassbackSetting        interface{}
		EndAt                       interface{}
		PublicSyllabus              bool
		PublicSyllabusToAuth        bool
		StorageQuotaMb              int
		IsPublicToAuthUsers         bool
		ApplyAssignmentGroupWeights bool
		Calendar                    struct {
			Ics string `json:"ics"`
		}
		TimeZone    string
		Blueprint   bool
		Enrollments []struct {
			Type                           string `json:"type"`
			Role                           string `json:"role"`
			RoleID                         int    `json:"role_id"`
			UserID                         int    `json:"user_id"`
			EnrollmentState                string `json:"enrollment_state"`
			LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
		}
		HideFinalGrades                  bool
		WorkflowState                    string
		RestrictEnrollmentsToCourseDates bool
		OverriddenCourseVisibility       string
		Locale                           string
	}
	type args struct {
		r Requester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			course := &Course{
				ID:                               tt.fields.ID,
				Name:                             tt.fields.Name,
				AccountID:                        tt.fields.AccountID,
				UUID:                             tt.fields.UUID,
				StartAt:                          tt.fields.StartAt,
				GradingStandardID:                tt.fields.GradingStandardID,
				IsPublic:                         tt.fields.IsPublic,
				CreatedAt:                        tt.fields.CreatedAt,
				CourseCode:                       tt.fields.CourseCode,
				DefaultView:                      tt.fields.DefaultView,
				RootAccountID:                    tt.fields.RootAccountID,
				EnrollmentTermID:                 tt.fields.EnrollmentTermID,
				License:                          tt.fields.License,
				GradePassbackSetting:             tt.fields.GradePassbackSetting,
				EndAt:                            tt.fields.EndAt,
				PublicSyllabus:                   tt.fields.PublicSyllabus,
				PublicSyllabusToAuth:             tt.fields.PublicSyllabusToAuth,
				StorageQuotaMb:                   tt.fields.StorageQuotaMb,
				IsPublicToAuthUsers:              tt.fields.IsPublicToAuthUsers,
				ApplyAssignmentGroupWeights:      tt.fields.ApplyAssignmentGroupWeights,
				Calendar:                         tt.fields.Calendar,
				TimeZone:                         tt.fields.TimeZone,
				Blueprint:                        tt.fields.Blueprint,
				Enrollments:                      tt.fields.Enrollments,
				HideFinalGrades:                  tt.fields.HideFinalGrades,
				WorkflowState:                    tt.fields.WorkflowState,
				RestrictEnrollmentsToCourseDates: tt.fields.RestrictEnrollmentsToCourseDates,
				OverriddenCourseVisibility:       tt.fields.OverriddenCourseVisibility,
				Locale:                           tt.fields.Locale,
			}
			if err := course.getFiles(tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("getFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCourse_getModules(t *testing.T) {
	type fields struct {
		ID                          int
		Name                        string
		AccountID                   int
		UUID                        string
		StartAt                     time.Time
		GradingStandardID           interface{}
		IsPublic                    bool
		CreatedAt                   time.Time
		CourseCode                  string
		DefaultView                 string
		RootAccountID               int
		EnrollmentTermID            int
		License                     string
		GradePassbackSetting        interface{}
		EndAt                       interface{}
		PublicSyllabus              bool
		PublicSyllabusToAuth        bool
		StorageQuotaMb              int
		IsPublicToAuthUsers         bool
		ApplyAssignmentGroupWeights bool
		Calendar                    struct {
			Ics string `json:"ics"`
		}
		TimeZone    string
		Blueprint   bool
		Enrollments []struct {
			Type                           string `json:"type"`
			Role                           string `json:"role"`
			RoleID                         int    `json:"role_id"`
			UserID                         int    `json:"user_id"`
			EnrollmentState                string `json:"enrollment_state"`
			LimitPrivilegesToCourseSection bool   `json:"limit_privileges_to_course_section"`
		}
		HideFinalGrades                  bool
		WorkflowState                    string
		RestrictEnrollmentsToCourseDates bool
		OverriddenCourseVisibility       string
		Locale                           string
	}
	type args struct {
		r Requester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Module
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			course := &Course{
				ID:                               tt.fields.ID,
				Name:                             tt.fields.Name,
				AccountID:                        tt.fields.AccountID,
				UUID:                             tt.fields.UUID,
				StartAt:                          tt.fields.StartAt,
				GradingStandardID:                tt.fields.GradingStandardID,
				IsPublic:                         tt.fields.IsPublic,
				CreatedAt:                        tt.fields.CreatedAt,
				CourseCode:                       tt.fields.CourseCode,
				DefaultView:                      tt.fields.DefaultView,
				RootAccountID:                    tt.fields.RootAccountID,
				EnrollmentTermID:                 tt.fields.EnrollmentTermID,
				License:                          tt.fields.License,
				GradePassbackSetting:             tt.fields.GradePassbackSetting,
				EndAt:                            tt.fields.EndAt,
				PublicSyllabus:                   tt.fields.PublicSyllabus,
				PublicSyllabusToAuth:             tt.fields.PublicSyllabusToAuth,
				StorageQuotaMb:                   tt.fields.StorageQuotaMb,
				IsPublicToAuthUsers:              tt.fields.IsPublicToAuthUsers,
				ApplyAssignmentGroupWeights:      tt.fields.ApplyAssignmentGroupWeights,
				Calendar:                         tt.fields.Calendar,
				TimeZone:                         tt.fields.TimeZone,
				Blueprint:                        tt.fields.Blueprint,
				Enrollments:                      tt.fields.Enrollments,
				HideFinalGrades:                  tt.fields.HideFinalGrades,
				WorkflowState:                    tt.fields.WorkflowState,
				RestrictEnrollmentsToCourseDates: tt.fields.RestrictEnrollmentsToCourseDates,
				OverriddenCourseVisibility:       tt.fields.OverriddenCourseVisibility,
				Locale:                           tt.fields.Locale,
			}
			got, err := course.getModules(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getModules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getModules() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Download(t *testing.T) {
	type fields struct {
		ID                 int
		UUID               string
		FolderID           int
		DisplayName        string
		Filename           string
		UploadStatus       string
		ContentType        string
		URL                string
		Size               int
		CreatedAt          time.Time
		UpdatedAt          time.Time
		UnlockAt           interface{}
		Locked             bool
		Hidden             bool
		LockAt             interface{}
		HiddenForUser      bool
		ThumbnailURL       interface{}
		ModifiedAt         time.Time
		MimeClass          string
		MediaEntryID       interface{}
		LockedForUser      bool
		CanvadocSessionURL string
		CrocodocSessionURL interface{}
	}
	type args struct {
		course Course
		r      Requester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &File{
				ID:                 tt.fields.ID,
				UUID:               tt.fields.UUID,
				FolderID:           tt.fields.FolderID,
				DisplayName:        tt.fields.DisplayName,
				Filename:           tt.fields.Filename,
				UploadStatus:       tt.fields.UploadStatus,
				ContentType:        tt.fields.ContentType,
				URL:                tt.fields.URL,
				Size:               tt.fields.Size,
				CreatedAt:          tt.fields.CreatedAt,
				UpdatedAt:          tt.fields.UpdatedAt,
				UnlockAt:           tt.fields.UnlockAt,
				Locked:             tt.fields.Locked,
				Hidden:             tt.fields.Hidden,
				LockAt:             tt.fields.LockAt,
				HiddenForUser:      tt.fields.HiddenForUser,
				ThumbnailURL:       tt.fields.ThumbnailURL,
				ModifiedAt:         tt.fields.ModifiedAt,
				MimeClass:          tt.fields.MimeClass,
				MediaEntryID:       tt.fields.MediaEntryID,
				LockedForUser:      tt.fields.LockedForUser,
				CanvadocSessionURL: tt.fields.CanvadocSessionURL,
				CrocodocSessionURL: tt.fields.CrocodocSessionURL,
			}
			if err := file.Download(tt.args.course, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("Download() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFolder_getFiles(t *testing.T) {
	type fields struct {
		ID        int
		Title     string
		Position  int
		Indent    int
		Type      string
		ModuleID  int
		HTMLURL   string
		ContentID int
		URL       string
	}
	type args struct {
		r      Requester
		course Course
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			folder := &Folder{
				ID:        tt.fields.ID,
				Title:     tt.fields.Title,
				Position:  tt.fields.Position,
				Indent:    tt.fields.Indent,
				Type:      tt.fields.Type,
				ModuleID:  tt.fields.ModuleID,
				HTMLURL:   tt.fields.HTMLURL,
				ContentID: tt.fields.ContentID,
				URL:       tt.fields.URL,
			}
			if err := folder.getFiles(tt.args.r, tt.args.course); (err != nil) != tt.wantErr {
				t.Errorf("getFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestModule_getFolders(t *testing.T) {
	type fields struct {
		ID                        int
		Name                      string
		Position                  int
		UnlockAt                  interface{}
		RequireSequentialProgress bool
		PublishFinalGrade         bool
		PrerequisiteModuleIds     []interface{}
		State                     string
		CompletedAt               time.Time
		ItemsCount                int
		ItemsURL                  string
	}
	type args struct {
		r Requester
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Folder
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			module := &Module{
				ID:                        tt.fields.ID,
				Name:                      tt.fields.Name,
				Position:                  tt.fields.Position,
				UnlockAt:                  tt.fields.UnlockAt,
				RequireSequentialProgress: tt.fields.RequireSequentialProgress,
				PublishFinalGrade:         tt.fields.PublishFinalGrade,
				PrerequisiteModuleIds:     tt.fields.PrerequisiteModuleIds,
				State:                     tt.fields.State,
				CompletedAt:               tt.fields.CompletedAt,
				ItemsCount:                tt.fields.ItemsCount,
				ItemsURL:                  tt.fields.ItemsURL,
			}
			got, err := module.getFolders(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFolders() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFolders() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoFilesError_Error(t *testing.T) {
	type fields struct {
		Course string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NoFilesError{
				Course: tt.fields.Course,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNoModulesError_Error(t *testing.T) {
	type fields struct {
		Course string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &NoModulesError{
				Course: tt.fields.Course,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCourses(t *testing.T) {
	type args struct {
		r    Requester
		spec []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Course
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCourses(tt.args.r, tt.args.spec)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCourses() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCourses() got = %v, want %v", got, tt.want)
			}
		})
	}
}
