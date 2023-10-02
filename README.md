# Purpose  
Building a small docker image size website (Size) with (almost) static web page with UI.  

# Small Website with UI using Go 
A very small and minimal website built using Go.   


## view.html
This is a template which gets replaced at runtime whenever the URL is hit, Populates dates in table.  

# Learning Sources
GoLang tutorial: https://go.dev/doc/articles/wiki/  
Deploying Docker application using DockerFile https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker  
For reducing the size of the docker image: https://klotzandrew.com/blog/smallest-golang-docker-image  

# Docker command to run 
`docker build -t small-website-golang:latest .`  
`docker run --rm -p 9000:9000 --name go-docker-app small-website-golang:latest`

URL: localhost:9000/view/demoTitle
`demoTitle` can be replaced as required to print a different heading.

# Result: Small website of size ~10MB
