package main

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test scan function to return the path of the files that have golang
// Test if it contains the start and end
// If it has the start only and no end
// If it has the end only but no start
func Test_ScanFilesWithExt_happy_scenarios(t *testing.T) {
	parentPathPrefix := "test_input_files/golang_file_templates/"
	childPathPrefix := "child_dir/"
	var tests = map[string]struct {
		path                  string
		pathsTolookFor        []string
		expectedReturnedPaths []string
		expectedErr           error
	}{
		"scan_files_with_extensions_in_arr_in_parent_dir": {
			path:                  "./test_input_files/golang_file_templates",
			pathsTolookFor:        []string{".go"},
			expectedReturnedPaths: []string{parentPathPrefix + childPathPrefix + "child_go_file.go", parentPathPrefix + "with_end_only.go", parentPathPrefix + "with_start_and_end.go", parentPathPrefix + "with_start_only.go"},
			expectedErr:           nil,
		},
		"scan_files_with_extensions_in_arr_in_child_dir": {
			path:                  "./test_input_files/golang_file_templates/child_dir",
			pathsTolookFor:        []string{".go"},
			expectedReturnedPaths: []string{parentPathPrefix + childPathPrefix + "child_go_file.go"},
			expectedErr:           nil,
		},
		"scan_files_with_extensions_not_in_arr": {
			path:                  "./test_input_files/golang_file_templates",
			pathsTolookFor:        []string{".abc"},
			expectedReturnedPaths: nil,
			expectedErr:           nil,
		},
	}
	for testName, testUsecase := range tests {

		t.Run(testName, func(t *testing.T) {
			actualReturn, _ := ScanFilesWithExt(testUsecase.path, testUsecase.pathsTolookFor)
			assert.Equal(t, testUsecase.expectedReturnedPaths, actualReturn)
		})
	}

}

func Test_ScanFilesWithExt_non_happy_scenarios(t *testing.T) {
	var tests = map[string]struct {
		path                  string
		pathsTolookFor        []string
		expectedReturnedPaths []string
		expectedErr           error
	}{
		"scan_files_with_non_existent_path": {
			path:                  "./invalidPath",
			pathsTolookFor:        []string{".abc"},
			expectedReturnedPaths: nil,
			expectedErr:           &fs.PathError{},
		},
	}
	for testName, testUsecase := range tests {

		t.Run(testName, func(t *testing.T) {
			actualReturn, err := ScanFilesWithExt(testUsecase.path, testUsecase.pathsTolookFor)
			assert.Nil(t, actualReturn)
			assert.NotNil(t, err)
		})
	}

}

func Test_FileWithinExtensions(t *testing.T) {
	var tests = map[string]struct {
		filePath        string
		extensionsInput []string
		expectedOutput  bool
	}{
		"filePath_extension_exists_in_Input": {
			filePath:        "abc.go",
			extensionsInput: []string{".go", ".md"},
			expectedOutput:  true,
		},
		"filePath_extension_doesnt_exist_in_Input": {
			filePath:        "abc.go",
			extensionsInput: []string{".txt", ".md"},
			expectedOutput:  false,
		},
	}
	for testName, testUsecase := range tests {
		t.Run(testName, func(t *testing.T) {
			actualOutput := FileWithinExtensions(testUsecase.filePath, testUsecase.extensionsInput)
			assert.Equal(t, testUsecase.expectedOutput, actualOutput)
		})
	}
}
