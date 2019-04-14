package flags

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type Flags struct {
	user      string
	dbname    string
	dbpass    string
	hostname  string
	port      uint
	sslmode   bool
	Schema    string
	WhiteList []string
}

func Askpass(prompt string) (string, error) {
	fmt.Fprintf(os.Stderr, "%s: ", prompt)
	pw1, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Fprintf(os.Stderr, "\n")
	if err != nil {
		return "", err
	}
	return string(pw1), nil
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

	flagStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=%s", f.hostname, f.port, f.user, f.dbname, sslmode)
	if f.dbpass != "" {
		flagStr += fmt.Sprintf(" password='%s'", f.dbpass)
	}
	return flagStr
}

func Parse() Flags {
	user := flag.String("user", "", "username of postgres db")
	askForPass := flag.Bool("W", false, "ask for password")
	dbname := flag.String("dbname", "", "dbname for which you want to generate dot file")
	sslmode := flag.Bool("sslmode", false, "enable sslmode for postgres db connection")
	schema := flag.String("schema", "public", "schema name")
	hostname := flag.String("host", "localhost", "database host")
	port := flag.Uint("port", 5432, "database port")
	whitelist := flag.String("whitelist", "", "comma separated list of tables you want to generate dot file for")
	flag.Parse()

	tempFlags := Flags{
		dbname:   *dbname,
		user:     *user,
		hostname: *hostname,
		port:     *port,
		sslmode:  *sslmode,
		Schema:   *schema,
	}
	if len(*whitelist) > 0 {
		tempFlags.WhiteList = strings.Split(*whitelist, ",")
	}
	if *askForPass {
		p, err := Askpass(fmt.Sprintf("password for %s@%s", *user, *dbname))
		if err != nil {
			log.Fatalf("error reading password: %s", err.Error())
		}
		tempFlags.dbpass = p
	}

	return tempFlags
}
