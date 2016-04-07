package main

import ("fmt"; "os"; "log"; "bufio"; "io/ioutil")

func check_err(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var fname string
	// The filename can be either given as command line
	// argument or read in from stdin.
	if len(os.Args) >= 2 {
		fname = os.Args[1]
	} else {
		_, err := fmt.Println("Please specify the filename:")
		check_err(err)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		fname = scanner.Text()
		err = scanner.Err()
		check_err(err)
	}
	_, err := fmt.Printf(">>> writing to %s\n", fname)
	check_err(err)
	// Open file for appending and create, if it doesn't exist
	f, err := os.OpenFile(fname, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0666)
	check_err(err)
	defer f.Close()
	writer := bufio.NewWriter(f)
	
	b, err := ioutil.ReadAll(os.Stdin)
	check_err(err)
	_, err = writer.WriteString(string(b))
	check_err(err)
	err = writer.Flush()
	check_err(err)
}
