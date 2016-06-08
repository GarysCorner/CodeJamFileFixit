//File:		codejamstorecredit.go
//Author:	Gary Bezet
//Date:		2016-06-08
//Desc:		This program is designed to solve Google Code Jam "File Fix-it"
//Problem:	https://code.google.com/codejam/contest/635101/dashboard

package main

//imports
import (
		"fmt"
		"flag"
		"os"
		"bufio"
		"strconv"
		"strings"
	)
	
	
//global variables
var infileopt, outfileopt string
var infile, outfile *os.File
var totalcases int
	
	
func main() {

	defer infile.Close()
	defer outfile.Close()

	initflags()  //initialize the command line args
	
	openFiles() //open the files
	
	processFile()
}

//get the flags from command line
func initflags() {
	flag.StringVar(&infileopt, "if", "", "Input file (required)")
	flag.StringVar(&outfileopt, "of", "-", "Output file, defaults to stdout" )

	flag.Parse()

	if infileopt == "" {
		printErrln("You must supply an input file\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	

}

//print error to console
func printErrln( line ...interface{} ) {
	fmt.Fprintln( os.Stderr, line... )
}

func openFiles() {
	
	var err error

	infile, err = os.Open(infileopt)

	if err != nil {
		printErrln( "Error:  Could not open:  ", infileopt)
		printErrln( "\tError: ", err  )
		os.Exit(2)
	}

	if outfileopt == "-"  {
		outfile = os.Stdout
		outfileopt = "Stdout"
	} else {
		outfile, err = os.Create(outfileopt)

		if err != nil {
			printErrln( "Error:  Could not create:  ", outfileopt)
			printErrln( "\tError: ", err  )
			os.Exit(3)
		} 
	}

	printErrln("InFile:\t", infileopt)
	printErrln("OutFile:\t", outfileopt, "\n")
		
}

func processFile() {  //process the input file into data structure


	var err error	
	var line string
	reader := bufio.NewReader(infile)
	
	line, err = reader.ReadString('\n')
	if err != nil {
		printErrln("Couldn't read first line from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(4)
		
	}

	totalcases, err = strconv.Atoi( strings.TrimSpace( line  ) )
	if err != nil  { //if error reading number of cases
		printErrln("Couldn't read case numbers from:  ", infileopt)
		printErrln("\tError:  ", err )
		os.Exit(5)
	}

	printErrln("Cases: ", totalcases)
	
	for i := 1; i <= totalcases; i++  {  //increment through the cases

		var cpudirnum, newdirnum int //varables for storing current cases directories on computer and directories to be created	
		
		line, err = reader.ReadString('\n')
		dirnums := strings.Split(strings.TrimSpace( line ), " ")
		printErrln( len(dirnums) )
		cpudirnum, err = strconv.Atoi( dirnums[0] )
		newdirnum, err = strconv.Atoi( dirnums[1] )
		
		printErrln("\nCase#", i, "CpuDirs=", cpudirnum, "  NewDirs=", newdirnum)
		
		for c := 0; c < cpudirnum; c++ {
			line, err = reader.ReadString('\n')
		}
		
		for n := 0; n < cpudirnum; n++ {
			line, err = reader.ReadString('\n')
		}
	
	}
		
}

