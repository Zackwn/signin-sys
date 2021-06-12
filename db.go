package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Database struct {
	filepath string
	file     *os.File
	items    []*Person
}

func LoadItems(file *os.File) []*Person {
	var items []*Person
	parsePerson := func(p string) {
		info := strings.Split(p, "|")
		if len(info) != 2 {
			return
		}
		name := info[0]
		age, err := strconv.Atoi(info[1][:len(info[1])])
		if err != nil {
			log.Fatalln(err)
		}
		items = append(items, &Person{name: name, age: age})
	}
	buffer := make([]byte, 32*1024)
	for {
		bufferSize, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}

		p := ""
		for _, char := range buffer[:bufferSize] {
			if char == '\n' {
				parsePerson(p)
				p = ""
			} else {
				p += string(char)
			}
		}
		// parse miss buffered data slice
		if len(p) != 0 {
			buff := make([]byte, 1)
			for {
				_, err := file.Read(buff)
				if err == io.EOF {
					break
				} else if err != nil {
					log.Fatalln(err)
				}
				p += string(buff)
			}
			parsePerson(p[:len(p)-1]) // remove \n
		}
	}
	return items
}

func NewDB() *Database {
	db := new(Database)
	db.filepath = "db.txt"
	var err error
	db.file, err = os.OpenFile(db.filepath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		log.Fatalln(err)
	}
	db.items = LoadItems(db.file)
	return db
}

func (db *Database) Insert(p *Person) {
	strP := p.name + "|" + strconv.Itoa(p.age) + "\n"
	_, err := db.file.Write([]byte(strP))
	if err != nil {
		log.Fatalln(err)
	}
	db.items = append(db.items, p)
}

func (db *Database) List() []*Person {
	return db.items
}
