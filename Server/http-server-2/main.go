package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			method := strings.Fields(ln)[0]
			fmt.Println("WE PRINTED THIS - METHOD: ", method)

		} else {
			if ln == "" {
				break
			}
		}

		i++
	}

	body := `<!DOCTYPE html>
<html lang="en">
<head>
	<title>Project 2</title>
	<link rel="stylesheet" type="text/css" href="style.css">
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<link rel="icon" type="image/png" href="James Masterpiece.png">
</head>

<body>
	<header>
		<nav class = "main-nav">
			<a href="home.html">Home</a>
			<a href="projects.html">Portfolio</a>
			<a href="about.html">About</a>
			<a href="contact.html">Contact</a>

		</nav>
	</header>
		<div class="title">
			<h1>Mikey<span>Robby</span></h1>



			<p>I am a 21 year old Computer Engineering student at Fresno City College in Fresno, California.</p>
			<p>I have some programming experience in:</p>
			<ul>
				<li>Java</li>
				<li>Android</li>
				<li>HTML/CSS/JavaScript</li>
				<li>Python</li>

			</ul>
		</div>

		<div class="btn">
			<a href="projects.html">Projects</a>
			<a href="about.html">About</a>
			<a href="contact.html">Contact</a>
		</div>







</body>


</html>`
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func main() {
	server, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer server.Close()

	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatalln(err.Error())
		}
		go handleConn(conn)
	}
}
