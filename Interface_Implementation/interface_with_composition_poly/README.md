Interface with Composition and Polymorphism
🧩 Problem:
Design an application for logging different types of messages (ConsoleLogger, FileLogger, RemoteLogger) that implement a common Logger interface.
📝 Requirements:
Create a Logger interface with Log(message string).


Implement:


ConsoleLogger that logs to fmt.Println.


FileLogger that logs to a file (simulate by writing to a string or slice).


RemoteLogger that logs to a fake remote server (simulate with a print statement).


Write a function LogAll(loggers []Logger, message string) that sends the same message to all loggers.


✅ Output Example:
Console: Hello!
File: Hello!
Remote: Hello!






