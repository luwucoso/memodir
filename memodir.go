package main

import (
  "fmt" 
  "os"
  "log"
  "path/filepath"
  "strings"
  "bufio"
  "sort"
)

import "github.com/lithammer/fuzzysearch/fuzzy"

func main() {
    argsWithoutProg := os.Args[1:]

    if(argsWithoutProg[0] == ".add") {
      addCwd()
    } else if argsWithoutProg[0] == ".help" {
      fmt.Fprintln(os.Stderr, `memodir | v 0.0.1
Prerequisite: create ~/.config/memodir/default.txt
To show this help text: memodir .help
To add the cwd: memodir .add
To search: memodir <your query>
Example use: cd "$(memodir <your query>)"
                              `)
    } else {
      find(argsWithoutProg[0])
    }

}

func getBucketPath() string {
  home_dir, home_err := os.UserHomeDir()
  if home_err != nil {
    log.Fatal(home_err)
  }

  config_path := filepath.Join(home_dir, ".config/memodir")
  file_path := filepath.Join(config_path, "default.txt")
  return file_path
}

func find(search string) {
  file_path := getBucketPath()
  file, err := os.Open(file_path)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := strings.TrimSpace(scanner.Text())
    if text == "" {
      continue
    }

    lines = append(lines, text)
  }

  matched := fuzzy.RankFindNormalizedFold(search, lines)
  sort.Sort(matched)

  for i, m := range matched {
    fmt.Fprintln(os.Stderr, i, m.Target)
  }

  var i int
  fmt.Fprintln(os.Stderr, "Choose a match")
  fmt.Scan(&i)

  fmt.Print(matched[i].Target)
}

func addCwd() {
  cwd, err := os.Getwd()
  if err != nil {
    log.Println(err)
    return;
  }

  file_path := getBucketPath()

  fmt.Printf("Adding %s to config in %s\n", cwd, file_path)

  bucket_file, err := os.OpenFile(file_path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

  if err != nil {
    log.Fatal(err)
  }

  if _, err := bucket_file.Write([]byte("\n" + cwd)); err != nil {
    log.Fatal(err)
  }

  if err := bucket_file.Close(); err != nil {
    log.Fatal(err)
  }

}
