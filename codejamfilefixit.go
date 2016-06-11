//File:		codejamstorecredit.go
//Author:	Gary Bezet
//Date:		2016-06-08
//Desc:		This program is designed to solve Google Code Jam "File Fix-it"
//Problem:	https://code.google.com/codejam/contest/635101/dashboard
//Results:	(native linux) SmallProblem:8.099313ms		LargeProblem:46.189754ms
ms

package main

//imports
import (
		"fmt"
		"flag"
		"os"
		"bufio"
		"strconv"
		"strings"
		"time"
	)
	
	
//global variables
var infileopt, outfileopt string  //input and output filenames
var infile, outfile *os.File  //input and output file pointers
var totalcases int  //number of cases 


var testcases []Testcase
	

//structures
type Testcase struct {
	num int
	
	curdirs []string
	newdirs []string
	solution int
	
}


//program entry point
func main() {

	starttime := time.Now()  //start time for stats

	defer infile.Close()
	defer outfile.Close()

	initflags()  //initialize the command line args
	
	openFiles() //open the files
	
	processFile()
	
	for _, v := range testcases {  //solve all cases
		fmt.Fprintf(outfile, "Case #%d: %d\n", v.num, v.solve())
		
	}
		
	printErrln("done!  Processed", totalcases, " in ", time.Now().Sub(starttime))
	
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

	starttime := time.Now()

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

	printErrln("Total cases: ", totalcases)
	
	testcases = make( []Testcase, totalcases )  //allocate structure for cases
	
	for i := 0; i < totalcases; i++  {  //increment through the cases

		testcases[i].num = i+1

		var curdirnum, newdirnum int //varables for storing current cases directories on computer and directories to be created	
		
		line, err = reader.ReadString('\n') 
		dirnums := strings.Split(strings.TrimSpace( line ), " ") 
		
		curdirnum, err = strconv.Atoi( dirnums[0] )  //get number of current directories
		if err != nil { 
			printErrln("Error converting cpudirnum: ", err)
			os.Exit(6)
		}
		
		newdirnum, err = strconv.Atoi( dirnums[1] ) //get number of new directories
		if err != nil {
			printErrln("Error converting newdirnum: ", err)
			os.Exit(7)
		}
		
		//printErrln("Reading Case#", i + 1, "CpuDirs=", curdirnum, "  NewDirs=", newdirnum)  //for testing
		
		testcases[i].curdirs = make( []string, curdirnum)
		for c := 0; c < curdirnum; c++ {  //read all directories on CPU
			line, err = reader.ReadString('\n')
			if( err != nil ) {
				printErrln("Error reading case#", i+1, " cpudirnum#", c+1)
				os.Exit(8)
			}
			
			testcases[i].curdirs[c] = strings.TrimSpace(line)
			
		}
		
		testcases[i].newdirs = make( []string, newdirnum )
		for n := 0; n < newdirnum; n++ {  //read all new directories to be creaed
			line, err = reader.ReadString('\n')
			if( err != nil ) {
				printErrln("Error reading case#", i+1, " newdirnum#", n+1)
				os.Exit(9)
			}
			
			testcases[i].newdirs[n] = strings.TrimSpace( line )
			
		}
		

	}
	
	printErrln("Read", totalcases, "in", time.Now().Sub(starttime) ,"\n")  //output time taken to read all cases
		
}

type Dir struct {
	dirs map[string]*Dir
}

//solve the cases here!!!  Store solution in self.solution
func (self Testcase) solve() int {
	
	var root Dir  //root directory always created
	root.dirs = make(map[string]*Dir)  //root directory	
	
	
	var created, curdirnum int = 0, 0 //created directories

	
	for _, v := range self.curdirs {  //go through current dir lines
		dirstructure := strings.Split(v, "/")  //split directories
		createdir(&curdirnum, &root, dirstructure[1:])  //ignore first dir because null
	}
	
	

	for _, v := range self.newdirs {  //go through current dir lines
		
		dirstructure := strings.Split(v, "/")  //split directories
		createdir(&created, &root, dirstructure[1:])  //ignore first dir because null
	}

	printErrln("Solved#", self.num, " answer=", created) 
	
	return created
	

}

func createdir( created *int,  root *Dir, structure []string ) {  //create a directory if it exists and increment created otherwise move on to next level


	if _, ok := root.dirs[structure[0]]; ok == false {  //create directory if null and increment created

		
		var newdir Dir
		newdir.dirs = make(map[string]*Dir)		
		root.dirs[structure[0]] = &newdir
		
		(*created)++
		
	}
	

	
	if len( structure) > 1  {		
		createdir( created, root.dirs[structure[0]], structure[1:]) //go to next level
	}
	


}
