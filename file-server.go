package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	print(`
  ______   _   _             _____                                     
 |  ____| (_) | |           / ____|                                    
 | |__     _  | |   ___    | (___     ___   _ __  __   __   ___   _ __ 
 |  __|   | | | |  / _ \    \___ \   / _ \ | '__| \ \ / /  / _ \ | '__|
 | |      | | | | |  __/    ____) | |  __/ | |     \ V /  |  __/ | |   
 |_|      |_| |_|  \___|   |_____/   \___| |_|      \_/    \___| |_|
 		
 Developed by Prappo Prince
 Example Use : ./file-server -p=8000 -d="./files"
  `)
	port := flag.String("p", "8000", "port to serve on")
	directory := flag.String("d", ".", "the directory of static file to host")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}