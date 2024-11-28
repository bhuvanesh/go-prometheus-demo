#!/bin/bash
# install-hooks.sh

# Colors for output
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Create hooks directory if it doesn't exist
mkdir -p .git/hooks

# Copy commit-msg hook
cp scripts/git-hooks/commit-msg .git/hooks/
chmod +x .git/hooks/commit-msg

# Optional: Copy prepare-commit-msg hook for commit template
cat > .git/hooks/prepare-commit-msg << 'EOF'
#!/bin/bash

# Only add template if there's no commit message already
if [ -z "$(cat $1)" ]; then
  cat << 'END' > $1
# Commit Message Format:
# type(scope): subject
#
# Types:
#   feat     (new feature)
#   fix      (bug fix)
#   docs     (documentation)
#   style    (formatting)
#   refactor (refactoring)
#   perf     (performance)
#   test     (testing)
#   build    (build system)
#   ci       (CI/CD)
#   chore    (maintenance)
#   revert   (revert changes)
#
# Example:
#   feat(auth): add JWT authentication
#
END
fi
EOF

chmod +x .git/hooks/prepare-commit-msg

# Optional: Copy pre-push hook for branch naming convention
cat > .git/hooks/pre-push << 'EOF'
#!/bin/bash

# Regular expression for branch naming convention
BRANCH_REGEX="^(feature|bugfix|hotfix|release)/[a-z0-9-]+$"

# Get current branch name
BRANCH_NAME=$(git symbolic-ref --short HEAD)

if ! [[ $BRANCH_NAME =~ $BRANCH_REGEX ]]; then
    echo "ERROR: Branch name '$BRANCH_NAME' does not follow the convention:"
    echo "  feature/feature-name"
    echo "  bugfix/bug-description"
    echo "  hotfix/issue-description"
    echo "  release/version-number"
    exit 1
fi

exit 0
EOF

chmod +x .git/hooks/pre-push

echo -e "${GREEN}Git hooks installed successfully${NC}"