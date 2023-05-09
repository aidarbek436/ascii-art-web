# ascii-art-web
## Authors: Aidarbek, anargube
### Objectives

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version of your last project, ascii-art.

Your webpage must allow the use of the different banners:
- *shadow*
- *standard*
- *thinkertoy*

Implement the following HTTP endpoints:

- **GET** /: Sends HTML response, the main page.
- 1.1. **GET** Tip: go templates to receive and display data from the server.

- **POST** /ascii-art: that sends data to Go server (text and a banner)
- 2.1. **POST** Tip: use form and other types of tags to make the post request.\

The way you display the result from the POST is up to you. What we recommend are one of the following :

>   Display the result in the route /ascii-art after the POST is completed. So going from the home page to another page.
>   Or display the result of the POST in the home page. This way appending the results in the home page.

**The main page must have:**

- text input
- radio buttons, select object or anything else to switch between banners
- button, which sends a POST request to '/ascii-art' and outputs the result on the page.
### HTTP status code
Your endpoints must return appropriate HTTP status codes.

- OK (200), if everything went without errors.
- Not Found, if nothing is found, for example templates or banners.
- Bad Request, for incorrect requests.
- Internal Server Error, for unhandled errors.

## Usage

```
After cloning the repository, simply do a [go run main.go.] After that, connect to the localhost:4000 through your browser and start using the static website.
```

### Implementation details: algorithm

- After creating the server and using the default Mux, each connection gets handled by the handlers that are stored in the server folder.
- Each request is being checked for possible errors and if none are found, the parsed template with necessary information will be shows in the browser.
- When a client enters text inside the form and submits it, the POST form with an */ascii* url is being passed as a request to the server, which extracts the input values and puts them on a separate rendered template with the Ascii art logic.
- In order to check for form submission accuracy, ascii art has been modified to return a string with an error that will handle any ASCII ART related mistakes.