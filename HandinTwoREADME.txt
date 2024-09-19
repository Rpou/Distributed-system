							HandinTwo **README**

a) What are packages in your implementation? What data structure do you use to transmit data and meta-data?

	In our implementation, we send SEQ and ACK numbers between the servers. Those are our packages. We also send messages
	after the connection has been established. We use channels in GO to transmit data between the client and server.

b) Does your implementation use threads or processes? Why is it not realistic to use threads?

	We use threads in our implementation, as we use the GO language. It is not realistic in the real world, 
	because gorutines aren't good at handling message delay/faliures

c) In case the network changes the order in which messages are delivered, how would you handle message re-ordering?
	
	A solution to this problem could be sending a timestamp along with the package. Then the reciever could sort it
	based on that timestamp, instead of when the package was recieved.


d) In case messages can be delayed or lost, how does your implementation handle message loss?

	In our implementation, if the message is delayed, the whole process is just halted and waits for a message.
	If the message never reaches the server or client, the program will simply wait forever for a response. 
	This is not ideal in the real world...


e) Why is the 3-way handshake important? 
	
	It is important, because it initializes contact between the server and client. It is also used for security
	reasons	as it establish a secure connection between the server and client. 
	Furthermore it greatly increase the chances of reliable flow of data from client to server. 
	You could say that the handshake is the building of a bridge where the data can quickly walk over safely and fast.
	