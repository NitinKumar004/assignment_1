log analyzer cli (file health checker)

a simple command line tool written in Go to analyze log file by counting
number of information like

INFO , WARNING ERROR
Feature:
read a log file passed via command-line interface
count :
   INFO,WARNING,ERROR
calculate weightage of each log type

show total number of log lines

display timestamp of analysis
 USES:

  switch : use to count each log type classification'
  defer: for file closing the defer file.Close() this will close file after main file exceution


   


