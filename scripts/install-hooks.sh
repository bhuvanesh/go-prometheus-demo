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

echo -e "${GREEN}Git hooks installed successfully${NC}"
