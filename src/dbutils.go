package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"os"
)

func printDb(db map[string]string) {
	for k, v := range db {
		log.Println(k, v)
	}
}

func saveDb(db map[string]string) {
	buf := encodeDb(db)
	// save buf to file
	file, err := os.OpenFile(dbPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	file.Write(buf.Bytes())
	defer file.Close()
}

func existsDb() bool {
	_, err := os.Stat(dbPath)
	return err == nil
}

func loadDb() map[string]string {
	file, err := os.Open(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := &bytes.Buffer{}
	buf.ReadFrom(file)

	return decodeDb(buf)
}

func encodeDb(data any) *bytes.Buffer {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)

	err := enc.Encode(data)
	if err != nil {
		log.Fatal("encode error:", err)
	}

	return buf
}

func decodeDb(buf *bytes.Buffer) map[string]string {
	dec := gob.NewDecoder(buf)
	var data map[string]string

	err := dec.Decode(&data)
	if err != nil {
		log.Fatal("decode error:", err)
	}

	return data
}
