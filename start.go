package agollo

//start apollo
func Start(appConfig *AppConfig) error {
	return StartWithLogger(appConfig, nil)
}

func StartWithLogger(appConfig *AppConfig, loggerInterface LoggerInterface) error {
	return StartWithParams(appConfig, loggerInterface, nil)
}

func StartWithCache(appConfig *AppConfig, cacheInterface CacheInterface) error {
	return StartWithParams(appConfig, nil, cacheInterface)
}

func StartWithParams(appConfig *AppConfig, loggerInterface LoggerInterface, cacheInterface CacheInterface) error {
	if loggerInterface != nil {
		initLogger(loggerInterface)
	}
	if cacheInterface != nil {
		initCache(cacheInterface)
	}

	//初始化配置信息
	initConfig(func() (*AppConfig, error) {
		return appConfig, nil
	})
	//AllNotifications
	initAllNotifications()

	switch appConfig.Mode {
	case Local:
		//加载本地config
		config, _ := loadConfigFile(appConfig.BackupConfigPath)
		if config != nil {
			updateApolloConfig(config, false)
		}
		return nil
	default:
		//init server ip list
		go initServerIpList(appConfig)
		//first sync
		err := notifySyncConfigServices()
		//first sync fail then load config file
		if err != nil {
			config, _ := loadConfigFile(appConfig.BackupConfigPath)
			if config != nil {
				updateApolloConfig(config, false)
			}
		}
		//start long poll sync config
		go StartRefreshConfig(&NotifyConfigComponent{})

		logger.Info("agollo start finished , error:", err)
		return err
	}
}
