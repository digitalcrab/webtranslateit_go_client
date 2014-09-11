package webtranslateit_go_client

import (
	"path/filepath"
	"io/ioutil"
	"testing"
	"fmt"
	"os"
)

const TEST_PROJECT_TOKEN = "98e71ee45042066f1053ed900b4e8f4ec1f98451"

func TestNewWebTranslateIt(t *testing.T) {
	wti := NewWebTranslateIt(TEST_PROJECT_TOKEN)
	expected := fmt.Sprintf("%T", &WebTranslateIt{})

	if fmt.Sprintf("%T", wti) != expected {
		t.Errorf("Error while creating new instance of WebTranslateIt expected type %s but got %T", expected, wti)
	} else if wti.GetToken() != TEST_PROJECT_TOKEN {
		t.Errorf("Error creating new instance of WebTranslateIt expected token %q but got %q", TEST_PROJECT_TOKEN, wti.GetToken())
	} else if wti.GetApiUrl() != DEFAULT_API_URL {
		t.Errorf("Error creating new instance of WebTranslateIt expected API url %q but got %q", DEFAULT_API_URL, wti.GetApiUrl())
	}
}

func TestGetProject(t *testing.T) {
	wti := NewWebTranslateIt(TEST_PROJECT_TOKEN)
	project, err := wti.GetProject()

	if err != nil {
		t.Errorf("Error getting project: %v", err)
	} else if project.Name != "WebTranslateIt" {
		t.Errorf("Received wrong project from API, expected name WebTranslateIt but got %q", project.Name)
	}
}

func TestGetProjectWithError(t *testing.T) {
	wti := NewWebTranslateIt("empty")
	if project, err := wti.GetProject(); err == nil {
		t.Errorf("Request of Empty project should be an error, but got %v", project)
	}
}

func TestGetZip(t *testing.T) {
	wti := NewWebTranslateIt(TEST_PROJECT_TOKEN)
	project, err := wti.GetProject()

	if err != nil {
		t.Errorf("Error getting project: %v", err)
	} else if project.Name != "WebTranslateIt" {
		t.Errorf("Received wrong project from API, expected name WebTranslateIt but got %q", project.Name)
	} else if zipFile, err := project.ZipFile(); err != nil {
		t.Errorf("Error getting zip file of the project %q: %v", project.Name, err)
	} else if zipFile.Size() == 0 {
		t.Errorf("Zip file is empty")
	} else {
		var (
			err		error
			path	string
			file	*os.File
			data	map[string][]byte
		)
		// Save as *os.File
		file, err = ioutil.TempFile(os.TempDir(), "webtranslateit_go_client")
		if err != nil {
			t.Fatalf("Error creating temp file %v", err)
		}

		// Save file
		if err = zipFile.SaveToFile(file); err != nil {
			file.Close()
			t.Fatalf("Error saving zip file to *os.File %q %v", file.Name(), err)
		}
		file.Close()

		// Remove temp file
		if err = os.Remove(file.Name()); err != nil {
			t.Fatalf("Error removing temp file %q %v", file.Name(), err)
		}

		// Save to path
		path = filepath.Join(os.TempDir(), "webtranslateit_go_client_TestGetZip.zip")
		if err = zipFile.SaveToPath(path); err != nil {
			t.Fatalf("Error saving zip file to path %q %v", path, err)
		}

		// Remove temp file
		if err = os.Remove(path); err != nil {
			t.Fatalf("Error removing temp file %q %v", path, err)
		}

		data, err = zipFile.Extract()
		if err != nil {
			t.Errorf("Error extracting zip file %v", err)
		}

		for _, f := range project.ProjectFiles {
			if _, ok := data[f.Name]; !ok {
				t.Errorf("Among the extracted files no file %q", f.Name)
			}
		}
	}
}

func TestGetZipWithEmptyWti(t *testing.T) {
	project := &Project{}
	if _, err := project.ZipFile(); err != WtiNil {
		t.Errorf("Request of ZipFile of the project should be an error, but got %v", project)
	}
}
