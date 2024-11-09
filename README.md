Project URL: https://roadmap.sh/projects/task-tracker
This is CLI task tracker.
List cmd:
 - run command: go run main.go [flag] --[arg-1]=[value] --[arg-2]=[value]
   + flag = add / delete / update / list / mark-done / mark-in-progress
   + arg:
     * add: description #description is arg of add
     * delete: id
     * list: type = "todo" / "in-progress" / "done"
     * update: id description
     * mark-done: id
     * mark-in-progress: id
