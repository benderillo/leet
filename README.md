# leet
Solutions for leetcode programming challenges https://leetcode.com/problemset/top-interview-questions/

# How to run
Each file contains single challenge and single or sometimes multiple solutions.
Each file defines `main()` entry that contains some test code.
You can't however, run the `main()` app entrypoint outide of the `main` package.
So the easiest way to run the solution would be to rename the package at the top of the file to `main` and execute:

```
go run <filename.go>
```

In the future I may improve it by adding filename_test.go files for each of the challenge file with one test function that will execute tests. Then it will be easy to test challenges by running `go test` command.

But this future may never come...
