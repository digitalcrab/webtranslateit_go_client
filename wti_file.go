package webtranslateit_go_client

type File struct {
	Id				uint	`json:"id"`
	Name			string	`json:"name"`
	CreatedAt		string	`json:"created_at"`
	UpdatedAt		string	`json:"updated_at"`
	Hash			string	`json:"hash_file"`
	MasterFileId	uint	`json:"master_project_file_id,omitempty"`
	LocaleCode		string	`json:"locale_code"`

	project			*Project
}
