This is a API made for an interview at Teachable.


To run it at port 8080, you need to add a ".env" file inside "./src/config" folder with 2 variables: TEACHABLE_BASE_URL=https://developers.teachable.com/v1 and API_KEY=XXXXXXX.
Then you can simple build and run a docker image with:

docker build . -t teachable-test-carlos

docker run -p 8080:8080 teachable-test-carlos


The API has only one route which is GET /courses where you retrieve information of all published courses at the given school, for each course it output information about: Course name, Course Heading and Enrolled students

CURL for Request: curl --location 'localhost:8080/courses'
