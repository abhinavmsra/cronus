# cronus
Time Logger (JIRA+Harvest)

# How does it work??

- This script checks for task id in branch names like `feature/auth-CODE-123`
- If the task id could not be found, it exits gracefully. (Needed for branches like main, develop)
- If a task id is present in branch name, it fetches task details from JIRA & creates a new entry in Harvest

# Configuration

Create a `cronus.json` file at the root of git repo.

```json
{
  "jira": {
    "code": "XXXX",
    "organization": "XXXX",
    "email": "USER EMAIL",
    "accessToken": "USER ACCESS TOKEN"
  },
  "harvest": {
    "access_token": "HARVEST ACCESS TOKEN",
    "account_id": "USER ACCOUNT ID",
    "project_id": "PROJECT ID",
    "task_id": "TASK ID TO LOG TIME INTO"
  }
}
```

# Use

- **Manual**: Checkout to a desired branch & simply run `cronus`

- **AutoLog (with git-hooks)**: 

1. Create a `post-checkout` file in `.git/hooks`

2. Add following:
      
      ```bash
      #!/bin/sh
      #
      # Log Harvest Entry
      cronus
      ```
This will run the script on every git checkout & create a time entry each time you work on a new task.
