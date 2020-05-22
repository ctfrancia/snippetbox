# snippetbox

## Flags
- at any point in the directory to of the `main.go` file type `go run <PATH_TO_MAIN.GO_FILE> -help` a list will appear of all flags available.
They are:

`-addr=":<NUMBER>"` - will specify the port on which to run 

`-dsn="user:pass@/database?parseTime=true"` - parseTime will parse the time from SQL time and date field into Go time.Time objects