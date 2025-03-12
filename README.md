# ConcurrentRequestSender

A Go tool that sends HTTP requests concurrently to multiple endpoints with a configurable concurrency limit. It processes requests in batches, handling results and errors efficiently.

How to Use:

./ConcurrentRequestSender -num <number_of_requests> <url_1> <url_2> ... <url_n>
The app will send HTTP requests concurrently with a batch limit and print the results.
