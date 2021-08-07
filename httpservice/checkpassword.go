package pepinohttpservice

func (r *request) checkPassword(pwd string) bool {
	return pwd == r.dbHTTPService.config.Password
}
