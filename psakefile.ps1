task default -Depends list

task list {
    Get-PSakeScriptTasks
} -Description "List all tasks in the script"

task clean {
    go clean
} -Description "Clean up the project"

task deps {
    go mod download
    go mod tidy
} -Description "Download dependencies"

task fmt {
    go fmt ./...
} -Depends deps -Description "Format the code"

task vet {
    go vet ./...
} -Depends deps -Description "Run go vet"