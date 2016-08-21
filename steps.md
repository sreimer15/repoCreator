Step one get user credential by reading from stdout
it will use user.name

git config user.name

How does it know you are authenticated?

Basic Authentication

curl -u "username" https://api.github.com
(Usaername is sreimer15) then it makes you put in a password

(more likely than not just have you do it )



Create

Create a new repository for the authenticated user.

Use username to go to correct url

POST /user/repos

take name from the root directory name
-p flag for private repo (optional)

from the response object get the url
	"url": "https://api.github.com/repos/octocat/Hello-World",


// Then I need to figure out how to set remote origin from api 
which shouldn't be too bad

// maybe jsut make go run

git remote add origin "use saved url"


why not jsut run git push origin master