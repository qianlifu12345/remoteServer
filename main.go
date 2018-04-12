package main

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
)

//IPList is for listing out the domains .
type IPList struct {
	Ipaddress  string `json:"ipaddress,omitempty"`
	Privatekey string `json:"privatekey,omitempty"`
}

//ShellList is for listing out the domains .
type ShellList struct {
	Content string `json:"content,omitempty"`
}

func main() {
	if len(os.Args) != 4 {
		log.Fatalf("Usage: %s <user> <host:port> <shellid>", os.Args[0])
	}

	client, session, err := connectToHost(os.Args[1], os.Args[2])
	if err != nil {
		panic(err)
	}
	out, err := session.CombinedOutput(GetShellScript(os.Args[3]))
	if err != nil {
		panic(err)
	}
	fmt.Println(string(out))
	client.Close()
}

func connectToHost(user, host string) (*ssh.Client, *ssh.Session, error) {
	hostdata := strings.Split(host, ":")
	var hostip string

	if len(hostdata) != 2 {
		err := errors.New("cannot get host ip")
		hostip = ""
		return nil, nil, err

	}

	hostip = hostdata[0]
	fmt.Println(hostip)
	// fmt.Println(PublicKeyFile(hostip))

	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			PublicKeyFile(hostip),
		},
	}

	client, err := ssh.Dial("tcp", host, sshConfig)
	if err != nil {
		return nil, nil, err
	}

	session, err := client.NewSession()
	if err != nil {
		client.Close()
		return nil, nil, err
	}

	return client, session, nil
}

//PublicKeyFile is for implement the get PublicKeys
func PublicKeyFile(host string) ssh.AuthMethod {

	privatekey := GetPrivateKey(host)

	key, err := ssh.ParsePrivateKey([]byte(privatekey))
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//GetPrivateKey from database
func GetPrivateKey(ip string) string {
	db, err := sql.Open("mysql", "newuser:newuser@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		log.Errorf("Error connecting to database: %v", err)
	}
	query, err := db.Query(fmt.Sprintf("SELECT * FROM `remoteServer`.`key` WHERE `ipaddress` = '%s' ", ip))
	// fmt.Printf("SELECT * FROM `remoteServer`.`key` WHERE `ipaddress` = '%s' ", ip)
	if err != nil {
		log.Errorf("Error get privatekey the from db record: %v", err)
	}

	queryResult, err := praseQueryResult(query)
	if err != nil {
		log.Errorf("Error parse the database: %v", err)
	}
	if len(queryResult) != 1 {
		log.Errorf("privatekey not found")
	}
	// privatekey := strings.TrimSpace(queryResult[0].Privatekey)
	// fmt.Print("START")
	// fmt.Print(queryResult[0].Privatekey)
	// fmt.Print("END")
	return queryResult[0].Privatekey

}

//GetShellScript get from database
func GetShellScript(shellid string) string {
	db, err := sql.Open("mysql", "newuser:newuser@tcp(127.0.0.1:3306)/?charset=utf8")
	if err != nil {
		log.Errorf("Error connecting to database: %v", err)
	}
	query, err := db.Query(fmt.Sprintf("SELECT * FROM `remoteServer`.`script` WHERE `id` = %s", shellid))
	// fmt.Printf("SELECT * FROM `remoteServer`.`key` WHERE `ipaddress` = '%s' ", ip)
	if err != nil {
		log.Errorf("Error get privatekey the from db record: %v", err)
	}

	queryResult, err := praseQueryShellResult(query)
	if err != nil {
		log.Errorf("Error parse the database: %v", err)
	}
	if len(queryResult) != 1 {
		log.Errorf("privatekey not found")
	}
	return queryResult[0].Content

}

func praseQueryResult(query *sql.Rows) ([]IPList, error) {
	column, _ := query.Columns()
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}
	results := []IPList{}
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			return nil, err
		}
		row := IPList{}
		for k, v := range values {

			switch column[k] {
			case "ipaddress":
				row.Ipaddress = string(v)
			case "privatekey":
				row.Privatekey = string(v)
			default:

			}
		}
		results = append(results, row)

	}
	return results, nil
}

func praseQueryShellResult(query *sql.Rows) ([]ShellList, error) {
	column, _ := query.Columns()
	values := make([][]byte, len(column))
	scans := make([]interface{}, len(column))
	for i := range values {
		scans[i] = &values[i]
	}
	results := []ShellList{}
	for query.Next() {
		if err := query.Scan(scans...); err != nil {
			return nil, err
		}
		row := ShellList{}
		for k, v := range values {
			switch column[k] {
			case "content":
				row.Content = string(v)
			default:
			}
		}
		results = append(results, row)

	}
	return results, nil
}
