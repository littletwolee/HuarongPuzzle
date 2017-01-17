package models

import (
	"bytes"
	"encoding/gob"
	"io/ioutil"
	"os"
)

const (
	FILE_ROOT = "./data/"
)

func FileSave(data map[string]*Piece, filename string) error {
	b, err := Encode(data)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("./"+filename, b, 0666) //写入文件(字节数组)
	if err != nil {
		return err
	}
	return nil
}

func FileRead(filename string) error {
	var data map[string]*Piece
	b, err := ioutil.ReadFile("./" + filename)
	if err != nil {
		return err
	}
	if err := Decode(b, &data); err != nil {
		return err
	}
	PersonTable = data
	return nil
}

func FileDelete(filename string) error {
	return os.Remove("./" + filename)
}

func Encode(data interface{}) ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}
