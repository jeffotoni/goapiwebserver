# Go Server Front-End
A server Go to render html and serve the front end to communicate with an API in the back end.
The server is an example of using libs html / template in Go, go-bindata was used to generate a .go with the templates in html that was created and thus we maintain the legibility of the templates in native Html and optimizing even more our server .

We removed the lib fmt from our front-end server, we used the lib bufio by typing straight into stdout.

The front-end server is independent and communicates with an API on the back end.

## API
we have two apis, a front-end server that renders html and communicates via rEST with backend server.
The backend is in another folder.

 - Front-End
 - Back-End