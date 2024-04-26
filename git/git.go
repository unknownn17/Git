package git

import (
	"fmt"
	"os"
	"os/exec"
)

func GetUserName() {
	// Fetch git username
	cmd := exec.Command("git", "config", "user.name")
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error fetching Git username:", err)
		return
	}
	username := string(out)

	// Writing username to a file
	usernameFile, err := os.OpenFile("username.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer usernameFile.Close()
	if _, err := usernameFile.WriteString(username); err != nil {
		fmt.Println("Error writing username to file:", err)
		return
	}

	// Command to fetch Git user email
	cmd = exec.Command("git", "config", "user.email")
	out, err = cmd.Output()
	if err != nil {
		fmt.Println("Error fetching Git user email:", err)
		return
	}
	userEmail := string(out)

	// Appending user email to the file
	if _, err := usernameFile.WriteString(userEmail); err != nil {
		fmt.Println("Error appending user email to file:", err)
		return
	}

	// Adding and committing changes
	if _, err := exec.Command("git", "add", ".").Output(); err != nil {
		fmt.Println("Error adding changes:", err)
		return
	}
	if _, err := exec.Command("git", "commit", "-m", "Added username and email").Output(); err != nil {
		fmt.Println("Error committing changes:", err)
		return
	}

	// Check if the feature/add-username branch exists
	branchExists := branchExists("feature/add-username")
	if !branchExists {
		//fmt.Println("Error: feature/add-username branch does not exist")
		return
	}

	// Switch to the feature/add-username branch
	if _, err := exec.Command("git", "checkout", "feature/add-username").Output(); err != nil {
		fmt.Println("Error switching to feature/add-username branch:", err)
		return
	}

	// Merge the changes from feature/add-user-email branch into feature/add-username branch
	if _, err := exec.Command("git", "merge", "feature/add-user-email").Output(); err != nil {
		fmt.Println("Error merging branches:", err)
		return
	}

	// Push changes to Github
	if _, err := exec.Command("git", "push", "origin", "feature/add-username").Output(); err != nil {
		fmt.Println("Error pushing changes to Github:", err)
		return
	}

	fmt.Println("All tasks completed successfully.")
}

// branchExists checks if the given branch exists
func branchExists(branchName string) bool {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branchName)
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}
