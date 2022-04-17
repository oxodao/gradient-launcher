package model

type Workspace struct {
	ID            *int64 `db:"id"`
	LauncherName  string `json:"launcher_name"`
	WorkspacePath string `json:"workspace_path"`

	RootURL string `json:"root_url"`

	Modpacks []Modpack `json:"-"`
}
