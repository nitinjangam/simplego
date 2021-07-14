# simplego
Simple GoLang micro-service - To-Do service

# Framework used - gorilla/mux
___________________________________________________________________
# Endpoint      |     Method      |     Description               |
___________________________________________________________________
/addtask        |     POST        |     Add task to task list     |
___________________________________________________________________
/removetask     |     DELETE      |     Remove task from list     |
___________________________________________________________________
/updatelist     |     PUT         |     Update task from list     |
___________________________________________________________________
/getlist/{name} |     GET         |     Get all tasks from list   |
___________________________________________________________________

# Request example : (POST/PUT/DELETE)
{
  "name"  : "user1",
  "tasks" : {
    "taskName" : "Task1",
    "status"   : "incomplete"
  }
}

# Response example : (POST/PUT/DELETE)
{
  "name"  : "user1",
  "tasks" : {
    "taskName" : "Task1",
    "status"   : "incomplete"
  }
}
