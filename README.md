# snippetbox

## Flags
- at any point in the directory to of the `main.go` file type `go run <PATH_TO_MAIN.GO_FILE> -help` a list will appear of all flags available.
They are:

`-addr=":<NUMBER>"` - will specify the port on which to run 

`-dsn="<user>:<password>@/database?parseTime=true"` - parseTime will parse the time from SQL time and date field into Go time.Time objects

`-secret="<YOUR_SECRE_HERE>"` - flag for setting your secret sessions key

## launch app

`$ go run ./cmd/web` will launch with default values. So if you have a specific values make sure you set the correct values, alternatively you can edit the defaults in the `main.go` file.

Also as I haven't yet created a script for creating the DB with the correct db/tables so they will need to be created.