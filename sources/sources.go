package sources

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type repo struct{
	// Declares the fields for a github repo
	owner string
	repo  string
}

type githubrepo interface{
	OwnerName()
	RepoName()
}

func ReadRepos (filePath string)(repoList []repo) {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// go over each line and split on the slash
		sp := strings.Split(scanner.Text(), "/")
		// append as a repo object to the array
		repoList = append(repoList, repo{owner: sp[0], repo: sp[1] })
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return repoList
}

func (r repo) OwnerName() string {
	return r.owner
}

func (r repo) RepoName() string {
	return r.repo
}