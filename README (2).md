# MP1
#Daniel Flores and Nicholas Hillis



----
# To Run
This program needs two instances of terminal to run correctly. In one instance type the following


```bash
cd server (to get to server folder)
go run unicastReceive.go
``` 
What this does is it opens our concurrent server which will listen to any number of clients

To access the client run our main go file
```bash
go run main.go
```
You will now be prompted to enter as many messages as there are instances in the config text file of processes
```bash
Example Code from client:
Send message: hello
Sent hello to (port) system time is "xxx"
```
The time here is a placeholder value and should display correctly after sending a message. We ended up using ports instead of processes because we could not find a way to get the process ids from the text file

```bash
Example Code from server:
Received hello from (port) system time is "xxx"
```
----
### Processes

Our concurrent TCP server took most of our time before the extension, so we did not have time to attempt importing the config or making the delay work but the lecture gave us ideas on how to approach those. We used the After function David Reps mentioned in lecture as the timeout was not working unless put to sleep which was advised against in the problem set
We were not able to figure out a way to get the read file function to work line by line from the config file so we simulated it in a sort of brute force method as Professor Tseng suggested if we had no other options.
### Sources

The following sites were used to help create our code:
```text
https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/#test-the-concurrent-tcp-server
https://opensource.com/article/18/5/building-concurrent-tcp-server-go
https://blog.golang.org/concurrency-timeouts 
https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/applevelprotocols/simple_example.html
http://zetcode.com/golang/readfile/ 
https://github.com/ashe7kh/TCP-Email-Program 
https://stackoverflow.com/questions/16465705/how-to-handle-configuration-in-go 
https://medium.com/@onexlab.io/golang-config-file-best-practise-d27d6a97a65a 
https://github.com/tkanos/gonfig 
https://golang.org/pkg/encoding/gob/ 
https://github.com/matthieuberger/go-tcp-serverhttps://github.com/matthieuberger/go-tcp-server 
https://stackoverflow.com/questions/39491435/how-to-import-structs-from-another-package-in-go