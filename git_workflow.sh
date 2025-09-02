# ./git_workflow.sh "test: testing the git workflow script"
#if permission denied, run: chmod +x git_workflow.sh

#!/bin/bash
# Exit immediately if a command exits with a non-zero status.
set -e

# Check if a commit message is provided
if [ -z "$1" ]; then
  echo "Error: Commit message is required."
  echo "Usage: ./git_workflow.sh \"Your commit message\""
  exit 1
fi

COMMIT_MESSAGE=$1

# Get the current branch name
CURRENT_BRANCH=$(git symbolic-ref --short HEAD)
echo "Current branch is $CURRENT_BRANCH"

# 1. Add all changes
echo "Staging changes..."
git add .

# 2. Commit the changes
echo "Committing with message: \"$COMMIT_MESSAGE\""
git commit -m "$COMMIT_MESSAGE"

# 3. Push the current branch to origin
echo "Pushing branch $CURRENT_BRANCH to origin..."
git push origin "$CURRENT_BRANCH"

# 4. Checkout to test branch
echo "Checking out to 'test' branch..."
git checkout test

# 5. Pull the latest changes on test branch
echo "Pulling latest changes on 'test' branch..."
git pull

# 6. Checkout back to the original branch
echo "Checking out back to '$CURRENT_BRANCH'..."
git checkout "$CURRENT_BRANCH"

# 7. Rebase the current branch with test
echo "Rebasing '$CURRENT_BRANCH' with 'test'..."
if ! git rebase test; then
  echo ""
  echo "⚠️  Rebase encountered conflicts that need manual resolution."
  echo "Please resolve the conflicts manually and then run:"
  echo "   git add ."
  echo "   git rebase --continue"
  echo ""
  echo "After resolving conflicts, press Enter to continue the script..."
  read -r
  echo "Continuing with the workflow..."
fi

# 8. Force push the rebased branch
echo "Pushing rebased branch to origin with --force-with-lease..."
git push --force-with-lease origin "$CURRENT_BRANCH"

echo "Workflow completed successfully!" 