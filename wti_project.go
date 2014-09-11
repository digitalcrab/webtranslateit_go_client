package webtranslateit_go_client

import "fmt"

type Project struct {
	Id				uint		`json:"id"`
	Name			string		`json:"name"`
	CreatedAt		string		`json:"created_at"`
	UpdatedAt		string		`json:"updated_at"`
	SourceLocale	Locale		`json:"source_locale"`
	TargetLocales	[]Locale	`json:"target_locales"`
	ProjectFiles	[]File		`json:"project_files"`

	wti				*WebTranslateIt
}

func (self *Project) SetWti(wti *WebTranslateIt) *Project {
	self.wti = wti
	return self
}

func (self *Project) GetWti() *WebTranslateIt {
	return self.wti
}

func (self *Project) ZipFile() (ProjectZipFile, error) {
	if self.wti == nil {
		return nil, WtiNil
	}

	// Get zip file
	data, err := self.wti.requestGet(fmt.Sprintf("%sprojects/%s/zip_file", self.wti.GetApiUrl(), self.wti.GetToken()))
	if err != nil {
		return nil, err
	}

	// Create ZipFile structure
	return ProjectZipFile(data), nil
}
