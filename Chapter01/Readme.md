1. Modify the `echo` program to also print `os.Args[0]`, the name of the command
   that invoked it.

2. Modify the `echo` program to print the index and value of each of its
   arguments, one per line.

3. Experiment to measure the difference in running time between our potentially
   inefficient versions and the one that uses `strings.Join`. (Section 1.6
   illustrates part of the `time` package, and Section 11.4 shows how to write
   benchmark tests for systematic performance evaulation.)
