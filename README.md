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

so I tried to implement a binary file with `go` that you can execute in you shell to use it, just for fun. ( thanks jadi :) )

#### usage:
- first clone the repo
- change the mode of the repo so it can be executed  
`chmod +x /your/dir/than/contains/the/app/repo`
- cd to the app dir  
`cd /cloned_users_cli/app`
- build the binary of the app  
  `go build main.go`
- now you have the executable file in your dir
- you can use it like this:  
`./main clonedUsers`

#### flags:
I wanted it to be a little bit fun. so I made two flags for you that you can use  
- first one is `--user` that you can specify that you want the logged in users to print in stdout
- second one is `--timestamp` that you can specify for your users timestamps
- third is the help text with `--help` that shows the what this command do and what is it for 🤠
- i can make it more fun with a lot of flags and os signals if you liked it 🤠


#### cobra-cli
- i've user cobra-cli to make this command line interface and the license. so a big thank to spf13
- and more important, big thanks to jadi to make such a ADINE for me  
  #### love you jadi❤️