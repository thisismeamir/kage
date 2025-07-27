package atom

import "github.com/thisismeamir/kage/pkg/atom"

type ModuleModel struct {
	Model        ModuleMetadata     `json:"model"`
	Execution    ModuleExecution    `json:"execution"`
	Dependencies []ModuleDependency `json:"dependencies"`
	Graph        []Vertices         `json:"graph"`
	Attachments  []atom.AtomModel   `json:"attachments,omitempty"`
}

type ModuleMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     string `json:"version"`
	Identifier  string `json:"identifier"`
	Author      string `json:"author"`
	AuthorEmail string `json:"author_email"`
	AuthorURL   string `json:"author_url"`
	Manual      string `json:"manual"`
	Repository  string `json:"repository"`
}

type ModuleExecution struct {
	Priority     int                         `json:"priority"`
	Policy       string                      `json:"policy"`
	Timeout      int                         `json:"timeout"`
	Retry        ModuleExecutionRetry        `json:"retry"`
	Access       ModuleExecutionAccess       `json:"access"`
	PreProcess   string                      `json:"pre_process"`
	PostProcess  string                      `json:"post_process"`
	OnError      string                      `json:"on_error"`
	OnFailure    string                      `json:"on_failure"`
	Notification ModuleExecutionNotification `json:"notification"`
}

type Vertices struct {
	From string `json:"from"`
	To   string `json:"to"`
	Bond string `json:"bond"`
}

type ModuleDependency struct {
	Type       string `json:"type"`
	Identifier string `json:"identifier"`
	Map        []Bond `json:"map,omitempty"`
}

type Bond struct {
	From string `json:"from"`
	To   string `json:"to"`
}
type ModuleExecutionRetry struct {
	MaxAttempts int `json:"max_attempts"`
	Delay       int `json:"delay"`
}

type ModuleExecutionAccess struct {
	FileSystemAccess ModuleExecutionAccessFileSystem `json:"file_system"`
	WebAccess        ModuleExecutionAccessWeb        `json:"web"`
}

type ModuleExecutionAccessFileSystem struct {
	Read    bool `json:"read"`
	Write   bool `json:"write"`
	Execute bool `json:"execute"`
}

type ModuleExecutionAccessWeb struct {
	Enabled     bool     `json:"enabled"`
	AllowedUrls []string `json:"allowed_urls"`
	BlockedUrls []string `json:"blocked_urls"`
}

type ModuleExecutionNotification struct {
	Email    ModuleExecutionNotificationEmail    `json:"email"`
	Webhooks []string                            `json:"webhooks"`
	Desktops ModuleExecutionNotificationDesktop  `json:"desktops"`
	Telegram ModuleExecutionNotificationTelegram `json:"telegram"`
	Slack    ModuleExecutionNotificationSlack    `json:"slack"`
	Discord  ModuleExecutionNotificationDiscord  `json:"discord"`
	LogLevel string                              `json:"log_level"`
}

type ModuleExecutionNotificationEmail struct {
	Enabled    bool     `json:"enabled"`
	Recipients []string `json:"recipients"`
}

type ModuleExecutionNotificationDesktop struct {
	Enabled          bool   `json:"enabled"`
	NotificationType string `json:"notification_type"`
	ComputerName     string `json:"computer_name"`
}

type ModuleExecutionNotificationTelegram struct {
	Enabled bool   `json:"enabled"`
	ChatID  string `json:"chat_id"`
}

type ModuleExecutionNotificationSlack struct {
	Enabled bool   `json:"enabled"`
	Webhook string `json:"webhook"`
}

type ModuleExecutionNotificationDiscord struct {
	Enabled bool   `json:"enabled"`
	Webhook string `json:"webhook"`
}
