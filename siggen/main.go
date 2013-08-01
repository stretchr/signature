package main

import (
	"fmt"
	"github.com/stretchr/commander"
	"github.com/stretchr/signature"
	"github.com/stretchr/stew/objects"
	"strconv"
)

// Generates a 32 character random signature
func main() {
	commander.Go(func() {

		commander.Map(commander.DefaultCommand, "", "",
			func(args objects.Map) {
				fmt.Println(signature.RandomKey(32))
			})

		commander.Map("len length=(int)", "Key length",
			"Specify the length of the generated key",
			func(args objects.Map) {
				length, _ := strconv.Atoi(args.Get("length").(string))
				fmt.Println(signature.RandomKey(length))
			})

	})
}
