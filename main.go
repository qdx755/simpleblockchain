package main

func main() {
	//fmt.Println("Hello Blockchain!")
	bc := NewBlockChain()
	defer bc.db.Close()
	cli := CLI{bc}
	cli.Run()
}
