package configs

func GetMustEnvItemsForApp() []string {
	return []string{
		"FRONT_APP_URL",
		"DB_HOST",
		"DB_USER",
		"DB_PASSWORD",
		"DB_NAME",
		"DB_PORT",
		"TZ",
		"SERVER_PORT",
		"JWT_SECRET_KEY",
		"JWT_EXP",
		"LIMITER_MAX",
		"LIMITER_EXPIRATION",
	}
}
