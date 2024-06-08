### cloned version of `users` command in mac Os

#### brief intorduction to how it works:

as you may know, there is a file in /var/run/utmpx (in POSIX) and /var/run/utmp (in OpenBSD) that the `users` command read to taggle the information about the users that are logged in to the system.
as i'm in macOS, i'll show you the struct of each entry in utmpx file for each user that contains information such as username, id, line, host, PID, Type, Typestr and timestamp.

#### file's format for each entry:

```go
type Entry struct {
	User      string    `json:"user"`
	ID        string    `json:"id"`
	Line      string    `json:"line"`
	Host      string    `json:"host"`
	PID       int       `json:"pid"`
	Type      Type      `json:"type"`
	TypeStr   string    `json:"typestr"`
	Timestamp time.Time `json:"timestamp"`
}
```
as well there is a utmpx package in Go. that if you are interested in the interface and types
represented by the package you can visit [utmpx package repo](https://github.com/stephane-martin/skewer/tree/2081a449f6b4/sys/utmpx).  
trust me it worth it to have a look 🤝

so I tried to implement a binary file with `go` that you can execute and use it, just for fun.

#### usage:
- first clone the repo
- change the mode of the repo so it can be executed  
`chmod +x /your/dir/that/contains/the/app/repo`
- cd to the app dir  
`cd /cloned_users_cli/app`
- build the binary of the app  
  `go build main.go`
- now you have the executable file in your dir
- you can use it like this:  
`./main clonedUsers`
- or with the flags like so  
`./main clonedUsers --users --timestamp`

#### flags:
I wanted it to be a little bit fun. so I made two flags for you that you can use  
- first one is `--user` that you can specify that you want the logged in users to print in stdout
- second one is `--timestamp` that you can specify for your user's timestamps
- third is the help text with `--help` or the abbreviated version one, `-h` that shows what this command do and what is it for 🤠

#### usage of the flags:
- `./main clonedUsers --users --timestamp` will show you the users and also the timestamp of each.


#### Notes:
- i've used cobra-cli to make this command line interface and the license. so a big thank to spf13, and as always it worth it to take a look at [cobra package](https://github.com/spf13/cobra-cli).
- and more important, big thanks to jadi to make such a ADINE for me. [jadi's youtube](https://www.youtube.com/@JadiMirmirani) is the best OpenSource project(!!) that i've encounter 🤝

### it could get more fun and i try to make it better in my weekends. so let me know if it was fun for you too 🤝

#### again, thanks & love you jadi❤️