package initialize

import (
	"fmt"

	"go.uber.org/zap"
)

func loggerInit() {
	/*cfg := zap.NewDevelopmentConfig()
	logfile := fmt.Sprintf("./logfiles/%s.log", time.Now().Format(global.TimeTemplateSec))
	cfg.OutputPaths = []string{"stderr", logfile}
	logger, err := cfg.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))*/
	logger, err := zap.NewDevelopment(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		panic(fmt.Errorf("fatal error logger: %s", err.Error()))
	}
	zap.ReplaceGlobals(logger)
}
