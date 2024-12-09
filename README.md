Hand in 5 Replication
To run the program clone the repository, then open at least 3 terminals but we suggest 4.

cd to Replication/Replication and run the commands:

Teminal window 1: (primary server) just run "go run server.go" 
Teminal window 2: (backup server) just run "go run server.go --port=50052" 
Teminal window 3: (client server) just run "go run server.go --port=50053" 
Teminal window 4: (client server) just run "go run server.go --port=50055" 

The current program is set up to crash if you want to see if it works without the crash,
1: comment out the import package "os" on line 9
2: comment out 
        time.Sleep(10 * time.Second)
		os.Exit(1)
on line 422 and 423 and then run it like before

Lastly we had a strange problem where every time you run the program windows firewall complains and asks for pemission to run
Because of this you need to be quick. Have each terminal ready and quickly run them all before the popup window appears.