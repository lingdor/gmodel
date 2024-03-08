package mylog

type Logger interface{
	Printf(format string, v ...any)
	Println(v ...any)
}
