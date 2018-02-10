package easy_server

import(
	"log"
	"io"
	"os"
)

type EasyLogger struct{
     sysinfoLogger * log.Logger
     debugLogger * log.Logger
     warnLogger * log.Logger
     errorLogger * log.Logger
}

/*Global logger instance using os.Stdout by default*/
var Logger * EasyLogger = newEasyLogger(os.Stdout,os.Stdout,os.Stdout,os.Stdout)

/*Set new logger using self-specified files*/
func SetEasyLogger(sysfile,debugfile,warnfile,errfile io.Writer) {
     Logger = newEasyLogger(sysfile,debugfile,warnfile,errfile)
}


func newEasyLogger(sysfile,debugfile,warnfile,errfile io.Writer) * EasyLogger{
     el := EasyLogger{}
     if sysfile!= nil {
     	sLogger := log.New(sysfile,"[SYSINFO]",log.LstdFlags)
	el.sysinfoLogger = sLogger
     }

     if debugfile != nil{
     	dLogger := log.New(debugfile,"[DEBUG]",log.LstdFlags)
	el.debugLogger = dLogger
     }

     if warnfile != nil {
     	wLogger := log.New(warnfile,"[WARN]",log.LstdFlags)
	el.warnLogger = wLogger
     }

     if errfile != nil {
        eLogger := log.New(errfile,"[ERROR]",log.LstdFlags)
	el.errorLogger = eLogger
     }
     
    return &el
        
}

func (log *EasyLogger) SysLog(s ...interface{}){
     if log.sysinfoLogger != nil {
     	log.sysinfoLogger.Println(s)
     }
     
}

func (log *EasyLogger) DebugLog(s ...interface{}){
     if log.debugLogger != nil {
     	log.debugLogger.Println(s)
     }
     
}

func (log *EasyLogger) WarnLog(s ...interface{}){
     if log.warnLogger != nil {
     	log.warnLogger.Println(s)
     }
     
}

func (log *EasyLogger) ErrorLog(s ...interface{}){
     if log.errorLogger != nil {
     	log.errorLogger.Println(s)
     }
     
}



