package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v2"
)


type Model struct {
	Uplinks []Link `yaml:"uplinks"`
	Peers []Peer   `yaml:"peers"`
	ASN int 	   `yaml:"asn"`
	Loopback Addr  `yaml:"loopback"`
}

type Link struct {
	Name string   `yaml:"name"`
	Prefix string `yaml:"prefix"`
}

type Peer struct {
	IP string `yaml:"ip"`
	ASN int `yaml:"asn"`
}

type Addr struct {
	IP string `yaml:"ip"`
}


func main() {
	src, err := os.Open("input.yml")
	// 에러를 처리합니다.
	defer src.Close()
	d := yaml.NewDecoder(src)

	var input Model
	err = d.Decode(&input)
	// 에러를 처리합니다.
}