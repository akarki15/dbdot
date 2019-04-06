package flags

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	user      string
	dbname    string
	sslmode   bool
	WhiteList []string
}

// TODO(akarki15): Implement for multiple dbs
func (f Flags) DbToConnect() string {
	return "postgres"
}

func (f Flags) ConnString() string {
	sslmode := "disable"
	if f.sslmode {
		sslmode = "require"
	}
	flagStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", f.user, f.dbname, sslmode)
	return flagStr
}

func Parse() Flags {
	user := flag.String("user", "", "username of postgres db")
	dbname := flag.String("dbname", "", "dbname for which you want to generate dot file")
	sslmode := flag.Bool("sslmode", false, "enable sslmode for postgres db connection")
	whitelist := flag.String("whitelist", "", "comma separated list of tables you want to generate dot file for")
	flag.Parse()

	tempFlags := Flags{
		dbname:  *dbname,
		user:    *user,
		sslmode: *sslmode,
	}
	if len(*whitelist) > 0 {
		tempFlags.WhiteList = strings.Split(*whitelist, ",")
	}
	return tempFlags
}
