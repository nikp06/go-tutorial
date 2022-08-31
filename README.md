# go-tutorial
	Program for

- retrieving all the topics from the tutorial gobyexample,
- retrieving their html via http request,
- editing and checking it line by line and depending on whether it is code or comment
- writing the content to a go-file.
In the end all 80 topics are crawled and the comments and lines of code are written to an executable go file.
Program was written to make use of the knowledge gained whithin the tutorial especially
- the basics
- methods
- string formatting
- regular expressions
- errors
- goroutines
- waitgroups
Cool thing:
- by using goroutines the execution time is reduced from ~30 seconds (not using goroutines in the main loop) to ~1.3 seconds
What I would like to have used but didn't
- interfaces!
- channels!
- case and select!
