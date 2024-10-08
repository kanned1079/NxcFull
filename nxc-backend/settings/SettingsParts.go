package settings

type SystemSettingOptions struct {
	Site      Site      `json:"site"`
	Security  Security  `json:"security"`
	Frontend  Frontend  `json:"frontend"`
	Subscribe Subscribe `json:"subscription"`
	Server    Server    `json:"server"`
	Sendmail  Sendmail  `json:"sendmail"`
	Notice    Notice    `json:"notice"`
	Myapp     Myapp     `json:"myapp"`
}
type Site struct {
	AppName        string `json:"app_name"`
	AppDescription string `json:"app_description"`
	AppURL         string `json:"app_url"`
	ForceHTTPS     bool   `json:"force_https"`
	LogoURL        string `json:"logo_url"`
	SubscribeURL   int    `json:"subscribe_url"`
	TosURL         string `json:"tos_url"`
	StopRegister   bool   `json:"stop_register"`
	TrialTime      int    `json:"trial_time"`
	TrialSubscribe string `json:"trial_subscribe"`
	Currency       string `json:"currency"`
	CurrencySymbol string `json:"currency_symbol"`
}
type Security struct {
	EmailVerify           bool   `json:"email_verify"`
	EmailGmailLimitEnable bool   `json:"email_gmail_limit_enable"`
	SafeModeEnable        bool   `json:"safe_mode_enable"`
	SecurePath            string `json:"secure_path"`
	EmailWhitelistEnable  bool   `json:"email_whitelist_enable"`
	RecaptchaEnable       bool   `json:"recaptcha_enable"`
	IPRegisterLimitEnable bool   `json:"ip_register_limit_enable"`
	IPRegisterLimitTimes  int    `json:"ip_register_limit_times"`
	IPRegisterLockTime    int    `json:"ip_register_lock_time"`
}
type Frontend struct {
	FrontendThemeSidebar  bool   `json:"frontend_theme_sidebar"`
	FrontendThemeHeader   bool   `json:"frontend_theme_header"`
	FrontendTheme         string `json:"frontend_theme"`
	FrontendBackgroundURL string `json:"frontend_background_url"`
}
type Subscribe struct {
	UserModifyEnable bool `json:"user_modify_enable"`
	ShowInfoInSub    bool `json:"show_info_in_sub"`
}
type Server struct {
	ServerToken        string `json:"server_token"`
	ServerPullInterval int    `json:"server_pull_interval"`
	ServerPushInterval int    `json:"server_push_interval"`
}
type Sendmail struct {
	EmailHost        string `json:"email_host"`
	EmailPort        int    `json:"email_port"`
	EmailEncryption  string `json:"email_encryption"`
	EmailUsername    string `json:"email_username"`
	EmailPassword    string `json:"email_password"`
	EmailFromAddress string `json:"email_from_address"`
	EmailTemplate    string `json:"email_template"`
}
type Notice struct {
	NoticeName string `json:"notice_name"`
	BarkHost   string `json:"bark_host"`
	BarkGroup  string `json:"bark_group"`
}
type Myapp struct {
	WinDownload     string `json:"win_download"`
	OsxDownload     string `json:"osx_download"`
	AndroidDownload string `json:"android_download"`
}
