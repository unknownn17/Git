package main

import (
	"main.go/git"
)

// Requrements:
// 1. Create a new repository on Github.
// 2. Clone it to your local directory.
// 3. Create a package named `git` outside the `main` package.
// 4. Create a function named `GetUserName` inside the `git` package.
// 5. Create a new branch named `feature/add-username` and switch to it from the `main branch`.
// 6. Fetch the `git` username inside the `GetUserName` function (hint: using `exec.Command` method).
// 7. Write the username to a new file.
// 8. Add and commit all changes.
// 9. Create a new branch named `feature/add-user-email` and switch to it from `feature/add-username` branch.
// 10. Create a function named `GetUserEmail` and fetch the `git` user email using `exec.Command`.
// 11. Append the user email to the previously created file.
// 12. Add and commit all changes.
// 13. Switch back to the `feature/add-username` branch.
// 14. Merge the changes from `feature/add-user-email` branch into `feature/add-username` branch.
// 15. Push the `feature/add-username` branch to the Github repository.

func main() {
	git.GetUserName()
}
