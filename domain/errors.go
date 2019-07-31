package domain

type CreateIssueError struct {
	Msg string
}

func (e *CreateIssueError) Error() string {
	return "Issueの作成に失敗しました。"
}
