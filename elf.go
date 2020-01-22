package main

import (
	"log"

	"github.com/szaqal/GoMiscModules/elf"
)

func main() {
	_, err := elf.ParseELFHeader("/Users/pawel.ma/devel/bin/docker-util-linux")
	if err != nil {
		log.Fatal("Error reading file")
	}
}
